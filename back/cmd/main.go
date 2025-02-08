package main

import (
	"context"
	"os"

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
	"github.com/DongnutLa/stockio/internal/zshared/middlewares"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/server"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/rs/zerolog"
)

func main() {
	utils.LoadConfig()

	conn := shared_repositories.NewMongoDB(context.TODO())
	logger := zerolog.New(os.Stderr)

	// jwtKey := utils.GetConfig("jwt_key")

	//repositories
	userRepository := user_repositories.NewUserRepository(context.TODO(), "users", conn.Database, &logger)
	storeRepository := store_repositories.NewStoreRepository(context.TODO(), "stores", conn.Database, &logger)
	productRepository := product_repositories.NewProductRepository(context.TODO(), "products", conn.Database, &logger)
	productHistoryRepository := product_repositories.NewProductHistoryRepository(context.TODO(), "products", conn.Database, &logger)
	orderRepository := order_repositories.NewOrderRepository(context.TODO(), "orders", conn.Database, &logger)

	//services
	// jwtService := shared_services.NewJwtService([]byte(jwtKey), &logger)
	userService := user_services.NewUserService(context.TODO(), &logger, userRepository)
	storeService := store_services.NewStoreService(context.TODO(), &logger, storeRepository)
	productService := product_services.NewProductService(context.TODO(), &logger, productRepository, productHistoryRepository)
	orderService := order_services.NewOrderService(context.TODO(), &logger, orderRepository)

	//handlers
	userHandlers := user_handlers.NewUserHandler(userService)
	storeHandlers := store_handlers.NewStoreHandlers(storeService)
	productHandlers := product_handlers.NewProductHandlers(productService)
	orderHandlers := order_handlers.NewOrderHandlers(orderService)

	//Middlewares
	auth := middlewares.NewAuthMiddleware(&logger, true)

	//server
	httpServer := server.NewServer(
		userHandlers,
		storeHandlers,
		productHandlers,
		orderHandlers,
		auth,
	)
	httpServer.Initialize()
}
