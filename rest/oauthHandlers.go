package rest

import (
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/models"
	"github.com/valyala/fasthttp"
)

func (srv *Service) authorizeClient(ctx *fasthttp.RequestCtx) {
	clientID := string(ctx.QueryArgs().Peek("client_id"))
	clientSecret := string(ctx.QueryArgs().Peek("client_secret"))
	redirectURI := string(ctx.QueryArgs().Peek("redirect_uri"))
	_ = string(ctx.QueryArgs().Peek("scope"))

	authCode := models.AuthCode{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	if err := srv.authClient.GetAuthCode(&authCode); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	f := createOauthForm(authCode.AuthCode, redirectURI)

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/html")
	ctx.SetBody([]byte(f))
}

func (srv *Service) getOauthToken(ctx *fasthttp.RequestCtx) {
	username := string(ctx.PostArgs().Peek("username"))
	password := string(ctx.PostArgs().Peek("password"))
	redirectURI := string(ctx.PostArgs().Peek("redirect_uri"))

	credentials := models.AuthCredentials{
		Username: username,
		Password: password,
	}
	data, err := srv.authClient.StartSession(&credentials)
	if err != nil {
		srv.WriteError(ctx, err)
		return
	}

	var session models.Session
	if err := easyjson.Unmarshal(data, &session); err != nil {
		srverr := errs.NewServiceError(err)
		srv.WriteError(ctx, srverr)
		return
	}

	redirectURI += fmt.Sprintf("auth_token=%s&refresh_token=%s&user_id=%s",
		session.AuthToken, session.RefreshToken, session.UserID,
	)

	ctx.SetStatusCode(fasthttp.StatusFound)
	ctx.Response.Header.Set("Location", redirectURI)
	ctx.SetBody(data)
}

func createOauthForm(authCode string, redirectURI string) string {
	return `
<!DOCTYPE HTML>
<html>
 <head>
  <meta charset="utf-8">
  <title>Authorize application</title>
 </head>
 <body>

 <form action="/oauth/token" method="POST">
  <p>Username:</p>
  <p><input type="text" name="username"></p>
  <p>Password:</p>
  <p><input type="password" name="password"></p>
  <input type="hidden" name="auth_code" value="` + authCode + `">
  <input type="hidden" name="redirect_uri" value="` + redirectURI + `">
  <p><input type="submit"></p>
 </form>

 </body>
</html>
    `
}
