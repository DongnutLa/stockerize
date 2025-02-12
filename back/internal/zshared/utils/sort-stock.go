package utils

import product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"

func SortStock(a, b product_domain.Stock) int {
	// Manejar nulos (consideramos nil como "menor" que cualquier tiempo)
	switch {
	case a.CreatedAt == nil && b.CreatedAt == nil:
		return 0 // Iguales
	case a.CreatedAt == nil:
		return -1 // a va primero
	case b.CreatedAt == nil:
		return 1 // b va primero
	default:
		// Comparar tiempos
		aTime := *a.CreatedAt
		bTime := *b.CreatedAt

		switch {
		case aTime.Before(bTime):
			return -1
		case aTime.After(bTime):
			return 1
		default:
			return 0
		}
	}
}
