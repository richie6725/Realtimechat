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

// 允許的前端網址列表
var allowedOrigins = []string{
	"http://localhost:3000",
	"http://54.250.17.168:3000",
	//"https://another-frontend.com",
}

// 動態檢查請求的 Origin 是否在 `allowedOrigins` 清單中
func dynamicCORS(origin string) bool {
	for _, o := range allowedOrigins {
		if origin == o {
			return true
		}
	}
	return false
}

// 初始化#4
func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = mux.NewRouter()

	// CORS 設定
	corsHandler := handlers.CORS(
		handlers.AllowedOriginValidator(dynamicCORS), // 根據請求的 Origin 來決定是否允許；指定可存取的前端網址
		//handlers.AllowedOrigins([]string{"http://localhost:3000"}), //允許特定的幾個前端網址
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.ExposedHeaders([]string{"Content-Length"}),
		handlers.AllowCredentials(), // 允許攜帶憑證
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

	// 包裝 CORS 中介軟體；確保API都有被CORS處理過
	http.Handle("/", corsHandler(r))
}

// 初始化#5
func Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
