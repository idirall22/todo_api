package memory

import (
	"fmt"
	"log"
	"time"

	task "github.com/idirall22/todo_api/api"
)

// Memory store struct.
type Memory struct {
	tasks []task.Task
}

// NewMemoryStore create memory store.
func NewMemoryStore() *Memory {
	return &Memory{
		tasks: []task.Task{},
	}
}

// Create task
func (m *Memory) Create(form task.Form) (int64, error) {
	var id int64 = int64(len(m.tasks)) + 1
	task := task.Task{
		ID:          id,
		Title:       form.Title,
		Description: form.Description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	m.tasks = append(m.tasks, task)
	log.Println(task)
	return id, nil
}

// Get task
func (m *Memory) Get(id int64) (task.Task, error) {
	for _, task := range m.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return task.Task{}, fmt.Errorf("There is no task with id: %d", id)
}

// List task
func (m *Memory) List() []task.Task {
	return m.tasks
}

// Update task
func (m *Memory) Update(id int64, t task.Form) error {
	for index, ts := range m.tasks {
		if ts.ID == id {
			m.tasks[index].Title = t.Title
			m.tasks[index].Description = t.Description
			break
		}
	}
	return nil
}

// Delete task
func (m *Memory) Delete(id int64) error {
	for index, ts := range m.tasks {
		if ts.ID == id {
			m.tasks = append(m.tasks[:index], m.tasks[index+1:]...)
			break
		}
	}

	return nil
}

// ToggleDone task
func (m *Memory) ToggleDone(id int64) error {
	for index, ts := range m.tasks {
		if ts.ID == id {
			m.tasks[index].Done = !m.tasks[index].Done
		}
	}
	return nil
}
