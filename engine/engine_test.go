package engine

import (
	"math"
	"testing"

	"github.com/DTrader-store/formula-go/types"
)

// createTestData creates sample market data for testing
func createTestData() []*types.MarketData {
	return []*types.MarketData{
		types.NewMarketData(100, 105, 107, 99, 1000, 100000),
		types.NewMarketData(105, 103, 108, 102, 1100, 110000),
		types.NewMarketData(103, 107, 109, 101, 1200, 120000),
		types.NewMarketData(107, 110, 112, 106, 1300, 130000),
		types.NewMarketData(110, 108, 113, 107, 1400, 140000),
		types.NewMarketData(108, 111, 114, 107, 1500, 150000),
		types.NewMarketData(111, 109, 115, 108, 1600, 160000),
		types.NewMarketData(109, 112, 116, 108, 1700, 170000),
		types.NewMarketData(112, 115, 117, 110, 1800, 180000),
		types.NewMarketData(115, 113, 118, 112, 1900, 190000),
	}
}

func TestEngineSimpleExpression(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "CLOSE"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	if len(result.Outputs[0].Data) != len(marketData) {
		t.Errorf("Expected %d data points, got %d", len(marketData), len(result.Outputs[0].Data))
	}
}

func TestEngineMA(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "MA5 := MA(CLOSE, 5)"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	ma5 := result.Outputs[0]
	if ma5.Name != "MA5" {
		t.Errorf("Expected output name 'MA5', got '%s'", ma5.Name)
	}

	// First 4 values should be NaN
	for i := 0; i < 4; i++ {
		if !math.IsNaN(ma5.Data[i]) {
			t.Errorf("Expected NaN at index %d, got %f", i, ma5.Data[i])
		}
	}

	// 5th value should be average of first 5 closes
	expected := (105.0 + 103.0 + 107.0 + 110.0 + 108.0) / 5.0
	if math.Abs(ma5.Data[4]-expected) > 0.01 {
		t.Errorf("Expected MA5[4] = %.2f, got %.2f", expected, ma5.Data[4])
	}
}

func TestEngineArithmetic(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "DIFF := HIGH - LOW"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	diff := result.Outputs[0]
	for i := range marketData {
		expected := marketData[i].High - marketData[i].Low
		if math.Abs(diff.Data[i]-expected) > 0.01 {
			t.Errorf("Index %d: expected %.2f, got %.2f", i, expected, diff.Data[i])
		}
	}
}

func TestEngineComparison(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "SIGNAL := CLOSE > OPEN"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	signal := result.Outputs[0]
	for i := range marketData {
		expected := 0.0
		if marketData[i].Close > marketData[i].Open {
			expected = 1.0
		}
		if signal.Data[i] != expected {
			t.Errorf("Index %d: expected %.0f, got %.0f", i, expected, signal.Data[i])
		}
	}
}

func TestEngineMultipleStatements(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := `
		MA5 := MA(CLOSE, 5)
		MA10 := MA(CLOSE, 10)
		CROSS_SIGNAL := CROSS(MA5, MA10)
	`

	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 3 {
		t.Fatalf("Expected 3 outputs, got %d", len(result.Outputs))
	}

	names := []string{"MA5", "MA10", "CROSS_SIGNAL"}
	for i, output := range result.Outputs {
		if output.Name != names[i] {
			t.Errorf("Expected output name '%s', got '%s'", names[i], output.Name)
		}
	}
}

func TestEngineREF(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "PREV_CLOSE := REF(CLOSE, 1)"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	prevClose := result.Outputs[0]

	// First value should be NaN
	if !math.IsNaN(prevClose.Data[0]) {
		t.Errorf("Expected NaN at index 0, got %f", prevClose.Data[0])
	}

	// Check other values
	for i := 1; i < len(marketData); i++ {
		expected := marketData[i-1].Close
		if math.Abs(prevClose.Data[i]-expected) > 0.01 {
			t.Errorf("Index %d: expected %.2f, got %.2f", i, expected, prevClose.Data[i])
		}
	}
}

func TestEngineIF(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "RESULT := IF(CLOSE > OPEN, HIGH, LOW)"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	output := result.Outputs[0]
	for i := range marketData {
		expected := marketData[i].Low
		if marketData[i].Close > marketData[i].Open {
			expected = marketData[i].High
		}
		if math.Abs(output.Data[i]-expected) > 0.01 {
			t.Errorf("Index %d: expected %.2f, got %.2f", i, expected, output.Data[i])
		}
	}
}

func TestEngineHHVLLV(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := `
		HIGHEST := HHV(HIGH, 5)
		LOWEST := LLV(LOW, 5)
	`

	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 2 {
		t.Fatalf("Expected 2 outputs, got %d", len(result.Outputs))
	}

	// Verify HHV
	highest := result.Outputs[0]
	for i := 4; i < len(marketData); i++ {
		maxVal := marketData[i].High
		for j := 1; j < 5; j++ {
			if marketData[i-j].High > maxVal {
				maxVal = marketData[i-j].High
			}
		}
		if math.Abs(highest.Data[i]-maxVal) > 0.01 {
			t.Errorf("HHV at index %d: expected %.2f, got %.2f", i, maxVal, highest.Data[i])
		}
	}
}

func TestEngineEMA(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	formula := "EMA5 := EMA(CLOSE, 5)"
	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(result.Outputs))
	}

	ema5 := result.Outputs[0]
	if ema5.Name != "EMA5" {
		t.Errorf("Expected output name 'EMA5', got '%s'", ema5.Name)
	}

	// First value should equal first close
	if math.Abs(ema5.Data[0]-marketData[0].Close) > 0.01 {
		t.Errorf("Expected EMA5[0] = %.2f, got %.2f", marketData[0].Close, ema5.Data[0])
	}
}

func TestEngineComplexFormula(t *testing.T) {
	engine := NewFormulaEngine()
	marketData := createTestData()

	// MACD-like formula
	formula := `
		EMA12 := EMA(CLOSE, 5)
		EMA26 := EMA(CLOSE, 8)
		DIF := EMA12 - EMA26
		DEA := EMA(DIF, 3)
		MACD := (DIF - DEA) * 2
	`

	result, err := engine.Run(formula, marketData)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result.Outputs) != 5 {
		t.Fatalf("Expected 5 outputs, got %d", len(result.Outputs))
	}

	expectedNames := []string{"EMA12", "EMA26", "DIF", "DEA", "MACD"}
	for i, name := range expectedNames {
		if result.Outputs[i].Name != name {
			t.Errorf("Output %d: expected name '%s', got '%s'", i, name, result.Outputs[i].Name)
		}
	}
}
