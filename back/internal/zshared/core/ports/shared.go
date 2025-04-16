package shared_ports

import (
	"context"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
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
		quantity, price, gain float64,
		unit product_domain.ProductUnit,
		subUnit float64,
		productId, productName, productSku string,
		user *user_domain.User,
	)
}

type ISharedOrdersSummaryService interface {
	HandleOrdersSummary(c context.Context, payload map[string]interface{}, topic string)
	CalculateOrdersSummary(ctx context.Context, totals *order_domain.Totals, orderType order_domain.OrderType, paymentMethod order_domain.PaymentMethod)
	GetOrdersSummary(ctx context.Context) ([]shared_domain.OrderSummary, error)
}
