package tasks

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/use_cases/misc"
	"github.com/ichetiva/todo-golang/pkg/postgres"
	"github.com/ichetiva/todo-golang/pkg/postgres/models"
	"gorm.io/gorm"
)

type Controller struct {
	Config *config.Config
}

// @summary     Add task
// @description Add task to database with data which contains content
// @tags        tasks
// @accept      json
// @produce		json
// @router      /task/create [post]
// @param 		content body string true "Task content"
// @success 	200 {object} TaskResponse
// @failure 	400 {object} ErrorResponse "Non-valid data"
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	502 {object} ErrorResponse "Add task failed"
func (c *Controller) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var requestData TaskCreation
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Add task failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	task := models.Task{
		Content:   requestData.Content,
		Done:      false,
		UserRefer: user.ID,
	}
	db.Create(&task)
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]interface{}{
		"data": Task{
			ID:      task.ID,
			Content: task.Content,
			Done:    task.Done,
		},
	})
	_, _ = w.Write([]byte(data))
}

// @summary     Done task
// @description Mark task as done with data which contains id
// @tags        tasks
// @accept      json
// @router      /task/done [put]
// @param 		id body integer true "Task indentifer"
// @success 	200 {object} TaskResponse
// @failure 	400 {object} ErrorResponse "Non-valid data"
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	404 {object} ErrorResponse "Task not found"
// @failure 	502 {object} ErrorResponse "Mark task as done failed"
func (c *Controller) DoneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var requestData TaskMarkDone
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Mark task as done failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var task models.Task
	result = db.Where("id = ? and user_refer = ?", requestData.ID, user.ID).First(&task)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Task not found",
		})
		_, _ = w.Write([]byte(data))
	}
	task.Done = true
	db.Save(&task)

	data, _ := json.Marshal(map[string]interface{}{
		"data": Task{
			ID:      task.ID,
			Content: task.Content,
			Done:    task.Done,
		},
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Delete task
// @description Delete task from database with data which contains id
// @tags        tasks
// @accept      json
// @router      /task/delete [delete]
// @param 		id body integer true "Task identifer"
// @success 	200 {object} TaskResponse
// @failure 	400 {object} ErrorResponse "Non-valid data"
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	502 {object} ErrorResponse "Delete task failed"
func (c *Controller) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var requestData TaskDeletion
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Delete task failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	db.Where("id = ? and user_refer = ?", requestData.ID, user.ID).Delete(&models.Task{})

	data, _ := json.Marshal(map[string]interface{}{
		"ok": true,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Get task
// @description Get task from database
// @tags        tasks
// @accept      json
// @router      /task [get]
// @param 		id query integer true "Task identifer"
// @success 	200 {object} TaskResponse
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	404 {object} ErrorResponse "Task not found"
// @failure 	502 {object} ErrorResponse "Get task failed"
func (c *Controller) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	params := r.URL.Query()
	id := params.Get("id")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get task failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var task models.Task
	result = db.Where("id = ? and user_refer = ?", id, user.ID).First(&task)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Task not found",
		})
		_, _ = w.Write([]byte(data))
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": Task{
			ID:      task.ID,
			Content: task.Content,
			Done:    task.Done,
		},
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}

// @summary     Get all tasks
// @description Get all tasks from database
// @tags        tasks
// @accept      json
// @router      /task/all [get]
// @success 	200 {object} TaskListResponse
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	502 {object} ErrorResponse "Get all tasks failed"
func (c *Controller) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get all tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Where("user_refer = ?", user.ID).Find(&tasks)

	var tasksResponse []Task
	for i := 0; i < len(tasks); i++ {
		tasksResponse = append(
			tasksResponse,
			Task{
				ID:      tasks[i].ID,
				Content: tasks[i].Content,
				Done:    tasks[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": tasksResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

// @summary     Get done tasks
// @description Get tasks which marked as done
// @tags        tasks
// @accept      json
// @router      /task/done [get]
// @success 	200 {object} TaskListResponse
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	502 {object} ErrorResponse "Get done tasks failed"
func (c *Controller) GetDoneTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get done tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Where("done = ? and user_refer = ?", true, user.ID).Find(&tasks)

	var tasksResponse []Task
	for i := 0; i < len(tasks); i++ {
		tasksResponse = append(
			tasksResponse,
			Task{
				ID:      tasks[i].ID,
				Content: tasks[i].Content,
				Done:    tasks[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": tasksResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

// @summary     Get not done tasks
// @description Get tasks which not marked as done
// @tags        tasks
// @accept      json
// @router      /task/notDone [get]
// @success 	200 {object} TaskListResponse
// @failure 	403 {object} ErrorResponse "Non-valid session token"
// @failure 	502 {object} ErrorResponse "Get not done tasks failed"
func (c *Controller) GetNotDoneTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, err := misc.GetTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get not done tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var user User
	result := db.Table("users").Select("users.id").Joins(
		"left join sessions on sessions.user_refer = users.id",
	).Where("sessions.token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid session token",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Where("done = ? and user_refer = ?", false, user.ID).Find(&tasks)

	var tasksResponse []Task
	for i := 0; i < len(tasks); i++ {
		tasksResponse = append(
			tasksResponse,
			Task{
				ID:      tasks[i].ID,
				Content: tasks[i].Content,
				Done:    tasks[i].Done,
			},
		)
	}

	data, _ := json.Marshal(map[string]interface{}{
		"data": tasksResponse,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
