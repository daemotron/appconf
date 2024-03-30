# Appconf for Go

The appconf module is a lightweight [Go](https://go.dev) configuration solution. It supports

* setting defaults
* reading from JSON files
* reading from environment variables
* reading from command line flags

## Installation

```shell
go get github.com/daemotron/appconf
```

## What about Viper

[Viper](https://github.com/spf13/viper) is a highly sophisticated configuration solution, offering
far more flexibility. However, Viper also comes with a large amount of dependencies, while appconf
only uses packages from Go's standard library.

## Usage

```go
package main

import "github.com/daemotron/appconf"

func main() {
    conf := appconf.NewConf("MyApp")
}
```