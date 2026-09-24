package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"api-doc-example/internal/models"
)

type TodoHandler struct {
	todos  map[int]models.Todo
	nextID int
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// ListTodos godoc
// @Summary      Список задач
// @Description  Возвращает все задачи
// @Tags         todos
// @Produce      json
// @Success      200  {object}  models.SuccessResponse{data=[]models.Todo}
// @Router       /todos [get]
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	for _, todo := range h.todos {
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todos,
	})
}

// GetTodo godoc
// @Summary      Получить задачу по ID
// @Description  Возвращает одну задачу по её идентификатору
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "ID задачи"
// @Success      200  {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400  {string}  string  "Неверный ID"
// @Failure      404  {string}  string  "Задача не найдена"
// @Router       /todos/{id} [get]
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// CreateTodo godoc
// @Summary      Создать задачу
// @Description  Создаёт новую задачу
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.CreateTodoRequest  true  "Данные задачи"
// @Success      201   {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400   {string}  string  "Невалидный запрос"
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалидный запрос", http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		ID:          h.nextID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	h.todos[h.nextID] = todo
	h.nextID++

	log.Printf("Создана новая задача: %+v", todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// UpdateTodo godoc
// @Summary      Обновить задачу
// @Description  Частично обновляет задачу: любые из полей title, description, done
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      int                      true  "ID задачи"
// @Param        todo  body      models.UpdateTodoRequest true  "Обновляемые поля"
// @Success      200   {object}  models.SuccessResponse{data=models.Todo}
// @Failure      400   {string}  string  "Неверный ID или невалидный запрос"
// @Failure      404   {string}  string  "Задача не найдена"
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалидный запрос", http.StatusBadRequest)
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Done != nil {
		todo.Done = *req.Done
	}
	todo.UpdatedAt = time.Now()

	h.todos[id] = todo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// DeleteTodo godoc
// @Summary      Удалить задачу
// @Description  Удаляет задачу по её идентификатору
// @Tags         todos
// @Param        id   path  int  true  "ID задачи"
// @Success      204
// @Failure      400  {string}  string  "Неверный ID"
// @Failure      404  {string}  string  "Задача не найдена"
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.todos[id]; !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	delete(h.todos, id)

	w.WriteHeader(http.StatusNoContent)
}