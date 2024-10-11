package user

import (
	"context"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type CreateUserRes struct {
	ID       int64  `json:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type LoginUserRes struct {
	accessToken string
	ID          int64  `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
