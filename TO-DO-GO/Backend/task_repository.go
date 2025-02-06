package main


type TaskRepository struct {
	tasks []Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{tasks: []Task{}}
}

	func (r *TaskRepository) AddTask(task Task) {
		r.tasks = append (r.tasks, task)
}

func (r *TaskRepository) GetAllTasks() []Task {
	return r.tasks
}