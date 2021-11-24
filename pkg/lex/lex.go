package lex

import (
	"bytes"
	"log"
	"phrog/pkg/token"
)

type Lexer struct {
	inputStr     string
	position     int
	readPosition int
	curChar      byte
}

func New(input string) *Lexer {
	l := &Lexer{inputStr: input}
	// read the first char in so we don't get EOF
	l.readNextChar()
	return l
}

func throwInvalidInput() {
	log.Fatalf("could not parse input - invalid token.")
}

func (l *Lexer) readNextChar() {
	// if the next char in the input is > EOF, then set to 0
	if l.readPosition >= len(l.inputStr) {
		l.curChar = 0
	} else {
		// cur char is the next char
		l.curChar = l.inputStr[l.readPosition]
	}

	// position will point to the next char
	l.position = l.readPosition

	// increment readPosition
	l.readPosition += 1
}

/*

	ILLGEGAL = "ILLGEGAL" // unknown token
	EOF      = "EOF"      // end of file

	// Identifiers + literals
	ID   = "ID"
	INT  = "INT"
	CHAR = "CHAR"

	// Ops
	EQUAL    = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	GT       = ">"
	LT       = "<"
	NOT      = "!"

	// Delims
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LCURLEY = "{"
	RCURLEY = "}"

	// Kws
	RIBBIT = "RIBBIT" // function creation
	LEAP   = "LEAP"   // function call
	ASSIGN = "ASSIGN" // create variable

*/

// A-Z, a-z, _
func isAlpha(c byte) bool {
	return (c >= 97 && c <= 122) || (c >= 65 && c <= 90) || (c == 95)
}

func isNum(c byte) bool {
	return c >= 48 && c <= 57
}

func (l *Lexer) peekNext() byte {
	if l.readPosition >= len(l.inputStr) {
		return 0
	}
	return l.inputStr[l.readPosition]
}

// take in t s.t. we can do some error handling
func (l *Lexer) readNextBlock(t string) string {
	buff := bytes.NewBufferString("")

	// read until the next space, build buffer
	for (t == "KWID" && isAlpha(l.curChar)) ||
		(t == "NUMBER" && isNum(l.curChar)) {
		buff.WriteByte(l.curChar)
		l.readNextChar()
	}

	return buff.String()
}

func (l *Lexer) ParseKWID() token.Token {
	block := l.readNextBlock("KWID")

	var resultToken token.Token

	if tokenType, containsKw := token.Kws[block]; containsKw {
		resultToken.Type = tokenType
	} else {
		resultToken.Type = token.ID
	}

	resultToken.Literal = block

	return resultToken
}

func (l *Lexer) ParseNum() token.Token {
	block := l.readNextBlock("NUMBER")
	return token.Token{Type: token.INT, Literal: block}
}

func (l *Lexer) FetchNextToken() token.Token {
	var curToken token.Token
	// parse keywords and identifiers

	// skip whitespace
	for l.curChar == ' ' || l.curChar == '\t' || l.curChar == '\n' || l.curChar == '\r' {
		l.readNextChar()
	}

	if isAlpha(l.curChar) {
		curToken = l.ParseKWID()
	} else if isNum(l.curChar) {
		curToken = l.ParseNum()
	} else {
		switch l.curChar {
		case '=':
			curToken = token.CreateToken(token.EQUAL, l.curChar)
		case '+':
			curToken = token.CreateToken(token.PLUS, l.curChar)
		case '-':
			curToken = token.CreateToken(token.MINUS, l.curChar)
		case '*':
			curToken = token.CreateToken(token.MULTIPLY, l.curChar)
		case '/':
			curToken = token.CreateToken(token.DIVIDE, l.curChar)
		case '>':
			curToken = token.CreateToken(token.GT, l.curChar)
		case '<':
			curToken = token.CreateToken(token.LT, l.curChar)
		case '!':
			curToken = token.CreateToken(token.NOT, l.curChar)
		case ',':
			curToken = token.CreateToken(token.COMMA, l.curChar)
		case ';':
			curToken = token.CreateToken(token.SEMICOLON, l.curChar)
		case '(':
			curToken = token.CreateToken(token.LPAREN, l.curChar)
		case ')':
			curToken = token.CreateToken(token.RPAREN, l.curChar)
		case '{':
			curToken = token.CreateToken(token.LCURLEY, l.curChar)
		case '}':
			curToken = token.CreateToken(token.RCURLEY, l.curChar)
		case 0:
			curToken.Literal = ""
			curToken.Type = token.EOF
		}

		l.readNextChar()
	}

	return curToken
}
