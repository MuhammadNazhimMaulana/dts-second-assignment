package routers

import (
	"Assignment_2/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// Access Point
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrderAll)
	router.GET("/orders/:orderID", controllers.GetOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)
	return router
}
