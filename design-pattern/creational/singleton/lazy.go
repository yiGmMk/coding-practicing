package singleton

import (
	"sync"

	"github.com/google/uuid"
)

type singleton struct {
	Uid string
}

var (
	once     sync.Once
	instance *singleton
)

func LazySingleton() *singleton {
	once.Do(func() {
		instance = &singleton{Uid: uuid.NewString()}
	})
	return instance
}
