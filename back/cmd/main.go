package main

import (
	"context"
	"os"
	"sync"
	"time"

	order_services "github.com/DongnutLa/stockio/internal/order/core/services"
	order_handlers "github.com/DongnutLa/stockio/internal/order/handlers"
	order_repositories "github.com/DongnutLa/stockio/internal/order/repositories"
	product_services "github.com/DongnutLa/stockio/internal/product/core/services"
	product_handlers "github.com/DongnutLa/stockio/internal/product/handlers"
	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	store_services "github.com/DongnutLa/stockio/internal/store/core/services"
	store_handlers "github.com/DongnutLa/stockio/internal/store/handlers"
	store_repositories "github.com/DongnutLa/stockio/internal/store/repositories"
	user_services "github.com/DongnutLa/stockio/internal/user/core/services"
	user_handlers "github.com/DongnutLa/stockio/internal/user/handlers"
	user_repositories "github.com/DongnutLa/stockio/internal/user/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_services "github.com/DongnutLa/stockio/internal/zshared/core/services"
	shared_handlers "github.com/DongnutLa/stockio/internal/zshared/handlers"
	"github.com/DongnutLa/stockio/internal/zshared/middlewares"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/server"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/rs/zerolog"
)

var fiberLambda *fiberAdapter.FiberLambda
var logger zerolog.Logger

func init() {
	utils.LoadConfig()
	logger = zerolog.New(os.Stderr)
	ctx := context.Background()

	conn := shared_repositories.NewMongoDB(ctx)

	jwtKey := utils.GetConfig("jwt_key")

	//repositories
	userRepository := user_repositories.NewUserRepository(ctx, constants.USERS_COLLECTION, conn.Database, &logger)
	storeRepository := store_repositories.NewStoreRepository(ctx, constants.STORES_COLLECTION, conn.Database, &logger)
	productRepository := product_repositories.NewProductRepository(ctx, constants.PRODUCTS_COLLECTION, conn.Database, &logger)
	productHistoryRepository := product_repositories.NewProductHistoryRepository(ctx, constants.PRODUCTS_HISTORY_COLLECTION, conn.Database, &logger)
	orderRepository := order_repositories.NewOrderRepository(ctx, constants.ORDERS_COLLECTION, conn.Database, &logger)
	consecutiveRepository := shared_repositories.NewConsecutiveRepository(ctx, constants.CONSECUTIVES_COLLECTION, conn.Database, &logger)
	ordersSummaryRepository := shared_repositories.NewOrderSummaryRepository(ctx, constants.ORDERS_SUMMARY_COLLECTION, conn.Database, &logger)

	// Shared services
	sharedProductService := shared_services.NewSharedProductService(ctx, &logger, productRepository, productHistoryRepository)
	consecutiveService := shared_services.NewConsecutiveService(&logger, consecutiveRepository)
	ordersSummaryService := shared_services.NewSharedOrdersSummaryService(&logger, ordersSummaryRepository)

	isLambda := IsLambda()

	eventType := shared_ports.UseSNS
	if !isLambda {
		logger.Info().Msg("Starting local events handler...")
		eventType = shared_ports.UseBUS
		handler := shared_handlers.NewEventsHandler(ctx, &logger, sharedProductService, ordersSummaryService)

		msgNow := time.Now()
		shared_services.MessagingInit()
		var wg sync.WaitGroup
		defer wg.Wait()

		handler.Start(&wg)
		defer handler.Stop()
		logger.Info().Msgf("Messaging init time: %dms", time.Since(msgNow).Milliseconds())
	}

	messaging := shared_services.NewEventMessaging(&logger, eventType)

	// Services
	jwtService := shared_services.NewJwtService([]byte(jwtKey), &logger)
	userService := user_services.NewUserService(ctx, &logger, userRepository, jwtService)
	storeService := store_services.NewStoreService(ctx, &logger, storeRepository)
	productService := product_services.NewProductService(ctx, &logger, productRepository, productHistoryRepository, sharedProductService, messaging)
	orderService := order_services.NewOrderService(ctx, &logger, orderRepository, messaging, consecutiveService, ordersSummaryService)

	//handlers
	userHandlers := user_handlers.NewUserHandler(userService)
	storeHandlers := store_handlers.NewStoreHandlers(storeService)
	productHandlers := product_handlers.NewProductHandlers(productService)
	orderHandlers := order_handlers.NewOrderHandlers(orderService)

	//Middlewares
	auth := middlewares.NewAuthMiddleware(jwtService, userRepository, &logger, false)

	apiNow := time.Now()

	//server
	httpServer := server.NewServer(
		userHandlers,
		storeHandlers,
		productHandlers,
		orderHandlers,
		auth,
	)
	app := httpServer.Initialize()

	if isLambda {
		fiberLambda = fiberAdapter.New(app)
		logger.Info().Msgf("Api init time: %dms", time.Since(apiNow).Milliseconds())
	} else {
		app.Listen(":3000")
	}
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	logger.Info().Interface("request", request).Msg("[REQUEST] Incoming request")
	res, err := fiberLambda.ProxyWithContextV2(ctx, request)
	if err != nil {
		logger.Err(err).Msg("Error in proxy request")
		return nil, err
	}

	res.Headers["Access-Control-Allow-Methods"] = "*"
	res.Headers["Access-Control-Allow-Headers"] = "*"
	res.Headers["Access-Control-Allow-Origin"] = "*"
	res.Headers["Access-Control-Max-Age"] = "3600"
	res.Headers["Access-Control-Allow-Credentials"] = "true"

	return &res, nil
}

func main() {
	lambda.Start(Handler)
}

func IsLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}
