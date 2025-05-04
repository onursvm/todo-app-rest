package utils

import (
	"encoding/json"
	"os"
	"todo-app/config"
	"todo-app/models"
	"todo-app/storage"
)

var Data models.Data

func LoadData() error {
	file, err := os.ReadFile(config.DataFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Data)
	storage.Data.ToDos = Data.ToDos
	storage.Data.Steps = Data.Steps
	return err
}

func SaveData() error {
	Data.ToDos = storage.Data.ToDos
	Data.Steps = storage.Data.Steps
	dataBytes, err := json.MarshalIndent(Data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(config.DataFile, dataBytes, 0644)
}

func CalculatePercent(steps []models.Step) float64 {
	if len(steps) == 0 {
		return 0
	}
	done := 0
	for _, s := range steps {
		if s.Done {
			done++
		}
	}
	return float64(done) / float64(len(steps)) * 100
}
