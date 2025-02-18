package main

import (
	"Go_realtime_chat/server/db"
	"Go_realtime_chat/server/internal/user"
	"Go_realtime_chat/server/internal/ws"
	"Go_realtime_chat/server/router"
	"log"
)

func main() {

	//建立順序:
	//1.先搞定database的建立，以及database內的table
	//2.建立user_repository，也就是完成向repository填值的SQL命令
	//3.建立user_service來連接handler接受context後與repository之間的互動
	//4.建立user_handler來接受http傳送JSON的資料
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("could not connect to database,%s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	//1.定義Room,Hub等內容，再定義Newhub方法
	//2.定義newHandler，定義建立創建房間(CreatRoom)的handler function
	//3.設定handler JoinRoom，其中將加入者的資料廣播、以及在房間中建立使用者的ID，並且write read message
	//4.用goroutine來寫Run()當收到廣播、註冊資料後的動作
	//5.設定GetRooms，GetClinet function
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start(":8080")
}
