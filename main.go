package main

import (
	"github.com/Mensurui/goconcurrency/errors"
	"github.com/hashicorp/go-hclog"
)

func main() {
	hl := hclog.Default()
	raceCondition := errors.NewRaceCondition(hl)

	err := raceCondition.ConditionOne()
	if err != nil {
		hl.Log(hclog.NoLevel, "[ERROR]: %v", err)
	}

	err = raceCondition.ConditionTwo()
	if err != nil {
		hl.Log(hclog.NoLevel, "[ERROR]: %v", err)
	}
}
