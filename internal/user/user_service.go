package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang_auth/util"
)

const (
	secretKey = "secretasf123124215215reasdfweas412"
)

func CreateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	return token.SignedString([]byte(secretKey))
}

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

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (svc *userService) Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, svc.timeout)
	defer cancel()

	user, err := svc.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return nil, fmt.Errorf("wronge email or password")
	}

	ss, err := CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginUserRes{
		accessToken: ss,
		ID:          user.ID,
		Username:    user.Username,
	}, nil

}
