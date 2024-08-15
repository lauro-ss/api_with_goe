package handlers

import (
	"net/http"
	"time"

	"github.com/lauro-ss/api_with_goe/internal/service"
	"github.com/lauro-ss/api_with_goe/internal/utils"
)

func UserList(ur service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ur.List()
		if err != nil {
			utils.HttpStatus500(w)
		}
		utils.AsJson(w, users)
	}
}

func UserCreate(ur service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UserCreate"))
	}
}

func UserGet(ur service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.Write([]byte("UserGet"))
	}
}

func UserUpdate(ur service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UserUpdate"))
	}
}

func UserDelete(ur service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UserDelete"))
	}
}
