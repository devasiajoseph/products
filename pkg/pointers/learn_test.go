package pointers

import "testing"

func TestPointer(t *testing.T) {
	user := InitializePointer()
	if user.Email != "adoniaromal@gmail.com" {
		t.Errorf("Pointer not initialzed")
	}
}
