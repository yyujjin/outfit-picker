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

	type postItem struct {
		ItemId int `json:"itemId"`
		ItemName string `json:"itemName"`
		Category int    `json:"category"`
		Image    string `json:"image"`
	}

	var addItem postItem

	if err := c.BindJSON(&addItem); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request. Please provide valid data for clothing registration.",
		})
		return
	}
	
	err := itemsdb.InserItem(addItem.ItemName, addItem.Category, addItem.Image)

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

	type getItem struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Category string `json:"category"`
		Image string `json:"image"`
	}

	
	rows, err := itemsdb.SeleteItems() 

	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close() 

	item := []getItem{}

	for rows.Next() {

		var id int
		var name string
		var category string	
		var image string

		err := rows.Scan(&id, &name, &category, &image)
		if err != nil {
			log.Fatal(err)  
		}
		
		item = append(item, getItem{id, name, category, image})
		fmt.Println(id, name, category, image)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": item,
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

	err = itemsdb.DeleteItem(id)

	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

}