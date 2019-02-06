package rest

import (
	"github.com/fasthttp/router"
	"github.com/simple-music/gateway/clients"
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
	"github.com/simple-music/gateway/models"
	"github.com/simple-music/gateway/utils"
	"github.com/valyala/fasthttp"
)

type Service struct {
	handler fasthttp.RequestHandler

	logger *logs.Logger

	taskQueue    *utils.TaskQueue
	tokenManager *utils.TokenManager

	reqBodyErr    *errs.Error
	authTokenErr  *errs.Error
	permissionErr *errs.Error

	newUserValidator        *models.NewUserValidator
	musicianUpdateValidator *models.MusicianUpdateValidator

	authClient          *clients.AuthClient
	musiciansClient     *clients.MusiciansClient
	subscriptionsClient *clients.SubscriptionsClient
	avatarsClient       *clients.AvatarsClient
}

func NewService() *Service {
	srv := &Service{
		logger: common.Logger,

		reqBodyErr: errs.NewError(
			errs.InvalidFormat, "invalid request body",
		),
		authTokenErr: errs.NewError(
			errs.BadRequest, "invalid authorization token",
		),
		permissionErr: errs.NewError(
			errs.PermissionDenied, "permission denied",
		),

		newUserValidator:        models.NewNewUserValidator(),
		musicianUpdateValidator: models.NewMusicianUpdateValidator(),

		authClient:          clients.NewAuthClient(),
		musiciansClient:     clients.NewMusiciansClient(),
		subscriptionsClient: clients.NewSubscriptionsClient(),
		avatarsClient:       clients.NewAvatarsClient(),

		taskQueue:    common.TaskQueue,
		tokenManager: utils.NewTokenManager(),
	}

	r := router.New()

	r.POST("/users", srv.addUser)
	r.GET("/users", srv.findUser)
	r.GET("/users/:user", srv.getUser)
	r.PATCH("/users/:user", srv.WithAuth(srv.updateUser))
	r.DELETE("/users/:user", srv.WithAuth(srv.deleteUser))

	r.POST("/auth/session", srv.startSession)
	r.PATCH("/auth/session", srv.refreshSession)
	r.DELETE("/auth/session", srv.deleteSession)

	r.POST("/oauth/auth", srv.authorizeClient)
	r.POST("/oauth/token", srv.getOauthToken)

	r.GET("/users/:user/subscribers", srv.getSubscribers)
	r.GET("/users/:user/subscriptions", srv.getSubscriptions)
	r.POST("/users/:user/subscriptions/:subscription", srv.WithAuth(srv.addSubscription))
	r.GET("/users/:user/subscriptions/:subscription", srv.checkSubscription)
	r.DELETE("/users/:user/subscriptions/:subscription", srv.WithAuth(srv.deleteSubscription))

	r.POST("/users/:user/avatar", srv.WithAuth(srv.addAvatar))
	r.GET("/users/:user/avatar", srv.getAvatar)
	r.DELETE("/users/:user/avatar", srv.WithAuth(srv.deleteAvatar))

	srv.handler = srv.WithLogs(r.Handler)
	return srv
}

func (srv *Service) Run() error {
	addr := config.ServiceHost + ":" + config.ServicePort
	srv.logger.Info("starting service on " + addr)
	return fasthttp.ListenAndServe(addr, srv.handler)
}

func (srv *Service) Shutdown() error {
	return nil
}
