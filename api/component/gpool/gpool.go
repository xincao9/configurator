package gpool

import (
    "configurator/api/component/config"
    "configurator/api/component/logger"
    "github.com/panjf2000/ants/v2"
)

var (
	G *ants.Pool
)

func init() {
	var err error
	G, err = ants.NewPool(config.C.GetInt("gpool.size"))
	if err != nil {
		logger.L.Fatalf("Fatal ants new %v\n", err)
	}
}
