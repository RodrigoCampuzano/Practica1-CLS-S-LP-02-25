package main

import (
    "API/src/application"
    "API/src/infraestructure/controllers"
    "API/src/infraestructure/db"
    infraRepo "API/src/infraestructure/repositories"
    "API/src/infraestructure/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa la base de datos
    database := db.ConnectionDB()
    defer database.Close()

    // Crea los repos
    userRepo := infraRepo.NewUserRepository(database)
    orderRepo := infraRepo.NewOrderRepository(database)

    // crea los casos de uso para usuarios
    createUser := application.NewCreateUser(userRepo)
    getAllUsers := application.NewGetUser(userRepo)
    getUserByID := application.NewGetUserByID(userRepo)
    updateUser := application.NewUpdateUser(userRepo)
    deleteUserByID := application.NewDeleteUserByID(userRepo)
    deleteAllUsers := application.NewDeleteAllUser(userRepo)

    // crea los casos de uso para órdenes
    createOrder := application.NewCreateOrder(orderRepo)
    getAllOrders := application.NewGetAllOrders(orderRepo)
    getOrderByID := application.NewGetOrderByID(orderRepo)
    updateOrder := application.NewUpdateOrder(orderRepo)
    deleteOrderByID := application.NewDeleteOrderByID(orderRepo)
    deleteAllOrders := application.NewDeleteAllOrders(orderRepo)

    // se ccrean los controladores
    userController := controllers.NewUserController(createUser, getAllUsers, getUserByID, updateUser, deleteUserByID, deleteAllUsers)
    orderController := controllers.NewOrderController(createOrder, getAllOrders, getOrderByID, updateOrder, deleteOrderByID, deleteAllOrders)

    // se inizializa Gin
    router := gin.Default()

    // se configuran las rutas
    routes.SetupRoutes(router, userController)
    routes.SetupOrderRoutes(router, orderController)

    // Ejecutar el servidor principal en el puerto 8080
    go func() {
        router.Run(":8080")
    }()

    // Ejecutar short polling y long polling en goroutines en un puerto diferente
    go func() {
        pollingRouter := gin.Default()
        pollingRouter.GET("/users/shortpoll", userController.ShortPollUsers)
        pollingRouter.GET("/users/longpoll", userController.LongPollUsers)
        pollingRouter.GET("/orders/shortpoll", orderController.ShortPollOrders)
        pollingRouter.GET("/orders/longpoll", orderController.LongPollOrders)
        pollingRouter.Run(":8081")
    }()
    select {}
}