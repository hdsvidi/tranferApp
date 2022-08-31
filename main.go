
package main

import (
	"mncTestApp/config"
	"mncTestApp/delivery/middleware"
	"mncTestApp/logger"
	"mncTestApp/delivery/login"
)
import "github.com/gin-gonic/gin"

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (p *appServer) initHandlers() {
	p.routerEngine.Use(middleware.ErrorMiddleware())
}

func (p *appServer) Run() {
	p.initHandlers()
	logger.Log.Info().Msgf("Server run on %s", p.cfg.ApiConfig.Url)
	err := p.routerEngine.Run(p.cfg.ApiConfig.Url)
	if err != nil {
		logger.Log.Fatal().Msg("Server failed to run")
	}

}

func Server() AppServer {
	r := gin.Default()
	r.SetTrustedProxies([]string{""})
	r.GET("/login", login.Login())

	c := config.NewConfig()
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}


func main() {
	Server().Run()
}
