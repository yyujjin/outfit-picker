package itemsdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Closet struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Category int `json:"category"`
	Image string `json:"image"`
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

func InserItem(id uint, name string, category int, image string) error {

	db,err := ConnectDB()
	if err != nil {
		return err
	}

	closet := Closet{id,name,category,image}
	result := db.Create(&closet)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return err
}


func SeleteItems() ([]Closet, error) {

	db,err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	
	closets := []Closet{}
	db.Raw("SELECT closet.id, closet.name, cl.name as category, closet.image FROM closet join categorylist cl on closet.category=cl.id ").Scan(&closets)
	// rows, err := db.Query("SELECT closet.id, closet.name, cl.name as category, closet.image FROM closet join categorylist cl on closet.category=cl.id ")

	return closets, err
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