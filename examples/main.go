package main

import (
	"log"
	"net/http"

	"github.com/Alma-media/restful"
)

func main() {
	handler := static.NewHandler()

	handler.Init()

	http.Handle("/auth", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
