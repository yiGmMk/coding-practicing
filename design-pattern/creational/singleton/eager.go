package singleton

import "github.com/google/uuid"

var eagerInstance = &singleton{Uid: uuid.NewString()}

func EagerSingleton() *singleton {
	return eagerInstance
}
