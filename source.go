package met

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetSources() ([]Data, error) {
	endpoint := baseUrl + "/sources/v0.jsonld"
	id, err := getClientId()
	if err != nil {
		return []Data{}, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []Data{}, err
	}

	req.SetBasicAuth(id, "")

	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return []Data{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
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
