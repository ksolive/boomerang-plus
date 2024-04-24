package config

import (
	"testing"
)

func BenchmarkGetConfigFromFile(b *testing.B) {
	// TODO 改变filename
	filename := "/path/to/config.json"
	for i := 0; i < b.N; i++ {
		_, _ = getConfigFromFile(filename)
	}
}

func BenchmarkGetConfigFromArgs(b *testing.B) {
	args := []string{"arg1", "arg2", "arg3"}
	for i := 0; i < b.N; i++ {
		_, _ = getConfigFromArgs(args)
	}
}

func BenchmarkGetConfigFromUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getConfigFromUser()
	}
}
