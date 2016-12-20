package met

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Filter struct {
	Sources               []string
	ReferenceTime         string
	Elements              []string
	PerformanceCategories []string
	ExposureCategories    []string
	Fields                []string
}

func GetObservations(f Filter) ([]Data, error) {
	endpoint := baseUrl + "/observations/v0.jsonld?"

	u := createUrl(endpoint, f)

	body, err := get(u)
	if err != nil {
		return []Data{}, err
	}
	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Data{}, err
	}

	return response.Data, nil

}

func createUrl(endpoint string, f Filter) string {

	sources := strings.Join(f.Sources, ",")
	elements := strings.Join(f.Elements, ",")
	performanceCategories := strings.Join(f.PerformanceCategories, ",")
	exposureCategories := strings.Join(f.ExposureCategories, ",")
	fields := strings.Join(f.Fields, ",")

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

	return endpoint + v.Encode()
}
