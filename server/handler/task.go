package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type TaskCreate struct {
	Title string `json:"title"`
}

// とりあえずメモリに保存
var tasks = []Task{}

// GET /tasks
func (h *Handlers) GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func (h *Handlers) CreateTask(c echo.Context) error {
	var taskCreate TaskCreate
	if err := c.Bind(&taskCreate); err != nil {
		return err
	}
	task := Task{
		ID:    uuid.New(),
		Title: taskCreate.Title,
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func (h *Handlers) UpdateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID.String() == c.Param("id") {
			tasks[i].Title = task.Title
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

// DELETE /tasks/:id
func (h *Handlers) DeleteTask(c echo.Context) error {
	for i, t := range tasks {
		if t.ID.String() == c.Param("id") {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}
