package structures

import "time"

type Log struct {
	Datetime        time.Time `db:"datetime" json:"datetime" binding:"required"`
	ServiceMarketID int       `db:"service_market_id" json:"service_market_id" binding:"required"`
	Type            string    `db:"event_type" json:"event_type" binding:"required"`
}

type LogData struct {
	Datetime int64  `json:"datetime" binding:"required"`
	Type     string `json:"event_type" binding:"required"`
}

type LogInput struct {
	ServiceMarketID int       `json:"service_market_id" binding:"required"`
	Logs            []LogData `json:"logs" binding:"required"`
}

type LogOutput struct {
	ServiceMarketName string `json:"service_market_name"`
	ServiceMarketID   int    `json:"service_market_id" binding:"required"`
	Datetime          int64  `json:"datetime" binding:"required"`
	Type              string `json:"event_type" binding:"required"`
}
