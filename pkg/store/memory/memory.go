package memory

import (
	task "github.com/idirall22/products_api/api"
)

// Memory store struct.
type Memory struct {
	tasks []task.Task
}

// Create task
func (m *Memory) Create(form task.Form) (int64, error) {
	return 0, nil
}

// Get task
func (m *Memory) Get(id int64) task.Task {
	return task.Task{}
}

// List task
func (m *Memory) List() []task.Task {
	return []task.Task{}
}

// Update task
func (m *Memory) Update(task task.Task) error {
	return nil
}

// Delete task
func (m *Memory) Delete(id int64) error {
	return nil
}

// ToggleDone task
func (m *Memory) ToggleDone(id int64) error {
	return nil
}
