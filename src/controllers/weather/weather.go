package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

//c에 클라이언트가 요청한 정보가 다 담겨있고 함수안에서 제서
//라우터를 만들면 자동으로 파라미터를 넘겨받게 설계되어있음.
func GetWeather(c *gin.Context) {  
	  type Header struct {
		ResultCode string `json:"resultCode"`
		// ResultMsg string 
	  }

	  type Response struct {
		Header `json:"header"`
		// Body
	  }

	  //임베딩 
	  type weather struct {
		Response `json:"response"`
	  }


 //API를 받아오기 위해 총 3단계 필요

	//resp : API를 써서 결과를 받는것
	resp, err := http.Get("http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getVilageFcst?serviceKey=FkrjO6uyl2g8GMNLJMM5IMx5DOwDib0Zz8Ol4gsNyIzf4m1y9AeNSfQnSAbUZ%2Btpu%2FBxxz%2BMimxAkXvjgRg68w%3D%3D&numOfRows=10&pageNo=1&base_date=20240416&base_time=0500&nx=55&ny=127&dataType=JSON")

    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()

    // DATA : 받은 결과가 BODY로 넘어오니까 BODY에서 데이터를 읽는 작업
	
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", string(data))
 //읽어온 BODY 구조체에 넣기 
	var w weather
    err = json.Unmarshal(data, &w)
    if err != nil {
        panic(err)
    }
	c.JSON(http.StatusOK, w)
}

