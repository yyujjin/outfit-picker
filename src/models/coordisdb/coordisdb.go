package coordisdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InsertCoordi(data string, photo string, temperature int, weather int) error {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO coordi (date, photo,temperature,weather) VALUES (?,?,?,?) ",
	data, photo, temperature, weather)

	return err
}

func SelectCoordis(fisrt string)(*sql.Rows, error) {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM coordi WHERE date >= ? and date < DATE_ADD(?, INTERVAL 1 MONTH) ORDER BY id ASC;", first, first)

	return	rows, err
}

func DeleteCoordi(id int) (sql.Result,error) {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM coordi where id = ?",id) 

	return result, err
}