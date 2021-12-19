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
	"github.com/peltorator/psynder/internal/serviceimpl/tokenissuer"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func run_shelters() {
	const readTimeout = 10 * time.Second
	const writeTimeout = 10 * time.Second

	yamlConfigPath := "config_shelters.yml"
	flag.Parse()

	var cfg AppConfig
	err := cleanenv.ReadConfig(yamlConfigPath, &cfg)
	if err != nil {
		panic(err)
	}

	privateBytes, err := ioutil.ReadFile(cfg.JWT.KeyPath)
	if err != nil {
		panic(err)
	}
	publicBytes, err := ioutil.ReadFile(cfg.JWT.PublicKeyPath)
	if err != nil {
		panic(err)
	}
	tokenIssuer, err := tokenissuer.NewJWT(privateBytes, publicBytes, cfg.JWT.TokenDuration)
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Dbname)
	conn, err := postgres.New(connStr)
	if err != nil {
		panic(err)
	}

	var logger *zap.Logger
	if cfg.DevMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}

	accountRepo := postgres.NewAccountRepo(conn)
	shelterRepo := postgres.NewShelterRepo(conn)

	authService := authservice.New(accountRepo, tokenIssuer)
	shelterService := shelterservice.New(shelterRepo)

	api := httpapi.NewShelters(httpapi.ArgsShelters{
		DevMode:      cfg.DevMode,
		AuthService:  authService,
		ShelterService: shelterService,
		Logger:       logger.Sugar(),
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
