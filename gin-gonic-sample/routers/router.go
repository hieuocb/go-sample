package routers

import (
	v1 "gin-gonic-sample/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func AddDemoRouter() *gin.Engine {
	router := gin.Default()

	apiv1 := router.Group("/api/v1")

	apiv1.GET("/ping", v1.Hello)

	return router
}
