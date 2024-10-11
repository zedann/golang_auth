package user

import (
	"fmt"
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

func (h *userHandler) Login(c *gin.Context) {

	var u LoginUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Login(c.Request.Context(), &u)

	fmt.Println(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", user.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
