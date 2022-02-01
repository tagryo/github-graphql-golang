package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IGithubGraphqlClient interface {
	Execute(query string) (*http.Response, error)
}

type githubGraphqlClient struct {
	client   *http.Client
	endpoint string
	token    string
}

type requestBody struct {
	query string
}

func NewGithubGraphqlClient(endpoint string, githubAccessToken string) (IGithubGraphqlClient, error) {
	client := &http.Client{}
	return &githubGraphqlClient{
		client:   client,
		endpoint: endpoint,
		token:    githubAccessToken,
	}, nil
}

func (c *githubGraphqlClient) Execute(query string) (*http.Response, error) {
	reqBody := &requestBody{
		query: query,
	}
	jsonString, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, nil
}
