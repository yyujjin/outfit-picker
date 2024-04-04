package itemsdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InserItem(name string, category int, image string) error {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO closet (name, category,image) VALUES (?,?,?) ", name, category, image)

	return err
}


func SeleteItems() (*sql.Rows, error) {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	
	rows, err := db.Query("SELECT closet.id, closet.name, cl.name as category, closet.image FROM closet join categorylist cl on closet.category=cl.id ")

	return rows, err
}

func DeleteItem(id int) error {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM closet where id = ?",id) 

	return err

}