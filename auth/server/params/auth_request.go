package params

type CreateAuth struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string // auto generate with 4 characters
}

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
