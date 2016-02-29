package main

import (
	"fmt"
	"log"

	"./client/riot"
	"./db"
	"./models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	c := db.MATCH.C()
	var m models.Match
	err = c.Find(&models.Match{MatchID: match.MatchID}).One(&m)
	switch err {
	case nil:
		break
	case mgo.ErrNotFound:
		match.ID = bson.NewObjectId()
		if err := c.Insert(match); err != nil {
			log.Panic(err)
		}
		break
	default:
		log.Panic(err)
	}

	log.Print(m)

}
