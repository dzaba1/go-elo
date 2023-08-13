package main

type MatchesProvider interface {
	GetMatches() (map[int]*Match, error)
	DeleteMatch(id int) error
	AddMatch(match *Match) (int, error)
}
