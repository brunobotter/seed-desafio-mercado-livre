package main

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/routers"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")
	deps := configs.Init().ConfigAll()
	if deps == nil || deps.DB == nil || deps.Svc == nil {
		logger.Errorf("Config initialize with error")
		return
	}
	routers.Initialize(deps)
}
