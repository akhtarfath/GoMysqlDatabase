package GoDatabase

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestQuerySQL(t *testing.T) {
	db, ctx := GetConnection() // call the function
	defer func(db *sql.DB) {   // defer will execute the function after the function is finished
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	query := "SELECT * FROM customer"
	// rows will contain the result of the query
	// execute the query and check if there is an error
	rows, err := db.QueryContext(ctx, query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("")
	fmt.Println("Result of the query")
	for rows.Next() { // loop through the rows, similar to while loop
		var id, name string
		err = rows.Scan(&id, &name) // scan the result into the variable
		// rows.Scan is similar to rows.Next but it will return the value of the column
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ID:", id)
		fmt.Println("NAME:", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db, ctx := GetConnection() // call the function
	defer func(db *sql.DB) {   // defer will execute the function after the function is finished
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	// rows will contain the result of the query
	// execute the query and check if there is an error
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	for rows.Next() { // loop through the rows, similar to while loop
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		var err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // scan the result into the variable
		// rows.Scan is similar to rows.Next but it will return the value of the column
		if err != nil {
			panic(err.Error())
		}

		formatLayout := "Monday, 02 Jan 2006 15:04:05"
		fmt.Println("================================")
		fmt.Println("ID:", id)
		fmt.Println("NAME:", name)
		if email.Valid {
			fmt.Println("EMAIL:", email.String)
		} else {
			fmt.Println("EMAIL: -")
		}
		fmt.Println("BALANCE:", balance)
		fmt.Println("RATING:", rating)
		if birthDate.Valid {
			fmt.Println("BIRTH DATE:", birthDate.Time.Format(formatLayout))
		} else {
			fmt.Println("BIRTH DATE: -")
		}
		fmt.Println("MARRIED:", married)
		fmt.Println("CREATED AT:", createdAt.Format(formatLayout))
	}

	// close the rows after the loop is finished
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	if err != nil {
		panic(err.Error())
	}
}
