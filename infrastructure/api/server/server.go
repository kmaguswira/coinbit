package server

import (
	"github.com/kmaguswira/coinbit/application/config"
)

func Init() {
	r := SetupRouter()

	r.Run(config.GetConfig().Port)
}
