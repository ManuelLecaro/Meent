package di

import (
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() (*dig.Container, error) {
	Container = dig.New()

	//	configs := config.LoadConfiguration()

	/*	err = Container.Provide(server.NewApp)
		err = Container.Provide(collections.NewArticleCollection)
		err = Container.Provide(article.NewArticleService)

		err = Container.Provide(collections.NewCategoryCollection)
		err = Container.Provide(category.NewCategoryService)

		err = Container.Provide(collections.NewTransactionCollection)
		err = Container.Provide(transaction.NewTransactionService)

		err = Container.Provide(server.NewGraphHandler)
	*/
	return Container, nil
}
