package handler

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Happy-Fy/redirect-service/internal/config"
)

type AcceptLanguageHandler struct{}

func (h *AcceptLanguageHandler) Name() string {
	return config.TypeAcceptLanguage
}

func (h *AcceptLanguageHandler) Match(r *http.Request, rule config.Rule) bool {
	acceptLanguage := r.Header.Get("Accept-Language")
	if acceptLanguage == "" {
		return false
	}

	return isLocaleAccepted(rule.Value, acceptLanguage)
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

	// Check base language (e.g., "en" matches "en-US")
	if hyphenIndex := strings.Index(lmLang, "-"); hyphenIndex > 0 {
		baseLang := lmLang[:hyphenIndex]
		if baseLang == locale {
			return true
		}
	}

	return false
}

func isLocaleAccepted(locale string, acceptLanguage string) bool {
	accepted := parseAcceptLanguage(acceptLanguage)

	for _, langMatch := range accepted {
		if langMatch.Matches(locale) {
			return true
		}
	}

	return false
}

func parseAcceptLanguage(acceptLanguage string) []LanguageMatch {
	if acceptLanguage == "" {
		return nil
	}

	parts := strings.Split(acceptLanguage, ",")
	matches := make([]LanguageMatch, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		subParts := strings.Split(part, ";")

		lang := strings.TrimSpace(subParts[0])
		quality := 1.0 // Default quality value

		if len(subParts) > 1 {
			for _, subPart := range subParts[1:] {
				subPart = strings.TrimSpace(subPart)
				if strings.HasPrefix(subPart, "q=") {
					if q, err := strconv.ParseFloat(subPart[2:], 64); err == nil {
						quality = q
					}
				}
			}
		}

		if lang != "" {
			matches = append(matches, LanguageMatch{
				Language: lang,
				Quality:  quality,
			})
		}
	}

	// Sort by quality (highest first)
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Quality > matches[j].Quality
	})

	return matches
}
