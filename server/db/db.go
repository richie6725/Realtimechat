package db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type Database struct {
	db *sql.DB
}

// 初始化#1
func NewDatabase() (*Database, error) {

	connString := "sqlserver://sa:67256725@127.0.0.1:1433?database=websocket"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err.Error())
	}
	fmt.Println("Connected to the database successfully")

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

// 初始化#2
func (d *Database) GetDB() *sql.DB {
	return d.db
}
