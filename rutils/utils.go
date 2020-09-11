package rutils

import (
	"strconv"
)


func FormatUserChoice(playlists []map[string]interface{}) string {
	i := 1
	res := ""
	for _, v := range playlists {
		res += strconv.Itoa(i)+")"+v["name"].(string)+"\n"
		i++
	}	
	return res
}

func ValideUserResponse(userResponse string, pLenght int64) bool {
	nbr, err := strconv.ParseInt(userResponse, 10, 32)
	if err != nil {
		return false
	}else if nbr <= pLenght && nbr > 0 {
		return true
	}else {
		return false
	}
}