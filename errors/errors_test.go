package errors

import (
	"strings"
	"testing"
)

func TestNewFormulaError(t *testing.T) {
	err := NewFormulaError("test error")
	if err.Error() != "test error" {
		t.Errorf("Expected 'test error', got '%s'", err.Error())
	}
}

func TestNewLexerError(t *testing.T) {
	tests := []struct {
		name    string
		message string
		line    int
		column  int
		char    string
		expect  string
	}{
		{
			name:    "with char",
			message: "unexpected character",
			line:    1,
			column:  5,
			char:    "@",
			expect:  "Lexer error at line 1, column 5: unexpected character (char: @)",
		},
		{
			name:    "without char",
			message: "unexpected end of file",
			line:    2,
			column:  10,
			char:    "",
			expect:  "Lexer error at line 2, column 10: unexpected end of file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewLexerError(tt.message, tt.line, tt.column, tt.char)
			if err.Error() != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, err.Error())
			}
			if err.Line != tt.line {
				t.Errorf("Expected line %d, got %d", tt.line, err.Line)
			}
			if err.Column != tt.column {
				t.Errorf("Expected column %d, got %d", tt.column, err.Column)
			}
			if err.Char != tt.char {
				t.Errorf("Expected char '%s', got '%s'", tt.char, err.Char)
			}
		})
	}
}

func TestNewParserError(t *testing.T) {
	err := NewParserError("unexpected token", 3, 15)
	expected := "Parser error at line 3, column 15: unexpected token"
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
	if err.Line != 3 {
		t.Errorf("Expected line 3, got %d", err.Line)
	}
	if err.Column != 15 {
		t.Errorf("Expected column 15, got %d", err.Column)
	}
}

func TestNewRuntimeError(t *testing.T) {
	err := NewRuntimeError("division by zero")
	expected := "Runtime error: division by zero"
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestErrorTypes(t *testing.T) {
	// Test that all error types implement the error interface
	var _ error = &FormulaError{}
	var _ error = &LexerError{}
	var _ error = &ParserError{}
	var _ error = &RuntimeError{}
}

func TestErrorMessageFormat(t *testing.T) {
	err := NewLexerError("test", 1, 1, "x")
	if !strings.Contains(err.Error(), "Lexer error") {
		t.Error("Error message should contain 'Lexer error'")
	}
	if !strings.Contains(err.Error(), "line 1") {
		t.Error("Error message should contain line number")
	}
	if !strings.Contains(err.Error(), "column 1") {
		t.Error("Error message should contain column number")
	}
}
