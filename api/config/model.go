package config

import "time"

type Environment struct {
	Env             string `env:"ENV,default=dev"`
	Log             Log
	Key             Key
	Postgres        Postgres
	API             API
	SMTP            SMTP
	Cookie          Cookie
	CloudFlareImage CloudFlareImage
}

type Postgres struct {
	Host        string `env:"POSTGRES_HOST"`
	Port        int    `env:"POSTGRES_PORT"`
	User        string `env:"POSTGRES_USER"`
	Password    string `env:"POSTGRES_PASSWORD"`
	DBName      string `env:"POSTGRES_NAME"`
	DBSSLMode   string `env:"POSTGRES_SSL_MODE"`
	MaxConn     int    `env:"POSTGRES_MAX_CONN"`
	MaxIdle     int    `env:"POSTGRES_MAX_IDLE"`
	MaxLifeTime int    `env:"POSTGRES_MAX_LIFE_TIME"`
	Timeout     int    `env:"POSTGRES_TIMEOUT"`
}

type Log struct {
	Level string `env:"LOG_LEVEL,default=info"`
}

type Key struct {
	PrivateKey string `env:"KEY_ECDSA_PRIVATE"`
	PublicKey  string `env:"KEY_ECDSA_PUBLIC"`
}

type API struct {
	Port int `env:"API_PORT,default=8080"`
}

type SMTP struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	User     string `env:"SMTP_USER"`
	Password string `env:"SMTP_PASSWORD"`
}

type Cookie struct {
	Name string `env:"COOKIE_NAME,default=XPLife_id"`
}

type CloudFlareImage struct {
	Token   string        `env:"CLOUD_FLARE_IMAGE_API_TOKEN"`
	URL     string        `env:"CLOUD_FLARE_IMAGE_API_URL"`
	Timeout time.Duration `env:"CLOUD_FLARE_IMAGE_TIMEOUT,default=30s"`
}
