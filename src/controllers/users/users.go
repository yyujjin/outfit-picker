package users

import (
	"fmt"
	"net/http"
	"outfit-picker/src/models/authdb"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//회원가입 API
func SignUp(c *gin.Context) {
	
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

	count := authdb.GetUserCount(data.Id)

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
	
	err := authdb.InsertUser(data.Id, hash, data.Name, data.Birthday, data.PhoneNumber, data.Gender)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}




