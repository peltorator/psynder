package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"net/http"
	"psynder/internal/interface/httpapi"
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

	a := httpapi.New()

	server := http.Server{
		Addr: fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		Handler: a.Router(),
	}

	err = server.ListenAndServeTLS(cfg.TLS.CertPath, cfg.TLS.KeyPath)
	if err != nil {
		panic(err)
	}
}