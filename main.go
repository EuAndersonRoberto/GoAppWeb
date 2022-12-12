package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descrição string
	Preço           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Tenis", Descrição: "Da moda", Preço: 350.00, Quantidade: 2},
		{"Chuteira", "Do Neymar", 400.0, 4},
		{"SSD NVME2", "Muito rápido", 322., 1},
		{"Violão", "Tonante", 200, 5},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
