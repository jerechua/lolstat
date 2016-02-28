package models

type Match struct {
	MapID                 int64                  `json:"mapId"`
	MatchCreation         int64                  `json:"matchCreation"`
	MatchDuration         int64                  `json:"matchDuration"`
	MatchID               int64                  `json:"matchId"`
	MatchMode             string                 `json:"matchMode"`
	MatchType             string                 `json:"matchType"`
	MatchVersion          string                 `json:"matchVersion"`
	ParticipantIdentities []*ParticipantIdentity `json:"participantIdentities"`
	Participants          []*Participant         `json:"participants"`
	PlatformID            string                 `json:"platformId"`
	QueueType             string                 `json:"queueType"`
	Region                string                 `json:"region"`
	Season                string                 `json:"season"`
	Teams                 []*Team                `json:"teams"`
	Timeline              *Timeline              `json:"timeline"`
}

type Participant struct {
	ChampionID                int64                `json:"championId"`
	HighestAchievedSeasonTier string               `json:"highestAchievedSeasonTier"`
	Masteries                 []*Mastery           `json:"masteries"`
	ParticipantID             int64                `json:"participantId"`
	Runes                     []*Rune              `json:"runes"`
	Spell1ID                  int64                `json:"spell1Id"`
	Spell2ID                  int64                `json:"spell2Id"`
	Stats                     *ParticipantStats    `json:"stats"`
	TeamID                    int64                `json:"teamId"`
	Timeline                  *ParticipantTimeline `json:"timeline"`
}

type ParticipantIdentity struct {
	ParticipantID int    `json:"participantId"`
	Player        Player `json:"player"`
}

type Team struct {
	Bans                 []*BannedChampion `json:"bans"`
	BaronKills           int64             `json:"baronKills"`
	DominionVictoryScore int64             `json:"dominionVictoryScore"`
	DragonKills          int64             `json:"dragonKills"`
	FirstBaron           bool              `json:"firstBaron"`
	FirstBlood           bool              `json:"firstBlood"`
	FirstDragon          bool              `json:"firstDragon"`
	FirstInhibitor       bool              `json:"firstInhibitor"`
	FirstRiftHerald      bool              `json:"firstRiftHerald"`
	FirstTower           bool              `json:"firstTower"`
	InhibitorKills       int64             `json:"inhibitorKills"`
	RiftHeraldKills      int64             `json:"riftHeraldKills"`
	TeamID               int64             `json:"teamId"`
	TowerKills           int64             `json:"towerKills"`
	VilemawKills         int64             `json:"vilemawKills"`
	Winner               bool              `json:"winner"`
}

type Timeline struct {
	FrameInterval int64    `json:"frameInterval"`
	Frames        []*Frame `json:"frames"`
}

type Mastery struct {
	MasteryID int64 `json:"masteryId"`
	Rank      int64 `json:"rank"`
}

