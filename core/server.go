package core

import (
	"blog-backend/global"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	address := fmt.Sprintf(":%s", global.YAGAMI_CONFIG.App.Port)
	s := InitServer(address, r)

	time.Sleep(10 & time.Millisecond)

	fmt.Println("监听端口", address)

	s.ListenAndServe()

}
