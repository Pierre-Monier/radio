package main

import (
	"fmt"
	"io"
	"os"
)

// update update playlist/musics if possible, else it return a playlist map
func update(radioDir string) ([]map[string]interface{}, error) {
	playlists, err := getStuffsFromAPI("https://api-f-fstreaming.p-monier.fr/api/playlists")
	if err != nil {
		fmt.Println("Can't get the latest playlists from API")
		playlists, err = getFileSystemPlaylists(radioDir)
		if err != nil {
			return nil, err
		}
	} else {
		updatePlaylists(radioDir, playlists)
	}
	updateMusicsInPlaylist(radioDir, playlists)
	return playlists, nil
}

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

func updateMusicsInPlaylist(radioDir string, playlists []map[string]interface{}) {
	fmt.Println("Updating the playlists...")
	for _, pv := range playlists {
		musics, err := getStuffsFromAPI("https://api-f-fstreaming.p-monier.fr/api/playlists/" + pv["name"].(string))
		if err == nil {
			for _, v := range musics {
				filepath := radioDir + pv["name"].(string) + "/" + v["name"].(string)
				_, err := os.Stat(filepath)
				if os.IsNotExist(err) {
					fmt.Println("Uploading: " + v["name"].(string) + " in playlist: " + pv["name"].(string))
					music, err := getFilesFromDropbox(v["path"].(string))
					if err != nil {
						fmt.Println(err)
					} else {
						// Create the file
						defer music.Body.Close()
						out, err := os.Create(filepath)
						if err != nil {
							fmt.Println("Create")
							fmt.Println(err)
						}
						// write the body resp into it
						_, err = io.Copy(out, music.Body)
						if err != nil {
							fmt.Println("Copy")
							fmt.Println(err)
						}
					}
				}
			}
			fmt.Println("Updated")
		} else {
			fmt.Println(err)
		}
	}
}
