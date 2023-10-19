package routers

import (
	middlewares "gin-gonic-sample/middlewares"
	"gin-gonic-sample/routers/api/auth"
	v1 "gin-gonic-sample/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func AddDemoRouter() *gin.Engine {
	// router := gin.Default()
	router := gin.New()
	router.Use(middlewares.ErrorHanler)

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Hello)
		apiv1.POST("/consume", v1.ConsumeBalanceNoti)
	}

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/register", auth.Register)
	}

	return router
}
