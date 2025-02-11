package shared_services

import (
	"fmt"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
	"github.com/rs/zerolog"
)

// Bus is a ref to bus.Bus
var Bus *bus.Bus

// Monoton is an instance of monoton.Monoton
var Monoton monoton.Monoton

// Init inits the app config
func MessagingInit() {
	// Configure id generator
	node := uint64(1)
	initialTime := uint64(1577865600000)
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// Init an id generator
	var idGenerator bus.Next = m.Next

	// Create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// Register topics in here
	b.RegisterTopics(shared_domain.TopicList...)

	Bus = b
	Monoton = m

	fmt.Println("Initialized messaging & topics registered")
}

func NewEventMessaging(
	log *zerolog.Logger,
	eventType shared_ports.EventMessagingTypes,
) shared_ports.IEventMessaging {
	if eventType == shared_ports.UseBUS {
		return newBusMessaging(log)
	}

	if eventType == shared_ports.UseSNS {
		return newSnsMessaging(log)
	}

	return nil
}
