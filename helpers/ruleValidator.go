package helpers

import (
	"log"
	"os"
	"regexp"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	"gopkg.in/yaml.v3"
)

type RuleConfig struct {
	Name                string `yaml:"name"`
	Implementation      string `yaml:"implementation"`
	ValidMatchedRegex   string `yaml:"validMatchedRegex"`
	InvalidMatchedRegex string `yaml:"invalidMatchedRegex"`
}

var rules []func(string) bool

func LoadRulesFromConfig() {
	// Get configs from yml file
	f, err := os.ReadFile("./configs/validationRules.yml")
	if err != nil {
		log.Fatal(err)
	}
	var ruleConfigs []RuleConfig
	if err := yaml.Unmarshal(f, &ruleConfigs); err != nil {
		log.Fatal(err)
	}

	// Implement configs
	for _, ruleConfig := range ruleConfigs {
		if ruleConfig.Implementation != "" {
			i := interp.New(interp.Options{})
			i.Use(stdlib.Symbols)
			_, err := i.Eval(ruleConfig.Implementation)
			if err != nil {
				log.Fatal(err)
			}

			v, err := i.Eval("Check")
			if err != nil {
				log.Fatal(err)
			}

			check := v.Interface().(func(string) bool)
			rules = append(rules, check)
			continue
		}
		if ruleConfig.ValidMatchedRegex != "" {
			re := regexp.MustCompile(ruleConfig.ValidMatchedRegex)
			check := func(pin string) bool {
				return re.MatchString(pin)
			}
			rules = append(rules, check)
			continue
		}
		if ruleConfig.InvalidMatchedRegex != "" {
			re := regexp.MustCompile(ruleConfig.InvalidMatchedRegex)
			check := func(pin string) bool {
				return !re.MatchString(pin)
			}
			rules = append(rules, check)
			continue
		}
	}
}

func ValidatePin(pin string) bool {
	for _, rule := range rules {
		if !rule(pin) {
			return false
		}
	}
	return true
}
