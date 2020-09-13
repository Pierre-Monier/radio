package rutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

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

// GetplaylistData return a playlist chosen dir
func GetplaylistData(radioDir string, playlists []map[string]interface{}, userChoice string) (map[string]string, error) {
	i, err := strconv.ParseInt(userChoice, 10, 8)
	if err != nil {
		return nil, err
	}
	playlistData := map[string]string{
		"fullpath": radioDir + playlists[i-1]["name"].(string),
		"dir":      playlists[i-1]["name"].(string),
	}
	return playlistData, nil
}

// GetFileSystemPlaylists return already downloaded playlists if can't acces api
func GetFileSystemPlaylists(radioDir string) ([]map[string]interface{}, error) {
	var playlists []map[string]interface{}
	dir, err := ioutil.ReadDir(radioDir)
	if os.IsNotExist(err) {
		return nil, err
	}
	for _, v := range dir {
		if v.IsDir() {
			fmt.Println(v.Name())
			playlist := make(map[string]interface{})
			playlist["name"] = v.Name()
			playlists = append(playlists, playlist)
		}
	}
	return playlists, nil
}
