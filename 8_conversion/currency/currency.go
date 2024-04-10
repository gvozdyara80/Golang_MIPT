package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiURL = "https://freecurrencyapi.com/api/v2/latest?apikey=%s&base_currency=%s"
)

type CurrencyClient struct {
	APIKey string
	Client *http.Client
}

func NewCurrencyClient(apiKey string) *CurrencyClient {
	return &CurrencyClient{
		APIKey: apiKey,
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

type CurrencyResponse struct {
	Data map[string]float64 `json:"data"`
}

func (c *CurrencyClient) GetCurrencyRate(baseCurrency string) (map[string]float64, error) {
	url := fmt.Sprintf(apiURL, c.APIKey, baseCurrency)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currencyResp CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyResp); err != nil {
		return nil, err
	}

	return currencyResp.Data, nil
}