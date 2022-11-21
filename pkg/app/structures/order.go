package structures

type Order struct {
	ServiceMarketID int    `db:"service_market_id" binding:"required"`
	OrderNumber     string `db:"order_number" binding:"required"`
	CarBrand        string `db:"car_brand"`
	CarModel        string `db:"car_model"`
	CarNumber       string `db:"car_number"`
}

type OrderData struct {
	OrderNumber string `json:"order_number" binding:"required"`
	CarBrand    string `json:"car_brand"`
	CarModel    string `json:"car_model"`
	CarNumber   string `json:"car_number"`
}

type OrderList struct {
	ServiceMarketID int         `json:"service_market_id" binding:"required"`
	Orders          []OrderData `json:"orders"`
}
