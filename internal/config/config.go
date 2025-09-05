package config

import (
	"encoding/json"
	"errors"
	"os"
)

type (
	Config struct {
		Rules []Rule `yaml:"rules"`
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
	rules := cnf.Rules
	// simple bubble sort
	for i := 0; i < len(rules); i++ {
		for j := 0; j < len(rules)-i-1; j++ {
			if rules[j].Prio > rules[j+1].Prio {
				rules[j], rules[j+1] = rules[j+1], rules[j]
			}
		}
	}
	return rules
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
