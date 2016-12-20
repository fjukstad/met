package met

import "encoding/json"

func GetAllLocations() ([]Data, error) {
	u := baseUrl + "/locations/v0.jsonld"
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
