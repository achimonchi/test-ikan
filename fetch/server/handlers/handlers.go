package handlers

type Handlers struct {
	PingHandlers  PingHandlers
	FetchHandlers FetchHandlers
}

func NewHandlers(ping PingHandlers, fetch FetchHandlers) *Handlers {
	return &Handlers{
		PingHandlers:  ping,
		FetchHandlers: fetch,
	}
}
