package auth

import (
	"fmt"
	"net/http"

	"github.com/ArtDark/go-advanced/internal/config"
)

type AuthHandlerDeps struct {
	*config.Auth
}

type AuthHandler struct {
	*config.Auth
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	h := AuthHandler{
		deps.Auth,
	}
	mux.HandleFunc("GET /auth/login", h.Login())
	mux.HandleFunc("POST /auth/regster", h.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Login\nSecret: %s", h.Secret)))
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Register"))
	}
}
