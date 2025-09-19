package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Shizu-99/recipe-telegram/api"
	"github.com/gin-contrib/graceful"
)

func main() {
	//Ensure that Gin exits cleanly using Graceful
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	router, err := graceful.Default()
	if err != nil {
		panic(err)
	}
	defer router.Close()

	router.GET("/", api.Home)

	if err := router.RunWithContext(ctx); err != nil && err != context.Canceled {
		panic(err)
	}
}
