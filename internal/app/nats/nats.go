package nats

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/vladjong/L0/internal/app/cache"
	"github.com/vladjong/L0/internal/app/model"
	"github.com/vladjong/L0/internal/app/store"
)

type Stan struct {
	config *Config      // Конфиг для подписчика
	store  *store.Store // База данных
	cache  *cache.Cache // Кэш
}

// New ...
func New(config *Config, memoryCache *cache.Cache) *Stan {
	// Инициализируем в конструкторе конфиг и кэш
	return &Stan{
		config: config,
		cache:  memoryCache,
	}
}

// Конфигурирование базы данных
func (st *Stan) configureStore() error {
	s := store.New(st.config.Store)
	if err := s.Open(); err != nil {
		return err
	}
	st.store = s
	return nil
}

// Start ...
func (st *Stan) Start() error {
	// Инициализируем базу данных
	if err := st.configureStore(); err != nil {
		return err
	}

	// Подключение к nats-streaming
	log.Println("Connecting to nats-streaming-server")
	sc, err := stan.Connect(st.config.ClusterId, st.config.ClientId, stan.NatsURL(st.config.Host),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		return err
	}
	defer sc.Close()

	// Оформление подписки
	log.Println("Subscribing")
	sc.Subscribe(st.config.Subject, func(msg *stan.Msg) {
		if err := save(st.cache, st.store, msg.Data); err != nil {
			log.Println(err)
		}
	})
	log.Println("debug")
	Block()
	log.Println("debug1")

	return nil
}

func Block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}

// Сохранение полученного сообщения в базу данных и в кэш
func save(cache *cache.Cache, store *store.Store, m []byte) error {
	log.Println("Saving nats message")
	target := model.Order{}
	err := json.Unmarshal(m, &target)
	if err != nil {
		return err
	}
	log.Println("Saving message in db")
	p, err := store.Order().Create(&target)
	if err != nil {
		return err
	}
	log.Println("Print row in db")
	log.Println(p)
	log.Println("Saving data in cache")
	cache.Set(target.OrderId, target, 5*time.Minute)
	log.Println("Print data in cache")
	log.Println(cache)
	return nil
}
