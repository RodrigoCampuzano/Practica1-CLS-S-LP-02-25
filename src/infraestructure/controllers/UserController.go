package controllers

import (
    "API/src/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type UserController struct {
    createUser     *application.CreateUser
    getAllUsers    *application.GetUsers
    getUserByID    *application.GetUserByID
    updateUser     *application.UpdateUser
    deleteUserByID *application.DeleteUser
    deleteAllUsers *application.DeleteAllUser
}

func NewUserController(
    createUser *application.CreateUser,
    getAllUsers *application.GetUsers,
    getUserByID *application.GetUserByID,
    updateUser *application.UpdateUser,
    deleteUserByID *application.DeleteUser,
    deleteAllUsers *application.DeleteAllUser,
) *UserController {
    return &UserController{
        createUser:     createUser,
        getAllUsers:    getAllUsers,
        getUserByID:    getUserByID,
        updateUser:     updateUser,
        deleteUserByID: deleteUserByID,
        deleteAllUsers: deleteAllUsers,
    }
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var user struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := uc.createUser.Execute(user.Name, user.Email); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusCreated)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.getAllUsers.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    user, err := uc.getUserByID.Execute(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var user struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := uc.updateUser.Execute(int32(id), user.Name, user.Email); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := uc.deleteUserByID.Execute(int32(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}

func (uc *UserController) DeleteAllUsers(c *gin.Context) {
    if err := uc.deleteAllUsers.Execute(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}