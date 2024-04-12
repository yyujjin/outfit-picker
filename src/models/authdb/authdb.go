package authdb


import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Id uint        
	UserId string         
	Password string        
	Name string          
	Birthday string    
	Tel string 
	Gender  int   
  }

  func ConnectDB() (*gorm.DB,error) {
	password := os.Getenv("DB_password")
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 단수형 테이블명을 사용합니다. 기본적으로 GORM은 복수형 테이블명 규칙이 적용되는데 true로 설정하면 구조체 이름 그대로 테이블명을 생성합니다.
		  },
	})
	return db,err
}


func GetUserCount(userId string) int64 {
	db,err := ConnectDB()
	if err != nil {
		log.Fatal()
	}

	var count int64
	db.Model(&User{}).Table("user").Where("user_id = ?", userId).Count(&count)

	return count
}


func InsertUser(id uint,userId string, hash []byte, name string ,birthday string, phoneNumber string, gender int) error{
	db,err := ConnectDB()
	if err != nil {
		return err
	}

	userinfo := User{id,userId,string(hash),name,birthday,phoneNumber,gender}

	result := db.Create(&userinfo)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	fmt.Println("입력된 행의 갯수",result.RowsAffected)

	return err
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