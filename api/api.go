package api

import (
	"payload-app/api/controllers"
	"github.com/joho/godotenv"
	"math/rand"
	"time"
)

var server = controllers.Server{}

func Run() {
	_ = godotenv.Load()
	server.Init()

	rand.Seed(time.Now().Unix())

	server.Run()
}