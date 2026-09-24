package models

import "time"

// User represents a user account.
// @Description Пользователь системы.
type User struct {
	ID        int       `json:"id" example:"1"`
	Email     string    `json:"email" example:"user@example.com"`
	Name      string    `json:"name" example:"Иван Иванов"`
	CreatedAt time.Time `json:"created_at" example:"2025-01-01T12:00:00Z"`
}

// CreateUserRequest represents the request body for creating a user.
// @Description Тело запроса для создания пользователя.
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
	Name  string `json:"name" validate:"required,max=100" example:"Иван Иванов"`
}

// UpdateUserRequest represents the request body for updating a user.
// @Description Тело запроса для частичного обновления пользователя. Все поля опциональны.
type UpdateUserRequest struct {
	Email *string `json:"email" validate:"omitempty,email" example:"new@example.com"`
	Name  *string `json:"name" validate:"omitempty,max=100" example:"Пётр Петров"`
}

// ErrorResponse represents the standard error format.
// @Description Стандартный ответ об ошибке.
type ErrorResponse struct {
	Error   string `json:"error" example:"invalid_id"`
	Message string `json:"message" example:"Неверный ID"`
}

// SuccessResponse is a generic successful response wrapper.
// @Description Обёртка успешного ответа.
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
}