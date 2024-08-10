package network

import "github.com/gin-gonic/gin"

type CommonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 서비스별로 반환하는 응답 메시지와 결과를 다르게 하기 위함.
type DataResponse interface {
	GetMessage() string
	GetResult() interface{}
}

// register 유틸 함수들
func (n *Network) RegisterGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) RegisterPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) RegisterUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) RegisterDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

// Response 형태를 맞추기 위한 util 함수 목록
func (n *Network) OkResponse(c *gin.Context, data DataResponse) {
	c.JSON(200, CommonResponse{
		Status:  200,
		Message: data.GetMessage(),
		Data:    data.GetResult(),
	})
}

func (n *Network) CreatedResponse(c *gin.Context, data DataResponse) {
	c.JSON(201, CommonResponse{
		Status:  201,
		Message: data.GetMessage(),
		Data:    data.GetResult(),
	})
}

func (n *Network) FailResponse(c *gin.Context, data DataResponse) {
	c.JSON(400, CommonResponse{
		Status:  400,
		Message: data.GetMessage(),
		Data:    data.GetResult(),
	})
}
