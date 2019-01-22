package rest

import (
	"github.com/fasthttp/router"
	"github.com/simple-music/gateway/clients"
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
	"github.com/valyala/fasthttp"
)

type Service struct {
	handler fasthttp.RequestHandler

	logger *logs.Logger

	reqBodyErr *errs.Error

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

		authClient:          clients.NewAuthClient(),
		musiciansClient:     clients.NewMusiciansClient(),
		subscriptionsClient: clients.NewSubscriptionsClient(),
		avatarsClient:       clients.NewAvatarsClient(),
	}

	r := router.New()

	r.POST("/users", srv.addUser)
	r.GET("/users", nil)
	r.GET("/users/:user", nil)
	r.PATCH("/users/:user", nil)
	r.DELETE("/users/:user", nil)

	r.POST("/auth/session", nil)
	r.PATCH("/auth/session", nil)
	r.DELETE("/auth/session", nil)

	r.GET("/users/:user/subscribers", nil)
	r.GET("/users/:user/subscriptions", nil)
	r.POST("/users/:user/subscribers/:subscriber", nil)
	r.DELETE("/users/:user/subscribers/:subscriber", nil)

	r.POST("/users/:user/avatar", srv.addAvatar)
	r.GET("/users/:user/avatar", srv.getAvatar)
	r.DELETE("/users/:user/avatar", srv.deleteAvatar)

	srv.handler = r.Handler
	return srv
}

func (srv *Service) Run() error {
	addr := config.ServiceHost + ":" + config.ServicePort
	srv.logger.Info("starting service on " + addr)
	return fasthttp.ListenAndServe(addr, srv.handler)
}

func (srv *Service) Shutdown() error {
	return nil //TODO
}
