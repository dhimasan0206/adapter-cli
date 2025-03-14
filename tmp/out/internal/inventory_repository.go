package internal

import "github.com/dhimasan0206/adapter/sdk/api"

type inventoryRepository struct{}

func NewInventoryRepository() api.InventoryRepository {
	return &inventoryRepository{}
}
