// Copyright (c) 2022 OTTO GmbH & Co KG
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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/otto-de/traefik-config-validator/pkg/validation"
)

var (
	configFile                  string
	configFileSchema            string
	fileProviderConfigDirectory string
	fileProviderConfigSchema    string

	versionFlag bool
	version     = "custom-build"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.LUTC | log.Llongfile)
}

func main() {
	flag.StringVar(&configFile, "cfg", "", "traefik config file")
	flag.StringVar(&configFileSchema, "cfgschema", validation.TraefikJsonSchemaRef, "traefik config file schema")
	flag.StringVar(&fileProviderConfigDirectory, "cfgdir", "", "traefik file provider directory")
	flag.StringVar(&fileProviderConfigSchema, "cfgdirschema", validation.TraefikFileProviderJsonSchemaRef, "traefik file provider config schema")
	flag.BoolVar(&versionFlag, "version", false, "show version and exit")

	flag.Parse()

	if versionFlag {
		fmt.Printf("traefik-config-validator (version %s)\n", version)
		os.Exit(1)
	}

	if configFile != "" {
		log.Printf("validating config file %s\n", configFile)
		if err := validation.ValidateJsonSchema(configFile, configFileSchema); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("skipping validation of config")
	}

	if fileProviderConfigDirectory != "" {
		log.Printf("validating static file provider configs from directory %s\n", fileProviderConfigDirectory)
		if err := validation.ValidateFileProviderConfig(fileProviderConfigDirectory, fileProviderConfigSchema); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("skipping validation of static file provider configs")
	}
}
