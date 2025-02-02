package routes

import (
	"API/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderController *controllers.OrderController) {
    router.POST("/orders",orderController.CreateOrder) 
    router.GET("/orders/all", orderController.GetAllOrders)
    router.GET("/orders/get/:id", orderController.GetOrderByID)
    router.PUT("/orders/update/:id", orderController.UpdateOrder)
    router.DELETE("/orders/delete/:id", orderController.DeleteOrder)
    router.DELETE("/orders/deleteall", orderController.DeleteAllOrders)
}