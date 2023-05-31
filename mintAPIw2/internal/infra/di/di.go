package di

import (
	"mintapi/internal/core/event"
	"mintapi/internal/core/ticket"
	"mintapi/internal/infra/config"
	"mintapi/internal/infra/quicknode"
	"mintapi/internal/infra/rest"
	"mintapi/internal/infra/rest/handler"

	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() (*dig.Container, error) {
	Container = dig.New()

	Container = dig.New()

	configs := config.LoadConfiguration()

	err := Container.Provide(func() *quicknode.EthBaseClient {
		return quicknode.NewEthBaseClient(configs.GetString("QUICKNODE_URL"))
	})

	err = Container.Provide(event.NewEventService)

	err = Container.Provide(ticket.NewTicketService)

	err = Container.Provide(handler.NewEvent)
	err = Container.Provide(handler.NewTicket)

	err = Container.Provide(func() *rest.Server {
		return rest.NewServer(configs.GetString("APP_NAME"))
	})

	return Container, err
}
