package network

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
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

func (s *Network) RegisterRoutes() {
	s.HealthCheck()
	s.GetWeatherInfo()
}
