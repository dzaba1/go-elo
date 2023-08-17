package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const DateTimeLayout = "02.01.2006 15:04"

type csvMatchesProvider struct {
	filename string
	mu       *sync.Mutex
}

func NewCsvMatchesProvider(filename string) MatchesProvider {
	return &csvMatchesProvider{
		filename: filename,
		mu:       &sync.Mutex{},
	}
}

func deserializeRecord(line []string) (*Match, error) {
	var rec Match

	for columnNum, cell := range line {
		if columnNum == 0 {
			value, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			rec.Id = &value
		} else if columnNum == 1 {
			value, err := time.Parse(DateTimeLayout, cell)
			if err != nil {
				return nil, err
			}
			rec.DateTime = value
		} else if columnNum == 2 {
			rec.LeftPlayer = cell
		} else if columnNum == 3 {
			rec.RightPlayer = cell
		} else if columnNum == 4 {
			value, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			rec.LeftPlayerScore = value
		} else if columnNum == 5 {
			value, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			rec.RightPlayerScore = value
		}
	}

	return &rec, nil
}

func serializeRecord(match *Match) []string {
	data := make([]string, 6)

	data[0] = strconv.Itoa(*match.Id)
	data[1] = match.DateTime.Format(DateTimeLayout)
	data[2] = strings.TrimSpace(match.LeftPlayer)
	data[3] = strings.TrimSpace(match.RightPlayer)
	data[4] = strconv.Itoa(match.LeftPlayerScore)
	data[5] = strconv.Itoa(match.RightPlayerScore)

	return data
}

func deserialize(data [][]string) ([]*Match, error) {
	result := []*Match{}

	for lineNum, line := range data {
		if lineNum > 0 { // omit header line
			rec, err := deserializeRecord(line)
			if err != nil {
				return nil, err
			}
			result = append(result, rec)
		}
	}

	return result, nil
}

func serialize(matches []*Match) [][]string {
	result := make([][]string, len(matches))
	for i, match := range matches {
		data := serializeRecord(match)
		result[i] = data
	}
	return result
}

func (p *csvMatchesProvider) deserializeList() ([]*Match, error) {
	log.Printf("Opening %s for read.\n", p.filename)
	f, err := os.Open(p.filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return deserialize(data)
}

func (p *csvMatchesProvider) serializeList(matches []*Match) error {
	log.Printf("Opening %s for create.\n", p.filename)
	f, err := os.Create(p.filename)
	if err != nil {
		return err
	}

	defer f.Close()

	csvWriter := csv.NewWriter(f)
	records := serialize(matches)
	header := []string{"Id", "DateTime", "LeftPlayer", "RightPlayer", "LeftPlayerScore", "RightPlayerScore"}
	data := [][]string{header}
	data = append(data, records...)

	err = csvWriter.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *csvMatchesProvider) serialize(matches map[int]*Match) error {
	values := GetValues(matches)

	return p.serializeList(values)
}

func (p *csvMatchesProvider) getMatchesInternal() (map[int]*Match, error) {
	matchesList, err := p.deserializeList()
	if err != nil {
		return nil, err
	}

	result := make(map[int]*Match)
	for _, match := range matchesList {
		result[*match.Id] = match
	}
	return result, nil
}

func (p *csvMatchesProvider) GetMatches() (map[int]*Match, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.getMatchesInternal()
}

func (p *csvMatchesProvider) DeleteMatch(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	matches, err := p.getMatchesInternal()
	if err != nil {
		return err
	}

	delete(matches, id)
	err = p.serialize(matches)
	if err != nil {
		return err
	}

	log.Printf("The match with ID %d deleted.\n", id)
	return nil
}

func getNextId(matches map[int]*Match) int {
	max := 0
	for key := range matches {
		if key > max {
			max = key
		}
	}

	return max + 1
}

func (p *csvMatchesProvider) AddMatch(match *Match) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	matches, err := p.getMatchesInternal()
	if err != nil {
		return 0, err
	}

	newId := getNextId(matches)
	match.Id = &newId
	matches[newId] = match
	err = p.serialize(matches)
	if err != nil {
		return 0, err
	}

	log.Printf("New match with ID %d added.\n", newId)
	return newId, nil
}
