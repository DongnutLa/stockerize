package main

import (
	"context"
	"os"

	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_services "github.com/DongnutLa/stockio/internal/zshared/core/services"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger
var sharedProductService shared_ports.ISharedProductService

func init() {
	ctx := context.TODO()
	utils.LoadConfig()
	conn := shared_repositories.NewMongoDB(ctx)

	logger = zerolog.New(os.Stderr)

	productRepository := product_repositories.NewProductRepository(context.TODO(), constants.PRODUCTS_COLLECTION, conn.Database, &logger)
	productHistoryRepository := product_repositories.NewProductHistoryRepository(context.TODO(), constants.PRODUCTS_HISTORY_COLLECTION, conn.Database, &logger)

	sharedProductService = shared_services.NewSharedProductService(ctx, &logger, productRepository, productHistoryRepository)
}

func Handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		event := utils.JsonToStruct[shared_domain.MessageEvent](record.SNS.Message)
		logger.Info().Interface("event", event).Msg("[EVENT] Got event")

		switch event.EventTopic {
		case shared_domain.HandleStockTopic:
			sharedProductService.HandleStock(ctx, event.Data, string(event.EventTopic))
		}
	}
}

func main() {
	lambda.Start(Handler)
}
