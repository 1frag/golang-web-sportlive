package structs

type KindSportData struct {
	Title   string
	AsRowKS string
	AsAdj   string
}

type RespGet struct {
	Obj     string
	Kind    string
	GameID  int64
	EventID int
}

type RespPost struct {
	Result int
}

type Detail struct {
	Id    int
	Type  string
	Alias string
}

type State struct {
	GameID     int64
	GameName   string
	GameKindID int64
	Team1      int64
	Team2      int64
	CurEvent   int64
}
