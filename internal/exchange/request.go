package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

var APP_ID string

// request performs an http request to obtain the exchange rates of the
// quote currency from the Open Exchange Rates API.
func request(logger *zap.Logger, quantity float64, quote []byte) (float64, int) {
	// Create an HTTP client
	c := &http.Client{Timeout: 10 * time.Second}

	// Obtain authentication token
	if APP_ID := os.Getenv("CC_APP_ID"); APP_ID == "" {
		logger.Error("exchange: APP_ID not set as environemnt variable")
		return 0, http.StatusUnauthorized
	}

	// Create HTTP request
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", os.Getenv("CC_APP_ID"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("exchange: couldn't create an HTTP request: ", zap.Error(err))
		return 0, http.StatusBadRequest
	}

	// Set application header
	req.Header.Add("accept", "application/json")

	// Perform the HTTP request
	resp, err := c.Do(req)
	if err != nil {
		logger.Error("exchange: couldn't perform an HTTP request: ", zap.Error(err))
		return 0, http.StatusBadRequest
	}

	// Read body from the response
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("exchange: couldn't read from the response body: ", zap.Error(err))
		return 0, http.StatusInternalServerError
	}

	var rates map[string]interface{}
	if err := json.Unmarshal(body, &rates); err != nil {
		logger.Error("exchange: couldn't unmarshal json body into rates: ", zap.Error(err))
		return 0, http.StatusInternalServerError
	}

	// TODO : from the rates["rates"] extract the selected quote and apply the formula for the going rate
	// TODO : use regex support from standard library to finde the selected quote currency
	// https://www.honeybadger.io/blog/a-definitive-guide-to-regular-expressions-in-go/
	// TODO: type assertion might be necessary to convert the map[string]interface{}
	// map[AED:3.672725 AFN:72.50001 ALL:91.680889 AMD:394.500358 ANG:1.804457

	return 100, http.StatusOK
}
