package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/peltorator/psynder/internal/api/httpapi"
	"github.com/peltorator/psynder/internal/repo/postgres"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"github.com/peltorator/psynder/internal/serviceimpl/tokenissuer"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func getAuthService() (*authservice.AuthService, AppConfig) {

	yamlConfigPath := "config-accounts.yml"
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

	accountRepo := postgres.NewAccountRepo(conn)
	authService := authservice.New(accountRepo, tokenIssuer)
	return authService, cfg
}

func run_auth(a *authservice.AuthService, cfg AppConfig) {
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

	api := httpapi.NewAcccounts(httpapi.ArgsAccounts{
		DevMode:     cfg.DevMode,
		AuthService: a,
		Logger:      logger.Sugar(),
	})

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on %v...\n", addr)
	server := http.Server{
		Addr:         addr,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		Handler:      api.RouterAccounts(),
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
