package main

import (
	"log"

	"server/db"
	"server/router"
)

func init() {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db.Init()

	port := ":8080"
	log.Printf("[START] server port: %s\n", port)

	echoRouter := router.NewEchoRouter()
	echoRouter.Logger.Fatal(echoRouter.Start(port))
}
