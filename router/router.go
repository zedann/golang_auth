package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang_auth/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler user.UserHandler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func Start(addr string) error {
	if r == nil {
		return fmt.Errorf("router not initialized")
	}
	return r.Run(addr)
}
