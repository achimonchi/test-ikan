package handlers

type Handlers struct {
	PingHandlers PingHandlers
}

func NewHandlers(ping PingHandlers) *Handlers {
	return &Handlers{
		PingHandlers: ping,
	}
}
