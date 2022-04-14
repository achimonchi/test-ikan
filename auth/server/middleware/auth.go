package middleware

import (
	"auth/constants"
	"auth/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Auth struct {
	token utils.Token
}

func NewAuthMiddleware(token utils.Token) *Auth {
	return &Auth{
		token: token,
	}
}

func (a *Auth) Auth(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		bearerTokenHeader := r.Header.Get("Authorization")
		tokenHeader := strings.Split(bearerTokenHeader, "Bearer ")

		if len(tokenHeader) != 2 {
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"status": 401,
				"error":  "UNAUTHORIZED",
			})
			return
		}

		_, err := a.token.VerifyToken(tokenHeader[1])
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"status": 401,
				"error":  "UNAUTHORIZED",
			})
			return
		}

		ctx := context.WithValue(r.Context(), constants.TOKEN, tokenHeader[1])
		r = r.WithContext(ctx)

		next(rw, r, p)
	}
}
