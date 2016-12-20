package met_test

import (
	"testing"

	"github.com/fjukstad/met"
)

func TestHello(t *testing.T) {
	text, err := met.SecureHello()
	if err != nil {
		t.Error(err)
		return
	}

	if text != "Hello to you too, securely!\n" {
		t.Error(text)
		return
	}
}

func TestGetAllSources(t *testing.T) {
	_, err := met.GetAllSources()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetTromsoSources(t *testing.T) {
	f := met.Filter{
		Geometry: "nearest(POINT (18.958503 69.655747))",
	}

	data, err := met.GetSources(f)
	if err != nil {
		t.Error(err)
		return
	}

	if len(data) > 1 {
		t.Error("Multiple sources for tromso")
		return
	}

	tromso := data[0]

	if tromso.Id != "SN90450" {
		t.Error("Tromsø source does not have the correct id")
		return
	}

	// Could also check the other fields as well.

}

func TestGetAllElements(t *testing.T) {
	_, err := met.GetAllElements()
	if err != nil {
		t.Error(err)
		return
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
		return
	}
}

func TestGetAllLocations(t *testing.T) {
	_, err := met.GetAllLocations()
	if err != nil {
		t.Error(err)
		return
	}
}
