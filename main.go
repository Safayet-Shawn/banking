package main

import (
	"github.com/Safayet-Shawn/banking/app"
	"github.com/Safayet-Shawn/banking/logger"
)

func main() {
	// log.Println("Starting Banking Application.......")
	logger.Info("Starting Banking Application.......")
	app.Start()
}
