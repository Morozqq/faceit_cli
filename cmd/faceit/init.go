package main

import (
	"Faceit_Stats_CLI/cmd/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	APIKey        string `json:"api_key"`
	BaseURL       string `json:"base_url"`
	NameUrl       string `json:"name_url"`
	StatsEndPoint string `json:"stats_endpoint"`
}

func initializeApp() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var nickname string

	client := http.Client{}

	if len(os.Args) >= 2 {
		nickname = os.Args[1]
	} else {
		fmt.Println("Please provide a nickname as a command-line argument.")
		return
	}

	playerID, err := FindNickName(nickname)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", config.BaseURL+playerID+config.StatsEndPoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + config.APIKey},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error executing request:", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var matchData models.MatchData
	err = json.Unmarshal(body, &matchData)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	fmt.Printf("Statistics of, %s\n", nickname)
	printStatistics(matchData)
}
