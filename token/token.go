package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
return tok
	}

	return IDENT
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
