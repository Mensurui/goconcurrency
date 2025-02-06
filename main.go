package main

import (
	"sync"

	"github.com/Mensurui/goconcurrency/errors"
	"github.com/hashicorp/go-hclog"
)

func main() {
	hl := hclog.Default()
	memoryAccess := sync.Mutex{}
	raceCondition := errors.NewRaceCondition(hl)
	raceConditionSolution := errors.NewSync(hl, &memoryAccess)

	err := raceCondition.ConditionOne()
	if err != nil {
		hl.Log(hclog.NoLevel, "[ERROR]: %v", err)
	}

	err = raceCondition.ConditionTwo()
	if err != nil {
		hl.Log(hclog.NoLevel, "[ERROR]: %v", err)
	}

	err = raceConditionSolution.NonIdiomaticSolution()
	if err != nil {
		hl.Log(hclog.NoLevel, "[ERROR]: %v", err)
	}

}
