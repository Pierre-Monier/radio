package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"radio/rutils"
)

func main() {
	log.SetPrefix("radio: ")
	log.SetFlags(0)
	radioDir := "/home/yoon/Radio/"

	playlists, err := rutils.GetStuffsFromApi("https://api-f-fstreaming.p-monier.fr/api/playlists")
	// playlists, err := getPlaylistFromDropbox(radioDir)
	if err != nil {
		log.Fatal(err)
	}
	createDirForPlaylist(radioDir, playlists)
	fChoices := rutils.FormatUserChoice(playlists)
	userChoice := getUserChoice(fChoices, len(playlists))
	playlistDir := rutils.GetPlaylistDir(radioDir, playlists, userChoice)
	updateChoosenPlaylist(playlistDir, userChoice)
}

func createDirForPlaylist(radioDir string, playlists []map[string]interface{}) {
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

func getUserChoice(fChoices string, pLength int) string {
	fmt.Println("Select the playlist \n" + fChoices)
	match := false
	var userResponse string
	for !match {
		fmt.Scanln(&userResponse)
		match = rutils.ValideUserResponse(userResponse, int64(pLength))
	}
	return userResponse
}

func updateChoosenPlaylist(playlistDir string, userChoice string) error {
	fmt.Println("Updating the playlist...")

	musics, err := rutils.GetStuffsFromApi("https://api-f-fstreaming.p-monier.fr/api/playlists/" + userChoice)
	if err == nil {
		for _, v := range musics {
			filepath := playlistDir + "/" + v["name"].(string)
			_, err := os.Stat(filepath)
			if os.IsNotExist(err) {
				fmt.Println("Uploading: " + v["name"].(string))
				music, err := rutils.GetFilesFromDropbox(v["path"].(string))
				if err != nil {
					fmt.Println(err)
				}
				// Create the file
				out, err := os.Create(filepath)
				if err != nil {
					fmt.Println(err)
				}
				// write the body resp into it
				_, err = io.Copy(out, music.Body)
			}
		}
		fmt.Println("Updated")
	}
	return err
}
