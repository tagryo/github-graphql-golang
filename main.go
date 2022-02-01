package main

import (
	"fmt"
	"io/ioutil"
	"tagryo/github-graphql-golang/adapters"
)

type RequestBody struct {
	query string
}

func main() {
	const GITHUB_ACCESS_TOKEN = "TODO Replace with your token"
	const GITHUB_ENDPOINT = "https://api.github.com/graphql"

	c, _ := adapters.NewGithubGraphqlClient(GITHUB_ENDPOINT, GITHUB_ACCESS_TOKEN)
	resp, err := c.Execute("query { viewer { login } }")

	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
