package config

var App = map[string] string {

	// Application
	"name": Env().Get("app.name").(string),
	"host": Env().Get("app.host").(string),
	"mode": Env().Get("app.host").(string),

	// Snowflake
	"worker_id": 	Env().Get("snowflake.worker_id").(string), 	// 机器 id 默认 0
	"start_time": 	Env().Get("snowflake.start_time").(string), // 起点时间戳 默认 0 如果在程序跑了一段时间修改了这个值 可能会导致生成相同的 id

}