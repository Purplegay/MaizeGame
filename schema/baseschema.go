package schema

import "sync"

type baseSchema struct {
	lock sync.RWMutex
}