type ParticipantStats struct {
	Assists                         int64 `json:"assists"`
	ChampLevel                      int64 `json:"champLevel"`
	CombatPlayerScore               int64 `json:"combatPlayerScore"`
	Deaths                          int64 `json:"deaths"`
	DoubleKills                     int64 `json:"doubleKills"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	GoldEarned                      int64 `json:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent"`
	InhibitorKills                  int64 `json:"inhibitorKills"`
	Item0                           int64 `json:"item0"`
	Item1                           int64 `json:"item1"`
	Item2                           int64 `json:"item2"`
	Item3                           int64 `json:"item3"`
	Item4                           int64 `json:"item4"`
	Item5                           int64 `json:"item5"`
	Item6                           int64 `json:"item6"`
	KillingSprees                   int64 `json:"killingSprees"`
	Kills                           int64 `json:"kills"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike"`
	LargestKillingSpree             int64 `json:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                int64 `json:"magicDamageTaken"`
	MinionsKilled                   int64 `json:"minionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     int64 `json:"nodeCapture"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist"`
	NodeNeutralize                  int64 `json:"nodeNeutralize"`
	NodeNeutralizeAssist            int64 `json:"nodeNeutralizeAssist"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore"`
	PentaKills                      int64 `json:"pentaKills"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
	QuadraKills                     int64 `json:"quadraKills"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame"`
	TeamObjective                   int64 `json:"teamObjective"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalPlayerScore                int64 `json:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed"`
	TowerKills                      int64 `json:"towerKills"`
	TripleKills                     int64 `json:"tripleKills"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	UnrealKills                     int64 `json:"unrealKills"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame"`
	WardsKilled                     int64 `json:"wardsKilled"`
	WardsPlaced                     int64 `json:"wardsPlaced"`
	Winner                          bool  `json:"winner"`
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts *ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   *ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  *ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   *ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        *ParticipantTimelineData `json:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          *ParticipantTimelineData `json:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              *ParticipantTimelineData `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              *ParticipantTimelineData `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     *ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         *ParticipantTimelineData `json:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       *ParticipantTimelineData `json:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         *ParticipantTimelineData `json:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  *ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    *ParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                *ParticipantTimelineData `json:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    *ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      *ParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`
	Lane                            string                   `json:"lane"`
	Role                            string                   `json:"role"`
	TowerAssistsPerMinCounts        *ParticipantTimelineData `json:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          *ParticipantTimelineData `json:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          *ParticipantTimelineData `json:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      *ParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        *ParticipantTimelineData `json:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               *ParticipantTimelineData `json:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              *ParticipantTimelineData `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  *ParticipantTimelineData `json:"xpPerMinDeltas"`
}

type Rune struct {
	Rank   int64 `json:"rank"`
	RuneID int64 `json:"runeId"`
}

type Player struct {
	MatchHistoryUri string `json:"matchHistoryUri"`
	ProfileIcon     int64  `json:"profileIcon"`
	SummonerID      int64  `json:"summonerId"`
	SummonerName    string `json:"summonerName"`
}

type BannedChampion struct {
	ChampionID int64 `json:"championId"`
	PickTurn   int64 `json:"pickTurn"`
}

type Frame struct {
	Events            []*Event                     `json:"events"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames"`
	Timestamp         int64                        `json:"timestamp"`
}

type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty"`
	ThirtyToEnd    float64 `json:"thirtyToEnd"`
	TwentyToThirty float64 `json:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen"`
}

type Event struct {
	AscendedType            string    `json:"ascendedType"`
	AssistingParticipantIds []int64   `json:"assistingParticipantIds"`
	BuildingType            string    `json:"buildingType"`
	CreatorID               int64     `json:"creatorId"`
	EventType               string    `json:"eventType"`
	ItemAfter               int64     `json:"itemAfter"`
	ItemBefore              int64     `json:"itemBefore"`
	ItemID                  int64     `json:"itemId"`
	KillerID                int64     `json:"killerId"`
	LaneType                string    `json:"laneType"`
	LevelUpType             string    `json:"levelUpType"`
	MonsterType             string    `json:"monsterType"`
	ParticipantID           int64     `json:"participantId"`
	Point64Captured         string    `json:"point64Captured"`
	Position                *Position `json:"position"`
	SkillSlot               int64     `json:"skillSlot"`
	TeamID                  int64     `json:"teamId"`
	Timestamp               int64     `json:"timestamp"`
	TowerType               string    `json:"towerType"`
	VictimID                int64     `json:"victimId"`
	WardType                string    `json:"wardType"`
}

type ParticipantFrame struct {
	CurrentGold         int64     `json:"currentGold"`
	DominionScore       int64     `json:"dominionScore"`
	JungleMinionsKilled int64     `json:"jungleMinionsKilled"`
	Level               int64     `json:"level"`
	MinionsKilled       int64     `json:"minionsKilled"`
	ParticipantID       int64     `json:"participantId"`
	Position            *Position `json:"position"`
	TeamScore           int64     `json:"teamScore"`
	TotalGold           int64     `json:"totalGold"`
	XP                  int64     `json:"xp"`
}

type Position struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
