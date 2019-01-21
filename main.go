package main

import (
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/rest"
	"os"
	"os/signal"
)

func main() {
	logger := common.Logger

	srv := rest.NewService()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := srv.Run(); err != nil {
			logger.Error(err)
		}
		ch <- os.Interrupt
	}()
	defer func() {
		if err := srv.Shutdown(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch
}
