package config

var Redis = map[string] string {

	// Redis
	"host":     Env().Get("redis.host").(string),
	"port":     Env().Get("redis.port").(string),
	"password": Env().Get("redis.password").(string),
	
}