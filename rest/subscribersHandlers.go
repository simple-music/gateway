package rest

import (
	"github.com/simple-music/gateway/models"
	"github.com/valyala/fasthttp"
)

func (srv *Service) addSubscription(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)
	subscription := ctx.UserValue("subscription").(string)

	var musician models.Musician
	if err := srv.musiciansClient.GetMusician(user, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}
	if err := srv.musiciansClient.GetMusician(subscription, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	if err := srv.subscriptionsClient.AddSubscription(user, subscription); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *Service) deleteSubscription(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)
	subscription := ctx.UserValue("subscription").(string)

	if err := srv.subscriptionsClient.DeleteSubscription(user, subscription); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *Service) getSubscribers(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)

	pageIndex := ctx.QueryArgs().GetUintOrZero("page")
	pageSize := ctx.QueryArgs().GetUintOrZero("size")

	content, err := srv.subscriptionsClient.GetSubscribers(user, pageIndex, pageSize)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(content)
}

func (srv *Service) getSubscriptions(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)

	pageIndex := ctx.QueryArgs().GetUintOrZero("page")
	pageSize := ctx.QueryArgs().GetUintOrZero("size")

	content, err := srv.subscriptionsClient.GetSubscriptions(user, pageIndex, pageSize)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(content)
}