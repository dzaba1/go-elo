package main

import (
	"log"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())

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
