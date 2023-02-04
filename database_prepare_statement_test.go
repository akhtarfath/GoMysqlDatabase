package GoDatabase

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments (title, comment) VALUES (?, ?)")
	// prepare a statement for inserting data
	if err != nil {
		panic(err.Error())
	}

	for i := 1; i <= 10; i++ {
		result, err := stmt.ExecContext(ctx, "TEST"+strconv.Itoa(i), "Test Comment "+strconv.Itoa(i)+"!")
		// insert a new record, using stmt as the prepared statement
		if err != nil {
			panic(err.Error())
		}
		lastInsertId, err := result.LastInsertId() // get the last insert id
		fmt.Println("Last Insert ID:", lastInsertId)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close() // close the statement
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
}
