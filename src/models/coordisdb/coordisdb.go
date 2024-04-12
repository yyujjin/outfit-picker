package coordisdb

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Coordi struct {
	Id uint
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



func InsertCoordi(id uint,date string, photo string, temperature int, weather int) error {

	db,err := ConnectDB()
	if err != nil {
		return err
	}

	coordi := Coordi{id,date,photo,temperature,weather}
	
	//coordi에 저장돼 있는 데이터로 db를 만들겠다.  
	result := db.Create(&coordi)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	fmt.Println("입력된 행의 갯수",result.RowsAffected)
	
	return err
}

func SelectCoordis(first string)([]Coordi) {
//TODO: first 사용해서 조건문 완성해야함 
	db,err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	coordis := []Coordi{}
	rows, err := db.Model(&Coordi{}).Where("weather=?",0).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// rows, err := db.Query("SELECT * FROM coordi WHERE date >= ? and date < DATE_ADD(?, INTERVAL 1 MONTH) ORDER BY id ASC;", first, first)

	for rows.Next() {
		var coordi Coordi
		db.ScanRows(rows,&coordi) //coordi라는 변수에 내가 보낸 구조체 형식대로 저장해
		coordis = append(coordis,Coordi(coordi))
	}

	fmt.Println(coordis)

	return	coordis
}


func DeleteCoordi(id int) *gorm.DB {

	db,err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// db.Where("id = ?",id).Delete(&Coordi)//구조체에 id필드가 없어서 에러발생 
	
	result := db.Delete(&Coordi{},id)
	// result, err := db.Exec("DELETE FROM coordi where id = ?",id) 

	return result
}