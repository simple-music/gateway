package common

import (
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/logs"
	"github.com/simple-music/gateway/utils"
)

var (
	Logger          *logs.Logger
	TaskQueue       *utils.TaskQueue
	DiscoveryClient *utils.DiscoveryClient
)

func init() {
	Logger = logs.NewLogger()

	TaskQueue = utils.NewTaskQueue(Logger)

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
