package structs

import "time"

type Detail struct {
	Id    uint
	Alias string
	Type  string
}

type Event struct {
	Id        uint
	Name      string
	Kindsport *KindSport
}

type KindSport struct {
	Id   uint
	Name string
}

type Game struct {
	Id        uint
	Date      time.Time
	KindSport *KindSport
	Tourney   *Tourney
}

type Tourney struct {
	Id   uint
	Name string
}

type Teamsintourney struct {
	Id      uint
	Team    *Team
	Tourney *Tourney
	Point   int
}

type Team struct {
	Id        uint
	Name      string
	KindSport *KindSport
}

type Participant struct {
	Id      uint
	Name    string
	Tourney *Tourney
}


