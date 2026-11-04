package rest_test

import (
	"errors"
	"testing"

	"github.com/MrMoneyInTheBank/rill/internal/rest"
)

func TestValidateProduct(t *testing.T) {
	cases := []struct {
		name    string
		product string
		wantErr error
	}{
		{"ValidProduct", "BTC-USD", nil},
		{"InvalidProduct", "abc", rest.ErrInvalidProduct},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err, _ := rest.ValidateProduct(tc.product)

			if !errors.Is(err, tc.wantErr) {
				t.Error("Want error %w; got %w", tc.wantErr, err)
			}
		})
	}
}
