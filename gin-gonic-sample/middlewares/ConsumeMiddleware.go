package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHanler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		log.Println(err)
	}
	c.JSON(http.StatusInternalServerError, "")
}
