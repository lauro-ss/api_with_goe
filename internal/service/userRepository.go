package service

import "github.com/lauro-ss/api_with_goe/internal/data"

type userRepository struct {
	db *data.Database
}

func NewUserRepository(db *data.Database) *userRepository {
	return &userRepository{db: db}
}
