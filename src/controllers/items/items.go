package items

import (
	"fmt"
	"log"
	"net/http"
	"outfit-picker/src/models/itemsdb"
	"strconv"
	"github.com/gin-gonic/gin"
)

//프론트엔드에서 사용자의 옷장에 아이템을 추가하기 위한 API
func AddItem (c *gin.Context){

	var addItem itemsdb.Closet

	if err := c.BindJSON(&addItem); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request. Please provide valid data for clothing registration.",
		})
		return
	}
	//클라이언트에서 itmeID를 안넘겨줘도 (=기본값0)
	//GORM에서 ID는 프라이머리키라서 자동으로 DB에 등록됨. 
	err := itemsdb.InserItem(addItem)

	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "추가가 완료되었습니다!",
	})
}

//자신의 옷장에 추가한 전체 의류 아이템을 확인하기 위한 API	
func GetClothingItems(c *gin.Context) {

	closets, err := itemsdb.SeleteItems() 

	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": closets,
	})
}


//사용자의 옷장에서 선택한 아이템을 제거하기 위한 API
func DeleteItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id, err)

	if err != nil {
		fmt.Println("경고")
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error",
			"message": "Invalid request. Please provide valid data for clothing registration.",
		})
		return
	}

	result := itemsdb.DeleteItem(uint(id))

	if result.Error != nil {
		log.Fatal(err)
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "해당하는 ID를 찾을 수 없습니다.",
		})
		return
	}
	
	c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

}