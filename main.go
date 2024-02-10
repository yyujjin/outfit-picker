package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/samber/lo"
    // lop "github.com/samber/lo/parallel"
)

type item struct {
	ItemId   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	Category string `json:"category"`
	Image    string `json:"image"`
}

var myCloset []*item

func findIndex(request int,a []int)int{
	for index,value := range a {
		if value == request {
			return index
		}
	}
	return -1 
}

func main() {
	r := gin.Default()

	

	r.GET("/practice", func(c *gin.Context) {
		var request = 3
		var a = []int {
			1,2,3,4,5,
		}
	
	
		foundValue, foundIndex, ok := lo.FindIndexOf(a, func(value int) bool {
			fmt.Println(value)
			return value == request
		})
		fmt.Println(foundValue, foundIndex, ok)
		c.JSON(http.StatusOK, gin.H{
			"index": foundIndex,//findIndex(request,a),
		})
	})


	r.POST("/addToMyCloset", func(c *gin.Context) {
		var addItem *item
		if err := c.BindJSON(&addItem); err != nil {
			fmt.Println(err)
			return
		}

		myCloset = append(myCloset, addItem)
		fmt.Println(myCloset)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Clothing item successfully registered.",
			"data":    addItem,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request. Please provide valid data for clothing registration.",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
