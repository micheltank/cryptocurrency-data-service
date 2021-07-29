package domain

import (
	"fmt"
	"strings"
)

type NetworkCode string

var supportedCoins = []string{"BTC", "LTC", "DOGE"}

func (n NetworkCode) IsSupported() bool {
	for _, coin := range supportedCoins {
		if coin == string(n) {
			return true
		}
	}
	return false
}

func (n NetworkCode) CheckIsSupported() error {
	if !n.IsSupported() {
		return NewError(nil, "Network code not supported", "error.validation", fmt.Sprintf("This network code isn't supported. Available: %s", strings.Join(supportedCoins, ", ")))
	}
	return nil
}
