package rutils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetStuffsFromApi return data from api
func GetStuffsFromApi(url string) ([]map[string]interface{}, error) {
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

// FormatUserChoice return a clean playlist selection
func FormatUserChoice(playlists []map[string]interface{}) string {
	i := 1
	res := ""
	for _, v := range playlists {
		res += strconv.Itoa(i) + ")" + v["name"].(string) + "\n"
		i++
	}
	return res
}

// ValideUserResponse check if the user response (of the playlist choice) is correct
func ValideUserResponse(userResponse string, pLenght int64) bool {
	nbr, err := strconv.ParseInt(userResponse, 10, 8)
	if err != nil {
		return false
	} else if nbr <= pLenght && nbr > 0 {
		return true
	} else {
		return false
	}
}

// GetPlaylistDir return a playlist chosen dir
func GetPlaylistDir(radioDir string, playlists []map[string]interface{}, userChoice string) string {
	i, err := strconv.ParseInt(userChoice, 10, 8)
	if err != nil {
		return ""
	} else {
		return radioDir + playlists[i-1]["name"].(string)
	}
}
