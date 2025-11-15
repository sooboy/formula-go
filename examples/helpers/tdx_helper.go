package helpers

import (
	"fmt"

	"github.com/DTrader-store/formula-go/types"
	"github.com/injoyai/tdx"
	"github.com/injoyai/tdx/protocol"
)

// TDXClient wraps the TDX connection for easy data fetching
type TDXClient struct {
	client *tdx.Client
}

// NewTDXClient creates a new TDX client with auto-reconnect
func NewTDXClient(addr string) (*TDXClient, error) {
	client, err := tdx.Dial(addr, tdx.WithRedial())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to TDX server: %w", err)
	}
	return &TDXClient{client: client}, nil
}

// Close closes the TDX connection
func (c *TDXClient) Close() error {
	return c.client.Close()
}

// GetMarketData fetches K-line data and converts to MarketData format
func (c *TDXClient) GetMarketData(code string, start uint16, count uint16) ([]*types.MarketData, error) {
	// Fetch K-line data from TDX
	klines, err := c.client.GetKlineDay(code, start, count)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch K-line data: %w", err)
	}

	// Convert to MarketData format
	marketData := make([]*types.MarketData, len(klines.List))
	for i, kline := range klines.List {
		marketData[i] = KlineToMarketData(kline)
	}

	return marketData, nil
}

// KlineToMarketData converts TDX Kline to formula MarketData
func KlineToMarketData(kline *protocol.Kline) *types.MarketData {
	return &types.MarketData{
		Open:   float64(kline.Open),
		High:   float64(kline.High),
		Low:    float64(kline.Low),
		Close:  float64(kline.Close),
		Volume: float64(kline.Volume),
		Amount: float64(kline.Amount),
	}
}

// DefaultTDXServer returns a reliable TDX server address
func DefaultTDXServer() string {
	return "124.71.187.122:7709" // Shanghai Huawei server
}
