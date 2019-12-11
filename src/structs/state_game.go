package structs

type State struct {
	GameID     int64
	GameName   string
	GameKindID int64
	Team1      int64
	Team2      int64
	CurEvent   int64
}
