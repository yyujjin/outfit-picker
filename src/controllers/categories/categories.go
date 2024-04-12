package categories

import (
	"net/http"
	"outfit-picker/src/models/categorydb"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {

	categoryList := categorydb.SelectCategories()
	
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": categoryList,
	})
}