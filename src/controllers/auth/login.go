package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	type login struct {
		Id       string `json:"id" binding:"required"`
		Password string `josn:"password" binding:"required"`
	}

	var data login

	if err := c.BindJSON(&data); err != nil {
		// fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 요청입니다. 올바른 데이터를 제공해주세요.",
		})
		return
	}

	var userPass string

	// !!! db 어쩌고 빨간색 에러뜨면 아래 코드 복사하셈
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// !!!!!!!


//변수를 만들어서 거기에 값을 스캔하고 그 스캔된 값으로 해쉬 검사 
//이거하기 
	// userPassword,err := login.CheckDuplicateID(data.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
		})
		return
	}

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

} 