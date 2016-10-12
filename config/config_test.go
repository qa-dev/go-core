package config

import (
	"testing"
)

type TestConfig struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

func TestSuccess(t *testing.T) {
	var config *TestConfig
	err := LoadFromFile("./valid.json", &config)

	if err != nil {
		t.Error("Expected no errors, got ", err)
	}

	if config.Property1 != "value1" {
		t.Error("Expected value1, got ", config.Property1)
	}

	if config.Property2 != "value2" {
		t.Error("Expected value2, got ", config.Property2)
	}
}

func TestFailInvalidJson(t *testing.T) {
	var config *TestConfig
	err := LoadFromFile("./invalid.json", &config)

	if err == nil {
		t.Error("Expected 'invalid character 'w' after object key', got ", "not errors")
	}
}

func TestFailEmptyFilePath(t *testing.T) {
	var config *TestConfig
	err := LoadFromFile("", &config)

	if err == nil {
		t.Error("Expected 'empty configuration file path', got ", "not errors")
	}
}