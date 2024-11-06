package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-practice/config"
)

func NewServer(router *gin.Engine) *http.Server {

	return &http.Server{
		Addr:         ":" + config.Config.GetString("server.port"),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Config.GetInt64("server.read_timeout")) * time.Millisecond,
		WriteTimeout: time.Duration(config.Config.GetInt64("server.write_timeout")) * time.Millisecond,
	}
}
