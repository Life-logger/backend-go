package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func InitializeDB() {
	dbInstance = newDBInstance()
}

func GetDBInstance() *sql.DB {
	return dbInstance
}

func newDBInstance() *sql.DB {
	dbConfig := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		fmt.Printf("db connecting something wrong...")
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)

	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		fmt.Printf("conn userError")
	}
	defer conn.Close()

	return db
}
