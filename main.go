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

	matches, err := client.MatchListSinceTime(25286733, 1454822389504)
	if err != nil {
		log.Panic(err)
	}
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

		match, err := getOrCreateMatchDetails(client, m.MatchID)
		if err != nil {
			log.Panic(err)
		}
		log.Print(match.MatchID)
	}
}

// getOrCreateMatchDetails will try to retrieve the match details from the datastore
// first. If it doesn't fine it, it will fetch the details directly from the
// Riot API.
func getOrCreateMatchDetails(client riot.RiotAPI, matchID int64) (*models.Match, error) {
	c := db.MATCH.C()
	var m *models.Match
	err := c.Find(&models.Match{MatchID: matchID}).One(&m)
	switch err {
	case nil:
		return m, nil
	case mgo.ErrNotFound:
		log.Printf("Match ID: %d is not in datastore, fetching from Riot", matchID)

		match, err := client.MatchByID(matchID)
		if err != nil {
			return nil, err
		}
		if match.MatchID != matchID {
			return nil, fmt.Errorf("Match ID is not the same. Expected ID: %d, but got %d", matchID, match.MatchID)
		}

		match.ID = bson.NewObjectId()
		if err := c.Insert(match); err != nil {
			return nil, err
		}
		return match, nil
	default:
		return nil, err
	}
}
