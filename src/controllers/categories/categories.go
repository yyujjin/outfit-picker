package categories

import (
	"net/http"
	"outfit-picker/src/models/auth/categorydb"

	"github.com/gin-gonic/gin"
)


func GetCategories(c *gin.Context) {

	categoryList := categorydb.GetCategoryList()

	c.JSON(http.StatusOK, gin.H{
		"data": categoryList,
	})
}