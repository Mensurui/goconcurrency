package errors

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
)

type RaceCondition struct {
	hl hclog.Logger
}

func NewRaceCondition(hl hclog.Logger) RaceCondition {
	return RaceCondition{
		hl: hl,
	}
}

func (rc *RaceCondition) ConditionOne() error {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("[DATA]: %v\n", data)
		rc.hl.Log(hclog.NoLevel, "[DATA]: %v", data)
	}
	return nil
}

func (rc *RaceCondition) ConditionTwo() error {
	var data int
	go func() {
		data++
	}()
	time.Sleep(3 * time.Second)
	if data == 0 {
		fmt.Printf("[DATA]: %v\n", data)
		rc.hl.Log(hclog.NoLevel, "[DATA]: %v", data)
	}
	fmt.Printf("[DATA]: %v\n", data)
	rc.hl.Log(hclog.NoLevel, "[DATA]: %v", data)
	return nil
}
