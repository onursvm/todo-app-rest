package models

type Data struct {
	Users []User `json:"users"`
	ToDos []ToDo `json:"todos"`
	Steps []Step `json:"steps"`
}
