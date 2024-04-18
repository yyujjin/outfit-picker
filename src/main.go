package main

import (
	"outfit-picker/src/controllers/auth"
	"outfit-picker/src/controllers/categories"
	"outfit-picker/src/controllers/coordis"
	"outfit-picker/src/controllers/items"
	"outfit-picker/src/controllers/users"
	"outfit-picker/src/controllers/weather"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 코드 구조
// 1. 컨트롤러 2. 서비스 3. 모델

func main() {
	r := gin.Default()

//회원가입 API
r.POST("/api/users", users.SignUp)

//로그인 API 
	r.POST("/api/login",auth.Login )

//자신의 옷장에 추가한 전체 의류 아이템을 확인하기 위한 API	
	r.GET("/api/items",items.GetClothingItems)

//프론트엔드에서 사용자의 옷장에 아이템을 추가하기 위한 API
	r.POST("/api/items",items.AddItem)
 
//사용자의 옷장에서 선택한 아이템을 제거하기 위한 API
	r.DELETE("/api/items/:id", items.DeleteItem )

//사용자가 착용한 코디 목록을 조회하는 API  
	r.GET("/api/coordis",coordis.GetCoordiLogs)

//사용자가 착용한 옷 사진을 업로드하고 이를 날짜,기온,날씨와 함께 기록하는 API
	r.POST("/api/coordis",coordis.LogCoordis)
	
//사용자의 코디 기록에서 해당하는 정보를 삭제하는 API
	r.DELETE("/api/coordis/:id",coordis.DeleteCoordiLog )

//카테고리 리스트를 프론트로 전달 
r.GET("/api/categories", categories.GetCategories)

r.GET("/api/weather", weather.GetWeather)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
