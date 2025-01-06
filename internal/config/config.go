package config

import (
	"fmt"
	"embed"
	"encoding/json"
	"github.com/google/uuid"

	"github.com/xeipuuv/gojsonschema"
)

var (
	ErrConfigNotValid = fmt.Errorf("configuration not valid")
)

//go:embed config.schema.json
var jsonSchemaFile embed.FS

func LoadServiceConfiguration(filePath string) (*Configuration, error) {
	jsonSchema, err := jsonSchemaFile.ReadFile("config.schema.json")
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConfigNotValid, err)
	}

	jsonConfig, err := readFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConfigNotValid, err)
	}

	if err = validateJSONConfig(jsonSchema, jsonConfig); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConfigNotValid, err)
	}

	var config *Configuration
	if err := json.Unmarshal(jsonConfig, &config); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConfigNotValid, err)
	}

	if config.ID == "" {
		config.ID = "sentinel-" + uuid.New().String()
	}

	return config, nil
}

func validateJSONConfig(schema, jsonConfig []byte) error {
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	documentLoader := gojsonschema.NewBytesLoader(jsonConfig)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("error validating: %s", err.Error())
	}
	if !result.Valid() {
		return fmt.Errorf("json schema validation errors: %s", result.Errors())
	}
	return nil
}
