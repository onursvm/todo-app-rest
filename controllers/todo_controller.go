package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"todo-app/models"
	"todo-app/storage"
	"todo-app/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetToDos(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	role := r.Context().Value("role").(string)

	var todos []models.ToDo
	for _, todo := range storage.Data.ToDos {
		if todo.DeletedAt == nil && (role == "admin" || todo.Username == username) {
			todoSteps := []models.Step{}
			for _, step := range storage.Data.Steps {
				if step.ToDoID == todo.ID && step.DeletedAt == nil {
					todoSteps = append(todoSteps, step)
				}
			}
			todo.Percent = utils.CalculatePercent(todoSteps)
			todos = append(todos, todo)
		}
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var todo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil || todo.Name == "" {
		http.Error(w, "Geçersiz giriş", http.StatusBadRequest)
		return
	}

	todo.ID = uuid.NewString()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	todo.Username = username
	todo.Percent = 0

	storage.Data.ToDos = append(storage.Data.ToDos, todo)
	utils.SaveData()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateToDo(w http.ResponseWriter, r *http.Request) {
	var updatedToDo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&updatedToDo)
	if err != nil || updatedToDo.ID == "" {
		http.Error(w, "Geçersiz giriş", http.StatusBadRequest)
		return
	}

	for i, todo := range storage.Data.ToDos {
		if todo.ID == updatedToDo.ID && todo.DeletedAt == nil {
			updatedToDo.CreatedAt = todo.CreatedAt
			updatedToDo.UpdatedAt = time.Now()
			updatedToDo.Username = todo.Username
			storage.Data.ToDos[i] = updatedToDo
			utils.SaveData()
			json.NewEncoder(w).Encode(updatedToDo)
			return
		}
	}

	http.Error(w, "TO-DO bulunamadı", http.StatusNotFound)
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	for i, todo := range storage.Data.ToDos {
		if todo.ID == todoID && todo.DeletedAt == nil {
			now := time.Now()
			storage.Data.ToDos[i].DeletedAt = &now
			utils.SaveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "TO-DO bulunamadı", http.StatusNotFound)
}
