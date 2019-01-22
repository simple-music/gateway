package rest

import (
	"github.com/simple-music/gateway/models"
	"github.com/valyala/fasthttp"
)

func (srv *Service) addUser(ctx *fasthttp.RequestCtx) {
	var newUser models.NewUser
	if err := srv.ReadBody(ctx, &newUser); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	//TODO validation

	var newMusician models.NewMusician
	newMusician.From(&newUser)

	var musician models.Musician
	if err := srv.musiciansClient.AddMusician(&newMusician, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	var newCredentials models.NewCredentials
	newCredentials.From(&newUser, &musician)

	if err := srv.authClient.AddCredentials(&newCredentials); err != nil {
		if err := srv.musiciansClient.DeleteMusician(musician.ID, true); err != nil {
			srv.WriteError(ctx, err)
			return
		}
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}
