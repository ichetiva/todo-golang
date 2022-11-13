package todo

import (
	"encoding/json"
	"net/http"
)

type Controller struct{}

func (c *Controller) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(map[string]interface{}{
		"message": "Cool! You're in good way",
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(data))
}
