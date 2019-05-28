package main

import (
	c "github.com/gs-open-provider/poc-go-uber-zap/internal/configs"
)

func main() {
	// Initialize Viper across the application
	c.InitializeViper()

	// Initialize Logger across the application
	c.InitializeZapCustomLogger()

	c.Log.Info("Logger Initiaized..")

}
