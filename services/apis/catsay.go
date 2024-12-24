package api

import (
	"encoding/json"
	"net/http"
)

func Catsay() (string, error) {
	type Json struct {
		Fact string `json:"fact"`
	}
	response, err := http.Get("https://catfact.ninja/fact?max_length=40")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var jsonResponse Json
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		return "", err
	}
	return jsonResponse.Fact, err
}
