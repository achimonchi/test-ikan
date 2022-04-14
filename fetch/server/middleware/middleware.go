package middleware

type Middleware struct {
	Trace *Trace
	Auth  *Auth
}

func NewMiddleware(Trace *Trace, Auth *Auth) *Middleware {
	return &Middleware{
		Trace: Trace,
		Auth:  Auth,
	}
}
