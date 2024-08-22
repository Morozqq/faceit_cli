package main

import (
	"Faceit_Stats_CLI/cmd/calculations"
	"Faceit_Stats_CLI/cmd/models"
	"fmt"
)

func printStatistics(matchData models.MatchData) {
	fmt.Println("Average Kills:", calculations.StatsAverage(matchData.Items, "Kills"))
	fmt.Println("Average Deaths:", calculations.StatsAverage(matchData.Items, "Deaths"))
	fmt.Println("Average Assists:", calculations.StatsAverage(matchData.Items, "Assists"))
	fmt.Printf("Average Headshots Percentage: %s %%\n", calculations.StatsAverage(matchData.Items, "HeadshotsPercentage"))
	fmt.Println("Average ADR:", calculations.StatsAverage(matchData.Items, "Adr"))
	fmt.Println("Average K/D Ratio:", calculations.StatsAverage(matchData.Items, "KdRatio"))
	fmt.Println("Average K/R Ratio:", calculations.StatsAverage(matchData.Items, "KrRatio"))
}
