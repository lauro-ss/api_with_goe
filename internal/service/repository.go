package service

import "github.com/lauro-ss/api_with_goe/internal/data"

type UserRepository interface {
	Create() (uint, error)
	Update() (uint, error)
	Get() (data.User, error)
	List() ([]data.User, error)
	Delete() (bool, error)
}
