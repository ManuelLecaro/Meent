package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mintapi/internal/infra/di"
	"mintapi/internal/infra/rest"

	"go.uber.org/zap"
)

func main() {
	initMessage()

	container, err := di.BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(app *rest.Server) {

		app.SetupRoutes()

		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		sugar := logger.Sugar()
		sugar.Info("Initializing app")

		app.RunApp()
	})

	if err != nil {
		log.Fatal(err)
	}
}

func initMessage() {
	content, err := ioutil.ReadFile("docs/name.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
