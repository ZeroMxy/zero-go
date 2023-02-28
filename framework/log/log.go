package log

import (
	"fmt"
	"log"
	"os"
	"zero-go/config"
	"github.com/ZeroMxy/fastgo"
)

type Log struct {}

var logger *log.Logger

func Initialize () *os.File {

	var log_folder_path = config.Log["path"]

	// 不存在则创建目录
	if !fastgo.Path_is_exist(log_folder_path) {
		os.MkdirAll(log_folder_path, os.ModePerm)
	}

	var log_name = log_folder_path + "system.log"
	
	var log_file, _ = os.OpenFile(log_name, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)

	logger = log.New(log_file, "[DEBUG] ", log.LstdFlags)
	
	logger.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	return log_file

}

func Info (log_info ...interface {}) {

	logger.SetPrefix("[INFO] ")

	logger.Output(2, fmt.Sprintln(log_info...))
	
	return

}

func Debug (log_info ...interface {}) {

	logger.SetPrefix("[DEBUG] ")

	logger.Output(2, fmt.Sprintln(log_info...))

	return

}

func Error (log_info ...interface {}) {

	logger.SetPrefix("[ERROR] ")

	logger.Output(2, fmt.Sprintln(log_info...))
	
	return

}

