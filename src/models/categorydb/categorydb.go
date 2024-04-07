package categorydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func SelectCategories() (*sql.Rows,error){

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM categorylist ORDER BY id ASC;")

	return rows, err
}