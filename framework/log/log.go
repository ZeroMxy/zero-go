package log

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"github.com/ZeroMxy/fastgo"
	"zero-go/config"
)

type Log struct {}

var logger *log.Logger

func Initialize (log_prefix string) *os.File {

	var log_model 		= config.Log["model"]
	var log_day, _ 		= strconv.Atoi(config.Log["day"])
	var log_folder_path = config.Log["path"]

	// 不存在则创建目录
	if (!fastgo.Path_is_exist(log_folder_path)) {
		os.MkdirAll(log_folder_path, os.ModePerm)
	}

	// 清理日志文件
	clear_log_file(log_model, log_folder_path, log_prefix, log_day)

	var log_name string

	switch (log_model) {
		case "single":
			log_name = log_folder_path + log_prefix + ".log"
			break
		case "daily":
			log_name = log_folder_path + log_prefix + "-" + fastgo.Format_date_time(fastgo.YMD, "") + ".log"
			break
		default :
			log_name = log_folder_path + log_prefix + ".log"
	}
	
	var log_file, _ = os.OpenFile(log_name, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)

	logger = log.New(log_file, "[DEBUG] ", log.LstdFlags)
	
	logger.SetFlags(log.Llongfile | log.Ltime | log.Ldate)

	return log_file

}

func Info (log_info ...interface {}) {

	Initialize("zero-go")

	logger.SetPrefix("[INFO] ")

	logger.Output(2, fmt.Sprintln(log_info...))
	
	return

}

func Debug (log_info ...interface {}) {

	Initialize("zeri-go")

	logger.SetPrefix("[DEBUG] ")

	logger.Output(2, fmt.Sprintln(log_info...))

	return

}

func Error (log_info ...interface {}) {

	Initialize("zero-go")

	logger.SetPrefix("[ERROR] ")

	logger.Output(2, fmt.Sprintln(log_info...))
	
	return

}

// 清理日志文件
func clear_log_file (log_model, log_path, log_prefix string, log_day int) {

	if (log_model == "daily" && log_day > 0) {
		var log_file_path = log_path + log_prefix + "-" + time.Now().AddDate(0, 0, -log_day).Format("2006-01-02") + ".log"
		
		// 存在则删除日志文件
		if (fastgo.Path_is_exist(log_file_path)) {
			os.Remove(log_file_path)
		}
	}

	return
	
}
