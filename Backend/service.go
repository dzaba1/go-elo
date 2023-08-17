package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
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

func setErrorResp(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"error": msg,
	})
}

func (s *serviceImpl) Elo(c *gin.Context) {
	log.Println("Elo")
	matches, err := s.matchesProvider.GetMatches()
	if err != nil {
		setErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	ranking := s.elo.CalculateRanking(matches)

	c.JSON(http.StatusOK, ranking)
}

func (s *serviceImpl) GetMatches(c *gin.Context) {
	log.Println("Get matches")
	matches, err := s.matchesProvider.GetMatches()
	if err != nil {
		setErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := GetValues(matches)

	c.JSON(http.StatusOK, result)
}

func (s *serviceImpl) DeleteMatch(c *gin.Context) {
	log.Println("Delete matches")
	idVal := c.Params.ByName("id")
	id, err := strconv.Atoi(idVal)
	if err != nil {
		setErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}

	err = s.matchesProvider.DeleteMatch(id)
	if err != nil {
		setErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func validateNewMatch(match *Match) *string {
	leftP := strings.TrimSpace(match.LeftPlayer)
	rightP := strings.TrimSpace(match.RightPlayer)

	if leftP == "" {
		msg := "Left player name is required"
		return &msg
	}

	if rightP == "" {
		msg := "Right player name is required"
		return &msg
	}

	if rightP == leftP {
		msg := "Players must be different"
		return &msg
	}

	return nil
}

func (s *serviceImpl) NewMatch(c *gin.Context) {
	log.Println("New match")
	var newMatch Match

	err := c.BindJSON(&newMatch)
	if err != nil {
		setErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}

	validationMsg := validateNewMatch(&newMatch)
	if validationMsg != nil {
		setErrorResp(c, http.StatusBadRequest, *validationMsg)
		return
	}

	_, err = s.matchesProvider.AddMatch(&newMatch)
	if err != nil {
		setErrorResp(c, http.StatusInternalServerError, err.Error())
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
