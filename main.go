package main

func main() {
	radioDir := "/home/yoon/Radio/"

	playlists, err := update(radioDir)
	if err != nil {
		panic("Can't find any playlist")
	}
	userChoice, err := getUsefullUserChoice(playlists)
	if err != nil {
		panic("Can't get chosen playlist")
	}
	start(radioDir + userChoice)
}
