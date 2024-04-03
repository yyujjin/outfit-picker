package items

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//프론트엔드에서 사용자의 옷장에 아이템을 추가하기 위한 API
func AddItem (c *gin.Context){
//이거 함수 바깥에 했더니 안되던데 그러면 안됨?
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

//id필요 없을 듯 
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
	//근데 result 는 무슨 값이지?
	result, err := db.Exec("INSERT INTO closet (name, category,image) VALUES (?,?,?) ", addItem.ItemName, addItem.Category, addItem.Image)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "추가가 완료되었습니다!",
	})
}

//자신의 옷장에 추가한 전체 의류 아이템을 확인하기 위한 API	
func GetClothingItems(c *gin.Context) {

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	type getItem struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Category string `json:"category"`
		Image string `json:"image"`
	}

	var id int
	var name string
	var category string	
	var image string
	rows, err := db.Query("SELECT closet.id, closet.name, cl.name as category, closet.image FROM closet join categorylist cl on closet.category=cl.id ")
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	item := []getItem{}

	// 
	for rows.Next() {
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

	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//시맨틱은 이렇게 써야지만 되는거 
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id, err)
	if err != nil {
		fmt.Println("경고")
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error",
			"message": "Invalid request. Please provide valid data for clothing registration.",
		})
		return
	}

	result, err := db.Exec("DELETE FROM closet where id = ?",id) 
	fmt.Println(result)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

}