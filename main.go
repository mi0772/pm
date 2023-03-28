/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/joho/godotenv"
	"log"
	"mi0772/pm/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
