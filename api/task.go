package task

import "time"

//Task struct
type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// Form struct
type Form struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
