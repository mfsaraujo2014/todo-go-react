package router

import (
	"github.com/gorilla/mux"
	"github.com/mfsaraujo2014/todo-go-react/src/router/rotas"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
