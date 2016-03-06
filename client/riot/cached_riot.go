package riot

import (
	"fmt"
	"log"

	"../../db"
	"../../models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// cachchedRiotAPI implements RiotAPI with the addition of database caching.
type cachedRiotAPI struct {
	client *riotAPI
}

func NewCachedNAClient() (RiotAPI, error) {
	return newCachcedClient("na")
}

func newCachcedClient(region string) (*cachedRiotAPI, error) {
	client, err := newClient(region)
	if err != nil {
		return nil, err
	}
	return &cachedRiotAPI{client}, nil
}

func (api *cachedRiotAPI) Champions() []*Champion {
	return api.client.Champions()
}

func (api *cachedRiotAPI) ChampionByID(ID int) *Champion {
	return api.client.ChampionByID(ID)
}

func (api *cachedRiotAPI) SummonersByName(name string) (*models.Summoner, error) {
	return api.client.SummonersByName(name)
}

func (api *cachedRiotAPI) MatchListForSummonerID(ID int64) ([]*models.SummonerMatch, error) {
	gorm := db.GORM.DB()
	var latestMatch models.SummonerMatch
	gorm.Where(&models.SummonerMatch{SummonerID: ID}).Order("timestamp desc").First(&latestMatch)
	if latestMatch.MatchID == 0 {
		// Match List does not exist for this person, so lets just pull all that data
		// down and save it.
		matchList, err := api.client.MatchListForSummonerID(ID)
		if err != nil {
			return nil, err
		}
		for _, m := range matchList {
			if err := db.GORM.Create(m); err != nil {
				// Just return the err here, since presumably we get the list sorted.
				// so it will just fix itself. Maybe.
				return nil, err
			}
		}
		return matchList, err
	} else {
		var cachedMatches []*models.SummonerMatch
		gorm.Where(&models.SummonerMatch{SummonerID: ID}).Order("timestamp desc").Find(&cachedMatches)
		// Person has been found, lets get their latest match and update their profile.
		matchList, err := api.client.MatchListSinceTime(ID, latestMatch.Timestamp)
		if err != nil {
			return nil, err
		}
		for _, m := range matchList {
			if err := db.GORM.Create(m); err != nil {
				return nil, err
			}
		}
		return append(cachedMatches, matchList...), err
	}
}

func (api *cachedRiotAPI) MatchListSinceTime(ID, beginTime int64) ([]*models.SummonerMatch, error) {
	return api.client.MatchListSinceTime(ID, beginTime)
}

func (api *cachedRiotAPI) MatchByID(matchID int64) (*models.Match, error) {
	c := db.MATCH.C()
	var m *models.Match
	err := c.Find(&models.Match{MatchID: matchID}).One(&m)
	switch err {
	case nil:
		log.Printf("MatchID: %d found in datastore", matchID)
		return m, nil
	case mgo.ErrNotFound:
		log.Printf("Match ID: %d is not in datastore, fetching from Riot", matchID)

		match, err := api.client.MatchByID(matchID)
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
	}
	return nil, err
}
