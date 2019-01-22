package rest

import (
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/models"
	"github.com/simple-music/gateway/utils"
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

func (srv *Service) findUser(ctx *fasthttp.RequestCtx) {
	username := string(ctx.QueryArgs().Peek("username"))

	var musician models.Musician
	if err := srv.musiciansClient.FindMusician(username, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	var status models.SubscriptionsStatus
	if err := srv.subscriptionsClient.GetStatus(musician.ID, &status); err != nil {
		if err.Type == errs.InternalError {
			srv.logger.Error(err.NestedErr)
		}
	}

	var userFull models.UserFull
	userFull.From(&musician, &status)

	srv.WriteJSON(ctx, fasthttp.StatusOK, &userFull)
}

func (srv *Service) getUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("user").(string)

	var musician models.Musician
	if err := srv.musiciansClient.GetMusician(id, &musician); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	var status models.SubscriptionsStatus
	if err := srv.subscriptionsClient.GetStatus(id, &status); err != nil {
		if err.Type == errs.InternalError {
			srv.logger.Error(err.NestedErr)
		}
	}

	var userFull models.UserFull
	userFull.From(&musician, &status)

	srv.WriteJSON(ctx, fasthttp.StatusOK, &userFull)
}

func (srv *Service) updateUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("user").(string)

	authID := string(ctx.Request.Header.Peek("X-User-ID"))
	if authID != id {
		srv.WriteError(ctx, srv.permissionErr)
		return
	}

	var update models.MusicianUpdate
	if err := srv.ReadBody(ctx, &update); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	//TODO validation

	if err := srv.musiciansClient.UpdateMusician(id, &update); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *Service) deleteUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("user").(string)

	authID := string(ctx.Request.Header.Peek("X-User-ID"))
	if authID != id {
		srv.WriteError(ctx, srv.permissionErr)
		return
	}

	if err := srv.musiciansClient.DeleteMusician(id, true); err != nil {
		srv.WriteError(ctx, err)
		return
	}

	srv.taskQueue.AddTask(func() *utils.Task {
		return &utils.Task{
			TaskFunc: func() bool {
				err := srv.authClient.DeleteCredentials(id)
				if err.Type == errs.InternalError {
					srv.logger.Error(err.NestedErr)
					return false
				}
				return true
			},
		}
	}())

	srv.taskQueue.AddTask(func() *utils.Task {
		return &utils.Task{
			TaskFunc: func() bool {
				err := srv.subscriptionsClient.DeleteUser(id)
				if err.Type == errs.InternalError {
					srv.logger.Error(err.NestedErr)
					return false
				}
				return true
			},
		}
	}())

	srv.taskQueue.AddTask(func() *utils.Task {
		return &utils.Task{
			TaskFunc: func() bool {
				err := srv.avatarsClient.DeleteAvatar(id)
				if err.Type == errs.InternalError {
					srv.logger.Error(err.NestedErr)
					return false
				}
				return true
			},
		}
	}())

	ctx.SetStatusCode(fasthttp.StatusNoContent)
}
