package post

import (
	"blog/internal/applications/handlers"
	"blog/pkg/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	basePost   = "/api/posts/"
	getPost    = "/api/posts/{id}"
	deletePost = "/api/posts/delete/{id}"
	createPost = "/api/posts/create/"
)

type handler struct {
	logger     *logrus.Logger
	repository PostRepository
}

func NewHandler(logger *logrus.Logger, repository PostRepository) handlers.Handler {
	return &handler{
		logger:     logger,
		repository: repository,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(basePost, h.handlerGetList).Methods("GET", "OPTIONS")
	router.HandleFunc(getPost, h.handlerGetPost).Methods("GET", "OPTIONS")
	router.HandleFunc(createPost, h.handlerCreatePost).Methods("POST", "OPTIONS")
	http.Handle(createPost, auth.AuthMiddleware(router))
	router.HandleFunc(deletePost, h.handlerDeletePost).Methods("DELETE", "OPTIONS")

}

func (h *handler) handlerGetList(writer http.ResponseWriter, request *http.Request) {
	posts, err := h.repository.GetList()
	if err != nil {
		// TODO: handler panic
	}
	h.respond(writer, http.StatusOK, posts)
}

func (h *handler) handlerGetPost(writer http.ResponseWriter, request *http.Request) {

}

func (h *handler) handlerCreatePost(writer http.ResponseWriter, request *http.Request) {
	h.logger.Info("Okey!")
	h.respond(writer, http.StatusOK, "Ok")
}

func (h *handler) handlerDeletePost(writer http.ResponseWriter, request *http.Request) {

}

func (h *handler) error(w http.ResponseWriter, code int, err error) {
	h.respond(w, code, map[string]string{"error": err.Error()})
}

func (h *handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
