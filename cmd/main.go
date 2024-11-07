package main

import (
	"fmt"
	"log"

	"github.com/phn00dev/go-crud/internal/app"
	"github.com/phn00dev/go-crud/internal/setup/constructor"
)

func main() {
	dependencies, err := app.GetDependencies()
	if err != nil {
		log.Println("dependencies error :", err)
		return
	}
	constructor.InitDependencies(dependencies)
	appRouter := app.NewApp(dependencies.Config)
	runServer := fmt.Sprintf("%s:%s", dependencies.Config.HttpConfig.HttpHost, dependencies.Config.HttpConfig.HttpPort)
	serverURL := fmt.Sprintf("http://%s", runServer)
	log.Println(serverURL)
	if err := appRouter.Run(runServer); err != nil {
		log.Println("server run error: ", err)
		return
	}
}
