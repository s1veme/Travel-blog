package user

import (
	"blog/internal/applications/handlers"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	baseUser = "/api/users/"
)

type handler struct {
	logger     *logrus.Logger
	repository UserRepository
}

func NewHandler(logger *logrus.Logger, repository UserRepository) handlers.Handler {
	return &handler{
		logger:     logger,
		repository: repository,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(baseUser, h.handlerUsersCreate).Methods("POST")
}

func (h *handler) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	u := &User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.repository.Create(u); err != nil {
		h.error(w, http.StatusUnprocessableEntity, err)
		return
	}

	u.Sanitize()
	h.respond(w, http.StatusCreated, u)
}

// TODO: delete hardcode
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
