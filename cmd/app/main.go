package main

import (
	"rewardapp/internal/infra/config"
	"fmt"
	"os"
)

func main() {
	
	config, err := config.Load()
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("REWARD APP server starting")
	fmt.Printf("Port: localhost:%s\n", config.AppConfig.Port)
}
