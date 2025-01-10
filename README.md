# Traefik Config Validator

![OSS Lifecycle](https://img.shields.io/osslifecycle?file_url=https%3A%2F%2Fraw.githubusercontent.com%2Fotto-de%2Ftraefik-config-validator%2Frefs%2Fheads%2Fmain%2FOSSMETADATA)

<p align="center">
<img src="https://github.com/traefik/traefik/raw/master/docs/content/assets/img/traefik.logo.png" alt="Traefik" title="Traefik" />
</p>

[![CI](https://github.com/otto-de/traefik-config-validator/actions/workflows/My_CI.yml/badge.svg)](https://github.com/otto-de/traefik-config-validator/actions/workflows/My_CI.yml)
[![Release](https://github.com/otto-de/traefik-config-validator/actions/workflows/My_Release.yml/badge.svg)](https://github.com/otto-de/traefik-config-validator/actions/workflows/My_Release.yml)

`traefik-config-validator` is a CLI tool to (syntactically) validate your [Traefik](https://doc.traefik.io/traefik) configuration files to ensure bad configurations are being captured before hitting your production system or during development.
It can be used either as a developer tool on your machine or in CI/CD pipelines and has been used in production at OTTO.

It will be validated against the following JSON schemas hosted on schemastore:

- https://github.com/SchemaStore/schemastore/blob/master/src/schemas/json/traefik-v2.json
- https://github.com/SchemaStore/schemastore/blob/master/src/schemas/json/traefik-v2-file-provider.json

## Features

- Validate Traefik configuration (`traefik.yml`)
- Validate Traefik [file provider configuration](https://doc.traefik.io/traefik/providers/file/)

## Known Limitations

- Currently, only YAML configuration is supported while Traefik supports TOML, JSON, environment variables and the CLI for configuration
- Cannot recursively scan files for file provided configurations

## Installation

```
go get github.com/otto-de/traefik-config-validator
```

or via Docker

```
docker pull ghcr.io/otto-de/traefik-config-validator:latest
```

## Usage

```
traefik-config-validator -cfg <traefik.yml> -cfgdir <static-file-provider-root>
```

## Developer Guide

All steps in development can be performed either through the CI/CD pipeline with GitHub Actions or locally. See below for instructions on how to do it.

### Setting Up

We use [pre-commit](https://pre-commit.com/) for validating commits before pushing them. In addition, you will need `go` installed and (optionally) Docker for build an image.

### Building

Build via `make build` to build via golang build tools or `docker build --build-arg VERSION=<myversion>` to build via Docker.

### Testing

Currently, we are lacking good unit tests. Feel free to add them!

### Linting

Run `make lint` to lint the Golang code.

### Releasing

We use [Semantic Versioning](https://semver.org/). Each git tag should be named according to `v{major}.{minor}.{patch}`.

## Contributing

If you'd like to contribute to the project, refer to the [contributing documentation](CONTRIBUTING.md).

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md).
By participating in this project, you agree to abide by its terms.

## Credits

The gopher's logo of Traefik is licensed under the Creative Commons 3.0 Attributions license (see [traefik/traefik](https://github.com/traefik/traefik)).
