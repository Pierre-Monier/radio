package rweb

import (
	"net/http"
	"encoding/json"
)

// GetStuffsFromAPI return data from api
func GetStuffsFromAPI(url string) ([]map[string]interface{}, error) {
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

// GetFilesFromDropbox return files (music file) from dropbox
func GetFilesFromDropbox(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

