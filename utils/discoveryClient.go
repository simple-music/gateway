package utils

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/simple-music/gateway/consts"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
)

type DiscoveryClientConfig struct {
	DiscoveryServerHost string
	DiscoveryServerPort string
}

type DiscoveryClientComponents struct {
	Logger *logs.Logger
}

type DiscoveryClient struct {
	client     *api.Client
	config     DiscoveryClientConfig
	components DiscoveryClientComponents
}

func NewDiscoveryClient(conf DiscoveryClientConfig, comp DiscoveryClientComponents) *DiscoveryClient {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return &DiscoveryClient{
		client:     client,
		config:     conf,
		components: comp,
	}
}

func (c *DiscoveryClient) ResolveName(name string) (string, *errs.Error) {
	srvList, err := c.client.Agent().Services()
	if err != nil {
		return consts.EmptyString, errs.NewServiceError(err)
	}

	for _, srv := range srvList {
		if srv.Service == name {
			return fmt.Sprintf("%s:%d", srv.Address, srv.Port), nil
		}
	}

	return consts.EmptyString, errs.NewServiceError(
		errors.New("service name not resolved"),
	)
}
