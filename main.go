package main

import (
	"Golang-Banking/app"
	"Golang-Banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
