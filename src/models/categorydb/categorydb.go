package categorydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type category struct {
	Id   int
	Name string
}

func GetCategoryList() []category {

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	categoryList := []category{}

	rows, err := db.Query("SELECT * FROM categorylist ORDER BY id ASC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() 
	

	for rows.Next() {

		var id int
		var name string

		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		categoryList = append(categoryList, category{id, name})
		fmt.Println(id, name)
	}

	return categoryList

}