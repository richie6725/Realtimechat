## 使用 Golang 作為建置階段
#FROM golang:1.23.2 AS builder
#
## 設定工作目錄
#WORKDIR /app
#
## 複製 go.mod 和 go.sum 以加速依賴下載
#COPY go.mod go.sum ./
#
## 下載所有依賴
#RUN go mod download
#
## 複製專案的所有程式碼
#COPY . .
#
## 建置 Golang 應用程式
#RUN go build -o /go_realtime_chat server/cmd/main.go
#
## 使用較小的執行環境
#FROM alpine:latest
#
## 設定工作目錄
#WORKDIR /root/
#
## 複製建置好的執行檔
#COPY --from=builder /go_realtime_chat .
#
## 指定容器啟動時的指令
#CMD ["./go_realtime_chat"]

#V2
# 使用 Golang 作為建置階段
#FROM golang:1.23.2 AS builder
#
#WORKDIR /build
#
#COPY go.mod go.sum ./
#RUN go mod download
#COPY . .
#RUN go build -o /go_realtime_chat server/cmd/main.go
#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /build/go_realtime_chat ./go_realtime_chat
#CMD ["./go_realtime_chat"]

#V3
# Build stage
FROM golang:1.23.2-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 安裝必要的系統依賴
RUN apk add --no-cache gcc musl-dev

# 複製 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製整個專案
COPY . .

# 編譯應用
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./server/cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# 安裝必要的運行時依賴
RUN apk add --no-cache ca-certificates

# 從 builder stage 複製編譯好的執行檔
COPY --from=builder /app/main .

# 複製資料庫遷移文件
COPY --from=builder /app/server/db/migrations ./server/db/migrations

# 設定時區（選用）
RUN apk add --no-cache tzdata
ENV TZ=Asia/Taipei

# 暴露應用端口（根據你的應用需求修改）
EXPOSE 8080

# 運行應用
CMD ["./main"]