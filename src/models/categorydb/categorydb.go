package categorydb

import (

	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Category struct {
	Id   uint `json:"id"`
	Name string `json:"name"`
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

func SelectCategories() []Category{

	db,err := ConnectDB()
	
	if err != nil {
		log.Fatal(err)
	}
	categoryList := []Category{}
	
	db.Model(&Category{}).Table("categoryList").Order("id asc").Find(&categoryList)
	
	return	categoryList
}