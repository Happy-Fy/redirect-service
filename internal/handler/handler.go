package handler

import (
	"net/http"

	"github.com/Happy-Fy/redirect-service/internal/config"
)

type Handler interface {
	Name() string
	Match(r *http.Request, rule config.Rule) bool
}

func RedirectHandlers() map[string]Handler {
	h := &[]Handler{
		&AcceptLanguageHandler{},
		&FallbackHandler{},
	}

	m := make(map[string]Handler)
	for _, handler := range *h {
		m[handler.Name()] = handler
	}
	return m
}
