package main

import "github.com/seshoo/my-todo/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
