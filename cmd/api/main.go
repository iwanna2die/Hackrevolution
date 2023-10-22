package main

import (
	"HackRevolution/config"
	"HackRevolution/internal/server"
	"HackRevolution/pkg/postgres"
	"HackRevolution/utils"
	"log"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @host 10.200.164.12:8080
// @BasePath /

func main() {
	os.Setenv("ACCESS_TOKEN_SECRET", "secret_parohod")
	os.Setenv("REFRESH_TOKEN_SECRET", "secret_parohod-secret")
	os.Setenv("ACCESS_TIME", "5")
	os.Setenv("REFRESH_TIME", "24")
	log.Println("Starting api server")
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	err = postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Printf("Postgres connected, Status: %#v", postgres.DB().SDB.Stats())
	}

	defer postgres.DB().SDB.Close()
	srv := server.NewServer()
	server.Add_handlers(&srv)
	if err = srv.Run("10.200.164.12:8080"); err != nil {
		log.Fatal(err)
	}
}
