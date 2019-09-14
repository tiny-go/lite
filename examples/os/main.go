package main

import (
	"log"
	"net/http"

	"github.com/tiny-go/codec/driver"
	"github.com/tiny-go/config"
	"github.com/tiny-go/lite"
	local "github.com/tiny-go/lite/examples/os/config"
	"github.com/tiny-go/lite/examples/os/exec"

	// register command module
	_ "github.com/tiny-go/lite/examples/os/exec"
	// register codecs
	_ "github.com/tiny-go/codec/driver/json"
	_ "github.com/tiny-go/codec/driver/text"
	_ "github.com/tiny-go/codec/driver/xml"
)

func main() {
	// set text/plain codec as a default
	driver.Default("text/plain")
	// load config
	conf := new(local.Config)
	if err := config.Init(conf, "demo"); err != nil {
		log.Fatal(err)
	}
	// create new handler
	handler := lite.NewHandler()
	// map config to the handler to make it available for all of the controllers
	handler.Map(conf)

	handler.Map(exec.Commands{
		"ls",
		"pwd",
		"poweroff",
		"sensors",
		"whoami",
	})
	// register modules
	lite.Modules(func(alias string, module lite.Module) bool {
		handler.Use(alias, module)
		return true
	})
	// start HTTP server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
