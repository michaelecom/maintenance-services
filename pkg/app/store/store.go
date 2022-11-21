package store

import (
	"github.com/jmoiron/sqlx"
	"rimeks.ru/services/pkg/app/structures"
)

type OngoingMaintenance interface {
	CreateOrder(order structures.Order) (int, error)
	GetAllOrders() ([]structures.Order, error)
	GetAllOrdersByServiceMarketID(id int) ([]structures.Order, error)
	UpdateOrder(order structures.Order) error
	DeleteOrder(order structures.Order) error

	Clear() error
}

type Store struct {
	OngoingMaintenance
}

func New(db *sqlx.DB) *Store {
	return &Store{
		OngoingMaintenance: NewOngoingMaintenancePostgres(db),
	}
}
