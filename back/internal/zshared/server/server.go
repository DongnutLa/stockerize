package server

import (
	"log"

	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	//We will add every new Handler here
	userHandlers    user_ports.UserHandlers
	storeHandlers   store_ports.StoreHandlers
	productHandlers product_ports.ProductHandlers
	orderHandlers   order_ports.OrderHandlers
	//middlewares ports.Middlewares
}

func NewServer(
	userHandlers user_ports.UserHandlers,
	storeHandlers store_ports.StoreHandlers,
	productHandlers product_ports.ProductHandlers,
	orderHandlers order_ports.OrderHandlers,
) *Server {
	return &Server{
		userHandlers:    userHandlers,
		storeHandlers:   storeHandlers,
		productHandlers: productHandlers,
		orderHandlers:   orderHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	usersRoute := v1.Group("/users")
	usersRoute.Get("", s.userHandlers.ListUsers)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
