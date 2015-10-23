package main

import (
	yaml "gopkg.in/yaml.v2"
)

type TravisYml struct {
	BeforeInstall []string
	Install       []string
	Script        []string
}

func NewTravisYml(file string) (t *TravisYml) {
	if err := yaml.Unmarshal([]byte(dsata), t); err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	return t
}
