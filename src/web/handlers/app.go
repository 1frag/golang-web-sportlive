package handlers

import (
	"fmt"
	"log"
	"net/http"
	"structs"
)

func LoginMeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Print(username, password)

	log.Print(DB.Model(&structs.User{}).
		Where("username = '?' and password = '?'", username, password).
		Select("id, status").QueryExpr())

	rows, err := DB.Model(&structs.User{}).
		Where("username = ? and password = ?", username, password).
		Select("id, status").Rows()

	if err != nil {
		log.Print("w:", err)
		return
	}

	for rows.Next() {
		data := new(structs.User)
		err = rows.Scan(&data.Id, &data.Status)
		if err != nil {
			log.Print(err)
			return
		}
		_, _ = w.Write([]byte(fmt.Sprintf("{\"resp\": %d}", data.Status)))
		return
	}
	_, _ = w.Write([]byte("{\"resp\": 401}"))

}

func HandlerRowSql(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}

	right := r.FormValue("right")
	query := r.FormValue("query")
	pref := ""
	if right == "2" {
		pref = `begin transaction;
				set transaction read write; `
	} else {
		pref = `begin transaction;
				set transaction read only; `
	}
	log.Print(pref + query + "; end;")

	DB.Exec(pref + query + "; end;")
	_, _ = w.Write([]byte("{\"status\": \"Not Implemented\"}"))

}