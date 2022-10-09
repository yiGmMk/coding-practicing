package singleton

import "github.com/google/uuid"

var eagerInstance = &Singleton{Uid: uuid.NewString()}

func EagerSingleton() *Singleton {
	return eagerInstance
}
