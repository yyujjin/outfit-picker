package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"outfit-picker/src/models/login"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//회원가입 API
//err 1234
func SignUp(c *gin.Context) {
	
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
//생일 필수값 안해도 될것같음 , 폰넘버도
	type signup struct {
		Id string `json:"id" binding:"required"` 
		Password string `json:"password" binding:"required"` 
		Name string `json:"name" binding:"required"` 
		Birthday string `json:"birthday" binding:"required"`  
		PhoneNumber string `json:"phoneNumber" binding:"required"` 
		Gender int `json:"gender" ` 
	}

	var data signup
//err재할당 
	if err1 := c.BindJSON(&data); err1 != nil {
		fmt.Println(err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "필수 값을 입력해 주세요",
		})
		return
	}

	
	count := login.GetUserCount(data.Id)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id가 중복되었습니다.",
		})
		return
	}

	pass := []byte(data.Password)

	hash, err3 := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(hash))
	
	err4 := login.InsertUser(data.Id, hash, data.Name, data.Birthday, data.PhoneNumber, data.Gender)

	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}




