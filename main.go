package main

import (
	"GoAppWeb/routes"
	"net/http"
)

func main() {
	//db := conectaComBancoDeDados()
	//defer db.Close()
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
