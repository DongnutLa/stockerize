package shared_ports

import "context"

type ISharedProductService interface {
	HandleStock(c context.Context, payload map[string]interface{}, topic string)
}
