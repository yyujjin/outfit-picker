package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strconv"
)

type item struct {
	ItemId   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	Category int    `json:"category"`
	Image    string `json:"image"`
}

var myCloset []item
var itemId = 1

type category struct {
	id   int
	name string
}

var categoryList = []category{
	{1, "아우터"},
	{2, "상의"},
	{3, "하의"},
	{4, "신발"},
	{5, "가방"},
	{6, "악세서리"},
}

func main() {
	r := gin.Default()

	r.POST("/addToMyCloset", func(c *gin.Context) {
		var addItem item
		// item -> itemId 값을 넣어주고, itemID를 1 증가시켜
		if err := c.BindJSON(&addItem); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}

		addItem.ItemId = itemId

		myCloset = append(myCloset, addItem)
		fmt.Println(myCloset)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Clothing item successfully registered.",
			"data":    addItem,
		})

		itemId++
	})

	r.GET("/myCloset", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": myCloset,
		})
	})

	r.DELETE("/removeItem/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println(id, err)
		if err != nil {
			fmt.Println("경고")
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}
		item, index, ok := lo.FindIndexOf(myCloset, func(item item) bool {
			return item.ItemId == id
		})

		if ok == false {
			c.JSON(http.StatusNotFound, gin.H{
				"data": "id를 찾을 수 없습니다.",
			})
			return
		}
		myCloset = append(myCloset[:index], myCloset[index+1:]...)
		c.IndentedJSON(http.StatusOK, item)

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//옷장 아이템 삭제, 조회 API 만들기
