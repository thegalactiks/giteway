package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	ServeConfig   ServeConfig   `json:"serve" yaml:"serve"`
	LoggingConfig LoggingConfig `json:"logging" yaml:"logging"`
}

type ServeConfig struct {
	BaseUrl string `json:"base_url" yaml:"base_url"`
	Port    int    `json:"port" yaml:"port"`
	Cors    struct {
		Enabled          bool     `json:"enabled" yaml:"enabled"`
		AllowOrigins     []string `json:"allowed_origins" yaml:"allowed_origins"`
		AllowedMethods   []string `json:"allowed_methods" yaml:"allowed_methods"`
		AllowHeaders     []string `json:"allow_headers" yaml:"allow_headers"`
		ExposeHeaders    []string `json:"expose_headers" yaml:"expose_headers"`
		AllowCredentials bool     `json:"allow_credentials" yaml:"allow_credentials"`
	} `json:"cors" yaml:"cors"`
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}

type LoggingConfig struct {
	Level       int    `json:"level"`
	Encoding    string `json:"encoding"`
	Development bool   `json:"development"`
}

var k = koanf.New(".")

var _defaultPrefix = "GITEWAY_"

func New(configFilePath string) (*Config, error) {
	if configFilePath != "" {
		if _, err := os.Stat(configFilePath); err != nil {
			log.Fatalf("the configuration file has not been found on %s", configFilePath)

			return nil, err
		}
	}

	// load from default config
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)
	if err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	// load from env
	err = k.Load(env.Provider(_defaultPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, _defaultPrefix)), "_", ".", -1)
	}), nil)
	if err != nil {
		log.Printf("error loading config from env: %v", err)
	}

	// load from config file if exist
	if configFilePath != "" {
		path, err := filepath.Abs(configFilePath)
		if err != nil {
			log.Fatalf("failed to get absolute config path for %s: %v", configFilePath, err)
			return nil, err
		}
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
			return nil, err
		}
	}

	var cfg Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		log.Printf("failed to unmarshal with conf: %v", err)
		return nil, err
	}
	return &cfg, err
}
