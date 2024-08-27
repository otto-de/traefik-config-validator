// Copyright (c) 2024 OTTO GmbH & Co KG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/xeipuuv/gojsonschema"

	"github.com/otto-de/traefik-config-validator/pkg/utils"
)

const (
	TraefikJsonSchemaRef             = "https://raw.githubusercontent.com/SchemaStore/schemastore/master/src/schemas/json/traefik-v2.json"
	TraefikFileProviderJsonSchemaRef = "https://raw.githubusercontent.com/SchemaStore/schemastore/master/src/schemas/json/traefik-v2-file-provider.json"
)

// Validates the file provider configuration by reading a directory and running
// the validation via JSON Schema.
func ValidateFileProviderConfig(fileProviderConfigDirectory string, schemaRef string) (err error) {
	items, err := os.ReadDir(fileProviderConfigDirectory)
	if err != nil {
		return
	}
	for _, item := range items {
		if !item.IsDir() {
			if err = ValidateJsonSchema(path.Join(fileProviderConfigDirectory, item.Name()), schemaRef); err != nil {
				return
			}
		}
	}
	return
}

// Validates the configuration by reading a file and running
// the validation via JSON Schema.
func ValidateJsonSchema(configFile string, schemaRef string) (err error) {
	configFileContent, err := utils.ReadYamlFileAsJson(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %s", configFile, err)
	}
	schemaLoader := gojsonschema.NewReferenceLoader(schemaRef)
	documentLoader := gojsonschema.NewStringLoader(string(configFileContent))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return
	}
	if result.Valid() {
		log.Printf("The document [%s] is valid\n", configFile)
	} else {
		log.Printf("The document [%s] is not valid. see errors :\n", configFile)
		for _, desc := range result.Errors() {
			log.Printf("[X] %s\n", desc)
		}
		return errors.New("validation errors")
	}
	return
}
