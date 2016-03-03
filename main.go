package main

import (
	"fmt"
	"log"

	"./client/riot"
	"./db"
	"./models"
)

func main() {
	db.Init()

	client, err := riot.NewCachedNAClient()
	if err != nil {
		log.Panic(err)
	}

	for _, c := range client.Champions() {
		fmt.Printf("id: %d,  name: %s\n", c.ID, c.Name)
		break
	}

	fmt.Println(len(client.Champions()))

	res, err := client.SummonersByName("blooooop")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)

	fmt.Println(client.ChampionByID(99))

	// matches, err := client.MatchListSinceTime(25286733, 1454822389504)
	matches, err := client.MatchListForSummonerID(25286733)
	if err != nil {
		log.Panic(err)
	}
	log.Fatal(matches)
	for _, m := range matches {
		exists, err := db.GORM.Exists(&models.SummonerMatch{
			MatchID:    m.MatchID,
			SummonerID: m.SummonerID,
		})
		if err != nil {
			log.Panic(err)
		}
		if exists {
			log.Printf("%d already exists", m.MatchID)
		} else {
			if err := db.GORM.Create(m); err != nil {
				log.Panic(err)
			}
		}

		match, err := client.MatchByID(m.MatchID)
		if err != nil {
			log.Panic(err)
		}
		log.Print(match.MatchID)
	}
}
