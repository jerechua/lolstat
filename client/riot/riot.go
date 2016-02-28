package riot

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"../../models"
	"../../myhttp"
)

var (
	riotAPIKeyFlag = flag.String("riot_api_key", "6fc580c4-5f1e-4738-b07f-582e1e987c7c", "The Riot API key.")
)

func init() {
	flag.Parse()

	if *riotAPIKeyFlag == "" {
		panic("--riot_api_key is a required flag.")
	}
}

type RiotAPI struct {
	Name         string
	BaseURL      string
	APIKey       string
	Region       string
	ChampionsAPI *ChampionsAPI
}

type ChampionsAPI struct {
	champions []*Champion
}

func (api *ChampionsAPI) All() []*Champion {
	return api.champions
}

func (api *ChampionsAPI) ChampionByID(ID int) *Champion {
	for _, c := range api.champions {
		if c.ID == ID {
			return c
		}
	}
	return nil
}

type Champion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Summoner struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileIconId int64  `json:"profileIconId"`
}

func NAClient() (*RiotAPI, error) {
	return newClient("na")
}

func newClient(region string) (*RiotAPI, error) {
	riot := &RiotAPI{
		Name:    "Riot",
		BaseURL: "na.api.pvp.net",
		APIKey:  *riotAPIKeyFlag,
		Region:  region,
	}

	if err := riot.initChampionsAPI(); err != nil {
		return nil, err
	}

	return riot, nil
}

// initChampionsAPI initializes the champions.
func (r *RiotAPI) initChampionsAPI() error {
	uri := fmt.Sprintf("/api/lol/static-data/%s/v1.2/champion", r.Region)

	res, err := r.get(uri)
	if err != nil {
		return err
	}
	if err := res.StatusError(); err != nil {
		return err
	}

	// Riot's API returns champion names as a key, so just treat it as a generic
	// type, and loop over everything making them into "Champion"
	var champs interface{}
	err = decode(res.Body(), &champs)
	if err != nil {
		return err
	}

	var allChampions []*Champion
	reply := champs.(map[string]interface{})
	for _, i := range reply["data"].(map[string]interface{}) {
		var champ Champion
		err := remarshal(i, &champ)
		if err != nil {
			return err
		}
		allChampions = append(allChampions, &champ)
	}
	r.ChampionsAPI = &ChampionsAPI{allChampions}
	return nil
}

func (r *RiotAPI) SummonersByName(names ...string) ([]Summoner, error) {
	joined := strings.Join(names, ",")
	uri := fmt.Sprintf("/api/lol/%s/v1.4/summoner/by-name/%s", r.Region, joined)

	res, err := r.get(uri)
	if err != nil {
		return nil, err
	}

	var summs interface{}
	err = decode(res.Body(), &summs)
	if err != nil {
		return nil, err
	}

	var allSummoners []Summoner
	// _ is the summoner name
	for _, i := range summs.(map[string]interface{}) {
		var s Summoner
		err := remarshal(i, &s)
		if err != nil {
			return nil, err
		}
		allSummoners = append(allSummoners, s)
	}
	return allSummoners, nil
}

func (r *RiotAPI) MatchListForSummonerID(ID int64) ([]*models.SummonerMatch, error) {
	uri := fmt.Sprintf("/api/lol/%s/v2.2/matchlist/by-summoner/%d", r.Region, ID)

	res, err := r.get(uri)
	if err != nil {
		return nil, err
	}

	var matchList struct {
		Matches    []*models.SummonerMatch `json:"matches"`
		TotalGames int                     `json:"totalGames"`
		StartIndex int                     `json:startIndex`
		EndIndex   int                     `json:endIndex`
	}

	err = decode(res.Body(), &matchList)
	if err != nil {
		return nil, err
	}

	for _, m := range matchList.Matches {
		// This field is not in the JSON response, but kept here so we can reuse
		// this model for the DB.
		m.SummonerID = ID
	}

	return matchList.Matches, nil
}

// MatchByID retrieves the match from the Riot API. This call has a lot of
// returned values, consider just shoving this data to MongoDB?
func (r *RiotAPI) MatchByID(ID int64) (*models.Match, error) {
	uri := fmt.Sprintf("/api/lol/%s/v2.2/match/%d", r.Region, ID)

	res, err := r.get(uri, "includeTimeline", "true")
	if err != nil {
		return nil, err
	}

	var match models.Match
	err = decode(res.Body(), &match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}

func (r *RiotAPI) get(uri string, qp ...string) (*myhttp.Response, error) {
	if len(qp)%2 != 0 {
		return nil, fmt.Errorf("Must supply query parameters in key-value pairs.")
	}
	req := myhttp.NewRequest(r.BaseURL, uri)
	req.Secure()
	req.AddQueryParam("api_key", r.APIKey)
	for i := 0; i < len(qp); i += 2 {
		req.AddQueryParam(qp[i], qp[i+1])
	}
	return req.Get()
}

// decode takes the json body of the response and unmarshals it into a
// map[string]interface{}. The response can be casted by doing:
//
// res := decode(res.Body())
// res.(map[string]interface{})
func decode(b []byte, i interface{}) error {
	d := json.NewDecoder(bytes.NewBuffer(b))
	d.UseNumber()
	if err := d.Decode(&i); err != nil {
		return err
	}
	return nil
}

// remarshal will take interface i, which is a generic unmarshalled json object
// and s is the struct to put i into. expects the json response to be a map of
// "foo": {"bar": "baz"}, where "foo" is not a well defined field, and
// {"bar": "baz"} is a well defined field
func remarshal(i, s interface{}) error {
	// As ugly as it may be, since this was unmarshalled into a generic
	// interface, it would be easier to remarshal and unmarshal again into
	// the correct struct. Otherwise we'd have to manually populate it. Might
	// actually be better to do that..
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}
	return nil
}
