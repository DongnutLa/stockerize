package shared_ports

import (
	"context"
	"sync"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
)

type IEventMessaging interface {
	SendMessage(context.Context, *shared_domain.MessageEvent)
}

type EventMessagingTypes string

const (
	UseSNS EventMessagingTypes = "UseSNS"
	UseBUS EventMessagingTypes = "UseBUS"
)

type IEventMessagingHandler interface {
	Start(wg *sync.WaitGroup)
	Stop()
	Handler(wg *sync.WaitGroup)
}
