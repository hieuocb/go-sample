package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var response = ConsumeBalanceResponse{

// }

func ConsumeBalanceNoti(ctx *gin.Context) {
	response := ConsumeBalanceRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
