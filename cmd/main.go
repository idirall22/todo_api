package main

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	task "github.com/idirall22/todo_api/api"
	service "github.com/idirall22/todo_api/pkg"
	"github.com/idirall22/todo_api/pkg/store/memory"
)

func main() {
	app := fiber.New()
	service := service.NewService(memory.NewMemoryStore())
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// POST create a task
	// curl -d '{"title": "first task", "descritption": "simple"}' -H "Content-Type: application/json" -X POST 127.0.0.1:3000/api/v1/todos
	v1.Post("/todos", func(c *fiber.Ctx) error {
		form := task.Form{}
		if err := c.BodyParser(&form); err != nil {
			return c.SendStatus(400)
		}
		id, _ := service.CreateTask(form)

		return c.JSON(
			map[string]interface{}{
				"id": id,
			},
		)
	})

	// GET single task
	// curl -H "Content-Type: application/json" 127.0.0.1:3000/api/v1/todos/1
	v1.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.SendStatus(400)
		}
		task, err := service.GetTask(id)
		if err != nil {
			c = c.Status(404)
			return c.SendString(err.Error())
		}

		return c.JSON(task)
	})

	// List all tasks
	// curl -H "Content-Type: application/json" 127.0.0.1:3000/api/v1/todos
	v1.Get("/todos", func(c *fiber.Ctx) error {
		tasks := service.ListTask()
		return c.JSON(tasks)
	})

	// Update a task
	/*
		curl \
		-d '{"title": "first task updated", "descritption": "simple updated"}' \
		-H "Content-Type: application/json" \
		-X PUT \
		127.0.0.1:3000/api/v1/todos
	*/
	v1.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.SendStatus(400)
		}
		form := task.Form{}
		if err := json.Unmarshal(c.Body(), &form); err != nil {
			return c.SendStatus(400)
		}
		tasks := service.UpdateTask(id, form)
		return c.JSON(tasks)
	})

	// Delete a task
	/*
		curl \
		-H "Content-Type: application/json" \
		-X DELETE \
		127.0.0.1:3000/api/v1/todos/1
	*/
	v1.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.SendStatus(400)
		}
		service.DeleteTask(id)
		return c.SendStatus(204)
	})

	// Toggle done task
	/*
		curl \
		-H "Content-Type: application/json" \
		-X POST \
		127.0.0.1:3000/api/v1/todos/toggle/1
	*/
	v1.Post("/todos/toggle/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.SendStatus(400)
		}
		service.ToggleDoneTask(id)
		return c.SendStatus(204)
	})
	app.Listen(":3000")
}
