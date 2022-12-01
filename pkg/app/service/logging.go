package service

import (
	"time"

	"rimeks.ru/services/pkg/app/store"
	"rimeks.ru/services/pkg/app/structures"
)

type LoggingService struct {
	store store.Logging
}

func NewLoggingService(store store.Logging) *LoggingService {
	return &LoggingService{store: store}
}

func (s *LoggingService) CreateLog(list structures.LogInput) ([]int, error) {
	var ids []int

	_, err := s.store.GetServiceMarketNameByID(list.ServiceMarketID)
	if err != nil {
		return nil, err
	}

	for _, log := range list.Logs {
		id, err := s.store.CreateLog(structures.Log{
			Datetime:        time.Unix(log.Datetime, 0),
			ServiceMarketID: list.ServiceMarketID,
			Type:            log.Type,
		})
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (s *LoggingService) GetAllLogs() ([]structures.LogOutput, error) {
	allLogs, err := s.store.GetAllLogs()

	var allLogsList = []structures.LogOutput{}

	for _, log := range allLogs {
		ServiceMarketName, err := s.store.GetServiceMarketNameByID(log.ServiceMarketID)
		if err != nil {
			return nil, err
		}

		allLogsList = append(allLogsList, structures.LogOutput{
			ServiceMarketName: ServiceMarketName,
			ServiceMarketID:   log.ServiceMarketID,
			Datetime:          log.Datetime.Unix(),
			Type:              log.Type,
		})
	}

	return allLogsList, err
}

func (s *LoggingService) GetAllLogsByServiceMarketID(id int) ([]structures.LogOutput, error) {
	allLogs, err := s.store.GetAllLogsByServiceMarketID(id)

	var allLogsList = []structures.LogOutput{}

	for _, log := range allLogs {
		ServiceMarketName, err := s.store.GetServiceMarketNameByID(log.ServiceMarketID)
		if err != nil {
			return nil, err
		}

		allLogsList = append(allLogsList, structures.LogOutput{
			ServiceMarketName: ServiceMarketName,
			ServiceMarketID:   log.ServiceMarketID,
			Datetime:          log.Datetime.Unix(),
			Type:              log.Type,
		})
	}

	return allLogsList, err
}

func (s *LoggingService) ClearLogs() error {
	return s.store.ClearLogs()
}

func (s *LoggingService) GetServiceMarketNameByID(id int) (string, error) {
	return s.store.GetServiceMarketNameByID(id)
}
