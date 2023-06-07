package di

import (
	"mintapi/internal/core/event"
	"mintapi/internal/core/port/resource"
	"mintapi/internal/infra/config"
	"mintapi/internal/infra/ipfs"
	"mintapi/internal/infra/quicknode"
	"mintapi/internal/infra/quicknode/contract/meent"

	"mintapi/internal/infra/rest"
	"mintapi/internal/infra/rest/handler"

	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() (*dig.Container, error) {
	Container = dig.New()

	config.LoadConfiguration()

	err := Container.Provide(func() resource.Contract {
		ethClient := quicknode.NewEthBaseClient(config.GetQuickNodeConfigs())

		return meent.NewContractManager(*ethClient, config.GetContractAdressConfigs())
	})

	err = Container.Provide(ipfs.NewIPFSService)

	err = Container.Provide(event.NewEventService)

	err = Container.Provide(handler.NewEvent)

	err = Container.Provide(rest.NewHttpHandler)

	err = Container.Provide(rest.NewServer)

	return Container, err
}
