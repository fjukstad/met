package met

import "encoding/json"

func GetAllElements() ([]Data, error) {
	endpoint := baseUrl + "/elements/v0.jsonld"
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
