package config

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
)

type (
	Config struct {
		Rules    []Rule `yaml:"rules"`
		Fallback Rule   `yaml:"fallback"`
	}
	Rule struct {
		Prio   int    `yaml:"prio"`
		Type   string `yaml:"type"`
		Value  string `yaml:"value"`
		Target string `yaml:"target"`
	}
)

func Configs() (*Config, error) {
	var cnf Config
	cnfFile := getEnv(EnvRedirectConfigFile, "")
	cnfStr := getEnv(EnvRedirectConfig, "")
	if cnfFile != "" {
		data, err := os.ReadFile(cnfFile)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &cnf); err != nil {
			return nil, err
		}
	} else if cnfStr != "" {
		err := json.Unmarshal([]byte(cnfStr), &cnf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no config found")
	}

	rules := cnf.orderedRules()
	cnf.Rules = rules

	return &cnf, nil
}

func (cnf *Config) orderedRules() []Rule {
	rules := make([]Rule, len(cnf.Rules))
	copy(rules, cnf.Rules)

	// Use Go's built-in sort instead of bubble sort for O(n log n) performance
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Prio > rules[j].Prio
	})

	return rules
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
