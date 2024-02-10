package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type item struct {
	ItemId   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	Category int `json:"category"`
	Image    string `json:"image"`
}

var myCloset []item
var itemId = 1

type category struct {
	id int
	name string
}  

var categoryList = []category{
	{1,"아우터"},
	{2,"상의"},
	{3,"하의"},
	{4,"신발"},
	{5,"가방"},
	{6,"악세서리"},
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
			"data": addItem,
		})

		itemId++
		
	})

	r.GET("/myCloset", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": myCloset,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
