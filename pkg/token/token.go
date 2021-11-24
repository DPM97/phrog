package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLGEGAL = "ILLGEGAL" // unknown token
	EOF      = "EOF"      // end of file

	// Identifiers + literals
	ID   = "ID"
	INT  = "INT"
	CHAR = "CHAR"

	// Ops
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	NOT      = "!"

	// Comps
	EQ  = "=="
	NEQ = "!="
	GT  = ">"
	LT  = "<"

	// Delims
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LCURLEY = "{"
	RCURLEY = "}"

	// Kws
	FUNCTION = "FUNCTION"
	ASSIGN   = "ASSIGN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var Kws = map[string]TokenType{
	"function": FUNCTION,
	"assign":   ASSIGN,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

func CreateToken(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}
