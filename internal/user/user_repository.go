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
		return nil, err
	}

	user.ID = lastInsertedID

	return user, nil

}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	u := User{}
	query := "SELECT id , email , username , password FROM users WHERE email=$1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	return &u, nil

}
