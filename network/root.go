package network

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"io"
	"junction/config"
	"junction/db"
	"log"
	"net/http"
)

type Network struct {
	engine *gin.Engine
}

type MemberInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MemberResponse struct {
	Message    string
	MemberInfo *MemberInfo
}

func (u MemberResponse) GetMessage() string {
	return u.Message
}

func (u MemberResponse) GetResult() interface{} {
	return u.MemberInfo
}

type WeatherInfo struct {
	Weather   string `json:"weather"`
	Condition string `json:"condition"`
	Gap       int    `json:"gap"`
}

type WeatherResponse struct {
	Message     string
	WeatherInfo *WeatherInfo
}

func (w WeatherResponse) GetMessage() string {
	return w.Message
}

func (w WeatherResponse) GetResult() interface{} {
	return w.WeatherInfo
}

type CoordinateResponse struct {
	Message        string
	CoordinateInfo []*Coordinate
}

func (c CoordinateResponse) GetMessage() string {
	return c.Message
}

func (c CoordinateResponse) GetResult() interface{} {
	return c.CoordinateInfo
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

func NewNetwork(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			network := &Network{
				engine: r,
			}

			network.RegisterRoutes()
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			return nil
		},
	})
	return r
}

// TODO: 이후에 제거가 필요
func (s *Network) HealthCheck() {
	s.engine.GET("/ping", func(c *gin.Context) {
		mockData := &MemberInfo{
			Name: "이설희",
			Age:  20,
		}

		response := MemberResponse{
			Message:    "사용자 조회 성공",
			MemberInfo: mockData,
		}

		s.OkResponse(c, response)
	})
}

func (s *Network) GetWeatherInfo() {
	s.engine.GET("/weatherInfo", func(c *gin.Context) {
		result := &WeatherInfo{
			Weather:   "폭염",
			Condition: "cloudy",
			Gap:       33 - 32, // 오늘 기온과 어제 기온의 차이, TODO: 이후에 API로 기온 통신 필요
		}

		response := WeatherResponse{
			Message:     "날씨 조회 성공",
			WeatherInfo: result,
		}

		s.OkResponse(c, response)
	})
}

func (s *Network) GetBestRoutes() {
	s.engine.GET("/findPath", func(c *gin.Context) {
		conf := flag.String("findPathConf", "../config/config.toml", "config file not found")
		bestConf := config.NewConfig(*conf)
		client, err := db.NewDb(*conf)

		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		ctx := context.Background()

		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		baseUrl := bestConf.FindPathApi.Url
		longitude := c.Query("current_longitude")
		latitude := c.Query("current_latitude")
		walkingDuration := c.Query("walking_duration_sec")
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")

		requestUrl := fmt.Sprintf("%s?current_longitude=%s&current_latitude=%s&walking_duration_sec=%s&start_date=%s&end_date=%s",
			baseUrl, longitude, latitude, walkingDuration, startDate, endDate)

		resp, err := http.Get(requestUrl)

		if err != nil {
			log.Fatalf("GET 요청 실패: %v", err)
		}
		defer resp.Body.Close()

		// 응답 본문 읽기
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("서버에러: %s", err.Error())
		}

		// 응답 출력
		responseBody := string(body)

		var coordinates []Coordinate
		err = json.Unmarshal([]byte(responseBody), &coordinates)

		if err != nil {
			fmt.Errorf("error 발생~~!!: ", err.Error())
		}

		client.JupgingLog.Create().SetStartDate(startDate).SetEndDate(endDate).SetMemberID(1).SetLog(string(body))
		response := CoordinateResponse{
			Message:        "경로 조회 성공",
			CoordinateInfo: convertToPointerSlice(coordinates),
		}

		s.OkResponse(c, response)
	})
}

func (s *Network) RegisterRoutes() {
	s.HealthCheck()
	s.GetWeatherInfo()
	s.GetBestRoutes()
}

func convertToPointerSlice(coords []Coordinate) []*Coordinate {
	var result []*Coordinate
	for i := range coords {
		result = append(result, &coords[i])
	}
	return result
}
