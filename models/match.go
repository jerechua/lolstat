package models

import "gopkg.in/mgo.v2/bson"

type Match struct {
	ID                    bson.ObjectId          `bson:"_id,omitempty"`
	MapID                 int64                  `json:"mapId" bson:"mapId,omitempty"`
	MatchCreation         int64                  `json:"matchCreation" bson:"matchCreation,omitempty"`
	MatchDuration         int64                  `json:"matchDuration" bson:"matchDuration,omitempty"`
	MatchID               int64                  `json:"matchId" bson:"matchId,omitempty"`
	MatchMode             string                 `json:"matchMode" bson:"matchMode,omitempty"`
	MatchType             string                 `json:"matchType" bson:"matchType,omitempty"`
	MatchVersion          string                 `json:"matchVersion" bson:"matchVersion,omitempty"`
	ParticipantIdentities []*ParticipantIdentity `json:"participantIdentities" bson:"participantIdentities,omitempty"`
	Participants          []*Participant         `json:"participants" bson:"participants,omitempty"`
	PlatformID            string                 `json:"platformId" bson:"platformId,omitempty"`
	QueueType             string                 `json:"queueType" bson:"queueType,omitempty"`
	Region                string                 `json:"region" bson:"region,omitempty"`
	Season                string                 `json:"season" bson:"season,omitempty"`
	Teams                 []*Team                `json:"teams" bson:"teams,omitempty"`
	Timeline              *Timeline              `json:"timeline" bson:"timeline,omitempty"`
}

type Participant struct {
	ChampionID                int64                `json:"championId" bson:"championId,omitempty"`
	HighestAchievedSeasonTier string               `json:"highestAchievedSeasonTier" bson:"highestAchievedSeasonTier,omitempty"`
	Masteries                 []*Mastery           `json:"masteries" bson:"masteries,omitempty"`
	ParticipantID             int64                `json:"participantId" bson:"participantId,omitempty"`
	Runes                     []*Rune              `json:"runes" bson:"runes,omitempty"`
	Spell1ID                  int64                `json:"spell1Id" bson:"spell1Id,omitempty"`
	Spell2ID                  int64                `json:"spell2Id" bson:"spell2Id,omitempty"`
	Stats                     *ParticipantStats    `json:"stats" bson:"stats,omitempty"`
	TeamID                    int64                `json:"teamId" bson:"teamId,omitempty"`
	Timeline                  *ParticipantTimeline `json:"timeline" bson:"timeline,omitempty"`
}

type ParticipantIdentity struct {
	ParticipantID int    `json:"participantId" bson:"participantId,omitempty"`
	Player        Player `json:"player" bson:"player,omitempty"`
}

type Team struct {
	Bans                 []*BannedChampion `json:"bans" bson:"bans,omitempty"`
	BaronKills           int64             `json:"baronKills" bson:"baronKills,omitempty"`
	DominionVictoryScore int64             `json:"dominionVictoryScore" bson:"dominionVictoryScore,omitempty"`
	DragonKills          int64             `json:"dragonKills" bson:"dragonKills,omitempty"`
	FirstBaron           bool              `json:"firstBaron" bson:"firstBaron,omitempty"`
	FirstBlood           bool              `json:"firstBlood" bson:"firstBlood,omitempty"`
	FirstDragon          bool              `json:"firstDragon" bson:"firstDragon,omitempty"`
	FirstInhibitor       bool              `json:"firstInhibitor" bson:"firstInhibitor,omitempty"`
	FirstRiftHerald      bool              `json:"firstRiftHerald" bson:"firstRiftHerald,omitempty"`
	FirstTower           bool              `json:"firstTower" bson:"firstTower,omitempty"`
	InhibitorKills       int64             `json:"inhibitorKills" bson:"inhibitorKills,omitempty"`
	RiftHeraldKills      int64             `json:"riftHeraldKills" bson:"riftHeraldKills,omitempty"`
	TeamID               int64             `json:"teamId" bson:"teamId,omitempty"`
	TowerKills           int64             `json:"towerKills" bson:"towerKills,omitempty"`
	VilemawKills         int64             `json:"vilemawKills" bson:"vilemawKills,omitempty"`
	Winner               bool              `json:"winner" bson:"winner,omitempty"`
}

type Timeline struct {
	FrameInterval int64    `json:"frameInterval" bson:"frameInterval,omitempty"`
	Frames        []*Frame `json:"frames" bson:"frames,omitempty"`
}

type Mastery struct {
	MasteryID int64 `json:"masteryId" bson:"masteryId,omitempty"`
	Rank      int64 `json:"rank" bson:"rank,omitempty"`
}

