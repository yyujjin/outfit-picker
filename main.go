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
			Date string `json:"date" binding:"required"` 
			Photo string `json: "photo" binding:"required"`
			Temperature float32 `json:"temperature"`
			Weather int `json:"weather" binding:"required"`
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
		//위에 에러를 주석하지 않으면 밑에께 작동이 안되는데 위에껀
		//왜 작동을 하는가? 
		if err := c.ShouldBindJSON(&registerCoordi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "필수 입력값을 입력해주세요.",
			})
			return
		}

		result, err := db.Exec(
			"INSERT INTO coordi (date, photo,temperature,weather) VALUES (?,?,?,?) ", 
			registerCoordi.Date, registerCoordi.Photo, registerCoordi.Temperature,registerCoordi.Weather,
		)
		if err != nil {
			log.Fatal(err) // TODO 뭐하는 앤지 알아오기
			c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "Invalid request. Please provide valid data for clothing registration.",
					})
					return
		} 
		fmt.Println(result)
		//이 에러 처리는 굳이 해줄필요가 없나? 
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  "error",
		// 		"message": "Invalid request. Please provide valid data for clothing registration.",
		// 	})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "추가가 완료되었습니다!",
		})

	})

	r.GET("/api/coordis", func(c *gin.Context) {
		var id int
		var date string
		var photo string
		var temperature float32
		var weather int

		type getCoordi struct {
			Id int `json:"id"`
			Date string `json:"date"`
			Photo string `json:"photo"`
			Temperature float32 `json:"temperature"`
			Weather int `json:"weather"`
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
