package utils

import (
	"fmt"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
	"github.com/valyala/fasthttp"
)

type RestClientConfig struct {
	ServiceName string
}

type RestClientComponents struct {
	Logger          *logs.Logger
	DiscoveryClient *DiscoveryClient
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

func (c *RestClient) PerformRequest(req *fasthttp.Request, path string) (*fasthttp.Response, *errs.Error) {
	address, err := c.components.DiscoveryClient.ResolveName(c.config.ServiceName)
	if err != nil {
		return nil, err
	}

	req.SetRequestURI(fmt.Sprintf("http://%s%s", address, path))

	resp := fasthttp.AcquireResponse()
	reqErr := c.client.Do(req, resp)
	if reqErr != nil {
		return nil, errs.NewServiceError(reqErr)
	}

	return resp, nil
}
