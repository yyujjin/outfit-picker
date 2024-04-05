package categories

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD
	"outfit-picker/src/models/categorydb"
=======
	"os"
>>>>>>> parent of a4f4a86 (categories sql 모델 생성)

	"github.com/gin-gonic/gin"
)


func GetCategories(c *gin.Context) {

	password := os.Getenv("DB_password")
	fmt.Println(password)
	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/outfit-picker", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	type category struct {
		Id   int
		Name string
	}

	var id int
	var name string

	rows, err := db.Query("SELECT * FROM categorylist ORDER BY id ASC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	categoryList := []category{}

	for rows.Next() {
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