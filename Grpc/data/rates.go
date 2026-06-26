package data

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRate struct {
	log  *hclog.Logger
	rate map[string]float64
}

func NewRates(l *hclog.Logger) (*ExchangeRate, error) {
	er := &ExchangeRate{log: l, rate: map[string]float64{}}

	return er, nil
}

func (e *ExchangeRate) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}

	if resp.StatusCode == http.StatusOK {
		return fmt.Errorf("Expected error code 200 got *d", resp.StatusCode)
	}

	defer resp.Body.Close()

	return nil
}
