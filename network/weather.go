package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"junction/config"
	"junction/internal/pkg/weather"
	"log"
	"net/http"
)

type WeatherNetwork struct {
	engine *gin.Engine
	config *config.Config

	WeatherSvc *weather.WeatherService
}

func RegisterWeatherRoutes(r *gin.Engine) {
	r.GET("/weatherInfo", GetWeatherInfo)
}

// 쿼리 파라미터 요청 받는 처리가 필요하다.
func GetWeatherInfo(c *gin.Context) {
	//configPath := flag.String("config", ".././config/config.toml", "config file not found")
	api := &WeatherNetwork{
		config: config.NewConfig(*config.ConfigPath),
	}

	// url 가져오기
	url := api.config.WeatherApi.Url
	airportCode := api.config.WeatherApi.AirPortCode

	// 쿼리 파라미터 입력 받기
	baseDate := c.Query("baseDate")
	baseTime := c.Query("baseTime")

	requestUrl := fmt.Sprintf("%s&base_date=%s&base_time=%d&airPortCd=%s", url, baseDate, baseTime, airportCode)

	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Fatalf("GET 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("응답 읽기 실패: %v", err)
	}

	// 응답 출력
	fmt.Println(string(body))
}
