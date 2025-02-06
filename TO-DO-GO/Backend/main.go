package main 

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	repository = NewTaskRepository()
	mutex = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/tasks", handleTask)
	http.ListenAndServe(":8080", nil)
}

func handleTask ( w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {

		return 

	}

	switch r.Method {
	case http.MethodGet:
		tasks := repository.GetAllTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
		case http.MethodPost:
			var task Task
			if err:= json.NewDecoder(r.Body).Decode (&task); err !=nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mutex.Lock()
			repository.AddTask(task)
			mutex.Unlock()
			w.WriteHeader(http.StatusCreated)
	
	}			
}
