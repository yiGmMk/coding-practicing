package concurrency

import set "github.com/duke-git/lancet/v2/datastructure/set"

func SetNode() {
	s := set.NewSet[int](1, 2, 3)
	s.Add(1, 2, 3)

}
