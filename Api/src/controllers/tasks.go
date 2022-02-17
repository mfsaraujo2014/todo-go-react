package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	repository "github.com/mfsaraujo2014/todo-go-react/src/Repository"
	"github.com/mfsaraujo2014/todo-go-react/src/answers"
	"github.com/mfsaraujo2014/todo-go-react/src/db"
	"github.com/mfsaraujo2014/todo-go-react/src/models"
)

func CriarTask(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var task models.Task
	if erro = json.Unmarshal(corpoRequest, &task); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = task.Preparar(); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTasks(db)
	task.ID, erro = repositorio.Criar(task)
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusCreated, task)
}

func BuscarTasks(w http.ResponseWriter, r *http.Request) {
	title := strings.ToLower(r.URL.Query().Get("task"))
	db, erro := db.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTasks(db)
	tasks, erro := repositorio.Buscar(title)
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusOK, tasks)
}

func BuscarTask(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	taskID, erro := strconv.ParseUint(parametros["taskId"], 10, 64)
	if erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTasks(db)
	task, erro := repositorio.BuscarPorID(taskID)
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusOK, task)
}

func AtualizarTask(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	taskID, erro := strconv.ParseUint(parametros["taskId"], 10, 64)
	if erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var task models.Task
	if erro = json.Unmarshal(corpoRequest, &task); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = task.Preparar(); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTasks(db)
	if erro = repositorio.Atualizar(taskID, task); erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

func DeletarTask(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	taskID, erro := strconv.ParseUint(parametros["taskId"], 10, 64)
	if erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTasks(db)
	if erro = repositorio.Deletar(taskID); erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
