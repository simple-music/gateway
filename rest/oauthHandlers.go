package rest

import (
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
	user := ctx.UserValue("user").(string)
	subscription := ctx.UserValue("subscription").(string)

	if err := srv.subscriptionsClient.CheckSubscription(user, subscription); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
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
