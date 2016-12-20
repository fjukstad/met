package met

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	Country               string `json:"country"`
	SourceID              string `json:"sourceId"`
	Geometry              `json:"geometry"`
	Levels                []Level       `json:"levels"`
	ReferenceTime         time.Time     `json:"referenceTime"`
	Observations          []Observation `json:"observations"`
	ValidFrom             string        `json:"validFrom"`
	LegacyMetNoConvention `json:"legacyMetNoConvention"`
	CfConvention          `json:"cfConvention"`
}

type CfConvention struct {
	StandardName string `json:"standardName"`
	Unit         string `json:"unit"`
	Status       string `json:"status"`
}

type LegacyMetNoConvention struct {
	ElemCodes []string `json:"elemCodes"`
	Category  string   `json:"category"`
	Unit      string   `json:"unit"`
}

type Level struct {
	LevelType string `json:"levelType"`
	Value     int    `json:"value"`
	Unit      string `json:"unit"`
}

type Observation struct {
	ElementID           string  `json:"elementId"`
	Value               float64 `json:"value"`
	Unit                string  `json:"unit"`
	CodeTable           string  `json:"codeTable"`
	PerformanceCategory string  `json:"performanceCategory"`
	ExposureCategory    string  `json:"exposureCategory"`
	QualityCode         int     `json:"qualityCode"`
	DataVersion         string  `json:"dataVersion"`
}

type Geometry struct {
	Type         string    `json:"@type"`
	Coordinates  []float64 `json:"coordinates"`
	Interpolated bool      `json:"interpolated"`
}

type Filter struct {
	Sources               []string
	ReferenceTime         string
	Elements              []string
	PerformanceCategories []string
	ExposureCategories    []string
	Fields                []string
	Ids                   []string
	Types                 []string
	Geometry              string
	ValidTime             string
}

var baseUrl = "https://data.met.no"

func getClientId() (string, error) {
	id := os.Getenv("CLIENT_ID")
	if id == "" {
		return "", errors.New("CLIENT_ID not set")
	}
	return id, nil
}

func get(endpoint string) ([]byte, error) {
	id, err := getClientId()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(id, "")

	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func createUrl(endpoint string, f Filter) string {

	sources := strings.Join(f.Sources, ",")
	elements := strings.Join(f.Elements, ",")
	performanceCategories := strings.Join(f.PerformanceCategories, ",")
	exposureCategories := strings.Join(f.ExposureCategories, ",")
	fields := strings.Join(f.Fields, ",")
	ids := strings.Join(f.Fields, ",")
	types := strings.Join(f.Types, ",")

	v := url.Values{}

	if len(sources) > 0 {
		v.Add("sources", sources)
	}

	if len(elements) > 0 {
		v.Add("elements", elements)
	}

	if f.ReferenceTime != "" {
		v.Add("referencetime", f.ReferenceTime)
	}

	if len(performanceCategories) > 0 {
		v.Add("performancecategory", performanceCategories)
	}

	if len(exposureCategories) > 0 {
		v.Add("exposurecategory", exposureCategories)
	}

	if len(fields) > 0 {
		v.Add("fields", fields)
	}

	if len(ids) > 0 {
		v.Add("ids", ids)
	}

	if len(types) > 0 {
		v.Add("types", types)
	}

	if f.Geometry != "" {
		v.Add("geometry", f.Geometry)
	}

	if f.ValidTime != "" {
		v.Add("validtime", f.ValidTime)
	}

	return endpoint + v.Encode()
}
