package kite

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strings"
)

const (
	LimitOrderPercentage = 5
	TickSize             = 0.05
)

func (kite *Kite) PlaceOrderInKite(kOrder *KiteOrderPayload) (*KiteResponsePayload, error) {
	url := kite.BaseUrl + "/oms/orders/" + kOrder.Variety
	queries := make([]string, 0)
	typ := reflect.TypeOf(*kOrder)
	val := reflect.ValueOf(kOrder).Elem()
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		ft, _ := typ.FieldByName(fieldName)
		fv := val.FieldByName(fieldName)
		queries = append(queries, fmt.Sprintf("%v=%v", ft.Tag.Get("query"), fv))
	}
	payload := strings.Join(queries, "&")

	//Limit Order
	lastPrice, err := kite.GetLastPrice(kOrder.Exchange, kOrder.TradingSymbol)
	if err != nil {
		return nil, err
	}
	if lastPrice > 0 {
		kOrder.OrderType = "LIMIT"
		if kOrder.TransactionType == "BUY" {
			kOrder.Price = fmt.Sprintf("%v", math.Floor(lastPrice*(1-LimitOrderPercentage/100)/TickSize)*TickSize)
		}
		if kOrder.TransactionType == "SELL" {
			kOrder.Price = fmt.Sprintf("%v", math.Ceil(lastPrice*(1+LimitOrderPercentage/100)/TickSize)*TickSize)
		}
	}
	headers := make(map[string]string)
	headers["authorization"] = kite.Token
	headers["content-type"] = "application/x-www-form-urlencoded"
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var kiteResponse *KiteResponsePayload
	err = json.NewDecoder(response.Body).Decode(&kiteResponse)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 && kiteResponse.Data != nil && kiteResponse.Data.OrderId != "" {
		return kiteResponse, nil
	}
	return nil, errors.New(kiteResponse.Message)
}

func (kite *Kite) GetLastPrice(exchange string, tradingSymbol string) (float64, error) {
	url := kite.BaseUrl + "/oms/quote?i=" + exchange + ":" + tradingSymbol
	headers := make(map[string]string)
	headers["authorization"] = kite.Token
	headers["content-type"] = "application/x-www-form-urlencoded"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(``))
	if err != nil {
		return 0, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	var respData *KiteQuoteResponsePayload
	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return 0, err
	}
	return (*respData.Data)[exchange+":"+tradingSymbol].LastPrice, nil
}
