package models

import (
	"fmt"

	"../db"

	v "github.com/jerechua/validate"
)

func init() {
	db.RegisterModel(&SummonerMatch{})
}

type SummonerMatch struct {
	SummonerID int64  `sql:"unique_index:idx_summid_matchid";validate:"required"`
	Timestamp  int64  `json:"timestamp"`
	Champion   int    `json:"champion"`
	Region     string `json:"region"`
	Queue      string `json:"queue"`
	Season     string `json:"season"`
	MatchID    int64  `json:"matchId";sql:"unique_index:idx_summid_matchid";validate:"required"`
	// Role is either DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT
	Role       string `json:"role"`
	PlatformID string `json:"platformId"`
	Lane       string `json:"lane"`
}

// Create creates a new SummonerMatch model
func CreateSummonerMatch(sm *SummonerMatch) error {
	if err := v.Validate(sm); err != nil {
		return err
	}
	if err := db.DB.Create(sm).Error; err != nil {
		return fmt.Errorf("This match already exists.")
	}
	return nil
}
