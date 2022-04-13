package services

import (
	"auth/models"
	"auth/pkg/utils"
	"auth/server/params"

	"golang.org/x/crypto/bcrypt"
)

func makeModelFromParamCreate(req *params.CreateAuth) *models.Auth {
	newPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	return &models.Auth{
		ID:       utils.GenerateUUID().String(),
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: string(newPass),
	}
}
