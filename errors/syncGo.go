package errors

import (
	"fmt"
	"sync"

	"github.com/hashicorp/go-hclog"
)

type Sync struct {
	hl           hclog.Logger
	memoryAccess *sync.Mutex
}

func NewSync(hl hclog.Logger, memoryAccess *sync.Mutex) *Sync {
	return &Sync{
		hl:           hl,
		memoryAccess: memoryAccess,
	}
}

func (s *Sync) NonIdiomaticSolution() error {
	var value int
	go func() {
		s.memoryAccess.Lock()
		value++
		s.memoryAccess.Unlock()
	}()
	s.memoryAccess.Lock()
	if value == 0 {
		fmt.Println("value is equal to 0")
	} else {
		fmt.Printf("value is equal to: %v", value)
	}

	return nil
}
