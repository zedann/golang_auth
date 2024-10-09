package user

import "github.com/gin-gonic/gin"

type userHandler struct {
	UserService
}

func NewHandler(svc UserService) UserHandler {
	return &userHandler{
		UserService: svc,
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {

}
