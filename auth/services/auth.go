package services

import "auth/repositories"

type AuthServices struct {
	repo repositories.AuthRepo
}

func NewAuthServices(repo repositories.AuthRepo) *AuthServices {
	return &AuthServices{
		repo: repo,
	}
}
