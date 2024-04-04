package auth

import (
	"net/http"
	"outfit-picker/src/models/authdb"

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

	userPassword,err1 := authdb.GetPassword(data.Id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(data.Password)) != nil {
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