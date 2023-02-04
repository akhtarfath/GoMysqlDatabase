package GoDatabase

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	query := "INSERT INTO comments (title, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, "TEST", "Test Comment!")
	if err != nil {
		panic(err.Error())
	}

	lastInsertId, err := result.LastInsertId() // get the last insert id
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Last Insert ID:", lastInsertId)
}
