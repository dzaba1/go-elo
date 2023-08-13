package main

import "time"

type Match struct {
	Id               *int      `json:"id"`
	DateTime         time.Time `json:"dateTime"`
	LeftPlayer       string    `json:"leftPlayer"`
	RightPlayer      string    `json:"rightPlayer"`
	LeftPlayerScore  int       `json:"leftPlayerScore"`
	RightPlayerScore int       `json:"rightPlayerScore"`
}

func (m *Match) CompareByDateTime(other *Match) int {
	return m.DateTime.Compare(other.DateTime)
}
