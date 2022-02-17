package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/mfsaraujo2014/todo-go-react/src/config"
	"github.com/mfsaraujo2014/todo-go-react/src/router"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	config.Carregar()
	r := router.Gerar()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
}
