package app

import (
	"fmt"

	"github.com/seshoo/my-todo/internal/config"
)

func Run(configPath string) {
	cnf, err := config.Init(configPath)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("%+v\n", cnf)
}
