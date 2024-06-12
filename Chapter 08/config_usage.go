package main

import (
 "fmt"
 "singleton_config/config"
)

func main() {
 cfg := config.GetInstance()

 fmt.Println("Database connection string:", cfg["database"])
 fmt.Println("Redis connection string:", cfg["redis"])
}
