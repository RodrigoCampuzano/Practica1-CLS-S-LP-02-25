package routes

import (
    "API/src/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
    router.POST("/users", userController.CreateUser)
    router.GET("/users/all", userController.GetAllUsers)
    router.GET("/users/get/:id", userController.GetUserByID)
    router.PUT("/users/update/:id", userController.UpdateUser)
    router.DELETE("/users/delete/:id", userController.DeleteUser)
    router.DELETE("/users/deleteall", userController.DeleteAllUsers)
}