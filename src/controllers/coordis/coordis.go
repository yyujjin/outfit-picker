package coordis

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func oook () *sql.DB{
	//데이터베이스랑 연결된 상태를 db변수가 보관하고 있는거다. 
	//db 변수 역할 => db connection
	//함수로 쓰려면 db내보내서 변수에 저장 
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db

}


//사용자가 착용한 옷 사진을 업로드하고 이를 날짜,기온,날씨와 함께 기록하는 API
func LogCoordis(c *gin.Context) {

	oook()
	
	type coordi struct {
		Date string `json:"date" binding:"required"` 
		Photo string `json:"photo" binding:"required"`
		Temperature int `json:"temperature"`
		Weather int `json:"weather" binding:"required"`
	}

	var registerCoordi coordi

	if err := c.ShouldBindJSON(&registerCoordi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "필수 입력값을 입력해주세요.",
		})
		return
	}
	
//변수에 저장 
	db := oook()
	result, err := db.Exec(
		"INSERT INTO coordi (date, photo,temperature,weather) VALUES (?,?,?,?) ", 
		registerCoordi.Date, registerCoordi.Photo, registerCoordi.Temperature,registerCoordi.Weather,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
				})
				return
	} 
	fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "추가가 완료되었습니다!",
	})

}

//사용자가 착용한 코디 목록을 조회하는 API  
func GetCoordiLogs(c *gin.Context) {
	
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//식을 쓰는 건 프론트 함수를 써서 받는건 백엔드 
	month := c.Query("month")
	year := c.Query("year")

	fmt.Println(month,year)
	var id int
	var date string
	var photo string
	var temperature int
	var weather int

	type getCoordi struct {
		Id int `json:"id"`
		Date string `json:"date"`
		Photo string `json:"photo"`
		Temperature int `json:"temperature"`
		Weather int `json:"weather"`
	}

	// date >= '2024-02-01' and date <'2024-03-01'
	first := year + "-" + month + "-" + "01"
	fmt.Println(first)

	rows, err := db.Query("SELECT * FROM coordi WHERE date >= ? and date < DATE_ADD(?, INTERVAL 1 MONTH) ORDER BY id ASC;", first, first)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	}
	defer rows.Close() 
	coordiList := []getCoordi{}

	for rows.Next() {
		err := rows.Scan(&id, &date, &photo, &temperature,&weather)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
			})
			return
		}
		coordiList = append(coordiList, getCoordi{id,date,photo,temperature,weather})
	}
	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "해당하는 날짜의 코디가 없습니다.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": coordiList,
	})
}

//사용자의 코디 기록에서 해당하는 정보를 삭제하는 API
func DeleteCoordiLog(c *gin.Context) {

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
		
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id, err)
	if err != nil {
		return
	}

	result, err := db.Exec("DELETE FROM coordi where id = ?",id) 
	fmt.Println(result)
 
	// 지피티가 알려준건데 이거 맞을까?
	rowCount, _ := result.RowsAffected()
	if  rowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "해당하는 ID를 찾을 수 없습니다.",
		})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

}