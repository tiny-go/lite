package main

import (
	"log"
	"net/http"

	"github.com/Alma-media/restful"
	"github.com/tiny-go/codec/driver"
	// register auth module
	_ "github.com/Alma-media/restful/examples/auth"
	// register codecs
	_ "github.com/tiny-go/codec/driver/json"
	_ "github.com/tiny-go/codec/driver/xml"
)

func main() {
	// set JSON codec as a default
	driver.Default("application/json")

	handler := static.NewHandler()
	static.Modules(func(alias string, module static.Module) bool {
		handler.Use(alias, module)
		return true
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
