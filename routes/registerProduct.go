package routes

import (
	"net/http"
	"price-tracker/parsers"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Product string `json:"product"`
}

func RegisterProduct(context *gin.Context) {
	var body RequestBody

	context.BindJSON(&body)

	parsers.RegisterProductByName(body.Product)

	context.JSON(http.StatusOK, gin.H{"product": body.Product})
}
