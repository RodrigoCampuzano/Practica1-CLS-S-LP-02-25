package controllers

import (
	"API/src/Users/application"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func (uc *UserController) ShortPollUsers(c *gin.Context) {
	initialCount, err := uc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	initialLen := len(initialCount)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-ticker.C:
			currentCount, err := uc.getAllUsers.Execute()
			if err != nil {
				c.SSEvent("error", gin.H{"error": err.Error()})
				return false
			}
			currentLen := len(currentCount)
			if currentLen != initialLen {
				c.SSEvent("message", gin.H{"mensaje": "El número de usuarios ha cambiado", "de": initialLen, "a": currentLen})
				initialLen = currentLen
			} else {
				c.SSEvent("message", gin.H{"mensaje": "No hay cambios", "cantidad": currentLen})
			}
			return true
		}
	})
}

func (uc *UserController) LongPollUsers(c *gin.Context) {
	initialCount, err := uc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	initialLen := len(initialCount)

	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusOK, gin.H{"mensaje": "No se detectaron cambios"})
			return
		default:
			currentCount, err := uc.getAllUsers.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			currentLen := len(currentCount)
			if currentLen != initialLen {
				c.JSON(http.StatusOK, gin.H{"mensaje": "El número de usuarios ha cambiado", "de": initialLen, "a": currentLen})
				return
			}
			time.Sleep(5 * time.Second)
		}
	}
}
