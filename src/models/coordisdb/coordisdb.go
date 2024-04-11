package coordisdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Coordi struct {
	Date string
	Photo string
	Temperature int
	Weather int
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



func InsertCoordi(date string, photo string, temperature int, weather int) error {

	db,err := ConnectDB()
	
	if err != nil {
		return err
	}

	coordi := Coordi{date,photo,temperature,weather}
	
	//coordi에 저장돼 있는 데이터를 뽑아다가 db에 저장시키겠다. 
	result := db.Create(&coordi)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	fmt.Println("입력된 행의 갯수",result.RowsAffected)
	
	return err
}

func SelectCoordis(first string)(*sql.Rows, error) {

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