package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service
}

// 初始化#3
func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

//#POST #1

// handler func，用於跟http收取JSON資料，並傳給service業務
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u CreateUserReq

	// 從請求體解析 JSON
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// 調用業務邏輯層處理請求
	res, err := h.Service.CreateUser(r.Context(), &u) //接user_service的CreateUser
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// 設置回應頭並回傳POSTMAN
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user LoginUserReq
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	//跳轉到Service的Login
	u, err := h.Service.Login(r.Context(), &user)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	// 設置 Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    u.accessToken,
		MaxAge:   60 * 60, // 一小時
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
	})

	// 返回 JSON 響應
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// 清除 Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1, // 立即過期
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
	})

	// 返回登出成功的訊息
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "logout successful"})
}
