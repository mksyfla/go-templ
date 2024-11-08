package main

import (
	"context"
	templateSyntax "go-jwp-emading/templates/01-syntax-usage"
	templateComponent "go-jwp-emading/templates/02-core-concepts"
	"log"
	"net/http"
)

func syntaxUsage(mux *http.ServeMux) {
	attributes := templateSyntax.Attribute()
	expressions := templateSyntax.Expression()
	statements := templateSyntax.Statements()
	compositions := templateSyntax.Composition()
	button := templateSyntax.CssTempl()
	javascript := templateSyntax.Javascript()
	contextTempl := templateSyntax.ContextTempl()

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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		button.Render(r.Context(), w)
	})

	mux.HandleFunc("/js", func(w http.ResponseWriter, r *http.Request) {
		javascript.Render(r.Context(), w)
	})

	mux.HandleFunc("/ctx", func(w http.ResponseWriter, r *http.Request) {
		type ctxKey string
		var key ctxKey = "name"

		ctx := context.WithValue(r.Context(), key, "Kasyfil Context")
		contextTempl.Render(ctx, w)
	})
}

func coreConcepts(mux *http.ServeMux) {
	templateComponent.ComponentConcept()
}

func main() {
	empty := templateSyntax.Empty()

	mux := http.NewServeMux()

	syntaxUsage(mux)
	coreConcepts(mux)

	mux.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		empty.Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err.Error())
	}
}
