package utils

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func InitSnowflake() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	if sf == nil {
		panic("init snowflake failed")
	}
}

func GenerateOrderNo() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		panic(fmt.Sprintf("generate order no failed: %v", err))
	}
	return fmt.Sprintf("ORD-%d", id), nil
}
