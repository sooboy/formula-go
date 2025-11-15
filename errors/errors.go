// Package errors provides error types for the Formula-Go parser and interpreter
package errors

import "fmt"

// FormulaError is the base error type for all formula-related errors
type FormulaError struct {
	message string
}

func (e *FormulaError) Error() string {
	return e.message
}

// NewFormulaError creates a new FormulaError
func NewFormulaError(message string) *FormulaError {
	return &FormulaError{message: message}
}

// LexerError represents an error that occurred during lexical analysis (tokenization)
type LexerError struct {
	FormulaError
	Line   int
	Column int
	Char   string
}

// NewLexerError creates a new LexerError
func NewLexerError(message string, line, column int, char string) *LexerError {
	charPart := ""
	if char != "" {
		charPart = fmt.Sprintf(" (char: %s)", char)
	}
	fullMessage := fmt.Sprintf("Lexer error at line %d, column %d: %s%s", line, column, message, charPart)
	return &LexerError{
		FormulaError: FormulaError{message: fullMessage},
		Line:         line,
		Column:       column,
		Char:         char,
	}
}

// ParserError represents an error that occurred during parsing
type ParserError struct {
	FormulaError
	Line   int
	Column int
}

// NewParserError creates a new ParserError
func NewParserError(message string, line, column int) *ParserError {
	fullMessage := fmt.Sprintf("Parser error at line %d, column %d: %s", line, column, message)
	return &ParserError{
		FormulaError: FormulaError{message: fullMessage},
		Line:         line,
		Column:       column,
	}
}

// RuntimeError represents an error that occurred during runtime execution
type RuntimeError struct {
	FormulaError
}

// NewRuntimeError creates a new RuntimeError
func NewRuntimeError(message string) *RuntimeError {
	fullMessage := fmt.Sprintf("Runtime error: %s", message)
	return &RuntimeError{
		FormulaError: FormulaError{message: fullMessage},
	}
}
