package models

import "time"

// Todo represents a task item.
// @Description Задача пользователя.
type Todo struct {
	ID          int       `json:"id" example:"1"`
	UserID      int       `json:"user_id" example:"1"`
	Title       string    `json:"title" example:"Купить молоко"`
	Description string    `json:"description" example:"Не забыть про скидку"`
	Done        bool      `json:"done" example:"false"`
	CreatedAt   time.Time `json:"created_at" example:"2025-01-01T12:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-01-01T12:00:00Z"`
}

// CreateTodoRequest represents the request body for creating a todo.
// @Description Тело запроса для создания задачи.
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required" example:"1"`
	Title       string `json:"title" validate:"required,max=200" example:"Купить молоко"`
	Description string `json:"description" validate:"max=1000" example:"Не забыть про скидку"`
	Done        bool   `json:"done" example:"false"`
}

// UpdateTodoRequest represents the request body for updating a todo.
// @Description Тело запроса для частичного обновления задачи. Все поля опциональны.
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200" example:"Купить хлеб"`
	Description *string `json:"description" validate:"omitempty,max=1000" example:"Обновлённое описание"`
	Done        *bool   `json:"done" example:"true"`
}