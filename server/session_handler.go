package server

import (
	"context"
	"net/http"
	"strings"
)

type SessionHandler struct {
	server  *Server
	ctx     context.Context
	cancel  context.CancelFunc
	counter *counter
}

func (sh *SessionHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.String(), "ws") {
		sh.server.generateHandleWS(sh.ctx, sh.cancel, sh.counter)(rw, r)
	} else {
		sh.server.handleNewSession(rw, r)
	}
}

func NewSessionHandler(server *Server, ctx context.Context, cancel context.CancelFunc, counter *counter) *SessionHandler {
	return &SessionHandler{server, ctx, cancel, counter}
}
