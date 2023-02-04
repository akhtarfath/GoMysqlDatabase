package GoDatabase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db, _ := GetConnection() // call the function
	if db == nil {           // check if the connection is nil
		panic("Connection is nil")
	} else {
		fmt.Println("Connected to database")
	}
}
