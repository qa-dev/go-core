package config

import (
	"io/ioutil"
	"testing"
	"os"
	"log"
)

type TestConfig struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

func TestSuccess(t *testing.T) {
	var config *TestConfig

	tempFile := makeJsonFile("" +
		"{" +
			"\"property1\": \"value1\", " +
			"\"property2\": \"value2\"" +
		"}" +
	"");
	defer os.Remove(tempFile.Name()) // clean up

	err := LoadFromFile(tempFile.Name(), &config)

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

	tempFile := makeJsonFile("" +
		"{" +
			"\"property1\"w: \"value1\", " +
			"\"property2\": \"value2\"" +
		"}" +
	"");
	defer os.Remove(tempFile.Name()) // clean up

	err := LoadFromFile(tempFile.Name(), &config)

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

func makeJsonFile(jsonString string) (*os.File) {
	content := []byte(jsonString)
	tempFile, err := ioutil.TempFile("", "json")

	if err != nil {
		log.Fatal(err)
	}
	if _, err := tempFile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
	}

	return tempFile
}