package utils

import (
	"errors"
	"fetch/config"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Token interface {
	CreateToken(name, phone, role string) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type JWT struct {
	config    *config.Config
	secretKey string
	expiredAt time.Duration
}

type Payload struct {
	jwt.StandardClaims
	ID        uuid.UUID `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var minJwtKeyLength = 6

func NewToken(config *config.Config) Token {
	if len(config.JWT_KEY) < minJwtKeyLength {
		panic(fmt.Errorf("jwt secret key is less than %d character ", minJwtKeyLength))
	}

	jwtExpiry, err := strconv.Atoi(config.JWT_EXPIRY)
	if err != nil {
		panic(err)
	}
	return &JWT{
		secretKey: config.JWT_KEY,
		expiredAt: time.Duration(jwtExpiry),
		config:    config,
	}
}

func newPayload(name, phone, role string, expiredAt time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	payload := Payload{
		ID:        tokenId,
		Name:      name,
		Phone:     phone,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * expiredAt),
	}

	return &payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("invalid token")
	}
	return nil
}

func (t *JWT) CreateToken(name, phone, role string) (string, error) {
	payload, err := newPayload(name, phone, role, t.expiredAt)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(t.secretKey))
}

func (t *JWT) VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)
	if err != nil {
		if err.Error() == jwt.ErrSignatureInvalid.Error() {
			return nil, err
		}
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("token gagal")
	}
	return payload, nil
}
