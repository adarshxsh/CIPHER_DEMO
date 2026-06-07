package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)


var (
	
	Tasks []Task 
	mu1   sync.RWMutex
)

var jwttoken string = "thisshouldbestoreinenv"

func ServeHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ServeTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Add read lock for concurrency safety
	mu1.RLock()
	defer mu1.RUnlock()

	json.NewEncoder(w).Encode(Tasks)
	for _, task := range Tasks {
		fmt.Printf("Task ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
}

func ServeTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	mu1.RLock()
	defer mu1.RUnlock()

	for _, task := range Tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	// Add write lock for thread-safe slice manipulation
	mu1.Lock()
	defer mu1.Unlock()

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
	w.Header().Set("Content-Type", "application/json")
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	
	// Deprecated rand.Seed removed; Go handles seeding automatically now
	// If using Go versions older than 1.20, keep: rand.Seed(time.Now().UnixNano())
	task.ID = rand.Intn(1000)

	mu1.Lock()
	Tasks = append(Tasks, task)
	mu1.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu1.Lock()
	defer mu1.Unlock()

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
