package models

type Match struct {
	MapID                 int64                  `json:"mapId" bson:"mapId"`
	MatchCreation         int64                  `json:"matchCreation" bson:"matchCreation"`
	MatchDuration         int64                  `json:"matchDuration" bson:"matchDuration"`
	MatchID               int64                  `json:"matchId" bson:"matchId"`
	MatchMode             string                 `json:"matchMode" bson:"matchMode"`
	MatchType             string                 `json:"matchType" bson:"matchType"`
	MatchVersion          string                 `json:"matchVersion" bson:"matchVersion"`
	ParticipantIdentities []*ParticipantIdentity `json:"participantIdentities" bson:"participantIdentities"`
	Participants          []*Participant         `json:"participants" bson:"participants"`
	PlatformID            string                 `json:"platformId" bson:"platformId"`
	QueueType             string                 `json:"queueType" bson:"queueType"`
	Region                string                 `json:"region" bson:"region"`
	Season                string                 `json:"season" bson:"season"`
	Teams                 []*Team                `json:"teams" bson:"teams"`
	Timeline              *Timeline              `json:"timeline" bson:"timeline"`
}

type Participant struct {
	ChampionID                int64                `json:"championId" bson:"championId"`
	HighestAchievedSeasonTier string               `json:"highestAchievedSeasonTier" bson:"highestAchievedSeasonTier"`
	Masteries                 []*Mastery           `json:"masteries" bson:"masteries"`
	ParticipantID             int64                `json:"participantId" bson:"participantId"`
	Runes                     []*Rune              `json:"runes" bson:"runes"`
	Spell1ID                  int64                `json:"spell1Id" bson:"spell1Id"`
	Spell2ID                  int64                `json:"spell2Id" bson:"spell2Id"`
	Stats                     *ParticipantStats    `json:"stats" bson:"stats"`
	TeamID                    int64                `json:"teamId" bson:"teamId"`
	Timeline                  *ParticipantTimeline `json:"timeline" bson:"timeline"`
}

type ParticipantIdentity struct {
	ParticipantID int    `json:"participantId" bson:"participantId"`
	Player        Player `json:"player" bson:"player"`
}

type Team struct {
	Bans                 []*BannedChampion `json:"bans" bson:"bans"`
	BaronKills           int64             `json:"baronKills" bson:"baronKills"`
	DominionVictoryScore int64             `json:"dominionVictoryScore" bson:"dominionVictoryScore"`
	DragonKills          int64             `json:"dragonKills" bson:"dragonKills"`
	FirstBaron           bool              `json:"firstBaron" bson:"firstBaron"`
	FirstBlood           bool              `json:"firstBlood" bson:"firstBlood"`
	FirstDragon          bool              `json:"firstDragon" bson:"firstDragon"`
	FirstInhibitor       bool              `json:"firstInhibitor" bson:"firstInhibitor"`
	FirstRiftHerald      bool              `json:"firstRiftHerald" bson:"firstRiftHerald"`
	FirstTower           bool              `json:"firstTower" bson:"firstTower"`
	InhibitorKills       int64             `json:"inhibitorKills" bson:"inhibitorKills"`
	RiftHeraldKills      int64             `json:"riftHeraldKills" bson:"riftHeraldKills"`
	TeamID               int64             `json:"teamId" bson:"teamId"`
	TowerKills           int64             `json:"towerKills" bson:"towerKills"`
	VilemawKills         int64             `json:"vilemawKills" bson:"vilemawKills"`
	Winner               bool              `json:"winner" bson:"winner"`
}

type Timeline struct {
	FrameInterval int64    `json:"frameInterval" bson:"frameInterval"`
	Frames        []*Frame `json:"frames" bson:"frames"`
}

type Mastery struct {
	MasteryID int64 `json:"masteryId" bson:"masteryId"`
	Rank      int64 `json:"rank" bson:"rank"`
}

