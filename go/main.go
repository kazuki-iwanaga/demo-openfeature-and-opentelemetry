package main

import (
	"context"
	"fmt"

	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	// Register your feature flag provider
	openfeature.SetProvider(openfeature.NoopProvider{})
	// Create a new client
	client := openfeature.NewClient("app")
	// Evaluate your feature flag
	v2Enabled, _ := client.BooleanValue(
		context.Background(), "v2_enabled", true, openfeature.EvaluationContext{},
	)
	// Use the returned flag value
	if v2Enabled {
		fmt.Println("v2 is enabled")
	}
}
