package main

import (
	"log/slog"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/VeryHappy2/GoApi/docs"
	"github.com/VeryHappy2/GoApi/internal/config"
	"github.com/VeryHappy2/GoApi/internal/http-server/handlers/profile"
	"github.com/VeryHappy2/GoApi/internal/repositories"
	"github.com/VeryHappy2/GoApi/internal/storage/sqlite"
	"github.com/VeryHappy2/GoApi/pkg"
	"github.com/VeryHappy2/GoApi/router"
)

//	@title			Training App API
//	@version		1.0
//	@description	API Server for training

//	@host		localhost:8080
//	@BasePath	/

func main() {
	config := config.MustLoad()

	log := pkg.SetupLog(config.Env)
	log.Info("Starting logging")

	db, err := sqlite.New(config.StoragePath)

	if err != nil {
		log.Error("Failed to create storage", "error", err)
	}

	if db == nil {
		log.Info("Storage is nil")
	}
	p := profile.Profile{}

	var routes = []router.Route{
		{
			Method:      "POST",
			Path:        "/profile/add",
			HandlerFunc: p.New(log, repositories.ProfileRep{}, db),
		},
		{
			Method:      "POST",
			Path:        "/profile/getById/{id}",
			HandlerFunc: p.GetById(log, repositories.ProfileRep{}, db),
		},
	}

	r := router.NewRouter(routes)

	r.Mount("/swagger", httpSwagger.WrapHandler)

	log.Info("starting server", slog.String("address", config.Address))

	srv := &http.Server{
		Addr:         config.Address,
		Handler:      r,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("server was stopped")
}
