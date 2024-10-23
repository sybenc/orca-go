package idutils

import (
	"github.com/sony/sonyflake"
	"orca/pkg/errors"
	"os"
	"strconv"
	"time"
)

var Sonyflake *sonyflake.Sonyflake

func init() {
	machineID, err := getMachineIDFromEnv()
	if err != nil {
		panic("生成雪花算法ID时，获取机器ID失败：" + err.Error())
	}
	Sonyflake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2003, 11, 27, 0, 0, 0, 0, time.UTC),
		MachineID: machineID,
	})
}

// getMachineIDFromEnv 从环境变量中获取 MachineID
func getMachineIDFromEnv() (func() (uint16, error), error) {
	machineIDStr := os.Getenv("SONYFLAKE_MACHINE_ID")
	if machineIDStr == "" {
		return nil, errors.New("环境变量 SONYFLAKE_MACHINE_ID 未设置")
	}

	machineID, err := strconv.ParseUint(machineIDStr, 10, 16)
	if err != nil {
		return nil, errors.New("无效的 SONYFLAKE_MACHINE_ID，必须是 0-65535 之间的整数")
	}
	return func() (uint16, error) {
		return uint16(machineID), nil
	}, nil
}
