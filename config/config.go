package config

import (
	"github.com/simple-music/gateway/args"
	"github.com/simple-music/gateway/env"
)

var (
	ServiceHost = "0.0.0.0"
	ServicePort = args.GetString("service_port", "8080")

	RegistryHost = env.GetVar("REGISTRY_HOST", "127.0.0.1")
	RegistryPort = env.GetVar("REGISTRY_PORT", "8500")
)
