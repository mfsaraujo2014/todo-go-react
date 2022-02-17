package repository

import (
	"database/sql"
	"fmt"

	"github.com/mfsaraujo2014/todo-go-react/src/models"
)

type Tasks struct {
	db *sql.DB
}

func NovoRepositorioDeTasks(db *sql.DB) *Tasks {
	return &Tasks{db}
}

func (repositorio Tasks) Criar(task models.Task) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into tasks(title, completed) values(?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(task.Title, task.Completed)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Tasks) Buscar(title string) ([]models.Task, error) {
	title = fmt.Sprintf("%%%s%%", title) // %title%

	linhas, erro := repositorio.db.Query(
		"select id, title, completed from tasks where title LIKE ?",
		title,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tasks []models.Task

	for linhas.Next() {
		var task models.Task

		if erro = linhas.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
		); erro != nil {
			return nil, erro
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repositorio Tasks) BuscarPorID(ID uint64) (models.Task, error) {
	linhas, erro := repositorio.db.Query(
		"select id, title, completed from tasks where id = ?",
		ID,
	)
	if erro != nil {
		return models.Task{}, erro
	}
	defer linhas.Close()

	var task models.Task

	if linhas.Next() {
		if erro = linhas.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
		); erro != nil {
			return models.Task{}, erro
		}
	}

	return task, nil
}

func (repositorio Tasks) Atualizar(ID uint64, task models.Task) error {
	statement, erro := repositorio.db.Prepare(
		"update tasks set title = ?, completed = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(task.Title, task.Completed, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Tasks) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from tasks where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
