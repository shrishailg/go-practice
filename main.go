package main

import (
	"github.com/go-practice/config"
	"github.com/go-practice/logger"
	"github.com/go-practice/pkg/database"

	// files "github.com/go-practice/internal/files"
	router "github.com/go-practice/cmd/router"
)

func main() {
	//init configs
	config.LoadConfig()

	//to handle logs
	logger.Init()

	sqlDb, err := database.InitDb().DB()
	if err != nil {
		panic("DB initialisation failed")
	}

	defer sqlDb.Close()

	router.Init()

}
