package main

import (
	"go-jwp-emading/templates"
	"log"
	"net/http"
)

func main() {
	attributes := templates.Attribute()
	empty := templates.Empty()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		attributes.Render(r.Context(), w)
	})

	mux.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		empty.Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err.Error())
	}
}
