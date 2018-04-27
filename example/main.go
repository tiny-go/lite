package main

import (
	"log"
	"net/http"

	"github.com/tiny-go/codec/driver"
	"github.com/tiny-go/config"
	"github.com/tiny-go/lite"
	local "github.com/tiny-go/lite/example/config"
	// register auth module
	_ "github.com/tiny-go/lite/example/auth"
	// register codecs
	_ "github.com/tiny-go/codec/driver/json"
	_ "github.com/tiny-go/codec/driver/xml"
)

func main() {
	// set JSON codec as a default
	driver.Default("application/json")
	// load config
	conf := new(local.Config)
	if err := config.Init(conf, "demo"); err != nil {
		log.Fatal(err)
	}
	// create new handler
	handler := lite.NewHandler()
	// map config to the handler to make it available for all of the controllers
	handler.Map(conf)
	// "fake" some users (instead of passing database instance) and map as a dependency
	handler.Map(map[string]string{
		"admin@test.com": "admin",
		"user@test.com":  "user",
	})
	// register modules
	lite.Modules(func(alias string, module lite.Module) bool {
		handler.Use(alias, module)
		return true
	})
	// start HTTP server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
