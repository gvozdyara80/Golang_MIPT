package currency

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewCurrencyClient(t *testing.T) {
	apiKey := "testAPIKey"
	client := NewCurrencyClient(apiKey)

	if client.APIKey != apiKey {
		t.Errorf("APIKey not set correctly. Expected: %s, Got: %s", apiKey, client.APIKey)
	}

	if client.Client.Timeout != 10*time.Second {
		t.Errorf("HTTP client timeout not set correctly. Expected: 10s, Got: %s", client.Client.Timeout)
	}
}

func TestGetCurrencyRate(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := `{"data": {"USD": 1.0, "EUR": 0.85}}`
		w.Write([]byte(resp))
	}))
	defer mockServer.Close()

	c := CurrencyClient{
		Client:  mockServer.Client(),
		APIKey:  "test_key",
		BaseURL: mockServer.URL,
	}

	t.Run("Successful API Call", func(t *testing.T) {
		rates, err := c.GetCurrencyRate("USD")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedRates := map[string]float64{"USD": 1.0, "EUR": 0.85}
		if !reflect.DeepEqual(rates, expectedRates) {
			t.Errorf("Expected rates to be %v, got %v", expectedRates, rates)
		}
	})

	t.Run("Unsuccessful API Call", func(t *testing.T) {
		// Simulate API call failure
		c.BaseURL = "invalid_url"
		_, err := c.GetCurrencyRate("USD")
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})

	t.Run("Unsuccessful Decoding Response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			resp := `{"data": "invalid_response"}`
			w.Write([]byte(resp))
		}))
		defer mockServer.Close()

		c := CurrencyClient{
			Client:  mockServer.Client(),
			APIKey:  "test_key",
			BaseURL: mockServer.URL,
		}

		_, err := c.GetCurrencyRate("USD")
		if err == nil {
			t.Error("Expected an error due to decoding failure, got nil")
		}
	})
}
