package model

import (
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/model"
	"github.com/shopspring/decimal"
)

// OrdersRes ...
type OrdersRes struct {
	Data struct {
		List []struct {
			RootOrderID                 int64 `json:"rootOrderId"`
			OrderID                     int64 `json:"orderId"`
			configuration.Symbol        `json:"symbol"`
			configuration.Side          `json:"side"`
			configuration.OrderType     `json:"orderType"`
			configuration.ExecutionType `json:"executionType"`
			configuration.SettleType    `json:"settleType"`
			Size                        decimal.Decimal           `json:"size"`
			ExecutedSize                decimal.Decimal           `json:"executedSize"`
			Price                       decimal.Decimal           `json:"price"`
			LossCutPrice                decimal.Decimal           `json:"losscutPrice"`
			Status                      configuration.OrderStatus `json:"status"`
			configuration.CancelType    `json:"cancelType"`
			configuration.TimeInForce   `json:"timeInForce"`
			Timestamp                   time.Time `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	model.ResponseCommon
}
