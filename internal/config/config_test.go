package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cfg, err := New("")
	assert.NoError(t, err)

	// server configs
	equal(t, 8080, defaultConfig["serve.port"], cfg.ServeConfig.Port)
	// logging configs
	equal(t, -1, defaultConfig["logging.level"], cfg.LoggingConfig.Level)
	equal(t, "console", defaultConfig["logging.encoding"], cfg.LoggingConfig.Encoding)
	equal(t, true, defaultConfig["logging.development"], cfg.LoggingConfig.Development)
}

func TestLoadWithEnv(t *testing.T) {
	// given
	err := os.Setenv("KHARON_SERVE_PORT", "4000")
	assert.NoError(t, err)

	// when
	cfg, err := New("")

	// then
	assert.NoError(t, err)
	assert.Equal(t, 4000, cfg.ServeConfig.Port)
}

func TestLoadWithConfigFile(t *testing.T) {
	// given
	err := os.Setenv("KHARON_SERVE_PORT", "4000")
	assert.NoError(t, err)

	config := `
serve:
  port: 5000
`
	tempFile, err := ioutil.TempFile(os.TempDir(), "article-server-test")
	assert.NoError(t, err)
	fmt.Println("Create temp file::", tempFile.Name())
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(config)
	assert.NoError(t, err)

	// when
	cfg, err := New(tempFile.Name())

	// then
	assert.NoError(t, err)
	assert.Equal(t, 5000, cfg.ServeConfig.Port)
}

func TestMarshalJSON(t *testing.T) {
	conf, err := New("")
	assert.NoError(t, err)
	data, err := json.Marshal(conf)
	assert.NoError(t, err)

	var configMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &configMap))
	assert.True(t, strings.HasPrefix(configMap["db.dataSourceName"].(string), "root:****@tcp"))
	assert.Equal(t, "****", configMap["jwt.secret"])
}

func equal(t *testing.T, expected interface{}, values ...interface{}) {
	for _, v := range values {
		assert.EqualValues(t, expected, v)
	}
}

func equalDuration(t *testing.T, expected time.Duration, values ...interface{}) {
	for _, v := range values {
		if str, ok := v.(string); ok {
			d, err := time.ParseDuration(str)
			assert.NoError(t, err)
			assert.EqualValues(t, expected, d)
			continue
		}
		assert.EqualValues(t, expected, v)
	}
}
