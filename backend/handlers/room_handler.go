package handlers

import (
	"encoding/json"
	"log"
    "strconv"
	"net/http"
	"gorm.io/gorm"
	"backend/models"
)

type RoomHandler struct {
	DB *gorm.DB
}

func (h *RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var rooms []models.Room
	h.DB.Order("id asc").Find(&rooms)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(rooms)
}

func (h *RoomHandler) AddRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("AddRoom: invalid request body", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Println("AddRoom: received data:", room)

	if room.Floor == 0 || room.RoomType == "" || room.PricePerDay == 0 || room.Capacity == 0 {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}


	if room.ID == 0 {
		var maxID int
		h.DB.Model(&models.Room{}).Select("MAX(id)").Scan(&maxID)
		room.ID = uint(maxID + 1)
	}


	var existingRoom models.Room
	if err := h.DB.First(&existingRoom, room.ID).Error; err == nil {
		http.Error(w, "Room with this ID already exists", http.StatusConflict)
		return
	}

	if err := h.DB.Create(&room).Error; err != nil {
		log.Println("AddRoom: DB error", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/rooms/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var room models.Room
	if err := h.DB.First(&room, id).Error; err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	var updateData struct {
		Floor       int     `json:"floor"`
		RoomType    string  `json:"room_type"`
		PricePerDay float64 `json:"price_per_day"`
		Capacity    int     `json:"capacity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	room.Floor = updateData.Floor
	room.RoomType = updateData.RoomType
	room.PricePerDay = updateData.PricePerDay
	room.Capacity = updateData.Capacity

	if err := h.DB.Save(&room).Error; err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) GetPlacePrice(w http.ResponseWriter, r *http.Request) {
	floor, _ := strconv.Atoi(r.URL.Query().Get("floor"))
	roomID, _ := strconv.Atoi(r.URL.Query().Get("room_id"))
	var room models.Room
	result := h.DB.Where("floor = ? AND id = ?", floor, roomID).First(&room)
	if result.Error != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	pricePerPlace := room.PricePerDay / float64(room.Capacity)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]float64{"price_per_place": pricePerPlace})
}

func (h *RoomHandler) GetFreeRoomsAndPlaces(w http.ResponseWriter, r *http.Request) {
	var rooms []models.Room
	h.DB.Preload("Clients").Order("id asc").Find(&rooms)
	totalPlaces := 0
	occupiedPlaces := 0
	freeRooms := 0
	for _, room := range rooms {
		totalPlaces += room.Capacity
		occupiedPlaces += len(room.Clients)
		if len(room.Clients) == 0 {
			freeRooms++
		}
	}
	freePlaces := totalPlaces - occupiedPlaces
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]int{
		"free_places": freePlaces,
		"free_rooms":  freeRooms,
	})
}
