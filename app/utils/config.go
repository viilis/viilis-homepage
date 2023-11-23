package utils

import (
	"log"
	"os"
)

type ConfigStruct struct {
	Port string 
}

var Config ConfigStruct

func InitConfig() error {
	log.Println("Init config")

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("No env")
		port = "3000"
	}

	Config.Port = port

	return nil
}