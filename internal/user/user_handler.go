package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UserService
}

func NewUserHandler(svc UserService) UserHandler {
	return &userHandler{
		UserService: svc,
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := h.UserService.CreateUser(c.Request.Context(), &u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
