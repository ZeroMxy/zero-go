package config

import "github.com/akkuman/parseConfig"

func Env () *parseConfig.Config {

	var env = parseConfig.New("env.json")

	return &env

}