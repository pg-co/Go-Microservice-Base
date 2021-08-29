package main

import (
	"fmt"
	"api/api"
	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading environment: %v", err)
	} else {
		fmt.Println("Successfully loaded Environment")
	}
}


func main() {
	LoadEnvironment()
	api.Init()
}