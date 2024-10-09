package user

import (
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *User) (*User, error) {

	var lastInsertedID int64
	query := "INSERT INTO users (username , password , email) VALUES ($1,$2,$3) returning id;"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertedID)

	if err != nil {
		return &User{}, err
	}

	user.ID = lastInsertedID

	return user, nil

}
