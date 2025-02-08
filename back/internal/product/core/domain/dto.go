package product_domain

type CreateProductDTO struct {
	Name  string      `json:"name"`
	Sku   string      `json:"sku"`
	Stock *Stock      `json:"stock"`
	Unit  ProductUnit `json:"unit"`
}

type UpdateProductDTO struct {
	ID   string `json:"id"`
	Sku  string `json:"sku"`
	Name string `json:"name"`
}

type StockUpdateType string

const (
	StockAdd     StockUpdateType = "ADD"
	StockRelease StockUpdateType = "RELEASE"
	StockInfo    StockUpdateType = "INFO" // price/cost change
)

type UpdateStockDTO struct {
	ID         string          `json:"id"`
	UpdateType StockUpdateType `json:"updateType"`
	Cost       float64         `json:"cost"`
	Price      float64         `json:"price"`
	Quantity   int64           `json:"quantity"` // delta
}
