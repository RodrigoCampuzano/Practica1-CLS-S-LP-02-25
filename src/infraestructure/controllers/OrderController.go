package controllers

import (
    "API/src/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type OrderController struct {
    createOrder     *application.CreateOrder
    getAllOrders    *application.GetAllOrders
    getOrderByID    *application.GetOrderByID
    updateOrder     *application.UpdateOrder
    deleteOrderByID *application.DeleteOrder
    deleteAllOrders *application.DeleteAllOrders
}

func NewOrderController(
    createOrder *application.CreateOrder,
    getAllOrders *application.GetAllOrders,
    getOrderByID *application.GetOrderByID,
    updateOrder *application.UpdateOrder,
    deleteOrderByID *application.DeleteOrder,
    deleteAllOrders *application.DeleteAllOrders,
) *OrderController {
    return &OrderController{
        createOrder:     createOrder,
        getAllOrders:    getAllOrders,
        getOrderByID:    getOrderByID,
        updateOrder:     updateOrder,
        deleteOrderByID: deleteOrderByID,
        deleteAllOrders: deleteAllOrders,
    }
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
    var order struct {
        UserID int32   `json:"user_id"`
        Amount float32 `json:"amount"`
        Status string  `json:"status"`
    }
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := oc.createOrder.Execute(order.UserID, order.Amount, order.Status); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusCreated)
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
    orders, err := oc.getAllOrders.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) GetOrderByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    order, err := oc.getOrderByID.Execute(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if order == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var order struct {
        UserID int32   `json:"user_id"`
        Amount float32 `json:"amount"`
        Status string  `json:"status"`
    }
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := oc.updateOrder.Execute(int32(id), order.UserID, order.Amount, order.Status); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := oc.deleteOrderByID.Execute(int32(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}

func (oc *OrderController) DeleteAllOrders(c *gin.Context) {
    if err := oc.deleteAllOrders.Execute(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}