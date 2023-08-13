package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Ping(c *gin.Context)
	Elo(c *gin.Context)
	GetMatches(c *gin.Context)
	DeleteMatch(c *gin.Context)
	NewMatch(c *gin.Context)
	LeagueName(c *gin.Context)
}

type serviceImpl struct {
	elo             Elo
	matchesProvider MatchesProvider
}

func NewService(elo Elo, matchesProvider MatchesProvider) Service {
	return &serviceImpl{
		elo:             elo,
		matchesProvider: matchesProvider,
	}
}

func (s *serviceImpl) Ping(c *gin.Context) {
	log.Println("Ping")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *serviceImpl) Elo(c *gin.Context) {
	log.Println("Elo")
	matches, err := s.matchesProvider.GetMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ranking := s.elo.CalculateRanking(matches)

	c.JSON(http.StatusOK, ranking)
}

func (s *serviceImpl) GetMatches(c *gin.Context) {
	log.Println("Get matches")
	matches, err := s.matchesProvider.GetMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func (s *serviceImpl) DeleteMatch(c *gin.Context) {
	log.Println("Delete matches")
	idVal := c.Params.ByName("id")
	id, err := strconv.Atoi(idVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.matchesProvider.DeleteMatch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *serviceImpl) NewMatch(c *gin.Context) {
	log.Println("New match")
	var newMatch Match

	err := c.BindJSON(&newMatch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = s.matchesProvider.AddMatch(&newMatch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &newMatch)
}

func (s *serviceImpl) LeagueName(c *gin.Context) {
	env := os.Getenv("ELO_LEAGUE_NAME")
	c.JSON(http.StatusOK, gin.H{
		"name": env,
	})
}
