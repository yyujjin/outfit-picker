package auth

import (
	"errors"
	"fmt"
	"net/http"
	"outfit-picker/src/models/authdb"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var data authdb.User

	if err := c.BindJSON(&data); err != nil {
		// fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 요청입니다. 올바른 데이터를 제공해주세요.",
		})
		return
	}

	result,userPassword := authdb.GetPassword(data.UserId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
		})
		fmt.Println("찾는 행 없음 ")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(data.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "잘못된 로그인 정보입니다. 다시 시도해주세요.",
		})
		fmt.Println("로그인 실패")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "로그인 되었습니다!",
	})

} 