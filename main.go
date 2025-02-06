package main

import (
    orderApp "API/src/Orders/application"
    orderCtrl "API/src/Orders/infraestructure/controllers"
    infraRepoOrder "API/src/Orders/infraestructure/repositories"
    orderRoutes "API/src/Orders/infraestructure/routes"

    userApp "API/src/Users/application"
    userCtrl "API/src/Users/infraestructure/controllers"
    infraRepoUser "API/src/Users/infraestructure/repositories"
    userRoutes "API/src/Users/infraestructure/routes"

    "API/src/core/db"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa la base de datos
    database := db.ConnectionDB()
    defer database.Close()

    // Crea los repos
    userRepo := infraRepoUser.NewUserRepository(database)
    orderRepo := infraRepoOrder.NewOrderRepository(database)

    // crea los casos de uso para usuarios
    createUser := userApp.NewCreateUser(userRepo)
    getAllUsers := userApp.NewGetUser(userRepo)
    getUserByID := userApp.NewGetUserByID(userRepo)
    updateUser := userApp.NewUpdateUser(userRepo)
    deleteUserByID := userApp.NewDeleteUserByID(userRepo)
    deleteAllUsers := userApp.NewDeleteAllUser(userRepo)

    // crea los casos de uso para órdenes
    createOrder := orderApp.NewCreateOrder(orderRepo)
    getAllOrders := orderApp.NewGetAllOrders(orderRepo)
    getOrderByID := orderApp.NewGetOrderByID(orderRepo)
    updateOrder := orderApp.NewUpdateOrder(orderRepo)
    deleteOrderByID := orderApp.NewDeleteOrderByID(orderRepo)
    deleteAllOrders := orderApp.NewDeleteAllOrders(orderRepo)

    // se ccrean los controladores
    userController := userCtrl.NewUserController(createUser, getAllUsers, getUserByID, updateUser, deleteUserByID, deleteAllUsers)
    orderController := orderCtrl.NewOrderController(createOrder, getAllOrders, getOrderByID, updateOrder, deleteOrderByID, deleteAllOrders)

    // se inizializa Gin
    router := gin.Default()

    // se configuran las rutas
    userRoutes.SetupRoutes(router, userController)
    orderRoutes.SetupOrderRoutes(router, orderController)

    // Ejecutar el servidor principal en el puerto 8080
    go func() {
        router.Run(":8080")
    }()

    // Ejecutar short polling y long polling en un puerto diferente
    go func() {
        pollingRouter := gin.Default()
        pollingRouter.GET("/users/shortpoll", userController.ShortPollUsers)
        pollingRouter.GET("/users/longpoll", userController.LongPollUsers)
        pollingRouter.GET("/orders/shortpoll", orderController.ShortPollOrders)
        pollingRouter.GET("/orders/longpoll", orderController.LongPollOrders)
        pollingRouter.Run(":8081")
    }()

    // Mantener el programa principal en ejecución
    select {}
}