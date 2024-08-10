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

func (s *Network) RegisterRoutes() {
	s.HealthCheck()
	RegisterWeatherRoutes(s.engine)
}
