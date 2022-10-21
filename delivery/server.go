package delivery

import (
	"pinjol-perdana/config"
	"pinjol-perdana/delivery/controller"
	"pinjol-perdana/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		useCaseManager: usecase,
		engine:         ginEngine,
		host:           config.Url,
	}
}

func (a *appServer) initHandlers() {
	controller.NewAccountController(a.engine, a.useCaseManager.AccountUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run()
	if err != nil {
		panic(err)
	}
}
