package riot

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

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

type Champion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RiotAPI struct {
	Name      string
	BaseURL   string
	APIKey    string
	Region    string
	Champions []Champion
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
	if err = json.Unmarshal(res.Body(), &champs); err != nil {
		return err
	}

	var allChampions []Champion
	reply := champs.(map[string]interface{})
	for _, i := range reply["data"].(map[string]interface{}) {

		// As ugly as it may be, since this was unmarshalled into a generic
		// interface, it would be easier to remarshal and unmarshal again into
		// the correct struct. Otherwise we'd have to manually populate it. Might
		// actually be better to do that..
		b, err := json.Marshal(i)
		if err != nil {
			return err
		}

		var champ Champion
		if err = json.Unmarshal(b, &champ); err != nil {
			return err
		}
		allChampions = append(allChampions, champ)
	}
	r.Champions = allChampions
	return nil
}

func (r *RiotAPI) SummonerByName(names ...string) (*myhttp.Response, error) {
	joined := strings.Join(names, ",")
	uri := fmt.Sprintf("/api/lol/%s/v1.4/summoner/by-name/%s", r.Region, joined)

	return r.get(uri)
}

func (r *RiotAPI) get(uri string) (*myhttp.Response, error) {
	req := myhttp.NewRequest(r.BaseURL, uri)
	req.Secure()
	req.AddQueryParam("api_key", r.APIKey)
	return req.Get()
}
