package config

var Log = map[string] string {

	// Log
	"path":   Env().Get("log.path").(string), 		// 日志保存路径
	"model":  Env().Get("log.model").(string),      // single 单文件 daily 按日期切割文件
	"day":    Env().Get("log.day").(string),        // 日志保留天数
	"system": Env().Get("system").(string),         // api 请求日志 true 开启 false 关闭

}