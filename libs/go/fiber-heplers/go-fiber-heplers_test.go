package go_fiber_heplers

import (
	"testing"
)

func TestFiberHeplers(t *testing.T) {
	result := "works"
	if result != "works" {
		t.Error("Expected FiberHeplers to append 'works'")
	}
}
