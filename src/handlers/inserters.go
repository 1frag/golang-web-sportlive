package handlers

import (
	"encoding/json"
	"fmt"
	"globals"
	"net/http"
	"strconv"
	"using_db"
)

func HandlerIGame(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	t1 := globals.GetInt64(r, "team1")
	t2 := globals.GetInt64(r, "team2")
	date := r.FormValue("date")

	res, _ := DB.Query(fmt.Sprintf(`
		insert into Game (date, kindSport)
		values ('%s', %d)
		returning id;`,
		date, StateGame.GameKindID),
	)
	res.Next()
	_ = res.Scan(&StateGame.GameID)

	_, _ = DB.Query(fmt.Sprintf(`
		insert into TeamsInGames (game, team)
		values (%d, %d), (%d, %d);`,
		StateGame.GameID, t1, StateGame.GameID, t2),
	)

	_, _ = w.Write([]byte("OK"))

}

func HandlerIEvent(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	query := []byte(r.FormValue("apply"))
	time := r.FormValue("time")
	var clean = make(map[string]string)

	if err := json.Unmarshal(query, &clean); err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("Error: %q", err)))
		return
	}

	idEvent := using_db.AddHistoryEvent(time)

	for k, v := range clean {
		kasint64, _ := strconv.Atoi(k)
		using_db.InsertHistoryDetail(idEvent, int64(kasint64), v)
	}

	_, _ = w.Write([]byte("OK"))

}
