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
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	content, err := ioutil.ReadFile("docs/name.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	container, err := di.BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(app *rest.Server) {

		app.SetupRoutes()

		sugar.Info("Initializing app")

		app.RunApp()
	})

	if err != nil {
		log.Fatal(err)
	}
}
