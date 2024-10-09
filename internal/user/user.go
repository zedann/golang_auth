package user

import "context"

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

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}
type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
}

type UserHandler interface {
	UserService
}
