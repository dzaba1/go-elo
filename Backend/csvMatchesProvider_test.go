package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_csvMatchesProvider_GetMatches_WhenEmptyCsvFile_ThenEmptyData(t *testing.T) {
	filename := "test.csv"
	DeleteFileSafe(filename)
	defer DeleteFileSafe(filename)

	err := WriteAllFile(filename, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore")
	assert.Nil(t, err)

	sut := NewCsvMatchesProvider(filename)
	result, err := sut.GetMatches()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func Test_csvMatchesProvider_GetMatches_WhenCsvFile_ThenData(t *testing.T) {
	filename := "test.csv"
	DeleteFileSafe(filename)
	defer DeleteFileSafe(filename)

	err := WriteAllFile(filename, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore\n1,20.07.2023 10:11,Player1,Player2,2,0")
	assert.Nil(t, err)

	sut := NewCsvMatchesProvider(filename)
	result, err := sut.GetMatches()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
}

func Test_csvMatchesProvider_DeleteMatch_WhenEmptyCsvFile_ThenNothing(t *testing.T) {
	filename := "test.csv"
	DeleteFileSafe(filename)
	defer DeleteFileSafe(filename)

	err := WriteAllFile(filename, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore")
	assert.Nil(t, err)

	sut := NewCsvMatchesProvider(filename)
	err = sut.DeleteMatch(1)

	assert.Nil(t, err)
}

func Test_csvMatchesProvider_DeleteMatch_WhenCorrectId_ThenDeleted(t *testing.T) {
	filename := "test.csv"
	DeleteFileSafe(filename)
	defer DeleteFileSafe(filename)

	err := WriteAllFile(filename, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore\n1,20.07.2023 10:11,Player1,Player2,2,0")
	assert.Nil(t, err)

	sut := NewCsvMatchesProvider(filename)
	err = sut.DeleteMatch(1)
	assert.Nil(t, err)

	content, err := ReadAllFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore\n", content)
}

func Test_csvMatchesProvider_AddMatch_WhenEmptyCsvFile_ThenAdded(t *testing.T) {
	filename := "test.csv"
	DeleteFileSafe(filename)
	defer DeleteFileSafe(filename)

	err := WriteAllFile(filename, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore")
	assert.Nil(t, err)

	match := &Match{
		Id:               nil,
		DateTime:         time.Date(2023, 7, 20, 10, 11, 0, 0, time.Local),
		LeftPlayer:       "Player1",
		RightPlayer:      "Player2",
		LeftPlayerScore:  2,
		RightPlayerScore: 0,
	}

	sut := NewCsvMatchesProvider(filename)
	newId, err := sut.AddMatch(match)

	assert.Nil(t, err)
	assert.Equal(t, 1, newId)

	content, err := ReadAllFile(filename)
	assert.Nil(t, err)

	assert.Equal(t, "Id,DateTime,LeftPlayer,RightPlayer,LeftPlayerScore,RightPlayerScore\n1,20.07.2023 10:11,Player1,Player2,2,0\n", content)
}
