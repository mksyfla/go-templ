package main

import (
	"go-jwp-emading/templates"
	"log"
	"net/http"
)

func main() {
	empty := templates.Empty()

	attributes := templates.Attribute()
	expressions := templates.Expression()
	statements := templates.Statements()
	compositions := templates.Composition(templates.ComponentParam("Dynamic contents"))
	button := templates.Button("Test", "test test")

	mux := http.NewServeMux()

	mux.HandleFunc("/attr", func(w http.ResponseWriter, r *http.Request) {
		attributes.Render(r.Context(), w)
	})

	mux.HandleFunc("/exp", func(w http.ResponseWriter, r *http.Request) {
		expressions.Render(r.Context(), w)
	})

	mux.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		statements.Render(r.Context(), w)
	})

	mux.HandleFunc("/comp", func(w http.ResponseWriter, r *http.Request) {
		compositions.Render(r.Context(), w)
	})

	mux.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		empty.Render(r.Context(), w)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		button.Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err.Error())
	}
}
