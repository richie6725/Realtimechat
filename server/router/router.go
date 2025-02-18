package router

import (
	"Go_realtime_chat/server/internal/user"
	"Go_realtime_chat/server/internal/ws"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var r *mux.Router

// 初始化#4
func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = mux.NewRouter()

	// CORS 設定
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.ExposedHeaders([]string{"Content-Length"}),
		handlers.AllowCredentials(),
		handlers.MaxAge(int((12 * time.Hour).Seconds())),
	)

	// 路由設定
	api := r.PathPrefix("/").Subrouter()
	api.HandleFunc("/signup", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/login", userHandler.Login).Methods("POST")
	api.HandleFunc("/logout", userHandler.Logout).Methods("GET")

	api.HandleFunc("/ws/createRoom", wsHandler.CreateRoom).Methods("POST")
	api.HandleFunc("/ws/joinRoom/{roomId}", wsHandler.JoinRoom).Methods("GET")
	api.HandleFunc("/ws/getRooms", wsHandler.GetRooms).Methods("GET")
	api.HandleFunc("/ws/getClients/{roomId}", wsHandler.GetClients).Methods("GET")

	// 包裝 CORS 中介軟體
	http.Handle("/", corsHandler(r))
}

// 初始化#5
func Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
