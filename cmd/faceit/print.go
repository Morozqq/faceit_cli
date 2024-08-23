package main

import (
	"Faceit_Stats_CLI/cmd/calculations"
	"Faceit_Stats_CLI/cmd/models"
	"fmt"
)

func printStatistics(matchData models.MatchData) {

	fmt.Println(Green("Average Kills: ", calculations.StatsAverage(matchData.Items, "Kills")))
	fmt.Println(Red("Average Deaths:", calculations.StatsAverage(matchData.Items, "Deaths")))
	fmt.Println(Yellow("Average Assists:", calculations.StatsAverage(matchData.Items, "Assists")))
	fmt.Printf("Average Headshots Percentage: %s %%\n", calculations.StatsAverage(matchData.Items, "HeadshotsPercentage"))
	fmt.Println(Red("Average ADR:", calculations.StatsAverage(matchData.Items, "Adr")))
	fmt.Println(Green("Average K/D Ratio:", calculations.StatsAverage(matchData.Items, "KdRatio")))
	fmt.Println(Blue("Average K/R Ratio:", calculations.StatsAverage(matchData.Items, "KrRatio")))
}
