package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Hostname string `json:"hostname"`
	ClientIP string `json:"client_ip"`
}

func getHostname(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "HOSTNAME not set"
	}

	// Получаем IP адрес клиента
	clientIP := r.RemoteAddr

	response := Response{
		Hostname: hostname,
		ClientIP: clientIP,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Логируем IP адрес запроса
	log.Printf("Received request from %s", clientIP)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", getHostname)

	// Логируем запуск приложения
	log.Println("Application started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
