package kite

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var kite = &Kite{
	Token:   "enctoken <YOURTOKEN>",
	BaseUrl: "https://kite.zerodha.com",
}

func TestPlaceOrderInKite(t *testing.T) {
	response, err := kite.PlaceOrderInKite(&KiteOrderPayload{
		Exchange:          "NFO",
		TradingSymbol:     "NIFTY2391420000CE",
		TransactionType:   "BUY",
		Product:           "NRML",
		Quantity:          "250",
		Price:             "0",
		Variety:           "regular",
		OrderType:         "MARKET",
		Validity:          "DAY",
		DisclosedQuantity: "0",
		TriggerPrice:      "0",
		SquareOff:         "0",
		StopLoss:          "0",
		TrailingStopLoss:  "0",
	})
	log.WithFields(log.Fields{
		"response": response,
		"error":    err,
	}).Info("reponse")
}

func BenchmarkPlaceOrderInKite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kite.PlaceOrderInKite(&KiteOrderPayload{
			Exchange:          "NFO",
			TradingSymbol:     "NIFTY2391420000CE",
			TransactionType:   "BUY",
			Product:           "NRML",
			Quantity:          "250",
			Price:             "0",
			Variety:           "regular",
			OrderType:         "MARKET",
			Validity:          "DAY",
			DisclosedQuantity: "0",
			TriggerPrice:      "0",
			SquareOff:         "0",
			StopLoss:          "0",
			TrailingStopLoss:  "0",
		})
	}
}
