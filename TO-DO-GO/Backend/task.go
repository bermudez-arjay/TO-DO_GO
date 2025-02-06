package main

type Task struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

func (t *Task) Complete() {
	t.Completed = true
}