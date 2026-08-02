package controllers

import (
	"fmt"
	"net/http"
)

func Health() {
	http.HandleFunc("GET /health", respostaApi)
}

func respostaApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "ok"}`)

}
