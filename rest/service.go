package rest

import (
	"github.com/fasthttp/router"
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
	"github.com/valyala/fasthttp"
)

type ServiceConfig struct {
	Host string
	Port string
}

type ServiceComponents struct {
	Logger *logs.Logger
}

type Service struct {
	config     ServiceConfig
	components ServiceComponents
	handler    fasthttp.RequestHandler

	reqBodyErr *errs.Error
}

func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	cf.Host, cf.Port = config.ServiceHost, config.ServicePort

	srv := &Service{
		config:     cf,
		components: cp,

		reqBodyErr: errs.NewError(
			errs.InvalidFormat, "invalid request body",
		),
	}

	r := router.New()

	r.POST("/users", nil)
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

	r.POST("/users/:user/avatar", nil)
	r.GET("/users/:user/avatar", nil)
	r.DELETE("/users/:user/avatar", nil)

	r.GET("/users/:user/compositions", nil)
	r.POST("/users/:user/compositions/:composition", nil)
	r.GET("/users/:user/compositions/:composition", nil)
	r.DELETE("/users/:user/compositions/:composition", nil)

	srv.handler = r.Handler
	return srv
}

func (srv *Service) Run() error {
	addr := srv.config.Host + ":" + srv.config.Port
	srv.Logger().Info("starting service on " + addr)
	return fasthttp.ListenAndServe(addr, srv.handler)
}

func (srv *Service) Shutdown() error {
	return nil //TODO
}

func (srv *Service) Logger() *logs.Logger {
	return srv.components.Logger
}
