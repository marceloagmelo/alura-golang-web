package main

import (
	"net/http"

	"github.com/marceloagmelo/alura-golang-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
