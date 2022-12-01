package store

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"rimeks.ru/services/pkg/app/structures"
)

type OngoingMaintenancePostgres struct {
	db *sqlx.DB
}

func NewOngoingMaintenancePostgres(db *sqlx.DB) *OngoingMaintenancePostgres {
	return &OngoingMaintenancePostgres{db: db}
}

func (s *OngoingMaintenancePostgres) CreateOrder(order structures.Order) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (service_market_id, order_number, car_brand, car_model, car_number) VALUES ($1, $2, $3, $4, $5) RETURNING id", ongoingMaintenanceTable)
	row := s.db.QueryRow(query, order.ServiceMarketID, order.OrderNumber, order.CarBrand, order.CarModel, order.CarNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *OngoingMaintenancePostgres) GetAllOrders() ([]structures.Order, error) {
	var orders []structures.Order

	query := fmt.Sprintf("SELECT service_market_id, order_number, car_brand, car_model, car_number FROM %s", ongoingMaintenanceTable)
	err := s.db.Select(&orders, query)

	return orders, err
}

func (s *OngoingMaintenancePostgres) GetAllOrdersByServiceMarketID(id int) ([]structures.Order, error) {
	var orders []structures.Order

	query := fmt.Sprintf("SELECT service_market_id, order_number, car_brand, car_model, car_number FROM %s WHERE service_market_id = $1", ongoingMaintenanceTable)
	err := s.db.Select(&orders, query, id)

	return orders, err
}

func (s *OngoingMaintenancePostgres) UpdateOrder(order structures.Order) error {
	query := fmt.Sprintf("UPDATE %s SET car_brand = $3, car_model = $4, car_number = $5 WHERE service_market_id = $1 AND order_number = $2", ongoingMaintenanceTable)
	_, err := s.db.Exec(query, order.ServiceMarketID, order.OrderNumber, order.CarBrand, order.CarModel, order.CarNumber)

	return err
}

func (s *OngoingMaintenancePostgres) DeleteOrder(order structures.Order) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE service_market_id = $1 AND order_number = $2", ongoingMaintenanceTable)
	_, err := s.db.Exec(query, order.ServiceMarketID, order.OrderNumber)

	return err
}

func (s *OngoingMaintenancePostgres) ClearOrders() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	var query string

	query = fmt.Sprintf("DROP TABLE %s", ongoingMaintenanceTable)
	if _, err := s.db.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	var fields []string

	fields = append(fields, "id serial not null unique")
	fields = append(fields, "service_market_id int not null")
	fields = append(fields, "order_number varchar(30) not null unique")
	fields = append(fields, "car_brand varchar(128)")
	fields = append(fields, "car_model varchar(128)")
	fields = append(fields, "car_number varchar(10)")

	query = fmt.Sprintf("CREATE TABLE %s (%s)", ongoingMaintenanceTable, strings.Join(fields, ", "))
	if _, err := s.db.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
