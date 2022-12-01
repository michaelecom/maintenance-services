package store

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"rimeks.ru/services/pkg/app/structures"
)

type LoggingPostgres struct {
	db *sqlx.DB
}

func NewLoggingPostgres(db *sqlx.DB) *LoggingPostgres {
	return &LoggingPostgres{db: db}
}

func (s *LoggingPostgres) CreateLog(log structures.Log) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (datetime, service_market_id, event_type) VALUES ($1, $2, $3) RETURNING id", loggingTable)
	row := s.db.QueryRow(query, log.Datetime, log.ServiceMarketID, log.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LoggingPostgres) GetAllLogs() ([]structures.Log, error) {
	var logs []structures.Log

	query := fmt.Sprintf("SELECT datetime, service_market_id, event_type FROM %s", loggingTable)
	err := s.db.Select(&logs, query)

	return logs, err
}

func (s *LoggingPostgres) GetAllLogsByServiceMarketID(id int) ([]structures.Log, error) {
	var logs []structures.Log

	query := fmt.Sprintf("SELECT datetime, service_market_id, event_type FROM %s WHERE service_market_id = $1", loggingTable)
	err := s.db.Select(&logs, query, id)

	return logs, err
}

func (s *LoggingPostgres) ClearLogs() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	var query string

	query = fmt.Sprintf("DROP TABLE %s", loggingTable)
	if _, err := s.db.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	var fields []string

	fields = append(fields, "id serial not null unique")
	fields = append(fields, "datetime timestamp not null")
	fields = append(fields, "service_market_id int not null")
	fields = append(fields, "event_type text")

	query = fmt.Sprintf("CREATE TABLE %s (%s)", loggingTable, strings.Join(fields, ", "))
	if _, err := s.db.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *LoggingPostgres) GetServiceMarketNameByID(id int) (string, error) {
	var names []string

	query := fmt.Sprintf("SELECT service_market_name FROM %s WHERE service_market_id = $1", serviceMarketsTable)
	err := s.db.Select(&names, query, id)

	if len(names) != 0 {
		return names[0], err
	}

	return "", fmt.Errorf("не найден сервис - маркет с ID %d", id)
}
