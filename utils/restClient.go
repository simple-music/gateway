package utils

import (
	"github.com/simple-music/gateway/logs"
	"github.com/valyala/fasthttp"
)

type RestClientConfig struct {
	ServiceName string
}

type RestClientComponents struct {
	Logger *logs.Logger
}

type RestClient struct {
	client     fasthttp.Client
	config     RestClientConfig
	components RestClientComponents
}

func NewRestClient(cf RestClientConfig, cp RestClientComponents) *RestClient {
	return &RestClient{
		config:     cf,
		components: cp,
		client: fasthttp.Client{
			Name: cf.ServiceName + "-client",
		},
	}
}
