package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/peltorator/psynder/internal/api/httpapi"
	"github.com/peltorator/psynder/internal/repo/postgres"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"github.com/peltorator/psynder/internal/serviceimpl/shelterservice"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func getShelterService() (*shelterservice.ShelterService, AppConfig) {

	yamlConfigPath := "config-shelters.yml"
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

	shelterRepo := postgres.NewShelterRepo(conn)
	shelterService := shelterservice.New(shelterRepo)
	return shelterService, cfg
}

func run_shelters(a *authservice.AuthService, sh *shelterservice.ShelterService, cfg AppConfig) {
	const readTimeout = 10 * time.Second
	const writeTimeout = 10 * time.Second

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

	api := httpapi.NewShelters(httpapi.ArgsShelters{
		DevMode:        cfg.DevMode,
		AuthService:    a,
		ShelterService: sh,
		Logger:         logger.Sugar(),
	})

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on %v...\n", addr)
	server := http.Server{
		Addr:         addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      api.RouterShelters(),
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
