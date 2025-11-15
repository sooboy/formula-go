package interpreter

import (
	"fmt"
	"math"

	"github.com/DTrader-store/formula-go/errors"
	"github.com/DTrader-store/formula-go/types"
)

// fnMA implements Moving Average: MA(data, period)
func fnMA(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("MA requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("MA first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("MA second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("MA period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate MA
	for i := n - 1; i < len(data.Array); i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += data.Array[i-j]
		}
		result[i] = sum / float64(n)
	}

	return NewArrayValue(result), nil
}

// fnEMA implements Exponential Moving Average: EMA(data, period)
func fnEMA(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("EMA requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("EMA first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("EMA second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("EMA period must be between 1 and %d", len(data.Array)))
	}

	alpha := 2.0 / float64(n+1)
	result := make([]float64, len(data.Array))

	// First value is just the data point
	result[0] = data.Array[0]

	// Calculate EMA
	for i := 1; i < len(data.Array); i++ {
		result[i] = alpha*data.Array[i] + (1-alpha)*result[i-1]
	}

	return NewArrayValue(result), nil
}

// fnSUM implements Sum: SUM(data, period)
func fnSUM(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("SUM requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("SUM first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("SUM second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("SUM period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate SUM
	for i := n - 1; i < len(data.Array); i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += data.Array[i-j]
		}
		result[i] = sum
	}

	return NewArrayValue(result), nil
}

// fnMAX implements Max: MAX(a, b)
func fnMAX(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("MAX requires 2 arguments")
	}

	a, b := args[0], args[1]

	if !a.IsArray && !b.IsArray {
		return NewSingleValue(math.Max(a.Single, b.Single)), nil
	}

	if a.IsArray && b.IsArray {
		if len(a.Array) != len(b.Array) {
			return nil, errors.NewRuntimeError("MAX: array length mismatch")
		}
		result := make([]float64, len(a.Array))
		for i := range a.Array {
			result[i] = math.Max(a.Array[i], b.Array[i])
		}
		return NewArrayValue(result), nil
	}

	return nil, errors.NewRuntimeError("MAX: incompatible argument types")
}

// fnMIN implements Min: MIN(a, b)
func fnMIN(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("MIN requires 2 arguments")
	}

	a, b := args[0], args[1]

	if !a.IsArray && !b.IsArray {
		return NewSingleValue(math.Min(a.Single, b.Single)), nil
	}

	if a.IsArray && b.IsArray {
		if len(a.Array) != len(b.Array) {
			return nil, errors.NewRuntimeError("MIN: array length mismatch")
		}
		result := make([]float64, len(a.Array))
		for i := range a.Array {
			result[i] = math.Min(a.Array[i], b.Array[i])
		}
		return NewArrayValue(result), nil
	}

	return nil, errors.NewRuntimeError("MIN: incompatible argument types")
}

// fnABS implements Absolute value: ABS(value)
func fnABS(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 1 {
		return nil, errors.NewRuntimeError("ABS requires 1 argument")
	}

	val := args[0]

	if !val.IsArray {
		return NewSingleValue(math.Abs(val.Single)), nil
	}

	result := make([]float64, len(val.Array))
	for i, v := range val.Array {
		result[i] = math.Abs(v)
	}
	return NewArrayValue(result), nil
}

// fnSQRT implements Square root: SQRT(value)
func fnSQRT(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 1 {
		return nil, errors.NewRuntimeError("SQRT requires 1 argument")
	}

	val := args[0]

	if !val.IsArray {
		return NewSingleValue(math.Sqrt(val.Single)), nil
	}

	result := make([]float64, len(val.Array))
	for i, v := range val.Array {
		result[i] = math.Sqrt(v)
	}
	return NewArrayValue(result), nil
}

// fnREF implements Reference: REF(data, n) - reference data n periods ago
func fnREF(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("REF requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("REF first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("REF second argument must be a number")
	}

	n := int(period.Single)
	if n < 0 {
		return nil, errors.NewRuntimeError("REF period must be non-negative")
	}

	result := make([]float64, len(data.Array))

	// Fill first n values with NaN
	for i := 0; i < n && i < len(result); i++ {
		result[i] = math.NaN()
	}

	// Reference data
	for i := n; i < len(data.Array); i++ {
		result[i] = data.Array[i-n]
	}

	return NewArrayValue(result), nil
}

// fnHHV implements Highest High Value: HHV(data, period)
func fnHHV(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("HHV requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("HHV first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("HHV second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("HHV period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate HHV
	for i := n - 1; i < len(data.Array); i++ {
		maxVal := data.Array[i]
		for j := 1; j < n; j++ {
			if data.Array[i-j] > maxVal {
				maxVal = data.Array[i-j]
			}
		}
		result[i] = maxVal
	}

	return NewArrayValue(result), nil
}

// fnLLV implements Lowest Low Value: LLV(data, period)
func fnLLV(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("LLV requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("LLV first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("LLV second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("LLV period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate LLV
	for i := n - 1; i < len(data.Array); i++ {
		minVal := data.Array[i]
		for j := 1; j < n; j++ {
			if data.Array[i-j] < minVal {
				minVal = data.Array[i-j]
			}
		}
		result[i] = minVal
	}

	return NewArrayValue(result), nil
}

// fnIF implements conditional: IF(condition, trueValue, falseValue)
func fnIF(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 3 {
		return nil, errors.NewRuntimeError("IF requires 3 arguments")
	}

	cond := args[0]
	trueVal := args[1]
	falseVal := args[2]

	// Handle scalar condition
	if !cond.IsArray {
		if cond.Single != 0 {
			return trueVal, nil
		}
		return falseVal, nil
	}

	// Handle array condition
	if !trueVal.IsArray || !falseVal.IsArray {
		return nil, errors.NewRuntimeError("IF: when condition is array, both true/false values must be arrays")
	}

	if len(cond.Array) != len(trueVal.Array) || len(cond.Array) != len(falseVal.Array) {
		return nil, errors.NewRuntimeError("IF: array length mismatch")
	}

	result := make([]float64, len(cond.Array))
	for i := range cond.Array {
		if cond.Array[i] != 0 {
			result[i] = trueVal.Array[i]
		} else {
			result[i] = falseVal.Array[i]
		}
	}

	return NewArrayValue(result), nil
}

// fnCROSS implements cross detection: CROSS(a, b) - returns 1 when a crosses above b
func fnCROSS(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("CROSS requires 2 arguments")
	}

	a, b := args[0], args[1]

	if !a.IsArray || !b.IsArray {
		return nil, errors.NewRuntimeError("CROSS requires array arguments")
	}

	if len(a.Array) != len(b.Array) {
		return nil, errors.NewRuntimeError("CROSS: array length mismatch")
	}

	result := make([]float64, len(a.Array))
	result[0] = 0 // First element is always 0

	for i := 1; i < len(a.Array); i++ {
		if a.Array[i-1] <= b.Array[i-1] && a.Array[i] > b.Array[i] {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return NewArrayValue(result), nil
}
