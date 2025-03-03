# Golang Realtime Chat

基於 Golang 做為後端、 React 與 Tailwind CSS做為前端的即時聊天服務，提供使用者認證（註冊、登入、登出），與房間管理等功能，以及利用 WebSocket 實現即時聊天室功能。
並透過docker將專案部屬至AWS EC2上執行，並連接RDS達成雲端建置。

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

## Youtube影片連結:
https://youtu.be/ydCoqO49qCc

## 安裝與使用說明
- **設定資料庫**：
- 1. 確認安裝並啟動MySQL資料庫。
  2. 修改server/db/db.go 中的連線字串，依照您資料庫的參數（例如使用者、密碼、主機位址與資料庫名稱）。
  3. 執行srver/db/migrations 中的 SQL 腳本，以建立必要的資料表。
- **後端設定**：
  1. 安裝依賴套件，執行go mod download。
  2. CORS設定:在 server/router/router.go 中配置了 CORS，請新增自己的前端來源請求。
  3. 啟動:執行server/cmd/main.go
- **前端設定**：
  1. 安裝依賴套件:npm install
  2. 設定環境變數:請確認 constants 檔案中的URL根據後端位置調整對應URL。
  3. 啟動:npm run dev

