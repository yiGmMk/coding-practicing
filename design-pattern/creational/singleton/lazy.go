package singleton

import (
	"sync"

	"github.com/google/uuid"
)

type Singleton struct {
	Uid string
}

var (
	once     sync.Once
	instance *Singleton
)

func LazySingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{Uid: uuid.NewString()}
	})
	return instance
}
