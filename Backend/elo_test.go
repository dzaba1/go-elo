package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_eloImpl_CalculateExpected_WhenLeftAndRight_ThenValue(t *testing.T) {
	settings := &EloSettings{
		KFactor:    32,
		InitRating: 1000,
	}

	sut := NewElo(settings)

	result := sut.CalculateExpected(1000, 1000)
	assert.Equal(t, 0.5, result)
}

func Test_eloImpl_CalculateNewRating_WhenLeftAndRight_ThenValue(t *testing.T) {
	settings := &EloSettings{
		KFactor:    32,
		InitRating: 1000,
	}

	sut := NewElo(settings)

	result := sut.CalculateNewRating(1000, 1000, 1)
	assert.Equal(t, 1016.0, result)
}
