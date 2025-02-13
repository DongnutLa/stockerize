package order_domain

type Totals struct {
	Subtotal float64 `bson:"subtotal" json:"subtotal"`
	Discount float64 `bson:"discount" json:"discount"`
	Total    float64 `bson:"total" json:"total"`
}

func NewTotals(subtotal, discount, total float64) *Totals {
	return &Totals{
		Subtotal: subtotal,
		Discount: discount,
		Total:    total,
	}
}