type ParticipantStats struct {
	Assists                         int64 `json:"assists" bson:"assists,omitempty"`
	ChampLevel                      int64 `json:"champLevel" bson:"champLevel,omitempty"`
	CombatPlayerScore               int64 `json:"combatPlayerScore" bson:"combatPlayerScore,omitempty"`
	Deaths                          int64 `json:"deaths" bson:"deaths,omitempty"`
	DoubleKills                     int64 `json:"doubleKills" bson:"doubleKills,omitempty"`
	FirstBloodAssist                bool  `json:"firstBloodAssist" bson:"firstBloodAssist,omitempty"`
	FirstBloodKill                  bool  `json:"firstBloodKill" bson:"firstBloodKill,omitempty"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist" bson:"firstInhibitorAssist,omitempty"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill" bson:"firstInhibitorKill,omitempty"`
	FirstTowerAssist                bool  `json:"firstTowerAssist" bson:"firstTowerAssist,omitempty"`
	FirstTowerKill                  bool  `json:"firstTowerKill" bson:"firstTowerKill,omitempty"`
	GoldEarned                      int64 `json:"goldEarned" bson:"goldEarned,omitempty"`
	GoldSpent                       int64 `json:"goldSpent" bson:"goldSpent,omitempty"`
	InhibitorKills                  int64 `json:"inhibitorKills" bson:"inhibitorKills,omitempty"`
	Item0                           int64 `json:"item0" bson:"item0,omitempty"`
	Item1                           int64 `json:"item1" bson:"item1,omitempty"`
	Item2                           int64 `json:"item2" bson:"item2,omitempty"`
	Item3                           int64 `json:"item3" bson:"item3,omitempty"`
	Item4                           int64 `json:"item4" bson:"item4,omitempty"`
	Item5                           int64 `json:"item5" bson:"item5,omitempty"`
	Item6                           int64 `json:"item6" bson:"item6,omitempty"`
	KillingSprees                   int64 `json:"killingSprees" bson:"killingSprees,omitempty"`
	Kills                           int64 `json:"kills" bson:"kills,omitempty"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike" bson:"largestCriticalStrike,omitempty"`
	LargestKillingSpree             int64 `json:"largestKillingSpree" bson:"largestKillingSpree,omitempty"`
	LargestMultiKill                int64 `json:"largestMultiKill" bson:"largestMultiKill,omitempty"`
	MagicDamageDealt                int64 `json:"magicDamageDealt" bson:"magicDamageDealt,omitempty"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions" bson:"magicDamageDealtToChampions,omitempty"`
	MagicDamageTaken                int64 `json:"magicDamageTaken" bson:"magicDamageTaken,omitempty"`
	MinionsKilled                   int64 `json:"minionsKilled" bson:"minionsKilled,omitempty"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled" bson:"neutralMinionsKilled,omitempty"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle" bson:"neutralMinionsKilledEnemyJungle,omitempty"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle" bson:"neutralMinionsKilledTeamJungle,omitempty"`
	NodeCapture                     int64 `json:"nodeCapture" bson:"nodeCapture,omitempty"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist" bson:"nodeCaptureAssist,omitempty"`
	NodeNeutralize                  int64 `json:"nodeNeutralize" bson:"nodeNeutralize,omitempty"`
	NodeNeutralizeAssist            int64 `json:"nodeNeutralizeAssist" bson:"nodeNeutralizeAssist,omitempty"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore" bson:"objectivePlayerScore,omitempty"`
	PentaKills                      int64 `json:"pentaKills" bson:"pentaKills,omitempty"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt" bson:"physicalDamageDealt,omitempty"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions" bson:"physicalDamageDealtToChampions,omitempty"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken" bson:"physicalDamageTaken,omitempty"`
	QuadraKills                     int64 `json:"quadraKills" bson:"quadraKills,omitempty"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame" bson:"sightWardsBoughtInGame,omitempty"`
	TeamObjective                   int64 `json:"teamObjective" bson:"teamObjective,omitempty"`
	TotalDamageDealt                int64 `json:"totalDamageDealt" bson:"totalDamageDealt,omitempty"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions" bson:"totalDamageDealtToChampions,omitempty"`
	TotalDamageTaken                int64 `json:"totalDamageTaken" bson:"totalDamageTaken,omitempty"`
	TotalHeal                       int64 `json:"totalHeal" bson:"totalHeal,omitempty"`
	TotalPlayerScore                int64 `json:"totalPlayerScore" bson:"totalPlayerScore,omitempty"`
	TotalScoreRank                  int64 `json:"totalScoreRank" bson:"totalScoreRank,omitempty"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt" bson:"totalTimeCrowdControlDealt,omitempty"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed" bson:"totalUnitsHealed,omitempty"`
	TowerKills                      int64 `json:"towerKills" bson:"towerKills,omitempty"`
	TripleKills                     int64 `json:"tripleKills" bson:"tripleKills,omitempty"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt" bson:"trueDamageDealt,omitempty"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions" bson:"trueDamageDealtToChampions,omitempty"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken" bson:"trueDamageTaken,omitempty"`
	UnrealKills                     int64 `json:"unrealKills" bson:"unrealKills,omitempty"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame" bson:"visionWardsBoughtInGame,omitempty"`
	WardsKilled                     int64 `json:"wardsKilled" bson:"wardsKilled,omitempty"`
	WardsPlaced                     int64 `json:"wardsPlaced" bson:"wardsPlaced,omitempty"`
	Winner                          bool  `json:"winner" bson:"winner,omitempty"`
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts *ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts" bson:"ancientGolemAssistsPerMinCounts,omitempty"`
	AncientGolemKillsPerMinCounts   *ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts" bson:"ancientGolemKillsPerMinCounts,omitempty"`
	AssistedLaneDeathsPerMinDeltas  *ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas" bson:"assistedLaneDeathsPerMinDeltas,omitempty"`
	AssistedLaneKillsPerMinDeltas   *ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas" bson:"assistedLaneKillsPerMinDeltas,omitempty"`
	BaronAssistsPerMinCounts        *ParticipantTimelineData `json:"baronAssistsPerMinCounts" bson:"baronAssistsPerMinCounts,omitempty"`
	BaronKillsPerMinCounts          *ParticipantTimelineData `json:"baronKillsPerMinCounts" bson:"baronKillsPerMinCounts,omitempty"`
	CreepsPerMinDeltas              *ParticipantTimelineData `json:"creepsPerMinDeltas" bson:"creepsPerMinDeltas,omitempty"`
	CsDiffPerMinDeltas              *ParticipantTimelineData `json:"csDiffPerMinDeltas" bson:"csDiffPerMinDeltas,omitempty"`
	DamageTakenDiffPerMinDeltas     *ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas" bson:"damageTakenDiffPerMinDeltas,omitempty"`
	DamageTakenPerMinDeltas         *ParticipantTimelineData `json:"damageTakenPerMinDeltas" bson:"damageTakenPerMinDeltas,omitempty"`
	DragonAssistsPerMinCounts       *ParticipantTimelineData `json:"dragonAssistsPerMinCounts" bson:"dragonAssistsPerMinCounts,omitempty"`
	DragonKillsPerMinCounts         *ParticipantTimelineData `json:"dragonKillsPerMinCounts" bson:"dragonKillsPerMinCounts,omitempty"`
	ElderLizardAssistsPerMinCounts  *ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts" bson:"elderLizardAssistsPerMinCounts,omitempty"`
	ElderLizardKillsPerMinCounts    *ParticipantTimelineData `json:"elderLizardKillsPerMinCounts" bson:"elderLizardKillsPerMinCounts,omitempty"`
	GoldPerMinDeltas                *ParticipantTimelineData `json:"goldPerMinDeltas" bson:"goldPerMinDeltas,omitempty"`
	InhibitorAssistsPerMinCounts    *ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts" bson:"inhibitorAssistsPerMinCounts,omitempty"`
	InhibitorKillsPerMinCounts      *ParticipantTimelineData `json:"inhibitorKillsPerMinCounts" bson:"inhibitorKillsPerMinCounts,omitempty"`
	Lane                            string                   `json:"lane" bson:"lane,omitempty"`
	Role                            string                   `json:"role" bson:"role,omitempty"`
	TowerAssistsPerMinCounts        *ParticipantTimelineData `json:"towerAssistsPerMinCounts" bson:"towerAssistsPerMinCounts,omitempty"`
	TowerKillsPerMinCounts          *ParticipantTimelineData `json:"towerKillsPerMinCounts" bson:"towerKillsPerMinCounts,omitempty"`
	TowerKillsPerMinDeltas          *ParticipantTimelineData `json:"towerKillsPerMinDeltas" bson:"towerKillsPerMinDeltas,omitempty"`
	VilemawAssistsPerMinCounts      *ParticipantTimelineData `json:"vilemawAssistsPerMinCounts" bson:"vilemawAssistsPerMinCounts,omitempty"`
	VilemawKillsPerMinCounts        *ParticipantTimelineData `json:"vilemawKillsPerMinCounts" bson:"vilemawKillsPerMinCounts,omitempty"`
	WardsPerMinDeltas               *ParticipantTimelineData `json:"wardsPerMinDeltas" bson:"wardsPerMinDeltas,omitempty"`
	XpDiffPerMinDeltas              *ParticipantTimelineData `json:"xpDiffPerMinDeltas" bson:"xpDiffPerMinDeltas,omitempty"`
	XpPerMinDeltas                  *ParticipantTimelineData `json:"xpPerMinDeltas" bson:"xpPerMinDeltas,omitempty"`
}

