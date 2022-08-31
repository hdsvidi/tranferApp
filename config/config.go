package config

import (
	"fmt"
	"mncTestApp/logger"
	"mncTestApp/manager"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
type Manager struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type Config struct {
	Manager
	ApiConfig
	DbConfig
	LogLevel string
}

func (c Config) readConfigFile() Config {


	c.ApiConfig = ApiConfig{Url: ":8080"}
	c.DbConfig = DbConfig{
		Host:     "localhost",
		Port:     "5432",
		Name:     "data_user",
		User:     "vide",
		Password: "victoriacoren",
	}
	c.LogLevel = "debug"
	return c
}
func NewConfig() Config {
	cfg := Config{}
	cfg = cfg.readConfigFile()


	logger.New(cfg.LogLevel)
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbConfig.User, cfg.DbConfig.Password, cfg.DbConfig.Host, cfg.DbConfig.Port, cfg.DbConfig.Name)

	cfg.InfraManager = manager.NewInfra(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
