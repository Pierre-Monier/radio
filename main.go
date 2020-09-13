package main

import (
	"fmt"
)

func main() {
	radioDir := "/home/yoon/Radio/"

	playlists, err := getStuffsFromAPI("https://api-f-fstreaming.p-monier.fr/api/playlists")
	if err != nil {
		fmt.Println(err)
		// getting already created playlists
		playlists, err = getFileSystemPlaylists(radioDir)
		if err != nil {
			panic("Can't find any playlist")
		}
	} else {
		updatePlaylists(radioDir, playlists)
	}
	// fChoices := formatUserChoice(playlists)
	userChoice := getUserChoice(formatUserChoice(playlists), len(playlists))
	playlistData, err := getplaylistData(radioDir, playlists, userChoice)
	if err != nil {
		panic("Can't get the playlist to play")
	}

	updateChoosenPlaylist(playlistData)
}
