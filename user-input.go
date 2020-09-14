package main

import (
	"errors"
	"fmt"
	"strconv"
)

// getUserUsefullUserChoice return the playlist directory to play or error
func getUsefullUserChoice(playlists []map[string]interface{}) (string, error) {
	formatedUserChoice, err := formatUserChoice(playlists)
	if err != nil {
		return "", err
	}
	userChoice := getUserChoice(formatedUserChoice, len(playlists))
	i, err := strconv.ParseInt(userChoice, 10, 8)
	if err != nil {
		panic("Can't get userChoice index")
	}
	return playlists[i-1]["name"].(string), nil
}

// formatUserChoice return a clean playlist selection
func formatUserChoice(playlists []map[string]interface{}) (string, error) {
	i := 1
	res := ""
	for _, v := range playlists {
		res += strconv.Itoa(i) + ")" + v["name"].(string) + "\n"
		i++
	}
	if res == "" {
		return res, errors.New("Can't format user choice")
	}
	return res, nil
}

// valideUserResponse check if the user response (of the playlist choice) is correct
func valideUserResponse(userResponse string, pLenght int64) bool {
	nbr, err := strconv.ParseInt(userResponse, 10, 8)
	if err != nil {
		return false
	} else if nbr <= pLenght && nbr > 0 {
		return true
	} else {
		return false
	}
}

func getUserChoice(fChoices string, pLength int) string {
	fmt.Println("Select the playlist \n" + fChoices)
	match := false
	var userResponse string
	for !match {
		fmt.Scanln(&userResponse)
		match = valideUserResponse(userResponse, int64(pLength))
	}
	return userResponse
}
