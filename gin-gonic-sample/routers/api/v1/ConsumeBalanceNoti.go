package v1

import (
	"gin-gonic-sample/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumeBalanceNoti(ctx *gin.Context) {
	var request = new(models.ConsumeBalanceRequest)
	log.Printf("request: %s", ctx.Request.RemoteAddr)
	var response = new(models.ConsumeBalanceResponse)

	if err := ctx.BindJSON(&request); err != nil {
		response.Error = "400"
		response.ErrorMsg = "parse request body faild"
		log.Println("Parse Data failed:", err)
		ctx.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
