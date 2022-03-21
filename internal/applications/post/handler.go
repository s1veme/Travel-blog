package post

import (
	"blog/internal/applications/handlers"
	"blog/internal/applications/model"
	"blog/internal/applications/user"
	"blog/pkg/auth"
	"encoding/json"
	"fmt"
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
	logger         *logrus.Logger
	repository     PostRepository
	signingKey     string
	userRepository user.UserRepository
}

func NewHandler(logger *logrus.Logger, repository PostRepository,
	userRepository user.UserRepository, signingKey string) handlers.Handler {
	return &handler{
		logger:         logger,
		repository:     repository,
		signingKey:     signingKey,
		userRepository: userRepository,
	}
}

func (h *handler) Register(router *mux.Router) {
	authSubRouter := router.NewRoute().Subrouter()
	authMiddleware := auth.Register(h.signingKey)
	authSubRouter.Use(authMiddleware.Middleware)
	authSubRouter.HandleFunc(createPost, h.handlerCreatePost).Methods("POST", "OPTIONS")

	router.HandleFunc(basePost, h.handlerGetList).Methods("GET", "OPTIONS")
	router.HandleFunc(getPost, h.handlerGetPost).Methods("GET", "OPTIONS")
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

func (h *handler) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	email := r.Context().Value("user")
	user, _ := h.userRepository.FindByEmail(fmt.Sprintf("%v", email))

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		Owner:   user.ID,
	}

	err := h.repository.Create(post)
	if err != nil {
		h.respond(w, http.StatusBadRequest, err)
		return
	}
	h.respond(w, http.StatusCreated, post)
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
