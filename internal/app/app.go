package app

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/seshoo/my-todo/internal/config"
	"github.com/seshoo/my-todo/internal/handler"
	"github.com/seshoo/my-todo/internal/repository"
	"github.com/seshoo/my-todo/internal/repository/postgres"
	"github.com/seshoo/my-todo/internal/server"
	"github.com/seshoo/my-todo/internal/service"
	logger "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	cnf, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
	}

	fmt.Printf("%+v\n", cnf.Postgres)

	db, err := postgres.NewDb(cnf)
	if err != nil {
		logger.Error(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(service.Dependencies{
		Repository: repos,
	})
	handlers := handler.NewHandler(services)
	srv := server.NewServer(cnf, handlers.Init())

	// go func() {
	if err := srv.Run(); err != nil {
		logger.Errorf("error occurred while running http server: %s\n", err.Error())
	}
	// }()

	logger.Info("Server started")

}
