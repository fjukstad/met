package met

import "encoding/json"

func GetAllSources() ([]Data, error) {
	return GetSources(Filter{})
}

func GetSources(f Filter) ([]Data, error) {
	endpoint := baseUrl + "/sources/v0.jsonld?"

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
