package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/forum_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user_id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("User: %d - %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
