package main

import (
	"math"
	"sort"
)

type Elo interface {
	CalculateExpected(leftRating float64, rightRating float64) float64
	CalculateNewRating(leftRating float64, rightRating float64, scoreSign float64) float64
	CalculateRanking(matches map[int]*Match) []*PlayerRating
}

type eloImpl struct {
	settings *EloSettings
}

func NewElo(settings *EloSettings) Elo {
	return &eloImpl{
		settings: settings,
	}
}

func (e *eloImpl) CalculateExpected(leftRating float64, rightRating float64) float64 {
	exp := (leftRating - rightRating) / 400.0
	den := 1 + math.Pow(10, exp)
	return 1 / den
}

func (e *eloImpl) CalculateNewRating(leftRating float64, rightRating float64, scoreSign float64) float64 {
	expected := e.CalculateExpected(leftRating, rightRating)
	return leftRating + e.settings.KFactor*(scoreSign-expected)
}

func sortMatches(matches map[int]*Match) []*Match {
	sortedMatches := []*Match{}
	for _, value := range matches {
		sortedMatches = append(sortedMatches, value)
	}

	sort.Slice(sortedMatches, func(i, j int) bool {
		comp := sortedMatches[i].CompareByDateTime(sortedMatches[j])
		return comp < 0
	})
	return sortedMatches
}

func getPlayersFromMatches(matches []*Match) []string {
	dict := make(map[string]bool)

	for _, match := range matches {
		if !dict[match.LeftPlayer] {
			dict[match.LeftPlayer] = true
		}
		if !dict[match.RightPlayer] {
			dict[match.RightPlayer] = true
		}
	}

	keys := make([]string, len(dict))

	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}

	return keys
}

func (e *eloImpl) initRanking(players []string) map[string]*PlayerRating {
	dict := make(map[string]*PlayerRating)

	for _, player := range players {
		rating := &PlayerRating{
			Player: player,
			Rating: e.settings.InitRating,
		}
		dict[player] = rating
	}

	return dict
}

func getScoreSign(leftScore int, rightScore int) float64 {
	if leftScore > rightScore {
		return 1
	}

	if leftScore == rightScore {
		return 0.5
	}

	return 0
}

func (e *eloImpl) CalculateRanking(matches map[int]*Match) []*PlayerRating {
	sortedMatches := sortMatches(matches)
	allPlayers := getPlayersFromMatches(sortedMatches)
	rankingDict := e.initRanking(allPlayers)

	for _, match := range sortedMatches {
		leftRating := rankingDict[match.LeftPlayer].Rating
		rightRating := rankingDict[match.RightPlayer].Rating

		rankingDict[match.LeftPlayer].Rating = e.CalculateNewRating(leftRating, rightRating, getScoreSign(match.LeftPlayerScore, match.RightPlayerScore))
		rankingDict[match.RightPlayer].Rating = e.CalculateNewRating(rightRating, leftRating, getScoreSign(match.RightPlayerScore, match.LeftPlayerScore))
	}

	ranking := make([]*PlayerRating, len(rankingDict))

	i := 0
	for _, v := range rankingDict {
		ranking[i] = v
		i++
	}

	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].Rating > ranking[j].Rating
	})
	return ranking
}
