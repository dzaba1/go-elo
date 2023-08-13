package main

import (
	"os"
	"strconv"
)

type EloSettings struct {
	KFactor    float64
	InitRating float64
}

func GetEloSettingsFromEnvs() (*EloSettings, error) {
	kEnv := os.Getenv("ELO_KFACTOR")
	initEnv := os.Getenv("ELO_INIT_RATING")

	k, err := strconv.ParseFloat(kEnv, 64)
	if err != nil {
		return nil, err
	}

	init, err := strconv.ParseFloat(initEnv, 64)
	if err != nil {
		return nil, err
	}

	return &EloSettings{
		KFactor:    k,
		InitRating: init,
	}, nil
}
