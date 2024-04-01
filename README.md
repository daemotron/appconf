# Appconf for Go

![Build and Test State](https://github.com/daemotron/appconf/actions/workflows/go.yml/badge.svg?event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/daemotron/appconf.svg)](https://pkg.go.dev/github.com/daemotron/appconf)
![MIT License](https://img.shields.io/badge/license-MIT-green.svg)

The appconf module is a lightweight [Go](https://go.dev) configuration solution. It supports

* setting defaults
* reading from JSON files
* reading from environment variables
* reading from command line flags

## Installation

```shell
go get github.com/daemotron/appconf
```

## Usage

```go
package main

import (
    "github.com/daemotron/appconf"
    "log"
)

func main() {
    // initialize configuration context
    conf := appconf.NewConf("MyApp")
    
    // register configuration option
    err := conf.NewOption("foo", appconf.WithDefaultValue(appconf.StringValue("bar")))
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
}
```

## Conventions

The appconf module relies on several conventions in order to keep its interface
lean and simple.

### Configuration Files

The appconf module uses the provided *App Name* (as well as *App Author* and 
*App Version* if provided) to make assumptions about possible configuration
file paths and names.

It expects to find configuration files in one of the directories provided by

```go
conf.ConfigDirs(true)
```

and have a name of `config.json`, `conf.json` or `strings.ToLower(conf.Name) + ".json"`

Actually existing configuration files can be listed this way:

```go
confFiles, err := conf.ConfigFiles()
```

## What about Viper

[Viper](https://github.com/spf13/viper) is a highly sophisticated configuration solution, offering
far more flexibility. However, Viper also comes with a large amount of dependencies, while appconf
only uses packages from Go's standard library.