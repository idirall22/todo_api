package store

import (
	task "github.com/idirall22/todo_api/api"
)

// Store interface
type Store interface {
	Create(form task.Form) (int64, error)
	Get(id int64) task.Task
	List() []task.Task
	Update(task task.Task) error
	Delete(id int64) error
	ToggleDone(id int64) error
}
