package nats

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/nats-io/stan.go"
	"github.com/vladjong/L0/internal/app/cache"
	"github.com/vladjong/L0/internal/app/model"
	"github.com/vladjong/L0/internal/app/store"
)

type Nats struct {
	config *Config
	store  *store.Store
	cache  *cache.Cache
}

func New(config *Config, memoryCache *cache.Cache) *Nats {
	return &Nats{
		config: config,
		cache:  memoryCache,
	}
}

func (st *Nats) configureStore() error {
	s := store.New(st.config.Store)
	if err := s.Open(); err != nil {
		return err
	}
	st.store = s
	return nil
}

func (st *Nats) Start() error {
	if err := st.configureStore(); err != nil {
		return err
	}
	sc, err := stan.Connect(st.config.ClusterId, st.config.ClientId, stan.NatsURL(st.config.Host),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		return err
	}
	defer sc.Close()
	sc.Subscribe(st.config.Subject, func(msg *stan.Msg) {
		if err := save(st.cache, st.store, msg.Data); err != nil {
			log.Println(err)
		}
	})
	Block()
	return nil
}

func Block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}

func save(cache *cache.Cache, store *store.Store, m []byte) error {
	target := model.Order{}
	err := json.Unmarshal(m, &target)
	if err != nil {
		return err
	}
	log.Println("Saving message in db")
	target.Validate()
	err = store.Order().Create(&target)
	if err != nil {
		return err
	}
	log.Println("Saving message in cache")
	cache.Set(target.OrderId, target)
	return nil
}
