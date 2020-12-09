package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurhung/instagram-tidyup/models"
	"github.com/arthurhung/instagram-tidyup/pkg/logging"
	"github.com/arthurhung/instagram-tidyup/pkg/setting"
	"github.com/arthurhung/instagram-tidyup/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	logging.Setup()
	models.Setup()

}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
