package rest

import (
	"github.com/valyala/fasthttp"
)

func (srv *Service) addAvatar(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)
	content := ctx.Request.Body()

	err := srv.avatarsClient.AddAvatar(user, content)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *Service) getAvatar(ctx *fasthttp.RequestCtx) {
	//TODO
}

func (srv *Service) deleteAvatar(ctx *fasthttp.RequestCtx) {
	//TODO
}
