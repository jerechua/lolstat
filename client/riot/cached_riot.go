package riot

import (
	"../../models"
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

func (c *cachedRiotAPI) Champions() []*Champion {
	return c.client.Champions()
}
func (c *cachedRiotAPI) ChampionByID(ID int) *Champion {
	return c.client.ChampionByID(ID)
}

func (c *cachedRiotAPI) SummonersByName(names ...string) ([]Summoner, error) {
	return c.client.SummonersByName(names...)
}

func (c *cachedRiotAPI) MatchListForSummonerID(ID int64) ([]*models.SummonerMatch, error) {
	return c.client.MatchListForSummonerID(ID)
}

func (c *cachedRiotAPI) MatchListSinceTime(ID, beginTime int64) ([]*models.SummonerMatch, error) {
	return c.client.MatchListSinceTime(ID, beginTime)
}

func (c *cachedRiotAPI) MatchByID(ID int64) (*models.Match, error) {
	return c.client.MatchByID(ID)
}
