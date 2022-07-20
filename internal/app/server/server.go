package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/L0/internal/app/cache"
	"github.com/vladjong/L0/internal/app/store"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	cache  *cache.Cache
}

func New(config *Config, cache *cache.Cache) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		cache:  cache,
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	if err := s.configureCache(); err != nil {
		return err
	}
	s.logger.Info("starting server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	s.router.HandleFunc("/", s.handleMainPage())
	s.router.HandleFunc("/order", s.handleOrderPage())
	http.Handle("/", s.router)
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *Server) configureCache() error {
	tmp, err := s.store.Order().SelectAll()
	if err != nil {
		return err
	}
	for i := 0; i < len(tmp); i++ {
		s.cache.Set(tmp[i].OrderId, tmp[i])
	}
	log.Println("Number of arguments in the cache: ", s.cache.Len())
	return nil
}

func (s *Server) handleMainPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/html/search.html"))
		tmpl.Execute(w, nil)
	}
}

func (s *Server) handleOrderPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderUid := r.FormValue("order_id")
		tmp, err := s.cache.Get(orderUid)
		if !err {
			tmpl := template.Must(template.ParseFiles("templates/html/error.html"))
			tmpl.Execute(w, nil)
		} else {
			tmpl := template.Must(template.ParseFiles("templates/html/order.html"))
			tmpl.Execute(w, tmp)
		}
	}
}
