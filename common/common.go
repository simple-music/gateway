package common

import (
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/logs"
	"github.com/simple-music/gateway/utils"
)

var (
	Logger          *logs.Logger
	DiscoveryClient *utils.DiscoveryClient
)

func init() {
	Logger = logs.NewLogger()

	DiscoveryClient = utils.NewDiscoveryClient(
		utils.DiscoveryClientConfig{
			DiscoveryServerHost: config.DiscoveryHost,
			DiscoveryServerPort: config.DiscoveryPort,
		},
		utils.DiscoveryClientComponents{
			Logger: Logger,
		},
	)
}