type Rune struct {
	Rank   int64 `json:"rank" bson:"rank,omitempty"`
	RuneID int64 `json:"runeId" bson:"runeId,omitempty"`
}

type Player struct {
	MatchHistoryUri string `json:"matchHistoryUri" bson:"matchHistoryUri,omitempty"`
	ProfileIcon     int64  `json:"profileIcon" bson:"profileIcon,omitempty"`
	SummonerID      int64  `json:"summonerId" bson:"summonerId,omitempty"`
	SummonerName    string `json:"summonerName" bson:"summonerName,omitempty"`
}

type BannedChampion struct {
	ChampionID int64 `json:"championId" bson:"championId,omitempty"`
	PickTurn   int64 `json:"pickTurn" bson:"pickTurn,omitempty"`
}

type Frame struct {
	Events            []*Event                     `json:"events" bson:"events,omitempty"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames" bson:"participantFrames,omitempty"`
	Timestamp         int64                        `json:"timestamp" bson:"timestamp,omitempty"`
}

type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty" bson:"tenToTwenty,omitempty"`
	ThirtyToEnd    float64 `json:"thirtyToEnd" bson:"thirtyToEnd,omitempty"`
	TwentyToThirty float64 `json:"twentyToThirty" bson:"twentyToThirty,omitempty"`
	ZeroToTen      float64 `json:"zeroToTen" bson:"zeroToTen,omitempty"`
}

