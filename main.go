package main

import (
	"github.com/go-practice/config"
	"github.com/go-practice/internal/logger"

	// files "github.com/go-practice/internal/files"
	router "github.com/go-practice/cmd/router"
)

func main() {

	//init configs
	config.Init()

	//to handle logs
	logger.Init()

	router.Init()

}
