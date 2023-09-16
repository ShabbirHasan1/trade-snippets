package kite

type Kite struct {
	Token   string
	BaseUrl string
}

type KiteOrderPayload struct {
	Exchange          string `query:"exchange"`
	TradingSymbol     string `query:"tradingsymbol"`
	TransactionType   string `query:"transaction_type"`
	Product           string `query:"product"`
	Quantity          string `query:"quantity"`
	Price             string `query:"price"`
	Variety           string `query:"variety"`
	OrderType         string `query:"order_type"`
	Validity          string `query:"validity"`
	DisclosedQuantity string `query:"disclosed_quantity"`
	TriggerPrice      string `query:"trigger_price"`
	SquareOff         string `query:"squareoff"`
	StopLoss          string `query:"stoploss"`
	TrailingStopLoss  string `query:"trailing_stoploss"`
}

type KiteQuoteResponsePayload struct {
	Status    string `json:"error"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type"`
	Data      *map[string]struct {
		LastPrice float64 `json:"last_price"`
	} `json:"data"`
}

type KiteResponsePayload struct {
	Status  string `json:"error"`
	Message string `json:"message"`
	Data    *struct {
		OrderId string `json:"order_id"`
	} `json:"data"`
}
