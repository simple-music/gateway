package rest

import (
	"github.com/mailru/easyjson"
	"github.com/simple-music/gateway/errs"
	"github.com/valyala/fasthttp"
)

const (
	JsonType = "application/json"
)

func (srv *Service) ReadBody(ctx *fasthttp.RequestCtx, v easyjson.Unmarshaler) *errs.Error {
	if err := easyjson.Unmarshal(ctx.PostBody(), v); err != nil {
		return srv.reqBodyErr
	}
	return nil
}

func (srv *Service) WriteJSON(ctx *fasthttp.RequestCtx, status int, v easyjson.Marshaler) {
	b, _ := easyjson.Marshal(v)
	ctx.SetStatusCode(status)
	ctx.Response.Header.SetContentType(JsonType)
	ctx.Response.SetBody(b)
}

func (srv *Service) WriteError(ctx *fasthttp.RequestCtx, err *errs.Error) {
	status := fasthttp.StatusInternalServerError

	switch err.Type {
	case errs.NotFound:
		status = fasthttp.StatusNotFound
	case errs.BadRequest:
		status = fasthttp.StatusBadRequest
	case errs.Conflict:
		status = fasthttp.StatusConflict
	case errs.InvalidFormat:
		status = fasthttp.StatusUnprocessableEntity
	case errs.NotAuthorized:
		status = fasthttp.StatusUnauthorized
	case errs.PermissionDenied:
		status = fasthttp.StatusForbidden
	case errs.InternalError:
		srv.logger.Error(err.NestedErr)
	}

	ctx.SetStatusCode(status)
	ctx.Response.Header.SetContentType(JsonType)
	ctx.Response.SetBody(err.JSON)
}
