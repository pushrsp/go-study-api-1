package main

import (
	"github.com/pushrsp/banking/app"
	"github.com/pushrsp/banking/logger"
)

func main() {
	logger.Info("Start Application")
	app.Start()
}
