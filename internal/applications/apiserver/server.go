package apiserver

import (
	"blog/internal/applications/store"
	"blog/internal/applications/user"
	userRepository "blog/internal/applications/user/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	s.logger.Info("route registration successful")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	repository := userRepository.NewRepository(s.store, s.logger)
	userHandler := user.NewHandler(s.logger, repository)
	userHandler.Register(s.router)
}
