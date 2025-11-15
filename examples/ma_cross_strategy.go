package main

import (
	"fmt"
	"log"

	"github.com/DTrader-store/formula-go"
)

// MA Cross Strategy - Golden Cross and Death Cross Detection
// Golden Cross: Fast MA crosses above Slow MA (bullish signal)
// Death Cross: Fast MA crosses below Slow MA (bearish signal)

func main() {
	fmt.Println("=== MA Cross Strategy Example ===")
	fmt.Println("Detecting Golden Cross and Death Cross signals using real market data\n")

	// Initialize TDX client
	client, err := NewTDXClient(DefaultTDXServer())
	if err != nil {
		log.Fatalf("Failed to create TDX client: %v", err)
	}
	defer client.Close()

	// Fetch real market data for Ping An Bank (sz000001)
	stockCode := "sz000001"
	fmt.Printf("Fetching data for %s (Ping An Bank)...\n", stockCode)
	marketData, err := client.GetMarketData(stockCode, 0, 100)
	if err != nil {
		log.Fatalf("Failed to fetch market data: %v", err)
	}
	fmt.Printf("Fetched %d data points\n\n", len(marketData.Close))

	// Strategy 1: MA5 and MA10 Cross (Short-term)
	fmt.Println("--- Strategy 1: MA5/MA10 Cross (Short-term) ---")
	runMACrossStrategy(marketData, 5, 10)

	fmt.Println()

	// Strategy 2: MA10 and MA20 Cross (Medium-term)
	fmt.Println("--- Strategy 2: MA10/MA20 Cross (Medium-term) ---")
	runMACrossStrategy(marketData, 10, 20)

	fmt.Println()

	// Strategy 3: MA20 and MA60 Cross (Long-term)
	fmt.Println("--- Strategy 3: MA20/MA60 Cross (Long-term) ---")
	runMACrossStrategy(marketData, 20, 60)
}

func runMACrossStrategy(data *formula.MarketData, fastPeriod, slowPeriod int) {
	// Formula to detect MA cross
	formulaText := fmt.Sprintf(`
		FAST_MA: MA(CLOSE, %d);
		SLOW_MA: MA(CLOSE, %d);
		GOLDEN_CROSS: CROSS(FAST_MA, SLOW_MA);
		DEATH_CROSS: CROSS(SLOW_MA, FAST_MA);
	`, fastPeriod, slowPeriod)

	// Execute formula
	engine := formula.NewFormulaEngine()
	result, err := engine.Run(formulaText, data)
	if err != nil {
		log.Fatalf("Formula execution failed: %v", err)
	}

	// Analyze signals
	goldenCrosses := findSignals(result.Lines["GOLDEN_CROSS"].Data)
	deathCrosses := findSignals(result.Lines["DEATH_CROSS"].Data)

	fmt.Printf("Fast MA: %d, Slow MA: %d\n", fastPeriod, slowPeriod)
	fmt.Printf("Golden Crosses detected: %d\n", len(goldenCrosses))
	fmt.Printf("Death Crosses detected: %d\n", len(deathCrosses))

	if len(goldenCrosses) > 0 {
		fmt.Println("\nGolden Cross positions (bullish signals):")
		for _, pos := range goldenCrosses {
			if pos < len(data.Close) {
				fmt.Printf("  Position %d: Price = %.2f, Fast MA = %.2f, Slow MA = %.2f\n",
					pos,
					data.Close[pos],
					result.Lines["FAST_MA"].Data[pos],
					result.Lines["SLOW_MA"].Data[pos])
			}
		}
	}

	if len(deathCrosses) > 0 {
		fmt.Println("\nDeath Cross positions (bearish signals):")
		for _, pos := range deathCrosses {
			if pos < len(data.Close) {
				fmt.Printf("  Position %d: Price = %.2f, Fast MA = %.2f, Slow MA = %.2f\n",
					pos,
					data.Close[pos],
					result.Lines["FAST_MA"].Data[pos],
					result.Lines["SLOW_MA"].Data[pos])
			}
		}
	}

	// Show current status (last data point)
	lastIdx := len(data.Close) - 1
	if lastIdx >= 0 {
		fmt.Printf("\nCurrent Status (Last bar):\n")
		fmt.Printf("  Close Price: %.2f\n", data.Close[lastIdx])
		fmt.Printf("  Fast MA(%d): %.2f\n", fastPeriod, result.Lines["FAST_MA"].Data[lastIdx])
		fmt.Printf("  Slow MA(%d): %.2f\n", slowPeriod, result.Lines["SLOW_MA"].Data[lastIdx])

		if result.Lines["FAST_MA"].Data[lastIdx] > result.Lines["SLOW_MA"].Data[lastIdx] {
			fmt.Println("  Trend: BULLISH (Fast MA > Slow MA)")
		} else {
			fmt.Println("  Trend: BEARISH (Fast MA < Slow MA)")
		}
	}
}

// findSignals finds all positions where signal is 1 (cross detected)
func findSignals(data []float64) []int {
	var positions []int
	for i, val := range data {
		if val == 1.0 {
			positions = append(positions, i)
		}
	}
	return positions
}
