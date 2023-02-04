package GoDatabase

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSqLParameter(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	query := "INSERT INTO user (username, password) VALUES ('mfthnna', 'mfthnna')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert Success!")
}
