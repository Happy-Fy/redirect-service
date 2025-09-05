package handler

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/Happy-Fy/redirect-service/internal/config"
)

type AcceptLanguageHandler struct {
}

func (h *AcceptLanguageHandler) Name() string {
	return config.TypeAcceptLanguage
}

func (h *AcceptLanguageHandler) Match(r *http.Request, rule config.Rule) bool {
	acceptedLanguages := r.Header["Accept-Language"]
	if acceptedLanguages == nil || len(acceptedLanguages) == 0 {
		return false
	}
	acceptedLanguage := acceptedLanguages[0]

	isAccepted, err := isLocaleAccepted(rule.Value, acceptedLanguage)
	if err != nil {
		return false
	}

	if isAccepted {
		return true
	}

	return false
}

type LanguageMatch struct {
	Language string
	Quality  float64
}

func (lm LanguageMatch) Matches(locale string) bool {

	lmLang := strings.ToLower(lm.Language)
	locale = strings.ToLower(locale)

	if lmLang == locale {
		return true
	}

	if strings.Contains(lmLang, "-") {
		baseLang := strings.Split(lmLang, "-")[0]
		if baseLang == locale {
			return true
		}
	}

	return false
}

func isLocaleAccepted(locale string, acceptLanguage string) (bool, error) {
	accepted, err := parseAcceptLanguage(acceptLanguage)
	if err != nil {
		return false, err
	}

	for _, langMatch := range accepted {
		if langMatch.Matches(locale) {
			return true, nil
		}
	}

	return false, nil
}

func parseAcceptLanguage(acceptLanguage string) ([]LanguageMatch, error) {
	if acceptLanguage == "" {
		return nil, nil
	}

	parts := strings.Split(acceptLanguage, ",")
	matches := make([]LanguageMatch, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		subParts := strings.Split(part, ";")

		lang := subParts[0]
		quality := 1.0 // Default quality value

		if len(subParts) > 1 {
			qPart := subParts[1]
			if strings.HasPrefix(qPart, "q=") {
				fmt.Sscanf(qPart, "q=%f", &quality)
			}
		}

		matches = append(matches, LanguageMatch{
			Language: lang,
			Quality:  quality,
		})
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Quality > matches[j].Quality
	})

	return matches, nil
}
