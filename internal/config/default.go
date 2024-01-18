package config

var defaultConfig = map[string]interface{}{
	"serve.port":                   5000,
	"serve.cors.enabled":           false,
	"serve.cors.allowed_origins":   []string{"*"},
	"serve.cors.allowed_methods":   []string{"GET", "POST"},
	"serve.cors.allow_headers":     []string{"Authorization", "Content-Type", "Cookie"},
	"serve.cors.expose_headers":    []string{"Content-Type", "Set-Cookie"},
	"serve.cors.allow_credentials": true,
	"serve.timeout":                5000,

	"logging.level":       -1,
	"logging.encoding":    "console",
	"logging.development": true,
}
