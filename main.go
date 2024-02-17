package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
	"log"
	"net/http"
	"os"
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

	// 환경변수 읽기
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
		result, err := db.Exec("INSERT INTO closet (name, category,image) VALUES (?,?,?) ", addItem.ItemName, addItem.Category, addItem.Image)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		addItem.ItemId = itemId

		myCloset = append(myCloset, addItem)
		fmt.Println(myCloset)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Clothing item successfully registered.",
			"data":    addItem,
		})

		itemId++ //전역변수는 함수안에서 실행을 끝내도 값이 적용이된다.
	})

	r.GET("/items", func(c *gin.Context) {
		var id int
		var name string
		var category int
		var image string
		rows, err := db.Query("SELECT * FROM closet where id >= ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)
		item := []postItem{}

		for rows.Next() {
			err := rows.Scan(&id, &name, &category, &image)
			if err != nil {
				log.Fatal(err)
			}
			item = append(item, postItem{id, name, category, image})
			fmt.Println(id, name, category, image)
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": item,
		})
	})

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
