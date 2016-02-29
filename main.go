package main

import (
	"fmt"
	"log"

	"./client/riot"
	"./db"
)

func main() {
	db.Init()

	client, err := riot.NAClient()
	if err != nil {
		log.Panic(err)
	}

	for _, c := range client.ChampionsAPI.All() {
		fmt.Printf("id: %d,  name: %s\n", c.ID, c.Name)
		break
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
		fmt.Println(m.MatchID)
		exists, err := db.GORM.Exists(m)
		if err != nil {
			log.Print(err)
		}
		if exists {
			log.Printf("%d already exists", m.MatchID)
		}
		break
	}

	match, err := client.MatchByID(2090888305)
	if err != nil {
		log.Panic(err)
	}
	log.Print(match.MatchID)
	log.Print(match.MatchType)
	log.Print(match.Timeline)
}
