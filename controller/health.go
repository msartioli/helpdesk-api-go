package controller

import (
	"fmt"
	"net/http"
)

func ligacaoComServidor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "aplication/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}
func Rota() {
	http.HandleFunc("GET /health", ligacaoComServidor)
}
