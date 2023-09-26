package main

import "fmt"

func main1() {
	leagueTitles := make(map[string]int)

	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle": 0,
	}

	fmt.Printf("League titles: %v\nRecent head to head: %v\n", leagueTitles, recentHead2HeadWins)
}
