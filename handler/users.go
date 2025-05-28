package handlers

import (
	"github.com/gin-gonic/gin"
	"go-chain-kit/models"
	"go-chain-kit/services"
	"net/http"
)

type UserHandler struct {
	userService *services.Users
}

func NewUserHandler(usService *services.Users) UserHandler {
	return UserHandler{
		userService: usService,
	}
}

func (u UserHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var us models.UserReq
		if err := c.ShouldBindJSON(&us); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		id, err := u.userService.Create(&us)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"userID": id})
	}
}
