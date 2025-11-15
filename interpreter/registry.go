package interpreter

import (
	"fmt"
	"strings"

	"github.com/DTrader-store/formula-go/errors"
	"github.com/DTrader-store/formula-go/types"
)

// Function represents a built-in function
type Function func(args []*Value, marketData []*types.MarketData) (*Value, error)

// FunctionRegistry manages built-in functions
type FunctionRegistry struct {
	functions map[string]Function
}

// NewFunctionRegistry creates a new function registry
func NewFunctionRegistry() *FunctionRegistry {
	reg := &FunctionRegistry{
		functions: make(map[string]Function),
	}
	reg.registerBuiltinFunctions()
	return reg
}

// Register registers a function
func (r *FunctionRegistry) Register(name string, fn Function) {
	r.functions[strings.ToUpper(name)] = fn
}

// Call calls a registered function
func (r *FunctionRegistry) Call(name string, args []*Value, marketData []*types.MarketData) (*Value, error) {
	fn, exists := r.functions[strings.ToUpper(name)]
	if !exists {
		return nil, errors.NewRuntimeError(fmt.Sprintf("undefined function: %s", name))
	}
	return fn(args, marketData)
}

// registerBuiltinFunctions registers all built-in functions
func (r *FunctionRegistry) registerBuiltinFunctions() {
	// Mathematical functions
	r.Register("MA", fnMA)
	r.Register("EMA", fnEMA)
	r.Register("SUM", fnSUM)
	r.Register("MAX", fnMAX)
	r.Register("MIN", fnMIN)
	r.Register("ABS", fnABS)
	r.Register("SQRT", fnSQRT)

	// Reference functions
	r.Register("REF", fnREF)
	r.Register("HHV", fnHHV)
	r.Register("LLV", fnLLV)

	// Conditional functions
	r.Register("IF", fnIF)
	r.Register("CROSS", fnCROSS)

	// Phase 4: Additional functions
	r.Register("STD", fnSTD)
	r.Register("VAR", fnVAR)
	r.Register("SMA", fnSMA)
	r.Register("WMA", fnWMA)
	r.Register("COUNT", fnCOUNT)
	r.Register("EVERY", fnEVERY)
	r.Register("EXIST", fnEXIST)
	r.Register("BARSLAST", fnBARSLAST)
	r.Register("AVEDEV", fnAVEDEV)
	r.Register("FILTER", fnFILTER)
	r.Register("BETWEEN", fnBETWEEN)
}
