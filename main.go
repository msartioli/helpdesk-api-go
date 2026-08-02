package main

import(
	"net/http"
	"github.com/msartioli/helpdesk-api-go/controllers"
)

func main() {
	controllers.CriarRota()
	http.ListenAndServe(":8080", nil)
}
