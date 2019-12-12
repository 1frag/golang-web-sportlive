package structs

import "time"

type Detail struct {
	Id    uint
	Alias string
	Type  string
}

func (Detail) TableName() string {
	return "detail"
}

// --------------------------------------------------------

type Event struct {
	Id        uint
	Name      string
	Kindsport *KindSport
}

func (Event) TableName() string {
	return "event"
}

// --------------------------------------------------------

type KindSport struct {
	Id   uint
	Name string
}

func (KindSport) TableName() string {
	return "kindsport"
}

// --------------------------------------------------------

type Game struct {
	Id        uint
	Date      time.Time
	KindSport *KindSport
	Tourney   *Tourney
}

func (Game) TableName() string {
	return "game"
}

// --------------------------------------------------------

type Tourney struct {
	Id   uint
	Name string
}

func (Tourney) TableName() string {
	return "tourney"
}

// --------------------------------------------------------

type Teamsintourney struct {
	Id      uint
	Team    *Team
	Tourney *Tourney
	Point   int
}

func (Teamsintourney) TableName() string {
	return "teamsintourney"
}

// --------------------------------------------------------

type Team struct {
	Id        uint
	Name      string
	KindSport *KindSport
}

func (Team) TableName() string {
	return "team"
}

// --------------------------------------------------------
type Participant struct {
	Id      uint
	Name    string
	Tourney *Tourney
}

func (Participant) TableName() string {
	return "participant"
}

// --------------------------------------------------------

type Item struct {
	Id     uint
	detail *Detail
	event  *Event
}

func (Item) TableName() string {
	return "item"
}

// --------------------------------------------------------

type Historydetail struct {
	Id           uint
	Historyevent *Historyevent
	Item         *Item
	value        string
}

func (Historydetail) TableName() string {
	return "historydetail"
}

// --------------------------------------------------------

type Historyevent struct {
	Id    uint
	Game  *Game
	Time  time.Duration
	Event *Event
}

func (Historyevent) TableName() string {
	return "historyevent"
}

// --------------------------------------------------------

type User struct {
	Id       uint
	Username string
	Password string
	Status   int
}

func (User) TableName() string {
	return "user"
}

// --------------------------------------------------------
