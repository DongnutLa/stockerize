package product_domain

type CreateProductDTO struct {
	Name   string       `json:"name"`
	Sku    string       `json:"sku"`
	Stock  *StockCreate `json:"stock"`
	Unit   ProductUnit  `json:"unit"`
	Prices *[]Price     `json:"prices"`
}
type StockCreate struct {
	Cost      float64 `json:"cost"`
	Quantity  float64 `json:"quantity"` // delta
	UnitPrice Price   `json:"unitPrice"`
}

type UpdateProductDTO struct {
	ID     string      `json:"id"`
	Sku    string      `json:"sku"`
	Name   string      `json:"name"`
	Unit   ProductUnit `json:"unit"`
	Prices []Price     `json:"prices"`
}

type StockUpdateType string

const (
	StockIncrease StockUpdateType = "INCREASE"
	StockDecrease StockUpdateType = "DECREASE"
	StockInfo     StockUpdateType = "INFO" // price/cost change
	StockPurchase StockUpdateType = "PURCHASE"
	StockSale     StockUpdateType = "SALE"
)

type UpdateStockDTO struct {
	ID         string          `json:"id"`
	UpdateType StockUpdateType `json:"updateType"`
	StockCreate
}
