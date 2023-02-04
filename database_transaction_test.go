package GoDatabase

import (
	"database/sql"
	"fmt"
	"testing"
)

// database behavior is auto commit by default
// to disable auto commit, we can use the following code:
// Tx is a transaction. It is not used by the database/sql package itself.
func TestTransactions(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) { // defer the close until the end of the function
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)
	tx, err := db.Begin() // begin the transaction
	if err != nil {
		panic(err.Error())
	}
	// do something with the transaction
	query := "INSERT INTO comments (title, comment) VALUES (?, ?)"
	stmt, err := tx.PrepareContext(ctx, query) // prepare the statement
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.ExecContext(ctx, "TEST GUYS", "Test Comment Guys!")
	if err != nil {
		panic(err.Error())
	}
	lastInsertId, err := result.LastInsertId() // get the last insert id
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Last Insert ID:", lastInsertId) // Last Insert ID: 1

	errTx := tx.Rollback() // rollback the transaction
	if errTx != nil {
		panic(errTx.Error())
	}
}
