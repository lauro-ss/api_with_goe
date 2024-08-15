package handlers

import (
	"net/http"

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
