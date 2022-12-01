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

func (s *OngoingMaintenanceService) CreateOrder(list structures.OrderList) (int, error) {
	return s.store.CreateOrder(structures.Order{
		ServiceMarketID: list.ServiceMarketID,
		OrderNumber:     list.Orders[0].OrderNumber,
		CarBrand:        list.Orders[0].CarBrand,
		CarModel:        list.Orders[0].CarModel,
		CarNumber:       list.Orders[0].CarNumber,
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

func (s *OngoingMaintenanceService) UpdateOrder(list structures.OrderList) error {
	return s.store.UpdateOrder(structures.Order{
		ServiceMarketID: list.ServiceMarketID,
		OrderNumber:     list.Orders[0].OrderNumber,
		CarBrand:        list.Orders[0].CarBrand,
		CarModel:        list.Orders[0].CarModel,
		CarNumber:       list.Orders[0].CarNumber,
	})
}

func (s *OngoingMaintenanceService) DeleteOrder(list structures.OrderList) error {
	return s.store.DeleteOrder(structures.Order{
		ServiceMarketID: list.ServiceMarketID,
		OrderNumber:     list.Orders[0].OrderNumber,
		CarBrand:        list.Orders[0].CarBrand,
		CarModel:        list.Orders[0].CarModel,
		CarNumber:       list.Orders[0].CarNumber,
	})
}

func (s *OngoingMaintenanceService) ClearOrders() error {
	return s.store.ClearOrders()
}
