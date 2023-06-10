package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Env      EnvConfig
	Database DatabaseConfig
	QrEnv    QrConfig
}

type EnvConfig struct {
	AppName string `envconfig:"APP_NAME"`
	AppPort string `envconfig:"APP_PORT"`
}

type DatabaseConfig struct {
	Host         string `envconfig:"DB_HOST"`
	Port         string `envconfig:"DB_PORT"`
	DatabaseName string `envconfig:"DB_NAME"`
	Username     string `envconfig:"DB_USERNAME"`
	Password     string `envconfig:"DB_PASSWORD"`
}

type QrConfig struct {
	AmountLimit float64 `envconfig:"QR_AMOUNT_LIMIT"`
}

func (cfg *AppConfig) Init() {
	envconfig.MustProcess("", &cfg.Env)
	envconfig.MustProcess("", &cfg.Database)
}

func LoadAppConfig() *AppConfig {
	env, ok := os.LookupEnv("ENV")
	if ok && env != "" {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/../../.env.%v", basePath, env))
		if err != nil {
			err = godotenv.Load()
			if err != nil {
				panic(err)
			}
		}
	} else {
		godotenv.Load()
	}
	appCfg := AppConfig{}
	appCfg.Init()
	return &appCfg
}
