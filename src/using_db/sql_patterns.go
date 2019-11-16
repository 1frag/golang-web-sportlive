package using_db

import (
	"fmt"
	"globals"
)

var Db = globals.GetDataBase()
var StateGame = globals.GetStateGame()

func GetIdByKind(kind string) (int, error) {
	var id int
	var query = fmt.Sprintf(`select id
					from KindSport
					where name='%s'
					limit 1;`, kind)

	rows, err := Db.Query(query)

	if err != nil {
		return -1, err
	} else {
		rows.Next()
		err = rows.Scan(&id)

		if err != nil {
			return -1, err
		}
	}
	return id, nil
}

func AddHistoryEvent(time string) int64 {
	var id int64
	query := fmt.Sprintf(`
			insert into historyevent
			(game, event, time)
			values (%d, %d, '%s')
			returning id;`,
		StateGame.GameID, StateGame.CurEvent, time,
	)
	rows, _ := Db.Query(query)
	rows.Next()
	_ = rows.Scan(&id)
	return id
}

func InsertHistoryDetail(eventId int64, itemId int64, value string) {
	query := fmt.Sprintf(`
		insert into historydetail
		(historyEvent, item, value)
		values (%d, %d, '%s')
		returning id;`,
		eventId, itemId, value,
	)
	res, err := Db.Query(query)
	if err != nil {
		fmt.Printf("err: %q", err)
	} else {
		var lastid int64
		res.Next()
		err = res.Scan(&lastid)
		fmt.Printf("lastid: %d", lastid)
		if err != nil {
			fmt.Printf("err2: %q", err)
		}
	}
}
