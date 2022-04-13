package services

import (
	"auth/pkg/utils"
	"auth/repositories"
)

type AuthServices struct {
	repo  repositories.AuthRepo
	token utils.Token
}

func NewAuthServices(repo repositories.AuthRepo, token utils.Token) *AuthServices {
	return &AuthServices{
		repo:  repo,
		token: token,
	}
}
