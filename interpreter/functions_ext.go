package interpreter

import (
	"fmt"
	"math"

	"github.com/DTrader-store/formula-go/errors"
	"github.com/DTrader-store/formula-go/types"
)

// Additional built-in functions for Phase 4

// fnSTD implements Standard Deviation: STD(data, period)
func fnSTD(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("STD requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("STD first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("STD second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("STD period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate STD
	for i := n - 1; i < len(data.Array); i++ {
		// Calculate mean
		mean := 0.0
		for j := 0; j < n; j++ {
			mean += data.Array[i-j]
		}
		mean /= float64(n)

		// Calculate variance
		variance := 0.0
		for j := 0; j < n; j++ {
			diff := data.Array[i-j] - mean
			variance += diff * diff
		}
		variance /= float64(n)

		result[i] = math.Sqrt(variance)
	}

	return NewArrayValue(result), nil
}

// fnVAR implements Variance: VAR(data, period)
func fnVAR(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("VAR requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("VAR first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("VAR second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("VAR period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate VAR
	for i := n - 1; i < len(data.Array); i++ {
		// Calculate mean
		mean := 0.0
		for j := 0; j < n; j++ {
			mean += data.Array[i-j]
		}
		mean /= float64(n)

		// Calculate variance
		variance := 0.0
		for j := 0; j < n; j++ {
			diff := data.Array[i-j] - mean
			variance += diff * diff
		}
		result[i] = variance / float64(n)
	}

	return NewArrayValue(result), nil
}

// fnSMA implements Simple Moving Average (alias for MA): SMA(data, period)
func fnSMA(args []*Value, data []*types.MarketData) (*Value, error) {
	return fnMA(args, data)
}

// fnWMA implements Weighted Moving Average: WMA(data, period)
func fnWMA(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("WMA requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("WMA first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("WMA second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("WMA period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Calculate weight sum
	weightSum := float64(n * (n + 1) / 2)

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate WMA
	for i := n - 1; i < len(data.Array); i++ {
		weightedSum := 0.0
		for j := 0; j < n; j++ {
			weight := float64(n - j)
			weightedSum += data.Array[i-j] * weight
		}
		result[i] = weightedSum / weightSum
	}

	return NewArrayValue(result), nil
}

// fnCOUNT implements Count: COUNT(condition, period)
func fnCOUNT(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("COUNT requires 2 arguments")
	}

	condition := args[0]
	period := args[1]

	if !condition.IsArray {
		return nil, errors.NewRuntimeError("COUNT first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("COUNT second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(condition.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("COUNT period must be between 1 and %d", len(condition.Array)))
	}

	result := make([]float64, len(condition.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Count true conditions
	for i := n - 1; i < len(condition.Array); i++ {
		count := 0.0
		for j := 0; j < n; j++ {
			if condition.Array[i-j] != 0 {
				count++
			}
		}
		result[i] = count
	}

	return NewArrayValue(result), nil
}

// fnEVERY implements Every: EVERY(condition, period) - returns 1 if condition is true for all periods
func fnEVERY(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("EVERY requires 2 arguments")
	}

	condition := args[0]
	period := args[1]

	if !condition.IsArray {
		return nil, errors.NewRuntimeError("EVERY first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("EVERY second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(condition.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("EVERY period must be between 1 and %d", len(condition.Array)))
	}

	result := make([]float64, len(condition.Array))

	// Fill first n-1 values with 0
	for i := 0; i < n-1; i++ {
		result[i] = 0
	}

	// Check if every condition is true
	for i := n - 1; i < len(condition.Array); i++ {
		everyCond := true
		for j := 0; j < n; j++ {
			if condition.Array[i-j] == 0 {
				everyCond = false
				break
			}
		}
		if everyCond {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return NewArrayValue(result), nil
}

// fnEXIST implements Exist: EXIST(condition, period) - returns 1 if condition is true for any period
func fnEXIST(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("EXIST requires 2 arguments")
	}

	condition := args[0]
	period := args[1]

	if !condition.IsArray {
		return nil, errors.NewRuntimeError("EXIST first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("EXIST second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(condition.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("EXIST period must be between 1 and %d", len(condition.Array)))
	}

	result := make([]float64, len(condition.Array))

	// Fill first n-1 values with 0
	for i := 0; i < n-1; i++ {
		result[i] = 0
	}

	// Check if any condition is true
	for i := n - 1; i < len(condition.Array); i++ {
		exists := false
		for j := 0; j < n; j++ {
			if condition.Array[i-j] != 0 {
				exists = true
				break
			}
		}
		if exists {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return NewArrayValue(result), nil
}

// fnBARSLAST implements BarsLast: BARSLAST(condition) - returns bars since last true condition
func fnBARSLAST(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 1 {
		return nil, errors.NewRuntimeError("BARSLAST requires 1 argument")
	}

	condition := args[0]

	if !condition.IsArray {
		return nil, errors.NewRuntimeError("BARSLAST argument must be an array")
	}

	result := make([]float64, len(condition.Array))
	lastTrueIndex := -1

	for i := 0; i < len(condition.Array); i++ {
		if condition.Array[i] != 0 {
			lastTrueIndex = i
			result[i] = 0
		} else if lastTrueIndex >= 0 {
			result[i] = float64(i - lastTrueIndex)
		} else {
			result[i] = math.NaN()
		}
	}

	return NewArrayValue(result), nil
}

// fnAVEDEV implements Average Deviation: AVEDEV(data, period)
func fnAVEDEV(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("AVEDEV requires 2 arguments")
	}

	data := args[0]
	period := args[1]

	if !data.IsArray {
		return nil, errors.NewRuntimeError("AVEDEV first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("AVEDEV second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 || n > len(data.Array) {
		return nil, errors.NewRuntimeError(fmt.Sprintf("AVEDEV period must be between 1 and %d", len(data.Array)))
	}

	result := make([]float64, len(data.Array))

	// Fill first n-1 values with NaN
	for i := 0; i < n-1; i++ {
		result[i] = math.NaN()
	}

	// Calculate AVEDEV
	for i := n - 1; i < len(data.Array); i++ {
		// Calculate mean
		mean := 0.0
		for j := 0; j < n; j++ {
			mean += data.Array[i-j]
		}
		mean /= float64(n)

		// Calculate average deviation
		devSum := 0.0
		for j := 0; j < n; j++ {
			devSum += math.Abs(data.Array[i-j] - mean)
		}
		result[i] = devSum / float64(n)
	}

	return NewArrayValue(result), nil
}

// fnFILTER implements Filter: FILTER(condition, period) - filters signals
func fnFILTER(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 2 {
		return nil, errors.NewRuntimeError("FILTER requires 2 arguments")
	}

	condition := args[0]
	period := args[1]

	if !condition.IsArray {
		return nil, errors.NewRuntimeError("FILTER first argument must be an array")
	}
	if period.IsArray {
		return nil, errors.NewRuntimeError("FILTER second argument must be a number")
	}

	n := int(period.Single)
	if n <= 0 {
		return nil, errors.NewRuntimeError("FILTER period must be positive")
	}

	result := make([]float64, len(condition.Array))
	lastSignal := -n - 1 // Initialize to allow first signal

	for i := 0; i < len(condition.Array); i++ {
		if condition.Array[i] != 0 && (i-lastSignal) >= n {
			result[i] = 1
			lastSignal = i
		} else {
			result[i] = 0
		}
	}

	return NewArrayValue(result), nil
}

// fnBETWEEN implements Between: BETWEEN(value, lower, upper)
func fnBETWEEN(args []*Value, _ []*types.MarketData) (*Value, error) {
	if len(args) != 3 {
		return nil, errors.NewRuntimeError("BETWEEN requires 3 arguments")
	}

	value := args[0]
	lower := args[1]
	upper := args[2]

	// Handle scalar case
	if !value.IsArray && !lower.IsArray && !upper.IsArray {
		if value.Single >= lower.Single && value.Single <= upper.Single {
			return NewSingleValue(1), nil
		}
		return NewSingleValue(0), nil
	}

	// Handle array case
	if !value.IsArray {
		return nil, errors.NewRuntimeError("BETWEEN: value must be array when using array bounds")
	}

	result := make([]float64, len(value.Array))
	for i := range value.Array {
		lowerBound := lower.Single
		upperBound := upper.Single
		if lower.IsArray {
			lowerBound = lower.Array[i]
		}
		if upper.IsArray {
			upperBound = upper.Array[i]
		}

		if value.Array[i] >= lowerBound && value.Array[i] <= upperBound {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return NewArrayValue(result), nil
}
