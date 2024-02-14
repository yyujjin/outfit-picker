package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strconv"
)

type postItem struct {
	ItemId   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	Category int    `json:"category"`
	Image    string `json:"image"`
}

type getItem struct {
	ItemId   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	Category string `json:"category"`
	Image    string `json:"image"`
}

var myCloset []postItem
var itemId = 1
var userId = 1

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

type userInformation struct {
	Id       int    `json:"id"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   bool   `json:"gender"`
}

// true=male, flase=female
var userInformationList = []userInformation{}

func main() {
	r := gin.Default()

	r.POST("/items", func(c *gin.Context) {
		var addItem postItem
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

		itemId++ //함수안에서만 증가하는게 아닌 전역 변수에도 영향을 미치는 이유?
	})

	r.GET("/items", func(c *gin.Context) {
		fmt.Println(myCloset)
		getItemArr := []getItem{}
		for _, value := range myCloset {

			var categoryName string
			for _, categoryValue := range categoryList {
				if value.Category == categoryValue.id {
					categoryName = categoryValue.name
				}
			}

			getItemArr = append(getItemArr, getItem{
				ItemId:   value.ItemId,
				ItemName: value.ItemName,
				Category: categoryName,
				Image:    value.Image,
			})
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": getItemArr,
		})
	})

	//type intArray []int // 	별명만들기
	//
	//var a = []int{1, 2, 3}
	//var a = intArray{1, 2, 3}
	r.DELETE("/items/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println(id, err)
		if err != nil {
			fmt.Println("경고")
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}
		item, index, ok := lo.FindIndexOf(myCloset, func(item postItem) bool {
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

	r.POST("/users", func(c *gin.Context) {
		var data userInformation
		if err := c.BindJSON(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}
		data.Id = userId
		userInformationList = append(userInformationList, data)
		fmt.Println(userInformationList)
		c.IndentedJSON(http.StatusOK, data)
		userId++
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
