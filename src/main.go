package main

import (
	"fmt"
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	/* изменение состояния игры */
	http.HandleFunc("/api/", handlers.HandlerApi)

	/* дай страничку */
	http.HandleFunc("/A/", handlers.HandlerA)
	http.HandleFunc("/B/", handlers.HandlerB)
	http.HandleFunc("/C/", handlers.HandlerC)
	http.HandleFunc("/D/", handlers.HandlerD)
	http.HandleFunc("/H/", handlers.HandlerH)
	http.HandleFunc("/G/", handlers.HandlerG)
	http.HandleFunc("/J/", handlers.HandlerJ)

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
	http.HandleFunc("/insert/team/", handlers.HandlerITeam)

	/* статика */
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Здесь могла бы быть ваша реклама
	port := os.Getenv("PORT")
	fmt.Println("Serving webpage at http://0.0.0.0:8000/")
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Fatalf("Failed to listen: %q", err)
	}

}