type ParticipantStats struct {
	Assists                         int64 `json:"assists" bson:"assists"`
	ChampLevel                      int64 `json:"champLevel" bson:"champLevel"`
	CombatPlayerScore               int64 `json:"combatPlayerScore" bson:"combatPlayerScore"`
	Deaths                          int64 `json:"deaths" bson:"deaths"`
	DoubleKills                     int64 `json:"doubleKills" bson:"doubleKills"`
	FirstBloodAssist                bool  `json:"firstBloodAssist" bson:"firstBloodAssist"`
	FirstBloodKill                  bool  `json:"firstBloodKill" bson:"firstBloodKill"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist" bson:"firstInhibitorAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill" bson:"firstInhibitorKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist" bson:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill" bson:"firstTowerKill"`
	GoldEarned                      int64 `json:"goldEarned" bson:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent" bson:"goldSpent"`
	InhibitorKills                  int64 `json:"inhibitorKills" bson:"inhibitorKills"`
	Item0                           int64 `json:"item0" bson:"item0"`
	Item1                           int64 `json:"item1" bson:"item1"`
	Item2                           int64 `json:"item2" bson:"item2"`
	Item3                           int64 `json:"item3" bson:"item3"`
	Item4                           int64 `json:"item4" bson:"item4"`
	Item5                           int64 `json:"item5" bson:"item5"`
	Item6                           int64 `json:"item6" bson:"item6"`
	KillingSprees                   int64 `json:"killingSprees" bson:"killingSprees"`
	Kills                           int64 `json:"kills" bson:"kills"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike" bson:"largestCriticalStrike"`
	LargestKillingSpree             int64 `json:"largestKillingSpree" bson:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill" bson:"largestMultiKill"`
	MagicDamageDealt                int64 `json:"magicDamageDealt" bson:"magicDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions" bson:"magicDamageDealtToChampions"`
	MagicDamageTaken                int64 `json:"magicDamageTaken" bson:"magicDamageTaken"`
	MinionsKilled                   int64 `json:"minionsKilled" bson:"minionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled" bson:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle" bson:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle" bson:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     int64 `json:"nodeCapture" bson:"nodeCapture"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist" bson:"nodeCaptureAssist"`
	NodeNeutralize                  int64 `json:"nodeNeutralize" bson:"nodeNeutralize"`
	NodeNeutralizeAssist            int64 `json:"nodeNeutralizeAssist" bson:"nodeNeutralizeAssist"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore" bson:"objectivePlayerScore"`
	PentaKills                      int64 `json:"pentaKills" bson:"pentaKills"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt" bson:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions" bson:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken" bson:"physicalDamageTaken"`
	QuadraKills                     int64 `json:"quadraKills" bson:"quadraKills"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame" bson:"sightWardsBoughtInGame"`
	TeamObjective                   int64 `json:"teamObjective" bson:"teamObjective"`
	TotalDamageDealt                int64 `json:"totalDamageDealt" bson:"totalDamageDealt"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions" bson:"totalDamageDealtToChampions"`
	TotalDamageTaken                int64 `json:"totalDamageTaken" bson:"totalDamageTaken"`
	TotalHeal                       int64 `json:"totalHeal" bson:"totalHeal"`
	TotalPlayerScore                int64 `json:"totalPlayerScore" bson:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank" bson:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt" bson:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed" bson:"totalUnitsHealed"`
	TowerKills                      int64 `json:"towerKills" bson:"towerKills"`
	TripleKills                     int64 `json:"tripleKills" bson:"tripleKills"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt" bson:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions" bson:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken" bson:"trueDamageTaken"`
	UnrealKills                     int64 `json:"unrealKills" bson:"unrealKills"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame" bson:"visionWardsBoughtInGame"`
	WardsKilled                     int64 `json:"wardsKilled" bson:"wardsKilled"`
	WardsPlaced                     int64 `json:"wardsPlaced" bson:"wardsPlaced"`
	Winner                          bool  `json:"winner" bson:"winner"`
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts *ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts" bson:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   *ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts" bson:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  *ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas" bson:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   *ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas" bson:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        *ParticipantTimelineData `json:"baronAssistsPerMinCounts" bson:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          *ParticipantTimelineData `json:"baronKillsPerMinCounts" bson:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              *ParticipantTimelineData `json:"creepsPerMinDeltas" bson:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              *ParticipantTimelineData `json:"csDiffPerMinDeltas" bson:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     *ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas" bson:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         *ParticipantTimelineData `json:"damageTakenPerMinDeltas" bson:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       *ParticipantTimelineData `json:"dragonAssistsPerMinCounts" bson:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         *ParticipantTimelineData `json:"dragonKillsPerMinCounts" bson:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  *ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts" bson:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    *ParticipantTimelineData `json:"elderLizardKillsPerMinCounts" bson:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                *ParticipantTimelineData `json:"goldPerMinDeltas" bson:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    *ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts" bson:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      *ParticipantTimelineData `json:"inhibitorKillsPerMinCounts" bson:"inhibitorKillsPerMinCounts"`
	Lane                            string                   `json:"lane" bson:"lane"`
	Role                            string                   `json:"role" bson:"role"`
	TowerAssistsPerMinCounts        *ParticipantTimelineData `json:"towerAssistsPerMinCounts" bson:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          *ParticipantTimelineData `json:"towerKillsPerMinCounts" bson:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          *ParticipantTimelineData `json:"towerKillsPerMinDeltas" bson:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      *ParticipantTimelineData `json:"vilemawAssistsPerMinCounts" bson:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        *ParticipantTimelineData `json:"vilemawKillsPerMinCounts" bson:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               *ParticipantTimelineData `json:"wardsPerMinDeltas" bson:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              *ParticipantTimelineData `json:"xpDiffPerMinDeltas" bson:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  *ParticipantTimelineData `json:"xpPerMinDeltas" bson:"xpPerMinDeltas"`
}

type Rune struct {
	Rank   int64 `json:"rank" bson:"rank"`
	RuneID int64 `json:"runeId" bson:"runeId"`
}

type Player struct {
	MatchHistoryUri string `json:"matchHistoryUri" bson:"matchHistoryUri"`
	ProfileIcon     int64  `json:"profileIcon" bson:"profileIcon"`
	SummonerID      int64  `json:"summonerId" bson:"summonerId"`
	SummonerName    string `json:"summonerName" bson:"summonerName"`
}

type BannedChampion struct {
	ChampionID int64 `json:"championId" bson:"championId"`
	PickTurn   int64 `json:"pickTurn" bson:"pickTurn"`
}

type Frame struct {
	Events            []*Event                     `json:"events" bson:"events"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames" bson:"participantFrames"`
	Timestamp         int64                        `json:"timestamp" bson:"timestamp"`
}

