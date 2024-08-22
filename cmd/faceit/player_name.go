package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindNickName(nickname string) (string, error) {
	client := http.Client{}

	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return "", err
	}

	req, err := http.NewRequest("GET", config.NameUrl+nickname, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return "", err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer 3b5a9d07-4b77-4be4-bf5d-f9ff3d6cb7b6"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing name request")
		return "", err
	}

	defer res.Body.Close()

	body, ioError := io.ReadAll(res.Body)
	if ioError != nil {
		fmt.Println("Error reading response body")
		return "", ioError
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Player not found. Status code: %d", res.StatusCode)
	}

	var player map[string]interface{}
	jsonError := json.Unmarshal(body, &player)
	if jsonError != nil {
		return "", fmt.Errorf("Error unmarshalling player data: %s", jsonError)
	}

	// Check if the "player_id" field exists
	if playerID, ok := player["player_id"].(string); ok {
		return playerID, nil
	}

	return "", fmt.Errorf("Player not found.")
}
