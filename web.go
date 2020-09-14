package main

import (
	"encoding/json"
	"net/http"
)

// getStuffsFromAPI return data from api
func getStuffsFromAPI(url string) ([]map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// parse resposne
	var body []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// getFilesFromDropbox return files (music file) from dropbox
func getFilesFromDropbox(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
