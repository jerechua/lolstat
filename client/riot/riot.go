package riot

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"time"

	"../../models"
	"../../myhttp"
	rl "../../myhttp/ratelimit"
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

type RiotAPI interface {
	Champions() []*Champion
	ChampionByID(ID int) *Champion
	SummonersByName(names ...string) ([]Summoner, error)
	MatchListForSummonerID(ID int64) ([]*models.SummonerMatch, error)
	MatchByID(ID int64) (*models.Match, error)
}

// riotAPI implements RiotAPI
type riotAPI struct {
	Name      string
	BaseURL   string
	APIKey    string
	Region    string
	champions []*Champion

	rateLimitClient *rl.RateLimitClient
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

func NAClient() (RiotAPI, error) {
	return newClient("na")
}

func newClient(region string) (*riotAPI, error) {
	riot := &riotAPI{
		Name:    "Riot",
		BaseURL: "na.api.pvp.net",
		APIKey:  *riotAPIKeyFlag,
		Region:  region,
		rateLimitClient: &rl.RateLimitClient{
			Limits: []*rl.Limit{
				&rl.Limit{
					Requests: 10,
					Timeout:  10 * time.Second,
				},
				&rl.Limit{
					Requests: 500,
					Timeout:  10 * time.Minute,
				},
			},
		},
	}

	if err := riot.initChampionsAPI(); err != nil {
		return nil, err
	}

	return riot, nil
}

// initchampionsAPI initializes the champions.
func (r *riotAPI) initChampionsAPI() error {
	uri := fmt.Sprintf("/api/lol/static-data/%s/v1.2/champion", r.Region)

	res, err := r.get(myhttp.NewRequestBuilder().SetPath(uri))
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
	r.champions = allChampions
	return nil
}

func (r *riotAPI) Champions() []*Champion {
	return r.champions
}

func (r *riotAPI) ChampionByID(ID int) *Champion {
	for _, c := range r.champions {
		if c.ID == ID {
			return c
		}
	}
	return nil
}

func (r *riotAPI) SummonersByName(names ...string) ([]Summoner, error) {
	joined := strings.Join(names, ",")
	uri := fmt.Sprintf("/api/lol/%s/v1.4/summoner/by-name/%s", r.Region, joined)

	res, err := r.get(myhttp.NewRequestBuilder().SetPath(uri))
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

func (r *riotAPI) MatchListForSummonerID(ID int64) ([]*models.SummonerMatch, error) {
	uri := fmt.Sprintf("/api/lol/%s/v2.2/matchlist/by-summoner/%d", r.Region, ID)

	res, err := r.get(myhttp.NewRequestBuilder().SetPath(uri))
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
func (r *riotAPI) MatchByID(ID int64) (*models.Match, error) {
	uri := fmt.Sprintf("/api/lol/%s/v2.2/match/%d", r.Region, ID)

	// Never include timeline! Average size of response without timeline is ~40kb
	// with the timeline is ~265kb.
	rb := myhttp.NewRequestBuilder().
		SetPath(uri).
		AddQueryParam("includeTimeline", "false")

	res, err := r.get(rb)
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

func (r *riotAPI) get(rb *myhttp.RequestBuilder) (*myhttp.Response, error) {
	rb.SetHost(r.BaseURL).
		Secure().
		AddQueryParam("api_key", r.APIKey)

	req, err := rb.Build()
	if err != nil {
		return nil, err
	}
	cr := make(chan *myhttp.Response)
	ce := make(chan error)
	r.rateLimitClient.Get(req, cr, ce)
	res := <-cr
	err = <-ce

	if err != nil {
		return nil, err
	}
	if err := res.StatusError(); err != nil {
		return nil, err
	}
	return res, err
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
