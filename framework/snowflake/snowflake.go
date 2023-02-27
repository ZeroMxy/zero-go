package snowflake

import (
	"strconv"
	"sync"
	"time"

	"zero-go/config"
	"zero-go/framework/log"
)

type Snowflake struct {}

var (
	mutex           sync.Mutex
	timestamp    	int64 // 记录上一次 id 的时间戳
	machine_id     	int64 // 机器的 id
	number     		int64 // 当前毫秒已经生成的 id 序列号(从 0 开始累加) 1 毫秒内最多生成 4096 个 id
    start_time   	int64 // 如果在程序跑了一段时间修改了这个值 可能会导致生成相同的 id
	type_cast_err	error

	machine_bit  	uint8 = 10
    number_bit  	uint8 = 12
    machine_max   	int64 = -1 ^ (-1 << machine_bit)
    number_max   	int64 = -1 ^ (-1 << number_bit)
    time_shift   	uint8 = machine_bit + number_bit
    machine_shift 	uint8 = number_bit
)

// 生成 id 供外部调用
func Get_id () int64 {
	
	mutex.Lock()
	defer mutex.Unlock()

	machine_id, type_cast_err = strconv.ParseInt(config.App["machine_id"], 10, 64)
	start_time, type_cast_err = strconv.ParseInt(config.App["start_time"], 10, 64)
	
	if (type_cast_err != nil) {
		log.Error(type_cast_err)
	}

	return get_id()

}

// 生成 id
func get_id () int64 {

	var milliseconds = get_milliseconds()

	// 当前时间小于上次时间生成失败
	if (milliseconds < timestamp) {
		return 0
	}

	if (milliseconds == timestamp) {
		number++
		if (number > number_max) {
			for milliseconds <= timestamp {
				milliseconds = get_milliseconds()
			}
		}
	} else {
		number = 0
	}

	timestamp = milliseconds

	return int64((milliseconds - start_time) << time_shift | (machine_id << int64(machine_shift)) | (number))

}

// 获取当前时间
func get_milliseconds () int64 {

	return time.Now().UnixNano() / 1e6

}