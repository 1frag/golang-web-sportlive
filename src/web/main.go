package main

import (
	"fmt"
	"log"
	"net/http"
	"web/handlers"
)

func main() {

	/* для постов */
	http.HandleFunc("/api/", handlers.HandlerApi)

	/* дай страничку */
	http.HandleFunc("/A/", handlers.HandlerA)
	http.HandleFunc("/B/", handlers.HandlerB)
	http.HandleFunc("/C/", handlers.HandlerC)
	http.HandleFunc("/D/", handlers.HandlerD)
	http.HandleFunc("/H/", handlers.HandlerH)
	http.HandleFunc("/G/", handlers.HandlerG)

	/* селекты к базе */
	http.HandleFunc("/queries/teams/", handlers.HandlerQTeams)
	http.HandleFunc("/queries/events/", handlers.HandlerQEvents)
	http.HandleFunc("/queries/items/", handlers.HandlerQItems)
	http.HandleFunc("/queries/persons/", handlers.HandlerQPerson)
	http.HandleFunc("/queries/cur-teams/", handlers.HandlerQCurTeams)
	http.HandleFunc("/queries/games/", handlers.HandlerQGames)
	http.HandleFunc("/queries/history-events/", handlers.HandlerQHistoryEvent)

	/* инсерты в базу */
	http.HandleFunc("/insert/game/", handlers.HandlerIGame)
	http.HandleFunc("/insert/events/", handlers.HandlerIEvent)

	/* статика */
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Здесь могла бы быть ваша реклама
	fmt.Println("Serving webpage at http://0.0.0.0:8000/")
	err := http.ListenAndServe("0.0.0.0:8000", nil)

	if err != nil {
		log.Fatalf("Failed to listen: %q", err)
	}

}
