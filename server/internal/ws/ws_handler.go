package ws

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid create room payload"}`, http.StatusBadRequest)
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	// 設置回應頭並回傳POSTMAN
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, `{"error":"Invalid join room"}`, http.StatusBadRequest)
		return
	}

	// 使用 mux 獲取路徑參數
	roomID := mux.Vars(r)["roomId"]

	// 使用 r.URL.Query() 獲取查詢參數
	clientid := r.URL.Query().Get("userId")
	username := r.URL.Query().Get("username")

	// 初始化 Client
	cl := &Client{
		Conn:     conn, // 初始化連線物件（替換為實際連線）
		Message:  make(chan *Message, 10),
		ID:       clientid,
		RoomID:   roomID,
		Username: username,
	}

	// 初始化 Message
	msg := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl

	h.hub.Broadcast <- msg

	//加入後持續接收資料，並且另開一個goroutine來傳輸資料
	go cl.writeMessage()
	cl.readMessage(h.hub)

}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	// 設置回應頭並回傳POSTMAN
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []ClientRes
	roomId := mux.Vars(r)["roomId"]

	if _, ok := h.hub.Rooms[roomId]; !ok {
		http.Error(w, `{"error":"Room not found"}`, http.StatusNotFound)
		return
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
