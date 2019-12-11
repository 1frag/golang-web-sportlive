package using_db

import (
	//"fmt"
	"common"
	"structs"

	//"log"
)

var Db = common.GetDataBase()
var StateGame = common.GetStateGame()

func GetIdByKind(kind string) (int64, error) {
	//var id int
	//var query = fmt.Sprintf(`select id
	//				from KindSport
	//				where name='%s'
	//				limit 1;`, kind)
	row := new(structs.KindSport)
	err := Db.Model(row).
		Where("name = ?", kind).
		Limit(1).
		Select()
	//rows, err := Db.Query(query)
	//
	//if err != nil {
	//	return -1, err
	//} else {
	//	rows.Next()
	//	err = rows.Scan(&id)
	//
	//	if err != nil {
	//		return -1, err
	//	}
	//}
	return row.Id, err
}

func AddHistoryEvent(time string) int64 {
	//var id int64
	//query := fmt.Sprintf(`
	//		insert into historyevent
	//		(game, event, time)
	//		values (%d, %d, '%s')
	//		returning id;`,
	//	StateGame.GameID, StateGame.CurEvent, time,
	//)
	//rows, err := Db.Query(query)
	//if err != nil {
	//	log.Printf("Error$1: %q", err)
	//	return 0
	//}
	//rows.Next()
	//err = rows.Scan(&id)
	//if err != nil {
	//	log.Printf("Error$2: %q", err)
	//	return 0
	//}
	//return id
	return 0
}

func InsertHistoryDetail(eventId int64, itemId int64, value string) {
	//query := fmt.Sprintf(`
	//	insert into historydetail
	//	(historyEvent, item, value)
	//	values (%d, %d, '%s')
	//	returning id;`,
	//	eventId, itemId, value,
	//)
	//res, err := Db.Query(query)
	//if err != nil {
	//	fmt.Printf("err: %q", err)
	//} else {
	//	var lastid int64
	//	res.Next()
	//	err = res.Scan(&lastid)
	//	fmt.Printf("lastid: %d", lastid)
	//	if err != nil {
	//		fmt.Printf("err2: %q", err)
	//	}
	//}
}
