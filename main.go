package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"strconv"
)

type postItem struct {
	ItemId int `json:itemId`
	ItemName string `json:"itemName"`
	Category int    `json:"category"`
	Image    string `json:"image"`
}

type category struct {
	Id   int
	Name string
}

type userInformation struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   bool   `json:"gender"` // true=male, flase=female
}


func switchGenderToInt (gender bool) int {
	if gender == true {
		return 0
	}else{
		return 1
	}
}



func main() {
	r := gin.Default()

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.POST("/api/items", func(c *gin.Context) {
		var addItem postItem

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
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "추가가 완료되었습니다!",
		})
	})
 
	r.GET("/api/items", func(c *gin.Context) {
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
	})

	r.GET("/api/categories", func(c *gin.Context) {
		var id int
		var name string

		rows, err := db.Query("SELECT * FROM categorylist ORDER BY id ASC;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)
		categoryList := []category{}

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			categoryList = append(categoryList, category{id, name})
			fmt.Println(id, name)
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": categoryList,
		})
	})

	r.DELETE("/api/items/:id", func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println(id, err)
		if err != nil {
			fmt.Println("경고")
			c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}

		result, err := db.Exec("DELETE FROM closet where id = ?",id) //%c가아닌 %d 이유
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

	})

	
	r.POST("/api/users", func(c *gin.Context) {
		var data userInformation
		if err := c.BindJSON(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}

		a := switchGenderToInt(data.Gender) //함수만들기는 main밖에서 

		result, err := db.Exec("INSERT INTO user (user_id, password,name,gender) VALUES (?,?,?,?) ",
		 data.UserId, data.Password, data.Name, a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		c.IndentedJSON(http.StatusOK, data)
	})

	r.POST("/api/coordis", func(c *gin.Context) {
		type coordi struct {
			Date string `json:"date"`
			Photo string `json: "photo"`
			Temperature float32 `json:"temperature"`
			Weather int `json:"weather"`
		}

		var registerCoordi coordi

		if err := c.BindJSON(&registerCoordi); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request. Please provide valid data for clothing registration.",
			})
			return
		}

		result, err := db.Exec(
			"INSERT INTO coordi (date, photo,temperature,weather) VALUES (?,?,?,?) ", 
			registerCoordi.Date, registerCoordi.Photo, registerCoordi.Temperature,registerCoordi.Weather,
		)
		if err != nil {
			log.Fatal(err)
		} 
		fmt.Println(result)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "추가가 완료되었습니다!",
		})

		// TODO: 에러처리 해야함
	})

	r.GET("/api/coordis", func(c *gin.Context) {
		var id int
		var date string
		var photo string
		var temperature float32
		var weather int

		type getCoordi struct {
			// TODO: 소문자로 바꾸기
			Id int 
			Date string
			Photo string
			Temperature float32
			Weather int
		}

		rows, err := db.Query("SELECT * FROM coordi ORDER BY id ASC;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)
		coordiList := []getCoordi{}

		for rows.Next() {
			err := rows.Scan(&id, &date, &photo, &temperature,&weather)
			if err != nil {
				log.Fatal(err)
			}
			coordiList = append(coordiList, getCoordi{id,date,photo,temperature,weather})
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": coordiList,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
