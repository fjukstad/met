package met_test

import (
	"testing"

	"github.com/fjukstad/met"
)

func TestHello(t *testing.T) {
	text, err := met.SecureHello()
	if err != nil {
		t.Error(err)
	}

	if text != "Hello to you too, securely!\n" {
		t.Error(text)
	}
}

func TestGetAllSources(t *testing.T) {
	_, err := met.GetSources()
	if err != nil {
		t.Error(err)
	}
}
