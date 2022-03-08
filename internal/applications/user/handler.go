package user

import (
	"blog/internal/applications/handlers"
	"blog/internal/applications/model"
	"blog/pkg/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	baseUser    = "/api/users/"
	tokenCreate = "/api/auth/"
)

type handler struct {
	logger     *logrus.Logger
	repository UserRepository
	usecase    auth.UseCase
}

func NewHandler(logger *logrus.Logger, repository UserRepository, usecase auth.UseCase) handlers.Handler {
	return &handler{
		logger:     logger,
		repository: repository,
		usecase:    usecase,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(baseUser, h.handlerUsersCreate).Methods("POST", "OPTIONS")
	router.HandleFunc(tokenCreate, h.handlerTokenCreate).Methods("POST", "OPTIONS")
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

	u := &model.User{
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

func (h *handler) handlerTokenCreate(w http.ResponseWriter, r *http.Request) {
	req := new(model.User)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.repository.FindByEmail(req.Email)
	if err != nil || !u.ComparePassword(req.Password) {
		h.error(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.usecase.SignIn(req)

	if err != nil {
		if err == auth.ErrInvalidAccessToken {
			h.error(w, http.StatusBadRequest, err)
			return
		}

		if err == auth.ErrUserDoesNotExist {
			h.error(w, http.StatusBadRequest, err)
			return
		}

		h.error(w, http.StatusBadRequest, err)
		return
	}

	h.respond(w, http.StatusOK, token)
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
