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

func CreateStep(w http.ResponseWriter, r *http.Request) {
	var step models.Step
	err := json.NewDecoder(r.Body).Decode(&step)
	if err != nil || step.ToDoID == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	step.ID = uuid.NewString()
	step.CreatedAt = time.Now()
	step.UpdatedAt = time.Now()

	storage.Data.Steps = append(storage.Data.Steps, step)
	utils.SaveData()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(step)
}

func UpdateStep(w http.ResponseWriter, r *http.Request) {
	var updatedStep models.Step
	if err := json.NewDecoder(r.Body).Decode(&updatedStep); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	for i, step := range storage.Data.Steps {
		if step.ID == updatedStep.ID && step.DeletedAt == nil {
			updatedStep.CreatedAt = step.CreatedAt
			updatedStep.UpdatedAt = time.Now()
			storage.Data.Steps[i] = updatedStep
			utils.SaveData()
			json.NewEncoder(w).Encode(updatedStep)
			return
		}
	}

	http.Error(w, "Adım bulunamadı", http.StatusNotFound)
}

func DeleteStep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stepID := vars["id"]

	for i, step := range storage.Data.Steps {
		if step.ID == stepID && step.DeletedAt == nil {
			now := time.Now()
			storage.Data.Steps[i].DeletedAt = &now
			utils.SaveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Adım bulunamadı", http.StatusNotFound)
}
