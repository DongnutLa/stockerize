package product_ports

import (
	"context"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type IProductService interface {
	CreateProduct(ctx context.Context, productDto *product_domain.CreateProductDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError)
	UpdateProduct(ctx context.Context, productDto *product_domain.UpdateProductDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError)
	UpdateProductStock(ctx context.Context, productDto *product_domain.UpdateStockDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError)
}

type IProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	UpdateProductStock(c *fiber.Ctx) error
}
