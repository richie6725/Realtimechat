# Golang Realtime Chat 後端 Server

Chat Application Backend
這是一個使用 Golang 建置的即時聊天後端服務，提供使用者認證（註冊、登入、登出）以及利用 WebSocket 實現即時聊天室功能。

Chat Application Frontend
這個專案是基於 React 與 Tailwind CSS 所打造的即時聊天應用前端。前端透過 WebSocket 與 Golang 後端進行連線，實現聊天室、使用者登入與房間管理等功能。

## 特色
- **即時聊天**：使用 WebSocket 進行聊天室的即時通訊，支援聊天室建立、加入及廣播訊息。
- **使用者管理**：支援使用者註冊、登入與登出。
- **聊天室管理**：：支援創建房間、列出房間與加入聊天室。

## 使用相關技術
- **AWS雲端運算**：將前後端的Server建立在AWS EC2，連接RDS資料庫。
- **Docker化**：將前後端Server透過Docker部屬於AWS EC2上。
- **模組化架構**：內部分為 user、ws (WebSocket) 與 util 模組，架構清晰易於維護。
- **認證機制**：利用 JWT 產生驗證 token，並以 Cookie 方式管理 Session。
- **資料庫整合**：與 MySQL 資料庫連線，內建遷移腳本協助建立必要的資料表。
- **CORS 支援**：配置跨域請求，允許特定前端網址存取 API。
- **使用Tailwind**：採用 Tailwind CSS 打造跨裝置適用的使用者介面。




