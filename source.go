package met

import "encoding/json"

func GetAllSources() ([]Data, error) {
	endpoint := baseUrl + "/sources/v0.jsonld"

	body, err := get(endpoint)
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
