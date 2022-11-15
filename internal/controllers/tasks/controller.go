package tasks

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

// @summary     Add task
// @description Add task to database with data which contains content
// @tags        tasks
// @accept      json
// @produce		json
// @router      /task/create [post]
// @param 		content body string true "Task content"
// @success 	200 {object} TaskResponse
// @Failure 	400 {object} ErrorResponse "Non-valid data"
// @Failure 	502 {object} ErrorResponse "Add task failed"
func (c *Controller) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var taskCreation TaskCreation
	err := json.NewDecoder(r.Body).Decode(&taskCreation)
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

	task := models.Task{
		Content: taskCreation.Content,
		Done:    false,
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
// @Failure 	400 {object} ErrorResponse "Non-valid data"
// @Failure 	502 {object} ErrorResponse "Mark task as done failed"
func (c *Controller) DoneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var taskMarkDone TaskMarkDone
	err := json.NewDecoder(r.Body).Decode(&taskMarkDone)
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

	task := models.Task{}
	db.First(&task, taskMarkDone.ID)
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
// @Failure 	400 {object} ErrorResponse "Non-valid data"
// @Failure 	502 {object} ErrorResponse "Delete task failed"
func (c *Controller) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var taskDeletion TaskDeletion
	err := json.NewDecoder(r.Body).Decode(&taskDeletion)
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

	db.Delete(&models.Task{}, taskDeletion.ID)

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
// @Failure 	502 {object} ErrorResponse "Get task failed"
func (c *Controller) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	var task models.Task
	db.First(&task, id)

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
// @success 	200 {object} TaskResponse
// @Failure 	502 {object} ErrorResponse "Get all tasks failed"
func (c *Controller) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get all tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Find(&tasks)

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
// @success 	200 {object} TaskResponse
// @Failure 	502 {object} ErrorResponse "Get done tasks failed"
func (c *Controller) GetDoneTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get done tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Where("done = ?", true).Find(&tasks)

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
// @success 	200 {object} TaskResponse
// @Failure 	502 {object} ErrorResponse "Get not done tasks failed"
func (c *Controller) GetNotDoneTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Get not done tasks failed",
		})
		_, _ = w.Write([]byte(data))
	}

	var tasks []models.Task
	db.Where("done = ?", false).Find(&tasks)

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
