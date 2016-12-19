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
	_, err := met.GetAllSources()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllElements(t *testing.T) {
	_, err := met.GetAllElements()
	if err != nil {
		t.Error(err)
	}
}
