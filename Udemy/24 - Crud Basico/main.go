package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) // Create

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
