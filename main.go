package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"
)

//go:embed template/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "template/*"))

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Hello": "Hello World!!",
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
