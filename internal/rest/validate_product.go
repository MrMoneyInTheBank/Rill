package rest

import (
	"errors"
	"fmt"
	"net/http"
)

const restEndpoint = "https://api.coinbase.com/api/v3/brokerage/market/products/"

var ErrInvalidProduct = errors.New("Invalid Product")

func ValidateProduct(product string) (error, bool) {
	url := restEndpoint + product

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return err, false
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s", ErrInvalidProduct, product), false
	}

	return nil, true
}
