package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getplaylistData(radioDir string, playlists []map[string]interface{}, userChoice string) (map[string]string, error) {
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

// getFileSystemPlaylists return already downloaded playlists if can't acces api
func getFileSystemPlaylists(radioDir string) ([]map[string]interface{}, error) {
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
