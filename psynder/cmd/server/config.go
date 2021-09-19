package main

import "time"

type AppConfig struct {
	Server struct {
		Host string `yaml:"host" env:"SERVER_HOST"`
		Port string `yaml:"server" env:"SERVER_PORT"`
	} `yaml:"server"`
	TLS struct {
		CertPath string `yaml:"cert-path" env:"TLS_CERT_PATH"`
		KeyPath string `yaml:"key-path" env:"TLS_KEY_PATH"`
	} `yaml:"tls"`
	JWT struct {
		KeyPath string `yaml:"key-path" env:"JWT_KEY_PATH"`
		PublicKeyPath string `yaml:"public-key-path" env:"JWT_PUBLIC_KEY_PATH"`
		TokenDuration time.Duration `yaml:"token-duration" env:"JWT_TOKEN_DURATION"`
		CookieDuration time.Duration `yaml:"cookie-duration" env:"JWT_COOKIE_DURATION"`
		Issuer string `yaml:"issuer" env:"JWT_ISSUER"`
		URL string `yaml:"url" env:"JWT_URL"`
	} `yaml:"jwt"`
}