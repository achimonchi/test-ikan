package middleware

import (
	"auth/constants"
	"auth/pkg/utils"

	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Trace struct{}

func NewTraceMiddleware() *Trace {
	return &Trace{}
}

func (t *Trace) Trace(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		traceId := utils.GenerateUUID()

		ctx := context.WithValue(r.Context(), constants.TRACE_ID, traceId)

		r = r.WithContext(ctx)
		rw.Header().Add("Trace-Id", traceId.String())
		next(rw, r, p)
	}
}
