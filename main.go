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

type Items struct {
	Nome string
	Id   int
}

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Titulo": "Tião",
			"Items": []Items{
				{Nome: "Teste-deploy-backend", Id: 1},
				{Nome: "Teste-deploy-frontend", Id: 2},
			},
		}

		t.ExecuteTemplate(w, "index.html", data)
	})

	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
