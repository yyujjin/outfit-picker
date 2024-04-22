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

func GetUrl()string {
	serviceKey := os.Getenv("serviceKey")
	numOfRows := 10
	pageNo := 1
	base_date := "20240421" 
	base_time := "0500" 
	nx := 55
	ny := 127
	dataType := "JSON"
	
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
	
	err := godotenv.Load("../.env")

	if err != nil {
        log.Fatal("Error loading .env file")
    }

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

