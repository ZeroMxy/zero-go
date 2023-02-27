package config

var Database = map[string] string {

	// Database
	"drive":    Env().Get("database.drive").(string),
	"host":     Env().Get("database.host").(string),
	"port":     Env().Get("database.port").(string),
	"name":     Env().Get("database.name").(string),
	"username": Env().Get("database.username").(string),
	"password": Env().Get("database.password").(string),

}