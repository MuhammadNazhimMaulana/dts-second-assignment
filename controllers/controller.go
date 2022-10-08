package controllers

import (
	"Assignment_2/apimodels"
	"Assignment_2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var req apimodels.Request
	var res apimodels.Response

	// Check Error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := services.SaveOrder(req)

	// Jika gagal
	if err != nil {
		res.Status = "Create Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func GetOrderAll(ctx *gin.Context) {
	var res apimodels.ResponseGet

	res, err := services.AllOrder()

	// Jika gagal
	if err != nil {
		res.Status = "Get Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func UpdateOrder(ctx *gin.Context) {

}

func DeleteOrder(ctx *gin.Context) {

}
