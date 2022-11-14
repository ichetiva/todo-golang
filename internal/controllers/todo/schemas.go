package todo

type TodoCreation struct {
	Content string
}

type TodoMarkDone struct {
	ID uint
}

type TodoDeletion struct {
	TodoMarkDone
}

type TodoResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
