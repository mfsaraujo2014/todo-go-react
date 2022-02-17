package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mfsaraujo2014/todo-go-react/src/config"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
