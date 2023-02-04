package GoDatabase

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestExecQuerySQLInjection(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	username := "admin'; #"
	password := "salah"
	sqlQuery := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	// sql injection is possible because the query is not parameterized
	fmt.Println("SQL Query:", sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		var username string
		err = rows.Scan(&username) // scan the result into the variable, similar to rows.Next
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Login Success!", username)
	} else {
		fmt.Println("Login Failed!", username)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
}

func TestExecQuerySQLInjectionWithParameter(t *testing.T) {
	db, ctx := GetConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	username := "admin'; #"
	password := "salah"
	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" // ? is the placeholder
	// sql injection is not possible because the query is parameterized
	rows, err := db.QueryContext(ctx, sqlQuery, username, password) // pass the parameter to the query
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		var username string
		err = rows.Scan(&username) // scan the result into the variable, similar to rows.Next
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Login Success!", username)
	} else {
		fmt.Println("Login Failed!", username)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
}
