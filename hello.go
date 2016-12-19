package met

import (
	"io/ioutil"
	"net/http"
)

func SecureHello() (string, error) {
	endpoint := baseUrl + "/tests/secureHello"
	id, err := getClientId()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(id, "")

	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
