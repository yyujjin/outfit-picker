package coordis

import (
	"fmt"
	"net/http"
	"outfit-picker/src/models/coordisdb"
	"strconv"

	"github.com/gin-gonic/gin"
)

//사용자가 착용한 옷 사진을 업로드하고 이를 날짜,기온,날씨와 함께 기록하는 API
func LogCoordis(c *gin.Context) {
	
	type coordi struct {
		Date string `json:"date" binding:"required"` 
		Photo string `json:"photo" binding:"required"`
		Temperature int `json:"temperature"`
		Weather *int `json:"weather" binding:"required"`
	}

	data := &coordi{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "필수 입력값을 입력해주세요.",
		})
		return
	}
	
	err := coordisdb.InsertCoordi(data.Date,data.Photo,data.Temperature,*data.Weather)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "추가가 완료되었습니다!",
	})

}

//사용자가 착용한 코디 목록을 조회하는 API  
func GetCoordiLogs(c *gin.Context) {

	month := c.Query("month")
	year := c.Query("year")

	fmt.Println(month,year)
	
	// date >= '2024-02-01' and date <'2024-03-01'
	first := year + "-" + month + "-" + "01"
	fmt.Println(first)

	getCoordis := coordisdb.SelectCoordis(first) 
	if len(getCoordis) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "해당하는 날짜의 코디가 없습니다.",
		})
		return
	}
	
	c.IndentedJSON(http.StatusOK,getCoordis)
}

//사용자의 코디 기록에서 해당하는 정보를 삭제하는 API
func DeleteCoordiLog(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id, err)
	if err != nil {
		return
	}

	result, err := coordisdb.DeleteCoordi(id)
	
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

	c.JSON(http.StatusOK, "삭제가 완료되었습니다!")

}