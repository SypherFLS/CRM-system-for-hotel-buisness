package main

import (
	"log"
	"net/http"
	"backend/db"
	"backend/handlers"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func main() {
	db.Init()

	roomHandler := handlers.RoomHandler{DB: db.DB}
	clientHandler := handlers.ClientHandler{DB: db.DB}
	employeeHandler := handlers.EmployeeHandler{DB: db.DB}
	cleaningHandler := handlers.CleaningHandler{DB: db.DB}

	// Rooms
	http.HandleFunc("/api/rooms/place-price", enableCORS(roomHandler.GetPlacePrice))
	http.HandleFunc("/api/rooms/free", enableCORS(roomHandler.GetFreeRoomsAndPlaces))
	http.HandleFunc("/api/rooms", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			roomHandler.AddRoom(w, r)
		} else if r.Method == http.MethodGet {
			roomHandler.GetAllRooms(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/api/rooms/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			roomHandler.UpdateRoom(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	// Clients
	http.HandleFunc("/api/clients", enableCORS(clientHandler.GetAllClients))
	http.HandleFunc("/api/clients/by-city", enableCORS(clientHandler.GetClientsByCity))
	http.HandleFunc("/api/clients/in-single-rooms", enableCORS(clientHandler.GetClientsInSingleRooms))
	http.HandleFunc("/api/clients/total-payment", enableCORS(clientHandler.GetTotalPayment))
	http.HandleFunc("/api/clients/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			clientHandler.DeleteClient(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	// Employees
	http.HandleFunc("/api/employees", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			employeeHandler.GetAllEmployees(w, r)
		case http.MethodPost:
			employeeHandler.AddEmployee(w, r)
		case http.MethodDelete:
			employeeHandler.DeleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/api/employees/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			employeeHandler.UpdateEmployee(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	// Cleanings
	http.HandleFunc("/api/cleanings", enableCORS(cleaningHandler.GetCleanings))

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
