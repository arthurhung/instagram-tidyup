package app

import (
	"github.com/gin-gonic/gin"

	"github.com/arthurhung/instagram-tidyup/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	status := true
	if errCode != 200 {
		status = false
	}
	g.C.JSON(httpCode, Response{
		Status: status,
		Msg:    e.GetMsg(errCode),
		Data:   data,
	})
	return
}
