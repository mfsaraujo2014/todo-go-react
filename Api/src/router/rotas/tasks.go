package rotas

import (
	"net/http"

	"github.com/mfsaraujo2014/todo-go-react/src/controllers"
)

var rotasTasks = []Rota{
	{
		URI:                "/tasks",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTask,
		RequerAutenticacao: false,
	},
	{
		URI:                "/tasks",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTasks,
		RequerAutenticacao: true,
	},
	{
		URI:                "/tasks/{taskId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTask,
		RequerAutenticacao: true,
	},
	{
		URI:                "/tasks/{taskId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarTask,
		RequerAutenticacao: false,
	},
	{
		URI:                "/tasks/{taskId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarTask,
		RequerAutenticacao: false,
	},
}
