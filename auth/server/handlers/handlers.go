package handlers

type Handlers struct {
	Ping PingHandler
	Auth AuthHandler
}

func NewHandlers(ping PingHandler, auth AuthHandler) *Handlers {
	return &Handlers{
		Ping: ping,
		Auth: auth,
	}
}
