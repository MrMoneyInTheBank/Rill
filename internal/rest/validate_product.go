package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const restEndpoint = "https://api.coinbase.com/api/v3/brokerage/market/products/"

func ValidateProduct(product string) (error, bool) {
	url := restEndpoint + product
	reqCtx, reqCtxCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer reqCtxCancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Join(ErrBadReq, err), false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(reqCtx.Err(), context.DeadlineExceeded) {
			return ErrReqTimedOut, false
		}
		return errors.Join(ErrReqFailed, err), false
	}
	reqCtxCancel()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s. StatusCode: %d", ErrCantValidate, product, res.StatusCode), false
	}

	return nil, true
}
