package handlers

import (
	"common"
	"net/http"
	"using_db"
)

var DB = common.GetDataBase()
var StateGame = common.GetStateGame()

func HandlerApi(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	curState := r.FormValue("state")
	switch curState {
	case "InitGame":
		StateGame.GameName = r.FormValue("game")
		res, _ := using_db.GetIdByKind(StateGame.GameName)
		StateGame.GameKindID = int64(res)
		break
	case "ResetGame":
		StateGame.GameName = ""
		StateGame.GameKindID = 0
		break
	case "DateAndTeams":
		StateGame.Team1 = common.GetInt64(r, "team1")
		StateGame.Team2 = common.GetInt64(r, "team2")
		break
	case "ChangeEvent":
		StateGame.CurEvent = common.GetInt64(r, "event")
		break
	}
	_, _ = w.Write([]byte("OK"))
}
