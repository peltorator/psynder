package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"psynder/internal/interface/httpapi"
	"psynder/internal/interface/postgres/accountrepo"
	"psynder/internal/interface/postgres/swiperepo"
	"psynder/internal/service/token"
	"psynder/internal/usecases"
	"time"
)

func main() {
	const readTimeout = 10 * time.Second
	const writeTimeout = 10 * time.Second

	yamlConfigPath := flag.String("config", "config.yml", "path to the YAML config file")
	flag.Parse()

	var cfg AppConfig
	err := cleanenv.ReadConfig(*yamlConfigPath, &cfg)
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
	tokenIssuer, err := token.NewJwtHandler(privateBytes, publicBytes, cfg.JWT.TokenDuration)
	if err != nil {
		panic(err)
	}

	connStr := "user=postgres password=123 host=localhost dbname=postgres sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	a := httpapi.New(
		usecases.NewAccountUseCases(accountrepo.New(conn), tokenIssuer),
		usecases.NewSwipeUseCases(swiperepo.New(conn)),
		httpapi.NewJSONHandler())

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on %v...\n", addr)
	server := http.Server{
		Addr:         addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      a.Router(),
	}

	//err = server.ListenAndServe()
	err = server.ListenAndServeTLS(cfg.TLS.CertPath, cfg.TLS.KeyPath)
	if err != nil {
		panic(err)
	}
}
