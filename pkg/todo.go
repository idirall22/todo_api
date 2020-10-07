package service

import (
	task "github.com/idirall22/products_api/api"
	"github.com/idirall22/products_api/pkg/store"
)

// Service struct
type Service struct {
	store store.Store
}

// NewService create new todo service.
func NewService(store store.Store) *Service {
	return &Service{store}
}

// CreateTask task
func (s *Service) CreateTask(formTask task.Form) (int64, error) {
	return s.store.Create(formTask)
}

// GetTask single task
func (s *Service) GetTask(id int64) task.Task {
	return s.store.Get(id)
}

// ListTask Tasks
func (s *Service) ListTask() []task.Task {
	return s.store.List()
}

// UpdateTask single task
func (s *Service) UpdateTask(task task.Task) error {
	return s.store.Update(task)
}

// DeleteTask single task
func (s *Service) DeleteTask(id int64) error {
	return s.store.Delete(id)
}

// ToggleDoneTask single task
func (s *Service) ToggleDoneTask(id int64) error {
	return s.store.ToggleDone(id)
}
