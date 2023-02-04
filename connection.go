package GoDatabase

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() (*sql.DB, context.Context) {
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err.Error()) // panic will stop the execution of the program
	} else {
		fmt.Println("Connected to database")
	}

	db.SetConnMaxIdleTime(10)               // set the maximum idle time for the connection
	db.SetMaxOpenConns(100)                 // set the maximum number of open connections
	db.SetConnMaxIdleTime(5 * time.Second)  // set the maximum idle time for the connection
	db.SetConnMaxLifetime(60 * time.Second) // set the maximum lifetime for the connection
	return db, ctx
}
