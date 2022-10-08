package controllers

import (
	"Assignment_2/requests"
	"Assignment_2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var req requests.Request
	var res requests.Response

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
	var res requests.ResponseGet

	res, err := services.AllOrder()

	// Jika gagal
	if err != nil {
		res.Status = "Get All Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func GetOrder(ctx *gin.Context) {
	var res requests.ResponseGetOne

	// Getting id
	id, errstr := strconv.Atoi(ctx.Param("orderID"))

	if errstr != nil {
		res.Status = "Get Order Gagal"
		res.ResponseCode = "400"
	}

	res, err := services.Order(id)

	// Jika gagal
	if err != nil {
		res.Status = "Get Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func UpdateOrder(ctx *gin.Context) {
	var req requests.Request
	var res requests.Response

	// Getting id
	id, errstr := strconv.Atoi(ctx.Param("orderID"))

	if errstr != nil {
		res.Status = "Update Order Gagal"
		res.ResponseCode = "400"
	}

	// Check Error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := services.UpdateOrder(req, id)

	// Jika gagal
	if err != nil {
		res.Status = "Update Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func DeleteOrder(ctx *gin.Context) {
	var res requests.ResponseGetOne

	// Deleting id
	id, errstr := strconv.Atoi(ctx.Param("orderID"))

	if errstr != nil {
		res.Status = "Delete Order Gagal"
		res.ResponseCode = "400"
	}

	res, err := services.Delete(id)

	// Jika gagal
	if err != nil {
		res.Status = "Delete Order Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}
