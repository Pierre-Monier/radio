package main

import (
	"fmt"
	"strconv"
)

// formatUserChoice return a clean playlist selection
func formatUserChoice(playlists []map[string]interface{}) string {
	i := 1
	res := ""
	for _, v := range playlists {
		res += strconv.Itoa(i) + ")" + v["name"].(string) + "\n"
		i++
	}
	return res
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
