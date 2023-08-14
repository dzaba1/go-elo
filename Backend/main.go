package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	eloSettings, err := GetEloSettingsFromEnvs()
	if err != nil {
		log.Fatal(err)
	}

	elo := NewElo(eloSettings)
	matchesProvider := NewCsvMatchesProvider("db/matches.csv")
	service := NewService(elo, matchesProvider)

	r := gin.Default()
	r.GET("/ping", service.Ping)
	r.GET("/match", service.GetMatches)
	r.DELETE("/match/:id", service.DeleteMatch)
	r.POST("/match", service.NewMatch)
	r.GET("/elo", service.Elo)
	r.GET("/name", service.LeagueName)

	print("Starting the app.\n")
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
