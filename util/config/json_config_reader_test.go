package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfigFileExists(t *testing.T) {
	content := []byte(`{"port": "8080", "log_level": "INFO"}`)
	tmpfile, err := os.CreateTemp("", "test_config.json")
	assert.NoError(t, err, "Error creating temporary config file")
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		assert.Fail(t, "Error writing to temporary config file")
		return
	}
	tmpfile.Close()

	jc := JSONConfig{}
	config, err := jc.ReadConfig(tmpfile.Name())

	assert.NoError(t, err, "Expected no error for an existing config file")
	assert.NotNil(t, config, "Expected a valid config object")
	assert.Equal(t, "8080", config.Port, "Expected 'port' value to be '8080'")
	assert.Equal(t, "INFO", config.LogLevel, "Expected 'log_level' value to be 'INFO'")
}

func TestReadConfigFileNotExists(t *testing.T) {
	jc := JSONConfig{}
	_, err := jc.ReadConfig("non_existing_config.json")

	assert.Error(t, err, "Expected an error for a non-existing config file")
}
