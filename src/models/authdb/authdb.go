package authdb

//auth로해서 user에 있는거랑 login에 있는거랑 합쳐?
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id           int          
	User_id         string         
	Password        string        
	Name          string          
	Birthday      time.Time    
	Tel string 
	Gender  int   
  }

func GetUserCount(userId string) int64 {
	
	password := os.Getenv("DB_password")//환경변수를 불러오는 함수
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker?charset=utf8mb4&parseTime=True&loc=Local",password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
	}

	var count int64
	db.Model(&User{}).Table("user").Where("user_id = ?", userId).Count(&count)

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


func GetPassword(id string) (string, error) {
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var userPass string

	err = db.QueryRow("SELECT password FROM user WHERE user_id = ?", id).Scan(&userPass)

	return userPass, err
}