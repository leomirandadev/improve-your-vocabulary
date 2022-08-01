package hasher

import (
	"testing"
)

func TestGenerateAndCompare(t *testing.T) {
	h := NewBcryptHasher()

	firstPassword := "password"
	secondPassword := "password2"

	password1Hash, err := h.Generate(firstPassword)
	if err != nil {
		t.Error(err)
	}

	if err := h.Compare(password1Hash, firstPassword); err != nil {
		t.Error(err)
	}

	if err := h.Compare(password1Hash, secondPassword); err == nil {
		t.Error("password can't be equal")
	}
}
