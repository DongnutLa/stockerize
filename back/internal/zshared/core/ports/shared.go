package shared_ports

import (
	"context"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
)

type ISharedProductService interface {
	HandleStock(c context.Context, payload map[string]interface{}, topic string)
	UpdateProductStock(ctx context.Context, stockDto *product_domain.UpdateStockDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError)
	CreateHistory(
		ctx context.Context,
		hType product_domain.ProductHistoryType,
		quantity int64,
		price, gain float64,
		unit product_domain.ProductUnit,
		productId, productName, productSku string,
		user *user_domain.User,
	)
}
