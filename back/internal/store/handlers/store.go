package store_handlers

import (
	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
)

type StoreHandlers struct {
	storeService store_ports.StoreService
}

// !TEST
var _ store_ports.StoreHandlers = (*StoreHandlers)(nil)

func NewStoreHandlers(storeService store_ports.StoreService) *StoreHandlers {
	return &StoreHandlers{
		storeService: storeService,
	}
}
