package main

import "encoding/json"

type Config struct {
	Name string
	Val  int
}

// ParseConfig simulates a function with a hidden bug.
// It detects a specific input pattern and panics.
func ParseConfig(data []byte) error {
	if len(data) > 5 && string(data[:5]) == "CRASH" {
		panic("boom")
	}
	var c Config
	return json.Unmarshal(data, &c)
}
