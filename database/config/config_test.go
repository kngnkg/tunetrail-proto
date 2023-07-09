package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	wantPort := 9999
	t.Setenv("TUNETRAIL_DB_PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	assert.Equal(t, wantPort, got.DBPort)
	wantEnv := "dev"
	assert.Equal(t, wantEnv, got.Env)
}
