package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/routes"
	"todo-app/utils"
)

func main() {
	err := utils.LoadData()
	if err != nil {
		log.Fatal("Veri yüklenemedi:", err)
	}

	router := routes.RegisterRoutes()

	fmt.Println("Sunucu 8080 portunda başlatıldı...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
