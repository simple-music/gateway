package rest

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	BearerLength = len("Bearer: ")
)

func (srv *Service) WithLogs(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h(ctx)
		srv.logger.Info(fmt.Sprintf("%s %s %d",
			string(ctx.Method()), string(ctx.RequestURI()), ctx.Response.StatusCode()),
		)
	}
}

func (srv *Service) WithAuth(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))
		if len(token) <= BearerLength {
			srv.WriteError(ctx, srv.authTokenErr)
			return
		}

		userID, err := srv.tokenManager.ParseToken(token[BearerLength:])
		if err != nil {
			srv.WriteError(ctx, err)
			return
		}

		ctx.Request.Header.Set("X-User-ID", userID)
		h(ctx)
	}
}
