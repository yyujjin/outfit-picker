package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetUrl()string {
	getTime := time.Now() //현재날짜 가져와서 변수에 저장
	serviceKey := os.Getenv("serviceKey")
	numOfRows := 200
	pageNo := 1
	base_date := getTime.Format("20060102")//fotmat 반환값이 스트링이라 바로 변수에 저장 
	base_time := "0200" 
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
	// 하늘상태(SKY) 코드 : 맑음(1), 구름많음(3), 흐림(4)
	fmt.Println("하늘상태",w.Item[5].FcstValue)
	// 강수형태(PTY) 코드 : 없음(0), 비(1), 비/눈(2), 눈(3), 소나기(4) 
	fmt.Println("강수",w.Item[6].FcstValue)
	//TMN 일 최저기온 
	fmt.Println("최저기온",w.Item[48].FcstValue)
	//TMX 일 최고기온
	fmt.Println("최고기온",w.Item[157].FcstValue)
	
	c.JSON(http.StatusOK, gin.H {
		"하늘 상태" : w.Item[5].FcstValue,
		"강수 형태" : w.Item[6].FcstValue,
		"최저 기온" : w.Item[48].FcstValue,
		"최고 기온" : w.Item[157].FcstValue,
	})


	//최저기온, 최고기온 찾는 코드 
	// var tmni,tmxi int
	// for i := range w.Item {
	// 	if w.Item[i].Category=="TMN"{
	// 		tmni = i
	// 	}
	// 	if w.Item[i].Category=="TMX" {
	// 		tmxi = i
	// 	}	
	// }
	// fmt.Printf("tmni : %d , tmxi : %d",tmni,tmxi)
	
}

