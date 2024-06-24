package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vpascoalr/goeda/internal/bar"
	"github.com/vpascoalr/goeda/internal/foo"
)

func main() {
	go foo.Start()
	log.Println("started foo service")
	go bar.Start()
	log.Println("started bar service")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
