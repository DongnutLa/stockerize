package server

import (
	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	//We will add every new Handler here
	userHandlers    user_ports.IUserHandler
	storeHandlers   store_ports.IStoreHandler
	productHandlers product_ports.IProductHandler
	orderHandlers   order_ports.IOrderHandler
	authMw          fiber.Handler
}

func NewServer(
	userHandlers user_ports.IUserHandler,
	storeHandlers store_ports.IStoreHandler,
	productHandlers product_ports.IProductHandler,
	orderHandlers order_ports.IOrderHandler,
	authMw fiber.Handler,
) *Server {
	return &Server{
		userHandlers:    userHandlers,
		storeHandlers:   storeHandlers,
		productHandlers: productHandlers,
		orderHandlers:   orderHandlers,
		authMw:          authMw,
	}
}

func (s *Server) Initialize() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	v1.Route("/health-check", func(router fiber.Router) {
		router.Get("", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON("ok")
		})
	})

	userRoute := v1.Group("/user")
	userRoute.Get("/login/google", s.authMw, s.userHandlers.GoogleLogin)
	userRoute.Post("/login", s.userHandlers.Login)
	userRoute.Post("/", s.authMw, s.userHandlers.CreateUser)

	storeRoute := v1.Group("/store")
	storeRoute.Post("/", s.authMw, s.storeHandlers.CreateStore)

	productRoute := v1.Group("/product")
	productRoute.Get("/", s.authMw, s.productHandlers.SearchProducts)
	productRoute.Get("/stock", s.authMw, s.productHandlers.GetProductsStock)
	productRoute.Get("/:id", s.authMw, s.productHandlers.GetById)
	productRoute.Get("/:id/history", s.authMw, s.productHandlers.GetHistory)
	productRoute.Post("/", s.authMw, s.productHandlers.CreateProduct)
	productRoute.Patch("/", s.authMw, s.productHandlers.UpdateProduct)
	productRoute.Put("/stock", s.authMw, s.productHandlers.UpdateProductStock)

	orderRoute := v1.Group("/order")
	orderRoute.Get("/summary", s.authMw, s.orderHandlers.GetSummary)
	orderRoute.Get("/", s.authMw, s.orderHandlers.ListOrders)
	orderRoute.Get("/:id", s.authMw, s.orderHandlers.GetById)
	orderRoute.Post("/", s.authMw, s.orderHandlers.CreateOrder)
	orderRoute.Patch("/", s.authMw, s.orderHandlers.UpdateOrder)

	return app
}
