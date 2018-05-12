# lite

[![GoDoc][godoc-badge]][godoc-link]
[![License][license-badge]][license-link]
[![Build Status][circleci-badge]][circleci-link]
[![Report Card][report-badge]][report-link]
[![GoCover][cover-badge]][cover-link]

Simple tool for building RESTful APIs

### Usage
```go
package main

import (
	"log"
	"net/http"

	"github.com/tiny-go/lite"
	// register codecs
	"github.com/tiny-go/codec/driver"
	_ "github.com/tiny-go/codec/driver/json"
	_ "github.com/tiny-go/codec/driver/xml"
)

func main() {
	// set default codec
	driver.Default("application/json")
	// create new handler
	handler := lite.NewHandler()
	// map dependencies
	handler.Map(depA)
	handler.Map(depB)
	handler.Map(depC)
	// register modules
	handler.Use(aliasOne, moduleA)
	handler.Use(aliasTwo, moduleB)
	// start HTTP server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
```


[godoc-badge]: https://godoc.org/github.com/tiny-go/lite?status.svg
[godoc-link]: https://godoc.org/github.com/tiny-go/lite
[license-badge]: https://img.shields.io/:license-MIT-green.svg
[license-link]: https://opensource.org/licenses/MIT
[circleci-badge]: https://circleci.com/gh/tiny-go/lite.svg?style=shield
[circleci-link]: https://circleci.com/gh/tiny-go/lite
[report-badge]: https://goreportcard.com/badge/github.com/tiny-go/lite
[report-link]: https://goreportcard.com/report/github.com/tiny-go/lite
[cover-badge]: https://gocover.io/_badge/github.com/tiny-go/lite
[cover-link]: https://gocover.io/github.com/tiny-go/lite
