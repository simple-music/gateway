package rest

import (
	"github.com/simple-music/gateway/models"
	"github.com/valyala/fasthttp"
)

func (srv *Service) startSession(ctx *fasthttp.RequestCtx) {
	var authCredentials models.AuthCredentials
	if err := srv.ReadBody(ctx, &authCredentials); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	//TODO validation

	content, err := srv.authClient.StartSession(&authCredentials)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(content)
}

func (srv *Service) refreshSession(ctx *fasthttp.RequestCtx) {
	token := string(ctx.QueryArgs().Peek("refreshToken"))

	content, err := srv.authClient.RefreshSession(token)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(content)
}

func (srv *Service) deleteSession(ctx *fasthttp.RequestCtx) {
	token := string(ctx.QueryArgs().Peek("refreshToken"))

	if err := srv.authClient.DeleteSession(token); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
