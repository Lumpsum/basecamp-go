package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	BasecampUI = "https://3.basecampapi.com/4535691"
)

func getProjects() ([]Project, error) {
	f, err := os.Open("data/projects.json")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading io: %s", err)
	}

	var projects []Project
	if err := json.Unmarshal(b, &projects); err != nil {
		return nil, fmt.Errorf("error unmarshaling projects: %s", err)
	}

	return projects, nil
}

func sendRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects.json", BasecampUI), nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}

	req.Header.Add("User-Agent", "basecamp-go (rvergunst@xccelerated.io)")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", TOKEN))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	return body, nil
}
