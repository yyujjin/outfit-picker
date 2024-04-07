package categories

import (
	"fmt"
	"log"
	"net/http"
	"outfit-picker/src/models/categorydb"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {

	type category struct {
		Id   int
		Name string
	}

	rows,err := categorydb.SelectCategories()

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() 

	categoryList := []category{}

	for rows.Next() {

		var id int
		var name string

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
}