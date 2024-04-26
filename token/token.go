package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGEAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, x, y
	INT = "INT" //Dokumente und Einstellungen

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimeters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
