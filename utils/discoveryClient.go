package utils

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/simple-music/gateway/config"
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
	apiConfig := api.DefaultConfig()
	apiConfig.Address = config.DiscoveryHost + ":" + config.DiscoveryPort

	client, err := api.NewClient(apiConfig)
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
			address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
			c.components.Logger.Info(fmt.Sprintf("Resolved %s to %s", name, address))
			return address, nil
		}
	}

	return consts.EmptyString, errs.NewServiceError(
		errors.New("service name not resolved"),
	)
}
