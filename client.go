package godataclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type client struct {
	domain string
	apiKey string
	client *http.Client
}

// NewClients creates a client object and check if the information provided are not empty
func NewClient(domain, apiKey string) (*client, error) {
	switch {
	case domain == ``:
		return nil, fmt.Errorf(`error_no_domain`)
	case apiKey == ``:
		return nil, fmt.Errorf(`error_bno_api_key`)
	}

	return &client{
		domain: domain,
		apiKey: apiKey,
		client: &http.Client{},
	}, nil
}

func (c client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(`auth`, c.apiKey)
	req.Header.Add(`Content-Type`, `application/json`)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}

func (c client) post(endpoint string, payload string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add(`auth`, c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
