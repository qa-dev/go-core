package config

import (
	"io/ioutil"
	"testing"
	"os"
	"log"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

func TestSuccess(t *testing.T) {
	var config *TestConfig

	tempFile := makeJsonFile(`
		{
			"property1": "value1",
			"property2": "value2"
		}
	`);
	defer os.Remove(tempFile.Name()) // clean up

	err := LoadFromFile(tempFile.Name(), &config)

	assert.Nil(t, err)
	assert.Equal(t, "value1", config.Property1)
	assert.Equal(t, "value2", config.Property2)
}

func TestFailInvalidJson(t *testing.T) {
	var config *TestConfig

	tempFile := makeJsonFile(`
		{
			"property1"w: "value1",
			"property2": "value2"
		}
	`);
	defer os.Remove(tempFile.Name()) // clean up

	err := LoadFromFile(tempFile.Name(), &config)

	assert.Equal(t, "invalid character 'w' after object key", err.Error())

}

func TestFailEmptyFilePath(t *testing.T) {
	var config *TestConfig
	err := LoadFromFile("", &config)

	assert.Equal(t, "empty configuration file path", err.Error())
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