package login

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetUserCount(userId string) int {
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	err2 := db.QueryRow("SELECT count(*) FROM user WHERE user_id = ?", userId).Scan(&count) // id가 1인 학생을 조회
	if err2 != nil {
		log.Fatal(err)
	}

	return count
}


func InsertUser(id string, hash []byte, name string ,birthday string, phoneNumber string, gender int) error{
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err4 := db.Exec("INSERT INTO user (user_id, password,name,birthday,tel,gender) VALUES (?,?,?,?,?,?) ",
	id, hash, name, birthday, phoneNumber, gender)

	return err4
}


func CheckDuplicateID(id string) (string, error) {
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var userPassword string

	err = db.QueryRow("SELECT password FROM user WHERE user_id = ?", id).Scan(&userPassword)

	return userPassword, err
}