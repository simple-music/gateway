package rest

import (
	"github.com/simple-music/gateway/models"
	"github.com/valyala/fasthttp"
)

func (srv *Service) addAvatar(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)
	content := ctx.Request.Body()

	authID := string(ctx.Request.Header.Peek("X-User-ID"))
	if authID != user {
		srv.WriteError(ctx, srv.permissionErr)
		return
	}

	var musician models.Musician
	if err := srv.musiciansClient.GetMusician(user, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	err := srv.avatarsClient.AddAvatar(user, content)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *Service) getAvatar(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)

	content, err := srv.avatarsClient.GetAvatar(user)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("image/jpeg")
	ctx.SetBody(content)
}

func (srv *Service) deleteAvatar(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)

	authID := string(ctx.Request.Header.Peek("X-User-ID"))
	if authID != user {
		srv.WriteError(ctx, srv.permissionErr)
		return
	}

	if err := srv.avatarsClient.DeleteAvatar(user); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
