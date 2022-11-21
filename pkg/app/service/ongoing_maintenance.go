package service

import (
	"rimeks.ru/services/pkg/app/store"
	"rimeks.ru/services/pkg/app/structures"
)

type OngoingMaintenanceService struct {
	store store.OngoingMaintenance
}

func NewOngoingMaintenanceService(store store.OngoingMaintenance) *OngoingMaintenanceService {
	return &OngoingMaintenanceService{store: store}
}

func (s *OngoingMaintenanceService) CreateOrder(order structures.OrderList) (int, error) {
	return s.store.CreateOrder(structures.Order{
		ServiceMarketID: order.ServiceMarketID,
		OrderNumber:     order.Orders[0].OrderNumber,
		CarBrand:        order.Orders[0].CarBrand,
		CarModel:        order.Orders[0].CarModel,
		CarNumber:       order.Orders[0].CarNumber,
	})
}

func (s *OngoingMaintenanceService) GetAllOrders() ([]structures.OrderList, error) {
	allOrders, err := s.store.GetAllOrders()

	var allOrdersSorted = make(map[int][]structures.OrderData)
	var allOrdersList = []structures.OrderList{}

	for _, order := range allOrders {
		allOrdersSorted[order.ServiceMarketID] = append(
			allOrdersSorted[order.ServiceMarketID],
			structures.OrderData{
				OrderNumber: order.OrderNumber,
				CarBrand:    order.CarBrand,
				CarModel:    order.CarModel,
				CarNumber:   order.CarNumber,
			})
	}

	for ServiceMarketID, Orders := range allOrdersSorted {
		allOrdersList = append(allOrdersList, structures.OrderList{
			ServiceMarketID: ServiceMarketID,
			Orders:          Orders,
		})
	}

	return allOrdersList, err
}

func (s *OngoingMaintenanceService) GetAllOrdersByServiceMarketID(id int) ([]structures.OrderData, error) {
	allOrders, err := s.store.GetAllOrdersByServiceMarketID(id)

	var allOrdersList = []structures.OrderData{}

	for _, order := range allOrders {
		allOrdersList = append(
			allOrdersList,
			structures.OrderData{
				OrderNumber: order.OrderNumber,
				CarBrand:    order.CarBrand,
				CarModel:    order.CarModel,
				CarNumber:   order.CarNumber,
			})
	}

	return allOrdersList, err
}

func (s *OngoingMaintenanceService) UpdateOrder(order structures.OrderList) error {
	return s.store.UpdateOrder(structures.Order{
		ServiceMarketID: order.ServiceMarketID,
		OrderNumber:     order.Orders[0].OrderNumber,
		CarBrand:        order.Orders[0].CarBrand,
		CarModel:        order.Orders[0].CarModel,
		CarNumber:       order.Orders[0].CarNumber,
	})
}

func (s *OngoingMaintenanceService) DeleteOrder(order structures.OrderList) error {
	return s.store.DeleteOrder(structures.Order{
		ServiceMarketID: order.ServiceMarketID,
		OrderNumber:     order.Orders[0].OrderNumber,
		CarBrand:        order.Orders[0].CarBrand,
		CarModel:        order.Orders[0].CarModel,
		CarNumber:       order.Orders[0].CarNumber,
	})
}

func (s *OngoingMaintenanceService) Clear() error {
	return s.store.Clear()
}
