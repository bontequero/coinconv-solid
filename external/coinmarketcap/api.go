package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"convertor/entities"
)

type priceConvertor struct{}

func New() *priceConvertor {
	return &priceConvertor{}
}

func (c *priceConvertor) ConvertPrice(data entities.ConversionData) (entities.ConversionResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v2/tools/price-conversion", nil)
	if err != nil {
		return entities.ConversionResult{}, fmt.Errorf("cant create http request: %w", err)
	}

	q := url.Values{}
	amount := strconv.FormatFloat(data.Amount, 'E', -1, 64)
	symbol := string(data.From)
	convertTo := string(data.To)

	q.Add("amount", amount)
	q.Add("symbol", symbol)
	q.Add("convert", convertTo)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return entities.ConversionResult{}, fmt.Errorf("error sending request to server: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return entities.ConversionResult{}, fmt.Errorf("request failed, status=%v", resp.Status)
	}

	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return entities.ConversionResult{}, fmt.Errorf("invalid json in response: %w", err)
	}
	return entities.ConversionResult{Result: r.Data[symbol].Quote[convertTo].Price}, nil
}

type response struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}
