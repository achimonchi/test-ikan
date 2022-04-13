package services

import (
	"auth/server/params"
	"auth/server/views/web"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthServices) LoginByPhone(req *params.Login) (*web.LoginResponse, error) {
	isExist, err := s.repo.FindByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(isExist.Password), []byte(req.Password))
	if err != nil {
		return nil, sql.ErrNoRows
	}

	token, err := s.token.CreateToken(isExist.Name, isExist.Phone, isExist.Role)
	if err != nil {
		return nil, err
	}

	var res = web.LoginResponse{
		Token: token,
	}

	return &res, nil
}
