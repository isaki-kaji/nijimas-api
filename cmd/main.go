package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"go.uber.org/fx"

	"github.com/isaki-kaji/nijimas-api/api"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	"github.com/isaki-kaji/nijimas-api/api/route"
	"github.com/isaki-kaji/nijimas-api/application/service"
	"github.com/isaki-kaji/nijimas-api/configs"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

func NewConfig() (*configs.Config, error) {
	return configs.LoadConfig("environment/development")
}

func StartServer(lc fx.Lifecycle, config *configs.Config, server *api.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting server at %s", config.ServerAddress)
			if err := server.Start(); err != nil {
				log.Printf("Failed to start server: %v", err)
				return err
			}
			log.Printf("Server started successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(NewConfig),
		configs.Module,
		db.Module,
		service.Module,
		api.Module,
		controller.Module,
		route.Module,
		fx.Invoke(StartServer),
	)
	if err := app.Err(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	startCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
