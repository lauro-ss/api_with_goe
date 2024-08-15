package service

import "github.com/lauro-ss/api_with_goe/internal/data"

type UserRepository interface {
	Create(data.User) (uint, error)
	Update(data.User) (uint, error)
	Get(uint) (*data.User, error)
	List() ([]data.User, error)
	Delete(uint) (bool, error)
}
