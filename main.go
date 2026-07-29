package main

import (
	"net/http"

	"github.com/msartioli/helpdesk-api-go/controller"
)

func main() {
	conectarBanco()
	controller.Rota()
	http.ListenAndServe(":8080", nil)
}
