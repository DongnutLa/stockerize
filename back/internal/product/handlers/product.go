package product_handlers

import (
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
)

type ProductHandlers struct {
	productService product_ports.ProductService
}

// !TEST
var _ product_ports.ProductHandlers = (*ProductHandlers)(nil)

func NewProductHandlers(productService product_ports.ProductService) *ProductHandlers {
	return &ProductHandlers{
		productService: productService,
	}
}
