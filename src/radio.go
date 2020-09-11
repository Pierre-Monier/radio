package main

import(
	"fmt"    
	"net/http"
	"encoding/json"
	"os"
	"log"
	"radio/rutils"
)

func main(){
	log.SetPrefix("radio: ")
	log.SetFlags(0)
	radioDir := "/home/yoon/Radio/"

	playlists, err := getPlaylistFromDropbox(radioDir)
	if err != nil {
		log.Fatal(err)
	}
	createDirForPlaylist(radioDir, playlists)
	fChoices := rutils.FormatUserChoice(playlists)
	userChoice := getUserChoice(fChoices, len(playlists))
	fmt.Println(userChoice)
}

func getPlaylistFromDropbox(radioDir string) ([]map[string]interface{}, error) {
	// get all the playlist from api
	res, err := http.Get("https://api-f-fstreaming.p-monier.fr/api/playlists")
	if err != nil {
		fmt.Println("Something got wrong in the api call")
	}
	defer res.Body.Close()
	// parse resposne 

	var body []map[string]interface {}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		fmt.Println("Something got wrong in the api call")
		return nil, err
	}

	return body, nil
}

func createDirForPlaylist(radioDir string, playlists []map[string]interface{}){
    // test if their is an existing dir for each playlist, if not, create one
	for _, v := range playlists {
		filepath := radioDir+v["name"].(string)
		noDir := false
		fileinfo, err := os.Stat(filepath)
		// if no dir, create one, if file is a file (not a dir), remove file and create a dir
		if os.IsNotExist(err) {
			noDir = true
		}else if !fileinfo.IsDir() {
			os.Remove(filepath)
			noDir = true
		}

		if noDir {
			os.Mkdir(filepath, 0664)
		}
	}
}

func getUserChoice(fChoices string, pLength int) string {
	fmt.Println(pLength)
	fmt.Println("Select the playlist \n"+fChoices)
	match := false
	var userResponse string
	for !match {
		fmt.Scanln(&userResponse)
		match = rutils.ValideUserResponse(userResponse, int64(pLength))
	}
	return userResponse
}
