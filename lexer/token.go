package lexer

import "fmt"

// Token represents a single token in the formula lexer
type Token struct {
	Type   TokenType // The token type
	Value  string    // The token value
	Line   int       // The line number (1-indexed)
	Column int       // The column number (1-indexed)
}

// NewToken creates a new Token
func NewToken(tokenType TokenType, value string, line, column int) *Token {
	return &Token{
		Type:   tokenType,
		Value:  value,
		Line:   line,
		Column: column,
	}
}

// String returns a string representation of the token
func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %q, %d:%d)", t.Type, t.Value, t.Line, t.Column)
}

// Equals checks equality with another token
func (t *Token) Equals(other *Token) bool {
	if other == nil {
		return false
	}
	return t.Type == other.Type &&
		t.Value == other.Value &&
		t.Line == other.Line &&
		t.Column == other.Column
}
