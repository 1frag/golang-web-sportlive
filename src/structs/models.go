package structs

import "time"

type BaseModel struct {
	Id int64 `pg:"id"`
}

type KindSport struct {
	tableName struct{} `pg:"kindSport"`
	BaseModel
	Name string
}

type Tourney struct {
	TableName struct{} `pg:"tourney"`
	BaseModel
	Name string
}

type Game struct {
	TableName struct{} `pg:"game"`
	BaseModel
	date        *time.Time
	KindSportId *KindSport
	TourneyId   *Tourney
}

type Team struct {
	tableName struct{} `pg:"team"`
	BaseModel
	Name        string
	KindSportId *KindSport `pg:"kindsport,fk"`
}

type Event struct {
	TableName struct{} `pg:"event"`
	BaseModel
	Name        string
	KindSportId *KindSport
}

type Detail struct {
	TableName struct{} `pg:"detail"`
	BaseModel
	Alias string
	Type  string
}

type Item struct {
	TableName struct{} `pg:"item"`
	BaseModel
	DetailId *Detail
	EventId *Event
}

type HistoryEvent struct {
	TableName struct{} `pg:"historyEvent"`
	BaseModel
	GameId *Game
	EventId *Event
	Time *time.Time
}

type TeamsInGame struct {
	TableName struct{} `pg:"teamsInGame"`
	BaseModel
	GameId *Game
	TeamId *Team
}

type Participant struct {
	TableName struct{} `pg:"participant"`
	BaseModel
	Name string
	TeamId *Team
}

type TeamsInTourney struct {
	TableName struct{} `pg:"teamsInTourney"`
	BaseModel
	TeamId *Team
	TourneyId *Tourney
	Point int
}

type HistoryDetail struct {
	TableName struct{} `pg:"historyDetail"`
	BaseModel
	HistoryEventId *HistoryEvent
	ItemId *Item
	Value string
}
