package handler

import (
	"net/http"

	"github.com/Happy-Fy/redirect-service/internal/config"
)

type FallbackHandler struct {
}

func (h *FallbackHandler) Name() string {
	return config.TypeFallback
}

func (h *FallbackHandler) Match(r *http.Request, rule config.Rule) bool {
	_ = rule
	return true
}
