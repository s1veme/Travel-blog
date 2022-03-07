package apiserver

import (
	"blog/internal/applications/store"
	"blog/internal/applications/user"
	userRepository "blog/internal/applications/user/db"
	"blog/pkg/auth"
	"blog/pkg/auth/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type server struct {
	router      *mux.Router
	logger      *logrus.Logger
	store       store.Store
	authUseCase auth.UseCase
}

func newServer(store store.Store, config *Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	urepository := userRepository.NewRepository(s.store, s.logger)

	authUseCase := usecase.NewAuthorizer(
		urepository,
		config.HashSalt,
		[]byte(config.SigningKey),
		config.TokenTtl*time.Second,
	)
	s.authUseCase = authUseCase
	s.configureRouter(urepository)

	s.logger.Info("route registration successful")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter(urepository user.UserRepository) user.UserRepository {
	userHandler := user.NewHandler(s.logger, urepository, s.authUseCase)
	userHandler.Register(s.router)

	return urepository
}
