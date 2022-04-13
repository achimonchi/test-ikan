package web

type CreateAuthResponse struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
