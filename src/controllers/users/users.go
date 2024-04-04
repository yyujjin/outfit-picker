package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"outfit-picker/src/models/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//회원가입 API
func SignUp(c *gin.Context) {
	
	password := os.Getenv("DB_password")
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	type signup struct {
		Id string `json:"id" binding:"required"` 
		Password string `json:"password" binding:"required"` 
		Name string `json:"name" binding:"required"` 
		Birthday string `json:"birthday" binding:"required"`  
		PhoneNumber string `json:"phoneNumber" binding:"required"` 
		Gender int `json:"gender" ` 
	}

	var data signup

	if err = c.BindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "필수 값을 입력해 주세요",
		})
		return
	}

	
	count := auth.GetUserCount(data.Id)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id가 중복되었습니다.",
		})
		return
	}

	pass := []byte(data.Password)

	//TODO: 여기서는 err 재할당을 할수없는것인가? err말고 다른 변수명써야하나?
	hash, err3 := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(hash))
	
	err= auth.InsertUser(data.Id, hash, data.Name, data.Birthday, data.PhoneNumber, data.Gender)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}




