package auth

import (
	"gin-gonic-sample/helpers"
	"gin-gonic-sample/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type AuthController interface {
// 	Register(ctx *gin.Context)
// }

func Register(ctx *gin.Context) {
	var registerRequest models.RegisterRequest
	errRequest := ctx.ShouldBind(&registerRequest)
	log.Println(errRequest)

	response := helpers.BuildErrorResponse("Failed to process request", errRequest.Error(), helpers.EmptyObj{})

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
