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
	"golang.org/x/crypto/bcrypt"
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

// func EqualPassword(hashedPassword string, password string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password)) == nil
// }

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

	

	//TODO: 로그인 API 만들기 
	r.POST("/api/login", func(c *gin.Context) {
		type login struct {
			Id string `json : "id" binding:"required"` 
			Password string `josn : "password" binding:"required"`
		}

		var data login
	
		if err := c.BindJSON(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "잘못된 요청입니다. 올바른 데이터를 제공해주세요.",
			})
			return
		}
		
		var userPass string

// 해당되는 행이 없으면 에러반환
		err := db.QueryRow("SELECT password FROM user WHERE user_id = ?",data.Id).Scan(&userPass) 

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
			})
			return
		}
		

		//isUser := EqualPassword(userPass,data.Password)
		//반환값을 변수에 담아도 되고 그냥 해도 되고
		// 성공이 아니면 메세지 담아서 리턴
		if bcrypt.CompareHashAndPassword([]byte(userPass), []byte(data.Password)) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "로그인 되었습니다!",
		})

	})


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

		result, err := db.Exec("DELETE FROM closet where id = ?",id) 
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

	})

	r.POST("/api/users", func(c *gin.Context) {
		type signup struct {
			Id string `json:"id" binding:"required"` 
			Password string `json:"password" binding:"required"` 
			Name string `json:"name" binding:"required"` 
			Birthday string `json:"birthday" binding:"required"`  
			PhoneNumber string `json:"phoneNumber" binding:"required"` 
			Gender int `json:"gender" ` 
		}

		var data signup

		if err := c.BindJSON(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "필수 값을 입력해 주세요",
			})
			return
		}
	
		var count int
		err := db.QueryRow("SELECT count(*) FROM user WHERE user_id = ?",data.Id).Scan(&count) // id가 1인 학생을 조회
		if err != nil {
			log.Fatal(err)
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "id가 중복되었습니다.",
			})
			return
		}

		pass := []byte(data.Password)

		hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hash))
		
		result, err := db.Exec("INSERT INTO user (user_id, password,name,birthday,tel,gender) VALUES (?,?,?,?,?,?) ",
		 data.Id, hash, data.Name, data.Birthday, data.PhoneNumber, data.Gender)

		if err != nil {
			log.Fatal(err)
			fmt.Println(result)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
			})
			
		}
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
			log.Fatal(err) // TODO: 뭐하는 앤지 알아오기
			c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "Invalid request. Please provide valid data for clothing registration.",
					})
					return
		} 
		fmt.Println(result)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "추가가 완료되었습니다!",
		})

	})
	
	r.GET("/api/coordis", func(c *gin.Context) {

		month := c.Query("month")
		year := c.Query("year")

		fmt.Println(month,year)
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

		// date >= '2024-02-01' and date <'2024-03-01'
		first := year + "-" + month + "-" + "01"
		fmt.Println(first)

		rows, err := db.Query("SELECT * FROM coordi WHERE date >= ? and date < DATE_ADD(?, INTERVAL 1 MONTH) ORDER BY id ASC;", first, first)
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

	r.DELETE("/api/coordis/:id", func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println(id, err)
		if err != nil {
			fmt.Println("경고")
			return
		}

		result, err := db.Exec("DELETE FROM coordi where id = ?",id) 
		fmt.Println(result)
	 
		if err != nil {
			log.Fatal(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
			})
			
		}

		c.IndentedJSON(http.StatusOK, "삭제가 완료되었습니다!")

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
