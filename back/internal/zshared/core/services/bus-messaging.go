package shared_services

import (
	"context"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/mustafaturan/bus/v3"
	"github.com/rs/zerolog"
)

type BusMessaging struct {
	busMsg *bus.Bus
	logger *zerolog.Logger
}

func newBusMessaging(logger *zerolog.Logger) shared_ports.IEventMessaging {
	log := logger.With().Str("resource", "bus_messaging").Logger()

	return &BusMessaging{
		busMsg: Bus,
		logger: &log,
	}
}

func (b *BusMessaging) SendMessage(ctx context.Context, data *shared_domain.MessageEvent) {
	err := b.busMsg.Emit(ctx, string(data.EventTopic), data.Data)
	if err != nil {
		b.logger.Err(err).Interface("data", data).Msg("<BUS MESSAGING> Faild to send message")
	} else {
		b.logger.Log().Interface("data", data).Msg("<BUS MESSAGING> Message published successfully")
	}
}
