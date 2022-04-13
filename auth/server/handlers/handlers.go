package handlers

type Handlers struct {
	Ping PingHandler
}

func NewHandlers(ping PingHandler) *Handlers {
	return &Handlers{
		Ping: ping,
	}
}