type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty" bson:"tenToTwenty"`
	ThirtyToEnd    float64 `json:"thirtyToEnd" bson:"thirtyToEnd"`
	TwentyToThirty float64 `json:"twentyToThirty" bson:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen" bson:"zeroToTen"`
}

type Event struct {
	AscendedType            string    `json:"ascendedType" bson:"ascendedType"`
	AssistingParticipantIds []int64   `json:"assistingParticipantIds" bson:"assistingParticipantIds"`
	BuildingType            string    `json:"buildingType" bson:"buildingType"`
	CreatorID               int64     `json:"creatorId" bson:"creatorId"`
	EventType               string    `json:"eventType" bson:"eventType"`
	ItemAfter               int64     `json:"itemAfter" bson:"itemAfter"`
	ItemBefore              int64     `json:"itemBefore" bson:"itemBefore"`
	ItemID                  int64     `json:"itemId" bson:"itemId"`
	KillerID                int64     `json:"killerId" bson:"killerId"`
	LaneType                string    `json:"laneType" bson:"laneType"`
	LevelUpType             string    `json:"levelUpType" bson:"levelUpType"`
	MonsterType             string    `json:"monsterType" bson:"monsterType"`
	ParticipantID           int64     `json:"participantId" bson:"participantId"`
	Point64Captured         string    `json:"point64Captured" bson:"point64Captured"`
	Position                *Position `json:"position" bson:"position"`
	SkillSlot               int64     `json:"skillSlot" bson:"skillSlot"`
	TeamID                  int64     `json:"teamId" bson:"teamId"`
	Timestamp               int64     `json:"timestamp" bson:"timestamp"`
	TowerType               string    `json:"towerType" bson:"towerType"`
	VictimID                int64     `json:"victimId" bson:"victimId"`
	WardType                string    `json:"wardType" bson:"wardType"`
}

type ParticipantFrame struct {
	CurrentGold         int64     `json:"currentGold" bson:"currentGold"`
	DominionScore       int64     `json:"dominionScore" bson:"dominionScore"`
	JungleMinionsKilled int64     `json:"jungleMinionsKilled" bson:"jungleMinionsKilled"`
	Level               int64     `json:"level" bson:"level"`
	MinionsKilled       int64     `json:"minionsKilled" bson:"minionsKilled"`
	ParticipantID       int64     `json:"participantId" bson:"participantId"`
	Position            *Position `json:"position" bson:"position"`
	TeamScore           int64     `json:"teamScore" bson:"teamScore"`
	TotalGold           int64     `json:"totalGold" bson:"totalGold"`
	XP                  int64     `json:"xp" bson:"xp"`
}

type Position struct {
	X int64 `json:"x" bson:"x"`
	Y int64 `json:"y" bson:"y"`
}
