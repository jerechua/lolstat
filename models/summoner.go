package models

import (
	"../db/gorm"
)

func init() {
	gorm.RegisterModel(&SummonerMatch{})
}

type Summoner struct {
	ID            int64  `json:"id" sql:"unique_index"`
	Name          string `json:"name"`
	ProfileIconId int64  `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}
