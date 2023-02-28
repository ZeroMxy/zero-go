package config

var Log = map[string] string {

	// Log
	"path":   Env().Get("log.path").(string), 		// 日志保存路径

}