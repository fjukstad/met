package met

import (
	"errors"
	"os"
	"time"
)

type AutoGenerated struct {
	Context          string    `json:"@context"`
	Type             string    `json:"@type"`
	APIVersion       string    `json:"apiVersion"`
	License          string    `json:"license"`
	CreatedAt        time.Time `json:"createdAt"`
	QueryTime        string    `json:"queryTime"`
	CurrentItemCount int       `json:"currentItemCount"`
	ItemsPerPage     int       `json:"itemsPerPage"`
	Offset           int       `json:"offset"`
	TotalItemCount   int       `json:"totalItemCount"`
	NextLink         string    `json:"nextLink"`
	PreviousLink     string    `json:"previousLink"`
	CurrentLink      string    `json:"currentLink"`
	Data             []Data    `json:"data"`
}

type Data struct {
	SourceID      string `json:"sourceId"`
	Geometry      `json:"geometry"`
	Levels        string        `json:"levels"`
	ReferenceTime time.Time     `json:"referenceTime"`
	Observations  []Observation `json:"observations"`
}

type Observation struct {
	ElementID           string `json:"elementId"`
	Value               string `json:"value"`
	Unit                string `json:"unit"`
	CodeTable           string `json:"codeTable"`
	PerformanceCategory string `json:"performanceCategory"`
	ExposureCategory    string `json:"exposureCategory"`
	QualityCode         string `json:"qualityCode"`
	DataVersion         string `json:"dataVersion"`
}

type Geometry struct {
	Type         string `json:"@type"`
	Coordinates  string `json:"coordinates"`
	Interpolated bool   `json:"interpolated"`
}

var baseUrl = "https://data.met.no"

func getClientId() (string, error) {
	id := os.Getenv("CLIENT_ID")
	if id == "" {
		return "", errors.New("CLIENT_ID not set")
	}
	return id, nil
}
