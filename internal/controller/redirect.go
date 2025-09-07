package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/Happy-Fy/redirect-service/internal/config"
	"github.com/Happy-Fy/redirect-service/internal/handler"
)

type RedirectController struct {
	Config  *config.Config
	Handler map[string]handler.Handler
}

func NewRedirectHandler(cnf *config.Config) *RedirectController {
	return &RedirectController{
		Config:  cnf,
		Handler: handler.RedirectHandlers(),
	}
}

func (rc *RedirectController) Handle(w http.ResponseWriter, r *http.Request) {
	for _, rule := range rc.Config.Rules {
		for k, h := range rc.Handler {
			if k == rule.Type && h.Match(r, rule) {
				target := rc.replacePlaceholder(rule.Target, rc.placeholder(r))
				http.Redirect(w, r, target, http.StatusFound)
				return
			}
		}
	}

	if rc.Config.Fallback.Target != "" {
		target := rc.replacePlaceholder(rc.Config.Fallback.Target, rc.placeholder(r))
		http.Redirect(w, r, target, http.StatusFound)
		return
	}

	if _, err := w.Write([]byte("No matching redirect rule found")); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
func (rc *RedirectController) placeholder(r *http.Request) map[string]string {
	ph := map[string]string{}

	ph["{path}"] = strings.TrimPrefix(r.URL.Path, "/")

	return ph
}

func (rc *RedirectController) replacePlaceholder(target string, ph map[string]string) string {
	for k, v := range ph {
		target = strings.ReplaceAll(target, k, v)
	}
	return target
}
