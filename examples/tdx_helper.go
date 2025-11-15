package main

import (
	"fmt"
	"time"

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
	marketData := make([]*types.MarketData, len(klines))
	for i, kline := range klines {
		marketData[i] = KlineToMarketData(&kline)
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
		Time:   kline.Time,
	}
}

// PrintMarketData prints market data in a readable format
func PrintMarketData(data []*types.MarketData, limit int) {
	if limit > 0 && limit < len(data) {
		data = data[:limit]
	}

	fmt.Println("\nMarket Data:")
	fmt.Println("Time                 | Open      | High      | Low       | Close     | Volume")
	fmt.Println("--------------------------------------------------------------------------------")
	for _, d := range data {
		fmt.Printf("%s | %9.2f | %9.2f | %9.2f | %9.2f | %10.0f\n",
			d.Time.Format("2006-01-02 15:04:05"),
			d.Open, d.High, d.Low, d.Close, d.Volume)
	}
}

// PrintFormulaResult prints formula calculation results
func PrintFormulaResult(result *types.FormulaResult, limit int) {
	if result == nil {
		fmt.Println("No result")
		return
	}

	fmt.Printf("\nFormula Result (Type: %s):\n", result.Type)

	if result.Type == types.ResultTypeScalar {
		fmt.Printf("Value: %.4f\n", result.ScalarValue)
		return
	}

	if len(result.ArrayValue) == 0 {
		fmt.Println("Empty array result")
		return
	}

	if limit > 0 && limit < len(result.ArrayValue) {
		fmt.Printf("Showing last %d values:\n", limit)
		fmt.Println("Index | Value")
		fmt.Println("-------------------")
		for i := len(result.ArrayValue) - limit; i < len(result.ArrayValue); i++ {
			fmt.Printf("%5d | %10.4f\n", i, result.ArrayValue[i])
		}
	} else {
		fmt.Printf("All %d values:\n", len(result.ArrayValue))
		fmt.Println("Index | Value")
		fmt.Println("-------------------")
		for i, v := range result.ArrayValue {
			fmt.Printf("%5d | %10.4f\n", i, v)
		}
	}
}

// PrintSignals prints signal detection results with timestamps
func PrintSignals(signals []int, marketData []*types.MarketData, signalName string) {
	fmt.Printf("\n%s Signals:\n", signalName)
	fmt.Println("Date       | Price     | Signal")
	fmt.Println("--------------------------------------")

	for i, signal := range signals {
		if signal > 0 && i < len(marketData) {
			fmt.Printf("%s | %9.2f | %s\n",
				marketData[i].Time.Format("2006-01-02"),
				marketData[i].Close,
				signalName)
		}
	}
}

// DefaultTDXServer returns a reliable TDX server address
func DefaultTDXServer() string {
	return "124.71.187.122:7709" // Shanghai Huawei server
}
