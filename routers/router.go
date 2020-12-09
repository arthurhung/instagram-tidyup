package routers

import (
	"net/http"

	"github.com/arthurhung/instagram-tidyup/middleware/header_handler"
	v1 "github.com/arthurhung/instagram-tidyup/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(header_handler.HeaderHandlerMiddleware())

	r.LoadHTMLGlob("template/*")
	{
		//Statics HTML
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	apiv1 := r.Group("/api/v1")

	{
		apiv1.POST("/loginIG", v1.LoginIG)

	}

	return r
}
