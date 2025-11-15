package types

import "testing"

func TestNewMarketData(t *testing.T) {
	data := NewMarketData(100.0, 102.0, 105.0, 98.0, 10000, 1000000)
	if data.Open != 100.0 {
		t.Errorf("Expected open 100.0, got %f", data.Open)
	}
	if data.Close != 102.0 {
		t.Errorf("Expected close 102.0, got %f", data.Close)
	}
	if data.High != 105.0 {
		t.Errorf("Expected high 105.0, got %f", data.High)
	}
	if data.Low != 98.0 {
		t.Errorf("Expected low 98.0, got %f", data.Low)
	}
	if data.Volume != 10000 {
		t.Errorf("Expected volume 10000, got %f", data.Volume)
	}
	if data.Amount != 1000000 {
		t.Errorf("Expected amount 1000000, got %f", data.Amount)
	}
}

func TestMarketDataValidate(t *testing.T) {
	tests := []struct {
		name    string
		data    *MarketData
		isValid bool
	}{
		{
			name:    "valid data",
			data:    NewMarketData(100.0, 102.0, 105.0, 98.0, 10000, 1000000),
			isValid: true,
		},
		{
			name:    "high < low",
			data:    NewMarketData(100.0, 102.0, 95.0, 98.0, 10000, 1000000),
			isValid: false,
		},
		{
			name:    "negative volume",
			data:    NewMarketData(100.0, 102.0, 105.0, 98.0, -10000, 1000000),
			isValid: false,
		},
		{
			name:    "negative amount",
			data:    NewMarketData(100.0, 102.0, 105.0, 98.0, 10000, -1000000),
			isValid: false,
		},
		{
			name:    "high equals low",
			data:    NewMarketData(100.0, 100.0, 100.0, 100.0, 0, 0),
			isValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.data.Validate()
			if tt.isValid && err != nil {
				t.Errorf("Expected valid data, got error: %v", err)
			}
			if !tt.isValid && err == nil {
				t.Error("Expected validation error, got nil")
			}
		})
	}
}

func TestGetMarketDataLength(t *testing.T) {
	data := []*MarketData{
		NewMarketData(100, 102, 105, 98, 10000, 1000000),
		NewMarketData(102, 103, 106, 101, 12000, 1200000),
		NewMarketData(103, 101, 104, 100, 11000, 1100000),
	}

	length := GetMarketDataLength(data)
	if length != 3 {
		t.Errorf("Expected length 3, got %d", length)
	}

	emptyData := []*MarketData{}
	emptyLength := GetMarketDataLength(emptyData)
	if emptyLength != 0 {
		t.Errorf("Expected length 0, got %d", emptyLength)
	}
}
