package shared_handlers

import (
	"context"
	"fmt"
	"sync"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_services "github.com/DongnutLa/stockio/internal/zshared/core/services"
	"github.com/mustafaturan/bus/v3"
	"github.com/rs/zerolog"
)

var stockEvent chan bus.Event
var ordersSummaryEvent chan bus.Event
var cancel context.CancelFunc
var c context.Context

var stockWorker = "stock"
var ordersSummaryWorker = "ordersSummary"

type EventsHandler struct {
	SharedProductService       shared_ports.ISharedProductService
	SharedOrdersSummaryService shared_ports.ISharedOrdersSummaryService
	Logger                     *zerolog.Logger
}

func NewEventsHandler(
	ctx context.Context,
	log *zerolog.Logger,
	sharedProductService shared_ports.ISharedProductService,
	sharedOrdersSummaryService shared_ports.ISharedOrdersSummaryService,
) shared_ports.IEventMessagingHandler {
	return &EventsHandler{
		SharedProductService:       sharedProductService,
		Logger:                     log,
		SharedOrdersSummaryService: sharedOrdersSummaryService,
	}
}

func (h *EventsHandler) Start(wg *sync.WaitGroup) {
	c, cancel = context.WithCancel(context.Background())

	stockEvent = make(chan bus.Event)
	ordersSummaryEvent = make(chan bus.Event)

	// Handlers
	stockHandler := bus.Handler{Handle: func(_ context.Context, e bus.Event) {
		stockEvent <- e
	}, Matcher: string(shared_domain.HandleStockTopic)}
	ordersSumaryHandler := bus.Handler{Handle: func(_ context.Context, e bus.Event) {
		ordersSummaryEvent <- e
	}, Matcher: string(shared_domain.HandleOrdersSummary)}

	shared_services.Bus.RegisterHandler(stockWorker, stockHandler)
	shared_services.Bus.RegisterHandler(ordersSummaryWorker, ordersSumaryHandler)

	fmt.Printf("Registered handlers...\n")

	wg.Add(4)
	go h.Handler(wg)
}

// Stop deregisters handlers
func (h *EventsHandler) Stop() {
	defer fmt.Printf("Deregistered handlers...\n")
	shared_services.Bus.DeregisterHandler(stockWorker)
	shared_services.Bus.DeregisterHandler(ordersSummaryWorker)
	cancel()
}

func (h *EventsHandler) Handler(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-c.Done():
			return
		case e := <-stockEvent:
			h.SharedProductService.HandleStock(c, e.Data.(map[string]interface{}), e.Topic)
		case e := <-ordersSummaryEvent:
			h.SharedOrdersSummaryService.HandleOrdersSummary(c, e.Data.(map[string]interface{}), e.Topic)
		}
	}
}
