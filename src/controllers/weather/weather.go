package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//c에 클라이언트가 요청한 정보가 다 담겨있고 함수안에서 제서
//라우터를 만들면 자동으로 파라미터를 넘겨받게 설계되어있음.

//환경변수로 빼는건 중요한거만 빼기

func GetUrl()string {
	serviceKey := os.Getenv("serviceKey")
	numOfRows := 10
	pageNo := 1
	base_date := "20240421" //보통 날짜는 스트링으로 함
	base_time := "0500" //0으로 시작하는 숫자는 없어서 스트링임
	nx := 55
	ny := 127
	dataType := "JSON"
	//변수타입 맞춰서 넣기 %s %c

	//go규칙 : 파라미터 전달 시 괄호가 밑에 있으면 마지막에 쉼표를 넣어주기 
	return fmt.Sprintf(
		"http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getVilageFcst?&serviceKey=%s&numOfRows=%d&pageNo=%d&base_date=%s&base_time=%s&nx=%d&ny=%d&dataType=%s",
		serviceKey,numOfRows,pageNo,base_date,base_time,nx,ny,dataType,
	)
}

func GetWeather(c *gin.Context) {  

	  
	  type Item struct {
		BaseDate string `json:"baseDate"`
		BaseTime string `json:"baseTime"`
		Category string `json:"category"`
		FcstData string `json:"fcstData"`
		FcstTime string `json:"fcstTime"`
		FcstValue string `json:"fcstValue"`
		Nx int `json:"nx"`
		Ny int `json:"ny"`
	  }

	  type Items struct {
		Item []Item `json:"item"`
	  }

	  type Body struct {
		DataType string `json:"dataType"`
		Items `json:"items"`
		NumOfRows int `json:"numOfRows"`
		PageNo int `json:"pageNo"`
		TotalCount int `json:"totalCount"`
	  }

	  type Header struct {
		ResultCode string `json:"resultCode"`
		ResultMsg string `json:"resultMsg"`
	  }

	  type Response struct {
		Header `json:"header"`
		Body `json:"body"`
	  }

	  type Weather struct {
		Response `json:"response"`
	  }
	
//환경변수를 불러오는 함수-> 코드내에 변수 만들어서 불러옴 
//메인 을 기준으로 생각해서 밖으로 나가면됨. 
	err := godotenv.Load("../.env")

	if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	// fmt.Println("env [serviceKey]:", os.Getenv("serviceKey"))
	// 1. sprintf 함수를 이용해서 serviceKey 합치기
	// 2. 합친걸 url변수에 담고
	// 3. http.Get에 유알엘 변수 넣기
	//변수랑 문자랑 합쳐주는 함수 ->반환값은 스트링
	
	resp, err := http.Get(GetUrl())
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", string(data))

	var w Weather
    err = json.Unmarshal(data, &w)
    if err != nil {
        panic(err)
    }
	c.JSON(http.StatusOK, w)
}

