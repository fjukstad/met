package met

import "encoding/json"

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
