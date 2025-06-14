package configs

import (
	"os"

	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/data"
	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/domain/service"
	"github.com/spf13/viper"
)

var (
	logger *Logger
)

type Deps struct {
	Cfg *mapping.Config
	DB  contract.DataManager
	Svc contract.ServiceManager
}

func Init() *Deps {
	logger.Info("Loading configs")
	profile := os.Getenv("PROFILE")
	cfg, err := read(profile)
	if err != nil {
		logger.Errorf("Error to read config %v", err)
	}

	deps := &Deps{
		Cfg: cfg,
	}
	return deps
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func (deps *Deps) ConfigDB() *Deps {
	logger = GetLogger("mysql")
	db, err := data.Connect(deps.Cfg)
	if err != nil {
		logger.Errorf("Error to conect database %v", err)
		return nil
	}
	deps.DB = db
	return deps
}

func (deps *Deps) ConfigService() *Deps {
	if deps.DB == nil {
		logger.Errorf("Db was not initialize")
	}
	svc, err := service.New(service.ServiceDeps{
		Cfg: deps.Cfg,
		DB:  deps.DB,
	})
	if err != nil {
		logger.Errorf("Service error: %v", err)
	}
	deps.Svc = svc
	return deps
}

func (deps *Deps) ConfigAll() *Deps {
	return deps.ConfigDB().ConfigService()
}

func read(profile string) (*mapping.Config, error) {
	setupConfig(profile)
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Error to read config: %v", err)
		return nil, err
	}

	config := mapping.Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func setupConfig(profile string) {
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()
	viper.SetConfigName(profile)
	viper.AddConfigPath(".")
}
