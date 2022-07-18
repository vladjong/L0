package nats

import "github.com/vladjong/L0/internal/app/store"

type Config struct {
	ClusterId string
	ClientId  string
	Host      string
	Subject   string
	Store     *store.Config
}

func NewConfig() *Config {
	return &Config{
		Host:      "",
		ClusterId: "prod",
		Subject:   "test",
		ClientId:  "simple_pub",
		Store:     store.NewConfig(),
	}
}
