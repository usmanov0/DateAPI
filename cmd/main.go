package main

import (
	"fmt"
	"time"

	api "github.com/task_iman/api"
	_ "github.com/task_iman/api/docs"
	"github.com/task_iman/pkg/logger"
	token "github.com/task_iman/api/handlers"
)

type Tokentime struct {
	Time1 time.Duration
}

func main() {
	logger.Init()
	log := logger.GetLogger()
	log.Info("logger initialized")

	apiServer := api.New(api.RoutetOptions{
		Log: log,
	})

	tokenString, err := token.CreateToken(&token.Tokentime{
		Time1: time.Hour *24,
	})
	if err != nil {
		fmt.Print(">>>>>>>>>>", err)
	}
	fmt.Println("Token: ", tokenString)

	if err := apiServer.Run(fmt.Sprintf(":%s", "8081")); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
