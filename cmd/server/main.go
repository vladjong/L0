package main

import (
	"flag"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/vladjong/L0/internal/app/cache"
	"github.com/vladjong/L0/internal/app/nats"
	"github.com/vladjong/L0/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	log.Println("Inits cash")
	cache := cache.New(5*time.Minute, 10*time.Minute)
	log.Println("Init server")
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Init nats")
	configNats := nats.NewConfig()
	_, err = toml.DecodeFile(configPath, configNats)
	if err != nil {
		log.Fatal(err)
	}
	ns := nats.New(configNats, cache)
	log.Println("Starting nats")
	go func() {
		if err := ns.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Starting server")
	s := server.New(config, cache)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
