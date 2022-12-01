package service

import (
	"rimeks.ru/services/pkg/app/store"
	"rimeks.ru/services/pkg/app/structures"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type OngoingMaintenance interface {
	CreateOrder(order structures.OrderList) (int, error)
	GetAllOrders() ([]structures.OrderList, error)
	GetAllOrdersByServiceMarketID(id int) ([]structures.OrderData, error)
	UpdateOrder(order structures.OrderList) error
	DeleteOrder(order structures.OrderList) error

	ClearOrders() error
}

type Logging interface {
	CreateLog(log structures.LogInput) ([]int, error)
	GetAllLogs() ([]structures.LogOutput, error)
	GetAllLogsByServiceMarketID(id int) ([]structures.LogOutput, error)

	ClearLogs() error
	GetServiceMarketNameByID(id int) (string, error)
}

type Service struct {
	OngoingMaintenance
	Logging
}

func New(store *store.Store) *Service {
	return &Service{
		OngoingMaintenance: NewOngoingMaintenanceService(store.OngoingMaintenance),
		Logging:            NewLoggingService(store.Logging),
	}
}
