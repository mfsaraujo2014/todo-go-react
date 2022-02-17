package models

import (
	"errors"
)

type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}

type Task struct {
	ID        uint64 `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func (task *Task) Preparar() error {
	if erro := task.validar(); erro != nil {
		return erro
	}
	return nil
}

func (task *Task) validar() error {
	if task.Title == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}
	return nil
}
