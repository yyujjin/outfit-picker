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

type getCoordi struct {
	Id int `json:"id"`
	Date string `json:"date"`
	Photo string `json:"photo"`
	Temperature int `json:"temperature"`
	Weather int `json:"weather"`
}

func SelectCoordis(month string, year string) ([]getCoordi,error) {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// date >= '2024-02-01' and date <'2024-03-01'
	first := year + "-" + month + "-" + "01"
	fmt.Println(first)

	rows, err := db.Query("SELECT * FROM coordi WHERE date >= ? and date < DATE_ADD(?, INTERVAL 1 MONTH) ORDER BY id ASC;", first, first)
	if err != nil {
		log.Fatal()
	}
	defer rows.Close() 

	coordiList := []getCoordi{}

	for rows.Next() {

		var id int
		var date string
		var photo string
		var temperature int
		var weather int

		err := rows.Scan(&id, &date, &photo, &temperature,&weather)
		if err != nil {
			log.Fatal()
		}
		coordiList = append(coordiList, getCoordi{id,date,photo,temperature,weather})
	}
	
	return coordiList,err
}


func DeleteCoordi(id int) (sql.Result, error) {

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM coordi where id = ?",id) 
	fmt.Println(result)

	return result, err
 
}