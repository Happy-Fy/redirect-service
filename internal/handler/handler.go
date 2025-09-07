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
	handlers := []Handler{
		&AcceptLanguageHandler{},
		&FallbackHandler{},
	}

	handlerMap := make(map[string]Handler, len(handlers))
	for _, handler := range handlers {
		handlerMap[handler.Name()] = handler
	}
	return handlerMap
}
