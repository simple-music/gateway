package main

import (
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/logs"
	"github.com/simple-music/gateway/utils"
)

func main() {
	logger := logs.NewLogger()

	_ = utils.NewDiscoveryClient(
		utils.DiscoveryClientConfig{
			DiscoveryServerHost: config.DiscoveryHost,
			DiscoveryServerPort: config.DiscoveryPort,
		},
		utils.DiscoveryClientComponents{
			Logger: logger,
		},
	)
}
