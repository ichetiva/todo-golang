package todo

import (
	"encoding/json"
	"net/http"

	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/pkg/postgres"
	"github.com/ichetiva/todo-golang/pkg/postgres/models"
)

type Controller struct {
	Config *config.Config
}

// @summary     Add todo
// @description Add todo to database with data which contains content
// @tags        todo
// @accept      json
// @router      /todo/create [post]
func (c *Controller) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todoCreation TodoCreation
	err := json.NewDecoder(r.Body).Decode(&todoCreation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad request",
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	todo := models.Todo{
		Content: todoCreation.Content,
		Done:    false,
	}
	db.Create(&todo)
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]interface{}{
		"data": TodoResponse{
			ID:      todo.ID,
			Content: todo.Content,
			Done:    todo.Done,
		},
	})
	_, _ = w.Write([]byte(data))
}

// @summary     Done todo
// @description Mark todo as done with data which contains id
// @tags        todo
// @accept      json
// @router      /todo/done [put]
func (c *Controller) DoneTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todoMarkDone TodoMarkDone
	err := json.NewDecoder(r.Body).Decode(&todoMarkDone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad request",
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	todo := models.Todo{}
	db.First(&todo, todoMarkDone.ID)
	todo.Done = true
	db.Save(&todo)

	data, _ := json.Marshal(map[string]interface{}{
		"id":      todo.ID,
		"content": todo.Content,
		"done":    todo.Done,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Delete todo
// @description Delete todo from database with data which contains id
// @tags        todo
// @accept      json
// @router      /todo/delete [delete]
func (c *Controller) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todoDeletion TodoDeletion
	err := json.NewDecoder(r.Body).Decode(&todoDeletion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad request",
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	db.Delete(&models.Todo{}, todoDeletion.ID)

	data, _ := json.Marshal(map[string]interface{}{
		"ok": true,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Get todo
// @description Get todo from database
// @tags        todo
// @accept      json
// @router      /todo [get]
func (c *Controller) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	id := params.Get("id")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	var todo models.Todo
	db.First(&todo, id)

	data, _ := json.Marshal(map[string]interface{}{
		"data": TodoResponse{
			ID:      todo.ID,
			Content: todo.Content,
			Done:    todo.Done,
		},
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Get all todos
// @description Get all todos from database
// @tags        todo
// @accept      json
// @router      /todo/all [get]
func (c *Controller) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	var todos []models.Todo
	db.Find(&todos)

	var todoResponse []TodoResponse
	for i := 0; i < len(todos); i++ {
		todoResponse = append(
			todoResponse,
			TodoResponse{
				ID:      todos[i].ID,
				Content: todos[i].Content,
				Done:    todos[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": todoResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

// @summary     Get done todos
// @description Get todos which marked as done
// @tags        todo
// @accept      json
// @router      /todo/done [get]
func (c *Controller) GetDoneTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	var todos []models.Todo
	db.Where("done = ?", true).Find(&todos)

	var todoResponse []TodoResponse
	for i := 0; i < len(todos); i++ {
		todoResponse = append(
			todoResponse,
			TodoResponse{
				ID:      todos[i].ID,
				Content: todos[i].Content,
				Done:    todos[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": todoResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

// @summary     Get not done todos
// @description Get todos which not marked as done
// @tags        todo
// @accept      json
// @router      /todo/notDone [get]
func (c *Controller) GetNotDoneTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"error":   "Bad Gateway",
			"message": "Connect to database error, try later",
		})
		_, _ = w.Write([]byte(data))
	}

	var todos []models.Todo
	db.Where("done = ?", false).Find(&todos)

	var todoResponse []TodoResponse
	for i := 0; i < len(todos); i++ {
		todoResponse = append(
			todoResponse,
			TodoResponse{
				ID:      todos[i].ID,
				Content: todos[i].Content,
				Done:    todos[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": todoResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
