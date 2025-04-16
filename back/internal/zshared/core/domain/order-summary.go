package shared_domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderSummary struct {
	ID                   primitive.ObjectID     `bson:"_id" json:"_id"`
	OrderType            OrderType              `bson:"orderType" json:"orderType"`
	SummaryType          SummaryType            `bson:"summaryType" json:"summaryType"`
	PaymentMethod        PaymentMethod          `bson:"paymentMethod" json:"paymentMethod"`
	Count                uint                   `bson:"count" json:"count"`
	Total                float64                `bson:"total" json:"total"`
	Start                time.Time              `bson:"start" json:"start"`
	End                  time.Time              `bson:"end" json:"end"`
	UpdatedAt            *time.Time             `bson:"updatedAt" json:"updatedAt"`
	PaymentMethodDetails []PaymentMethodDetails `bson:"-" json:"paymentMethodDetails"`
}

type PaymentMethodDetails struct {
	PaymentMethod PaymentMethod `bson:"paymentMethod" json:"paymentMethod"`
	Count         uint          `bson:"count" json:"count"`
	Total         float64       `bson:"total" json:"total"`
}

type OrderType string

const (
	SaleType     ConsecutiveType = "SALE"
	PurchaseType ConsecutiveType = "PURCHASE"
)

type SummaryType string

const (
	DailySummaryType   SummaryType = "DAILY"
	WeeklySummaryType  SummaryType = "WEEKLY"
	MonthlySummaryType SummaryType = "MONTHLY"
)

type PaymentMethod string

const (
	CashPaymentMethod      PaymentMethod = "CASH"
	NequiPaymentMethod     PaymentMethod = "NEQUI"
	DaviplataPaymentMethod PaymentMethod = "DAVIPLATA"
)
