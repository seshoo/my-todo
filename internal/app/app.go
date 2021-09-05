package app

import (
	"fmt"

	"github.com/seshoo/my-todo/internal/config"
	logger "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	cnf, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
	}

	fmt.Printf("%+v\n", cnf)
}
