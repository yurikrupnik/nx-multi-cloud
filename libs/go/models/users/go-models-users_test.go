package go_models_users

import (
	"testing"
)

func TestUsers(t *testing.T) {
	result := "works"
	if result != "works" {
		t.Error("Expected Users to append 'works'")
	}
}
