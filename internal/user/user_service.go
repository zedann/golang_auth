package user

import (
	"context"
	"time"

	"github.com/golang_auth/util"
)

type userService struct {
	UserRepository
	timeout time.Duration
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
		time.Duration(2) * time.Second,
	}
}

func (svc *userService) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, svc.timeout)
	defer cancel()

	// hash password

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err := svc.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}
