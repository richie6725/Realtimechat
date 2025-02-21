## 使用官方 Node.js 版本
#FROM node:18-alpine
#
## 設定工作目錄
#WORKDIR /app
#
## 複製 package.json 和 package-lock.json
#COPY package*.json ./
#
## 安裝依賴
#RUN npm install
#
## 複製專案所有檔案
#COPY . .
#
## 編譯 Next.js 應用
#RUN npm run build
#
## 設定環境變數
#ENV PORT=3000
#
## 暴露 3000 端口
#EXPOSE 3000
#
## 啟動應用
#CMD ["npm", "run", "start"]


# 使用 Node.js 18 作為基礎映像
FROM node:18-alpine

# 設定工作目錄
WORKDIR /app

# 安裝必要的系統套件
RUN apk add --no-cache libc6-compat

# 複製 package.json 和 package-lock.json（如果存在）
COPY package*.json ./

# 安裝依賴
RUN npm install

# 複製其餘的專案檔案
COPY . .

# 建立生產環境版本
RUN npm run build

# 設定環境變數
ENV NODE_ENV=production
ENV PORT=3000

# 開放埠號
EXPOSE 3000

# 啟動應用程式
CMD ["npm", "start"]