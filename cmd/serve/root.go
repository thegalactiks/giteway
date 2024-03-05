package serve

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/timeout"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	otelmetric "go.opentelemetry.io/otel/sdk/metric"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/thegalactiks/giteway/api"
	"github.com/thegalactiks/giteway/internal/config"
	"github.com/thegalactiks/giteway/internal/logging"
	"github.com/thegalactiks/giteway/internal/otel"
)

func timeoutMiddleware(timeoutMS time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(timeoutMS*time.Millisecond),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(func(ctx *gin.Context) {
			ctx.String(http.StatusRequestTimeout, "timeout")
		}),
	)
}

func NewServeCmd(configFile string) (serveCmd *cobra.Command) {
	cfg, err := config.New(configFile)
	if err != nil {
		log.Fatal(err)
	}
	logging.SetConfig(&logging.Config{
		Encoding:    cfg.LoggingConfig.Encoding,
		Level:       zapcore.Level(cfg.LoggingConfig.Level),
		Development: cfg.LoggingConfig.Development,
	})
	defer logging.DefaultLogger().Sync()

	tp := otel.InitTracerProvider()
	mp := otel.InitMeterProvider()

	serveCmd = &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			app := fx.New(
				fx.Supply(cfg),
				fx.Supply(logging.DefaultLogger().Desugar()),
				fx.Supply(tp),
				fx.Supply(mp),
				fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
					return &fxevent.ZapLogger{Logger: log.Named("fx")}
				}),
				fx.Invoke(
					printAppInfo,
				),
				fx.Provide(
					api.NewHandler,
					newHTTPServer,
				),
				fx.Invoke(
					api.Routes,
					func(r *gin.Engine) {},
				),
			)
			app.Run()
		},
	}

	return serveCmd
}

func newHTTPServer(lc fx.Lifecycle, tp *oteltrace.TracerProvider, mp *otelmetric.MeterProvider, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	logger := otelzap.New(zap.NewExample())
	defer logger.Sync()

	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	otelzap.L().Info("replaced zap's global loggers")
	otelzap.Ctx(context.TODO()).Info("... and with context")

	r.Use(otelgin.Middleware("giteway"))
	r.Use(requestid.New())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	if cfg.ServeConfig.Cors.Enabled {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     cfg.ServeConfig.Cors.AllowOrigins,
			AllowMethods:     cfg.ServeConfig.Cors.AllowedMethods,
			AllowHeaders:     cfg.ServeConfig.Cors.AllowHeaders,
			ExposeHeaders:    cfg.ServeConfig.Cors.ExposeHeaders,
			AllowCredentials: cfg.ServeConfig.Cors.AllowCredentials,
		}))
	}
	r.Use(timeoutMiddleware(cfg.ServeConfig.Timeout))
	r.Use(gin.Recovery())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServeConfig.Port),
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.FromContext(ctx).Infof("starting server on port:%d", cfg.ServeConfig.Port)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logging.DefaultLogger().Errorw("failed to close http server", "err", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			tp.Shutdown(ctx)
			mp.Shutdown(ctx)
			logging.FromContext(ctx).Info("server shutdown")
			return srv.Shutdown(ctx)
		},
	})
	return r
}

func printAppInfo(cfg *config.Config) {
	b, _ := json.MarshalIndent(&cfg, "", "  ")
	logging.DefaultLogger().Infof("application information\n%s", string(b))
}
