package config

import "sync"

type Config map[string]string

var config Config
var once sync.Once

func initConfig() Config {
 config = make(Config)
 config["database"] = "user:pass@tcp(127.0.0.1:3306)/mydb"
 config["redis"] = "redis://localhost:6379?ConnectTimeout=5000"
 return config
}

func GetInstance() Config {
 once.Do(func() {
  config = initConfig() // Initialize the instance with YAML configuration.
 })
 return config
}
