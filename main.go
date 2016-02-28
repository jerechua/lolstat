package main

import (
	"fmt"
	"log"

	"./client/riot"

	_ "./db"
)

func main() {
	client, err := riot.NAClient()
	if err != nil {
		log.Panic(err)
	}

	for _, c := range client.ChampionsAPI.All() {
		fmt.Printf("id: %d,  name: %s\n", c.ID, c.Name)
	}

	fmt.Println(len(client.ChampionsAPI.All()))

	res, err := client.SummonersByName("blooooop")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)

	fmt.Println(client.ChampionsAPI.ChampionByID(99))

	matches, err := client.MatchListForSummonerID(25286733)
	if err != nil {
		log.Panic(err)
	}
	for _, m := range matches {
		fmt.Println(m.SummonerID)
	}
}
