package main

import (
	"discuss/internal/config"
	"discuss/internal/https_server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port

	go func() {
		if err := https_server.GE.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
			log.Fatal("Failed to start server:", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
