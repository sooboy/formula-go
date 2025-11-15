package types

// LineStyle represents line style configuration for output visualization
type LineStyle struct {
	Color     string // Color of the line (e.g., '#FF0000', 'red')
	LineWidth int    // Width of the line in pixels
	LineStyle string // Style of the line ('solid', 'dashed', 'dotted', etc.)
}

// OutputLine represents a single output line representing calculated data
type OutputLine struct {
	Name  string      // Name/identifier of the output line
	Data  []float64   // Data points for the output line
	Style *LineStyle  // Optional style configuration for visualization
}

// FormulaResult represents the result of formula calculation containing outputs and variables
type FormulaResult struct {
	Outputs   []*OutputLine       // Array of output lines from the formula calculation
	Variables map[string]float64  // Calculated variables and their values
}

// NewFormulaResult creates a new FormulaResult instance
func NewFormulaResult() *FormulaResult {
	return &FormulaResult{
		Outputs:   make([]*OutputLine, 0),
		Variables: make(map[string]float64),
	}
}

// AddOutput adds an output line to the result
func (f *FormulaResult) AddOutput(name string, data []float64, style *LineStyle) {
	f.Outputs = append(f.Outputs, &OutputLine{
		Name:  name,
		Data:  data,
		Style: style,
	})
}

// SetVariable sets a variable value in the result
func (f *FormulaResult) SetVariable(name string, value float64) {
	f.Variables[name] = value
}

// GetVariable gets a variable value from the result
func (f *FormulaResult) GetVariable(name string) (float64, bool) {
	value, exists := f.Variables[name]
	return value, exists
}
