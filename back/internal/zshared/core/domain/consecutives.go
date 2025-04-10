package shared_domain

type Consecutive struct {
	ID          ConsecutiveType `bson:"_id"`
	Sequence    int             `bson:"sequence"`
	Prefix      string          `bson:"prefix,omitempty"`
	Description string          `bson:"description,omitempty"`
}

type ConsecutiveType string

const (
	SaleConsecutive     ConsecutiveType = "SALE"
	PurchaseConsecutive ConsecutiveType = "PURCHASE"
)
