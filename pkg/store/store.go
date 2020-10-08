package store

import (
	task "github.com/idirall22/todo_api/api"
)

// Store interface
type Store interface {
	Create(form task.Form) (int64, error)
	Get(id int64) (task.Task, error)
	List() []task.Task
	Update(id int64, form task.Form) error
	Delete(id int64) error
	ToggleDone(id int64) error
}
