package handlers

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"globals"
	"log"
	"net/http"
)

func HandlerQTeams(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		Id   int64
		Name string
	}

	rows, _ := DB.Query(fmt.Sprintf(`
		select Team.id, Team.name
		from Team
		where Team.kindSport = %d;`,
		StateGame.GameKindID),
	)
	list := make([]*Resp, 0)

	if rows != nil {
		for rows.Next() {
			r := new(Resp)
			_ = rows.Scan(&r.Id, &r.Name)
			list = append(list, r)
		}
	}

	js, _ := json.Marshal(list)
	_, _ = w.Write(js)
}

func HandlerQEvents(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		Id   int64
		Name string
	}

	rows, _ := DB.Query(fmt.Sprintf(`
		select Event.id, Event.name
		from Event
		where Event.kindSport = %d;`,
		StateGame.GameKindID),
	)
	list := make([]*Resp, 0)

	if rows != nil {
		for rows.Next() {
			r := new(Resp)
			_ = rows.Scan(&r.Id, &r.Name)
			list = append(list, r)
		}
	}

	js, _ := json.Marshal(list)
	_, _ = w.Write(js)
}

func HandlerQItems(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		Id    int64
		Alias string
		Type  string
	}

	rows, _ := DB.Query(fmt.Sprintf(`
		select I.id, D.alias, D.type
		from Item I
		inner join Detail D on I.detail = D.id
		where I.event = %d;`,
		StateGame.CurEvent),
	)
	list := make([]*Resp, 0)

	if rows != nil {
		for rows.Next() {
			r := new(Resp)
			_ = rows.Scan(&r.Id, &r.Alias, &r.Type)
			list = append(list, r)
		}
	}

	js, _ := json.Marshal(list)
	_, _ = w.Write(js)
}

func HandlerQPerson(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		Id   int64
		Name string
		Team int64
	}

	rows, _ := DB.Query(fmt.Sprintf(`
		select P.id, P.name, P.team
		from Participant P
		where P.team = %d or P.team = %d;`,
		StateGame.Team1, StateGame.Team2),
	)

	list := make([]*Resp, 0)

	if rows != nil {
		for rows.Next() {
			r := new(Resp)
			_ = rows.Scan(&r.Id, &r.Name, &r.Team)
			list = append(list, r)
		}
	}

	js, _ := json.Marshal(list)
	_, _ = w.Write(js)
}

func HandlerQCurTeams(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		Team1 string
		Team2 string
	}

	rows, _ := DB.Query(fmt.Sprintf(`
		select T.id, T.name
		from Team T
		where T.id = %d or T.id = %d;`,
		StateGame.Team1, StateGame.Team2,
	))

	var t1n string
	var t2n string
	var t1i int64
	var t2i int64
	var js []byte

	rows.Next()
	_ = rows.Scan(&t1i, &t1n)

	rows.Next()
	_ = rows.Scan(&t2i, &t2n)

	if t1i != StateGame.Team1 {
		js, _ = json.Marshal(Resp{
			Team1: t2n,
			Team2: t1n,
		})
	} else {
		js, _ = json.Marshal(Resp{
			Team1: t1n,
			Team2: t2n,
		})
	}

	_, _ = w.Write(js)
}

func HandlerQGames(w http.ResponseWriter, r *http.Request) {
	type Resp struct {
		GameId   int64
		GameDate string
		Kind     string
		Teams    []string
	}

	rows, _ := DB.Query(fmt.Sprint(`
		select G.id, G.date, KS.name, array (
				select T.name from TeamsInGames TG
				inner join Team T on TG.team = T.id
				where TG.game = G.id
			)::text[] as Teams
		from Game G
		inner join KindSport KS on G.kindSport = KS.id
		order by G.date;`,
	))

	list := make([]*Resp, 0)

	for rows.Next() {
		resp := new(Resp)
		if err := rows.Scan(&resp.GameId,
			&resp.GameDate,
			&resp.Kind,
			(*pq.StringArray)(&resp.Teams));
			err != nil {
			log.Printf("err: %q", err)
			return
		}
		list = append(list, resp)
	}

	js, _ := json.Marshal(list)
	_, _ = w.Write(js)
}

func HandlerQHistoryEvent(w http.ResponseWriter, r *http.Request) {
	type AboutEvent struct {
		Id   int64
		Name string
		Time string
	}

	type AboutDetail struct {
		Type  string
		Value string
		Alias string
	}

	type Item struct {
		Event  AboutEvent
		Detail []AboutDetail
	}

	_ = r.ParseForm()

	if rows, err := DB.Query(fmt.Sprintf(`
		select HE.id as EventID, HE.time as EventTime, E.name as EventName,
			   D.type as DetailType, HD.value as Value, D.alias as DetailAlias
		from historydetail HD
		inner join historyevent HE on HD.historyEvent = HE.id
		inner join Event E on HE.event = E.id
		inner join Item I on HD.item = I.id
		inner join Detail D on I.detail = D.id
		where HE.game = %d
		order by HE.time;`, globals.GetInt64(r, "id"),
	)); err != nil {
		return
	} else {

		list := make([]*Item, 0)
		var lastEventId int64 = 0

		for rows.Next() {
			var detail = new(AboutDetail)
			var event = new(AboutEvent)

			err = rows.Scan(&event.Id, &event.Time, &event.Name,
				&detail.Type, &detail.Value, &detail.Alias)

			if err != nil {
				fmt.Printf("%q", err)
				return
			}

			if lastEventId == event.Id {
				item := list[0]
				item.Detail = append(item.Detail, *detail)

			} else {
				var item = Item{
					Event:  *event,
					Detail: make([]AboutDetail, 0),
				}
				item.Detail = append(item.Detail, *detail)
				list = append(list, &item)
				lastEventId = event.Id

			}
		}

		js, _ := json.Marshal(list)
		_, _ = w.Write(js)
	}
}
