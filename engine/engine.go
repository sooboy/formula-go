// Package engine provides the main FormulaEngine for parsing and executing formulas
package engine

import (
	"github.com/DTrader-store/formula-go/interpreter"
	"github.com/DTrader-store/formula-go/lexer"
	"github.com/DTrader-store/formula-go/parser"
	"github.com/DTrader-store/formula-go/parser/ast"
	"github.com/DTrader-store/formula-go/types"
)

// FormulaEngine is the main engine for compiling and executing formulas
type FormulaEngine struct {
	// You can add caching or other optimization features here
}

// NewFormulaEngine creates a new formula engine
func NewFormulaEngine() *FormulaEngine {
	return &FormulaEngine{}
}

// Compile compiles a formula string into an AST
func (e *FormulaEngine) Compile(formula string) (*ast.Program, error) {
	// Lexical analysis
	l := lexer.NewLexer(formula)
	tokens, err := l.Tokenize()
	if err != nil {
		return nil, err
	}

	// Syntax analysis
	p := parser.NewParser(tokens)
	program, err := p.Parse()
	if err != nil {
		return nil, err
	}

	return program, nil
}

// Execute executes a compiled program with market data
func (e *FormulaEngine) Execute(program *ast.Program, marketData []*types.MarketData) (*types.FormulaResult, error) {
	interp := interpreter.NewInterpreter(marketData)
	return interp.Execute(program)
}

// Run compiles and executes a formula in one step
func (e *FormulaEngine) Run(formula string, marketData []*types.MarketData) (*types.FormulaResult, error) {
	program, err := e.Compile(formula)
	if err != nil {
		return nil, err
	}

	return e.Execute(program, marketData)
}
