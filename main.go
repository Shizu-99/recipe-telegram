package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-contrib/graceful"

	"github.com/Shizu-99/recipe-telegram/api"
	"github.com/Shizu-99/recipe-telegram/database"
)

var (
	dataPath = filepath.Join(".", "files")
)

func main() {
	if envDataPath := os.Getenv("RECIPES_DATA_PATH"); envDataPath != "" {
		dataPath = envDataPath
	}
	if err := os.MkdirAll(dataPath, 0755); err != nil {
		panic(err)
	}

	if err := database.OpenDatabase(filepath.Join(dataPath, "recipes.sqlite3")); err != nil {
		panic(err)
	}
	defer database.CloseDatabase()

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