type Event struct {
	AscendedType            string    `json:"ascendedType" bson:"ascendedType,omitempty"`
	AssistingParticipantIds []int64   `json:"assistingParticipantIds" bson:"assistingParticipantIds,omitempty"`
	BuildingType            string    `json:"buildingType" bson:"buildingType,omitempty"`
	CreatorID               int64     `json:"creatorId" bson:"creatorId,omitempty"`
	EventType               string    `json:"eventType" bson:"eventType,omitempty"`
	ItemAfter               int64     `json:"itemAfter" bson:"itemAfter,omitempty"`
	ItemBefore              int64     `json:"itemBefore" bson:"itemBefore,omitempty"`
	ItemID                  int64     `json:"itemId" bson:"itemId,omitempty"`
	KillerID                int64     `json:"killerId" bson:"killerId,omitempty"`
	LaneType                string    `json:"laneType" bson:"laneType,omitempty"`
	LevelUpType             string    `json:"levelUpType" bson:"levelUpType,omitempty"`
	MonsterType             string    `json:"monsterType" bson:"monsterType,omitempty"`
	ParticipantID           int64     `json:"participantId" bson:"participantId,omitempty"`
	Point64Captured         string    `json:"point64Captured" bson:"point64Captured,omitempty"`
	Position                *Position `json:"position" bson:"position,omitempty"`
	SkillSlot               int64     `json:"skillSlot" bson:"skillSlot,omitempty"`
	TeamID                  int64     `json:"teamId" bson:"teamId,omitempty"`
	Timestamp               int64     `json:"timestamp" bson:"timestamp,omitempty"`
	TowerType               string    `json:"towerType" bson:"towerType,omitempty"`
	VictimID                int64     `json:"victimId" bson:"victimId,omitempty"`
	WardType                string    `json:"wardType" bson:"wardType,omitempty"`
}

type ParticipantFrame struct {
	CurrentGold         int64     `json:"currentGold" bson:"currentGold,omitempty"`
	DominionScore       int64     `json:"dominionScore" bson:"dominionScore,omitempty"`
	JungleMinionsKilled int64     `json:"jungleMinionsKilled" bson:"jungleMinionsKilled,omitempty"`
	Level               int64     `json:"level" bson:"level,omitempty"`
	MinionsKilled       int64     `json:"minionsKilled" bson:"minionsKilled,omitempty"`
	ParticipantID       int64     `json:"participantId" bson:"participantId,omitempty"`
	Position            *Position `json:"position" bson:"position,omitempty"`
	TeamScore           int64     `json:"teamScore" bson:"teamScore,omitempty"`
	TotalGold           int64     `json:"totalGold" bson:"totalGold,omitempty"`
	XP                  int64     `json:"xp" bson:"xp,omitempty"`
}

type Position struct {
	X int64 `json:"x" bson:"x,omitempty"`
	Y int64 `json:"y" bson:"y,omitempty"`
}
