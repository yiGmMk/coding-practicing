package redis

import (
	"fmt"
	"os"
	"strconv"
)

const (
	_stream   = "new_staff_stream"
	_group    = "new_staff_group"
	_consumer = "consumer:%s" // machine_id
)

// 消费者
func GenConsumerName(id ...int) (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", err
	}
	ids := ""
	if len(id) > 0 {
		ids = strconv.Itoa(id[0])
	}
	return fmt.Sprintf(_consumer, name) + ids, nil
}
