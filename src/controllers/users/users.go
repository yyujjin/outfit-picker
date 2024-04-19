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
	
	var data authdb.User

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "필수 값을 입력해 주세요",
		})
		return
	}

	count := authdb.GetUserCount(data.UserId)

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
	
	err := authdb.InsertUser(data.Id,data.UserId,[]byte(hash), data.Name, data.Birthday, data.Tel, data.Gender)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
		})
		fmt.Println("여기서 에러발생")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}




