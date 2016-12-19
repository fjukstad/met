package met

import (
	"errors"
	"os"
	"time"
)

type Response struct {
	Context          string    `json:"@context"`
	Type             string    `json:"@type"`
	APIVersion       string    `json:"apiVersion"`
	License          string    `json:"license"`
	CreatedAt        time.Time `json:"createdAt"`
	QueryTime        float64   `json:"queryTime"`
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
	Id            string `json:"id"`
	Name          string `json:"name"`
	Country       string `json:"country"`
	SourceID      string `json:"sourceId"`
	Geometry      `json:"geometry"`
	Levels        []Level       `json:"levels"`
	ReferenceTime time.Time     `json:"referenceTime"`
	Observations  []Observation `json:"observations"`
	ValidFrom     string        `json:"validFrom"`
}

type Level struct {
	LevelType string `json:"levelType"`
	Value     int    `json:"value"`
	Unit      string `json:"unit"`
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
	Type         string    `json:"@type"`
	Coordinates  []float64 `json:"coordinates"`
	Interpolated bool      `json:"interpolated"`
}

var baseUrl = "https://data.met.no"

func getClientId() (string, error) {
	id := os.Getenv("CLIENT_ID")
	if id == "" {
		return "", errors.New("CLIENT_ID not set")
	}
	return id, nil
}
