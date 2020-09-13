package main

import (
	"fmt"
	"io"
	"os"
)

func updatePlaylists(radioDir string, playlists []map[string]interface{}) {
	// test if their is an existing dir for each playlist, if not, create one
	for _, v := range playlists {
		filepath := radioDir + v["name"].(string)
		noDir := true
		fileinfo, err := os.Stat(filepath)
		// if no dir, create one, if file is a file (not a dir), remove file and create a dir
		if os.IsNotExist(err) {
			fmt.Println("Create the playlist " + v["name"].(string))
		} else if !fileinfo.IsDir() {
			os.Remove(filepath)
		} else if fileinfo.IsDir() {
			noDir = false
		}
		if noDir {
			os.Mkdir(filepath, 0774)
		}
	}
}

func updateChoosenPlaylist(playlistData map[string]string) {
	fmt.Println("Updating the playlist...")

	musics, err := getStuffsFromAPI("https://api-f-fstreaming.p-monier.fr/api/playlists/" + playlistData["dir"])
	if err == nil {
		for _, v := range musics {
			filepath := playlistData["fullpath"] + "/" + v["name"].(string)
			_, err := os.Stat(filepath)
			if os.IsNotExist(err) {
				fmt.Println("Uploading: " + v["name"].(string))
				music, err := getFilesFromDropbox(v["path"].(string))
				if err != nil {
					fmt.Println(err)
				} else {
					// Create the file
					out, err := os.Create(filepath)
					if err != nil {
						fmt.Println(err)
					}
					// write the body resp into it
					_, err = io.Copy(out, music.Body)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
		fmt.Println("Updated")
	}
}
