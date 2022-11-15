package tasks

type TaskCreation struct {
	Content string
}

type TaskMarkDone struct {
	ID uint
}

type TaskDeletion struct {
	TaskMarkDone
}

type Task struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type TaskResponse struct {
	Data Task `json:"data"`
}

type TaskListResponse struct {
	Data []Task `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
