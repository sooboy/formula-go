// Package types provides common types for market data and formula results
package types

import "fmt"

// MarketData represents a single market data record with OHLCV data
type MarketData struct {
	Open   float64 // Opening price
	Close  float64 // Closing price
	High   float64 // Highest price in the period
	Low    float64 // Lowest price in the period
	Volume float64 // Trading volume (number of shares/units traded)
	Amount float64 // Trading amount (volume * price, optional)
}

// NewMarketData creates a new MarketData instance
func NewMarketData(open, close, high, low, volume, amount float64) *MarketData {
	return &MarketData{
		Open:   open,
		Close:  close,
		High:   high,
		Low:    low,
		Volume: volume,
		Amount: amount,
	}
}

// Validate validates a MarketData object to ensure all required fields are present
// and have correct types and logical constraints
//
// Validation rules:
// - open, close, high, low must be numbers
// - volume must be a non-negative number
// - amount must be a non-negative number
// - high must be >= low
func (m *MarketData) Validate() error {
	// Validate logical constraints
	if m.High < m.Low {
		return fmt.Errorf("high (%f) must be >= low (%f)", m.High, m.Low)
	}

	// Volume must be non-negative
	if m.Volume < 0 {
		return fmt.Errorf("volume must be non-negative, got %f", m.Volume)
	}

	// Amount must be non-negative
	if m.Amount < 0 {
		return fmt.Errorf("amount must be non-negative, got %f", m.Amount)
	}

	return nil
}

// GetMarketDataLength returns the length of a MarketData slice
func GetMarketDataLength(data []*MarketData) int {
	return len(data)
}
