package main

import (
	"github.com/j4real2208/golang-db/controller"
	"github.com/j4real2208/golang-db/error"
)

func main() {
	logger := error.Getlogger()

	logger.Info("Welcome to pre-setup of Dev Site")
	error.Logger.Info("Starting Up Logger variable in Dev App")
	controller.InitHandlers()

	logger.Info("Closing of Dev Site")
	
}