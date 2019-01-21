package config

import (
	"github.com/simple-music/gateway/args"
	"github.com/simple-music/gateway/env"
)

var (
	ServiceHost = "0.0.0.0"
	ServicePort = args.GetString("service_port", "8080")

	DiscoveryHost = env.GetVar("SIMPLE_MUSIC_CONSUL_HOST", "127.0.0.1")
	DiscoveryPort = env.GetVar("SIMPLE_MUSIC_CONSUL_PORT", "8500")
)
