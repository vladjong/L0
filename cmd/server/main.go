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
	cache := cache.New(5*time.Minute, 10*time.Minute)
	flag.Parse()
	config := server.NewConfig()
	// Initialize Server subscriber
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Nats subscriber
	nsc := nats.NewConfig()
	_, err = toml.DecodeFile(configPath, nsc)
	if err != nil {
		log.Fatal(err)
	}
	ns := nats.New(nsc, cache)
	log.Println("Starting nats")
	go func() {
		if err := ns.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Starting web server")
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
