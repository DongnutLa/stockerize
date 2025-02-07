package order_domain

type Totals struct {
	Subtotal float64 `bson:"_id" json:"id"`
	Discount float64 `bson:"type" json:"type"`
	Total    float64 `bson:"quantity" json:"quantity"`
}

func NewTotals(subtotal, discount, total float64) *Totals {
	return &Totals{
		Subtotal: subtotal,
		Discount: discount,
		Total:    total,
	}
}
