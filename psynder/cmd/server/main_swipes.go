package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/peltorator/psynder/internal/api/httpapi"
	"github.com/peltorator/psynder/internal/repo/postgres"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"github.com/peltorator/psynder/internal/serviceimpl/swipeservice"
	"go.uber.org/zap"
	"net/http"
)

func getSwipeService() (*swipeservice.SwipeService, AppConfig) {

	yamlConfigPath := "config-swipes.yml"
	flag.Parse()

	var cfg AppConfig
	err := cleanenv.ReadConfig(yamlConfigPath, &cfg)
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Dbname)
	conn, err := postgres.New(connStr)
	if err != nil {
		panic(err)
	}

	psynaRepo := postgres.NewPsynaRepo(conn)
	likeRepo := postgres.NewLikeRepo(conn)

	swipeService := swipeservice.New(swipeservice.Args{
		Psynas: psynaRepo,
		Likes:  likeRepo,
	})
	return swipeService, cfg
}

func run_swipes(a *authservice.AuthService, sw *swipeservice.SwipeService, cfg AppConfig) {

	var logger *zap.Logger
	var err error
	if cfg.DevMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}

	api := httpapi.New(httpapi.Args{
		DevMode:      cfg.DevMode,
		AuthService:  a,
		SwipeService: sw,
		Logger:       logger.Sugar(),
	})

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on %v...\n", addr)
	server := http.Server{
		Addr:         addr,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		Handler:      api.Router(),
	}

	if cfg.TLS.Enable {
		err = server.ListenAndServeTLS(cfg.TLS.CertPath, cfg.TLS.KeyPath)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		panic(err)
	}
}
