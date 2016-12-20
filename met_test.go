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

func TestGetObservations(t *testing.T) {
	f := met.Filter{
		Sources:       []string{"SN18700"},
		ReferenceTime: "2005-07-01T00:00/2006-07-01T00:00",
		Elements:      []string{"max(air_temperature 1M)"},
	}

	_, err := met.GetObservations(f)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllLocations(t *testing.T) {
	_, err := met.GetAllLocations()
	if err != nil {
		t.Error(err)
	}
}
