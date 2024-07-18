package main

import (
	"fmt"
	"net/http"
)

// VICTOR DA SILVA DANTAS - 22351268
// GUILHERME TELES DA COSTA MOURA - 22353219

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
		})
	http.HandleFunc(
		"/welcome",
		func(w http.ResponseWriter, r *http.Request) {
			nome := r.URL.Query().Get("nome")
			fmt.Fprintf(w, "Ola, %s, seja bem-vindo a nossa pagina sobre Android !!", nome)
		})
	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "hello.html")
		})
	http.ListenAndServe(":8080", nil)
}