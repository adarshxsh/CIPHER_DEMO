package main

import (
	"encoding/json"
	"math/rand"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
)


var Tasks   []Task = []Task{{ID: 353, Title: "project selection", Description: "select the project for fun and learn", Status: "started"}, 
{ID: 817, Title: "meet lookup", Description: "Read the docs maintained till now ", Status: "not started"}, 
{ID: 361, Title: "cipher update", Description: "Create a server for basic crud operation", Status: "completed"}}


func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Serving Home Page")
	w.Write([]byte("<h1>Welcome to the Task Management API</h1>"))


}
func ServeTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tasks)

	for _, task := range Tasks {
		// w.Write([]byte(fmt.Sprintf("Task ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)))
		fmt.Printf("Task ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
		fmt.Fprintf(w, "Task ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
	
	
}


func ServeTask(w http.ResponseWriter, r *http.Request) {	
	vars := mux.Vars(r)
	id := vars["id"]

	for _, task := range Tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}		

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, task := range Tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}


func addTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	rand.Seed(time.Now().UnixNano())
	task.ID = rand.Intn(1000)
	Tasks = append(Tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)	
	w.Write([]byte(fmt.Sprintf("Task created with ID: %d", task.ID)))

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, task := range Tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			Tasks[i].Title = updatedTask.Title
			Tasks[i].Description = updatedTask.Description
			Tasks[i].Status = updatedTask.Status
			json.NewEncoder(w).Encode(Tasks[i])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}	


