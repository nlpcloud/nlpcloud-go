//go:build integration
// +build integration

package integration_test

import (
	"fmt"
	"os"
	"testing"
)

var APIToken string

func TestMain(tm *testing.M) {
	var set bool
	APIToken, set = os.LookupEnv("API_TOKEN")
	if !set {
		fmt.Println("`API_TOKEN` environment variable should be defined in order to achieve integration tests.")
		os.Exit(1)
	}

	os.Exit(tm.Run())
}
