package types

import "testing"

func TestNewFormulaResult(t *testing.T) {
	result := NewFormulaResult()
	if result == nil {
		t.Error("Expected non-nil result")
	}
	if result.Outputs == nil {
		t.Error("Expected non-nil outputs")
	}
	if result.Variables == nil {
		t.Error("Expected non-nil variables")
	}
	if len(result.Outputs) != 0 {
		t.Errorf("Expected empty outputs, got %d items", len(result.Outputs))
	}
	if len(result.Variables) != 0 {
		t.Errorf("Expected empty variables, got %d items", len(result.Variables))
	}
}

func TestAddOutput(t *testing.T) {
	result := NewFormulaResult()
	data := []float64{1.0, 2.0, 3.0}
	style := &LineStyle{
		Color:     "red",
		LineWidth: 2,
		LineStyle: "solid",
	}

	result.AddOutput("MA5", data, style)

	if len(result.Outputs) != 1 {
		t.Errorf("Expected 1 output, got %d", len(result.Outputs))
	}

	output := result.Outputs[0]
	if output.Name != "MA5" {
		t.Errorf("Expected name 'MA5', got '%s'", output.Name)
	}
	if len(output.Data) != 3 {
		t.Errorf("Expected data length 3, got %d", len(output.Data))
	}
	if output.Style.Color != "red" {
		t.Errorf("Expected color 'red', got '%s'", output.Style.Color)
	}
}

func TestSetVariable(t *testing.T) {
	result := NewFormulaResult()
	result.SetVariable("x", 10.5)
	result.SetVariable("y", 20.3)

	if len(result.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(result.Variables))
	}

	if result.Variables["x"] != 10.5 {
		t.Errorf("Expected x=10.5, got %f", result.Variables["x"])
	}
	if result.Variables["y"] != 20.3 {
		t.Errorf("Expected y=20.3, got %f", result.Variables["y"])
	}
}

func TestGetVariable(t *testing.T) {
	result := NewFormulaResult()
	result.SetVariable("x", 15.7)

	value, exists := result.GetVariable("x")
	if !exists {
		t.Error("Expected variable 'x' to exist")
	}
	if value != 15.7 {
		t.Errorf("Expected value 15.7, got %f", value)
	}

	_, exists = result.GetVariable("nonexistent")
	if exists {
		t.Error("Expected variable 'nonexistent' to not exist")
	}
}

func TestMultipleOutputs(t *testing.T) {
	result := NewFormulaResult()

	result.AddOutput("MA5", []float64{1, 2, 3}, nil)
	result.AddOutput("MA10", []float64{4, 5, 6}, nil)
	result.AddOutput("MA20", []float64{7, 8, 9}, nil)

	if len(result.Outputs) != 3 {
		t.Errorf("Expected 3 outputs, got %d", len(result.Outputs))
	}

	names := []string{"MA5", "MA10", "MA20"}
	for i, output := range result.Outputs {
		if output.Name != names[i] {
			t.Errorf("Expected output name '%s', got '%s'", names[i], output.Name)
		}
	}
}
