package services

import (
	"auth/server/params"
	"auth/server/views/web"
	"context"
	"fmt"

	"github.com/thanhpk/randstr"
)

func (s *AuthServices) CreateAuth(ctx context.Context, req *params.CreateAuth) (*web.CreateAuthResponse, error) {
	req.Password = randstr.String(4)

	auth := makeModelFromParamCreate(req)

	isExist, _ := s.repo.FindByPhone(auth.Phone)
	if isExist != nil {
		return nil, fmt.Errorf("user with phone - %s already registered", auth.Phone)
	}
	err := s.repo.Registry(auth)
	if err != nil {
		return nil, err
	}

	var response = web.CreateAuthResponse{
		Password: req.Password,
	}
	return &response, nil
}
