package controllers

import (
	"net/http"
	"fmt"
)

func CriarRota() {
	http.HandleFunc("GET /", respostaApi)
}

func respostaApi(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "ok"}`)

}