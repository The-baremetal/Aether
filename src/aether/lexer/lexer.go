package lexer

// The lexer of AetherC

/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"strings"
	"unicode"
)
type Lexer struct {
	input        string
	pos, readPos int
	ch           byte
	line         int
	column       int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}

	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	var tok Token
	switch l.ch {
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	case '=':
		if l.peekChar() == '=' {
		
			tok = l.newDoubleToken(EQ)
		} else {
			tok = l.newSingleToken(ASSIGN)
		}
	case '+':
		tok = l.newSingleToken(PLUS)
	case '-':
		tok = l.newSingleToken(MINUS)
	case '*':
		tok = l.newSingleToken(MULTIPLY)
	case '/':
		tok = l.newSingleToken(DIVIDE)
	case '%':
		tok = l.newSingleToken(MODULO)
	case '^':
		tok = l.newSingleToken(EXPONENT)
	case '>':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(GE)
		} else {
			tok = l.newSingleToken(GT)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(LE)
		} else {
			tok = l.newSingleToken(LT)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(NE)
		} else {
			tok = Token{Type: ILLEGAL, Literal: "!", Line: l.line, Column: l.column}
		}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch)}
		l.readChar()
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch)}
		l.readChar()
	case '{':
		tok = Token{Type: LBRACE, Literal: string(l.ch)}
		l.readChar()
	case '}':
		tok = Token{Type: RBRACE, Literal: string(l.ch)}
		l.readChar()
	case '[':
		tok = Token{Type: LBRACKET, Literal: string(l.ch)}
		l.readChar()
	case ']':
		tok = Token{Type: RBRACKET, Literal: string(l.ch)}
		l.readChar()
	case '.':
		if l.peekChar() == '.' {
			if l.peekAhead(2) == '.' {
				tok = l.newMultiCharToken(VAARG, 3)
			} else {
				tok = l.newDoubleToken(CONCAT)
			}
		} else {
			tok = l.newSingleToken(DOT)
		}
	case ':':
		tok = l.newSingleToken(COLON)
	case ',':
		tok = l.newSingleToken(COMMA)
	case '"', '\'':
		tok = l.readString(l.ch)
	default:
		if isLetter(l.ch) || l.ch == '_' {
			return l.readIdentifier()
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch)}
			l.readChar()
		}
	}
	return tok
}

func (l *Lexer) newSingleToken(t TokenType) Token {
	tok := Token{Type: t, Literal: string(l.ch), Line: l.line, Column: l.column}
	l.readChar()
	return tok
}

func (l *Lexer) newDoubleToken(t TokenType) Token {
	line := l.line
	column := l.column
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	l.readChar()
	return Token{Type: t, Literal: literal, Line: line, Column: column}
}

func (l *Lexer) newMultiCharToken(t TokenType, count int) Token {
	start := l.pos
	for i := 0; i < count; i++ {
		l.readChar()
	}
	return Token{Type: t, Literal: l.input[start:l.pos]}
}

func (l *Lexer) peekAhead(n int) byte {
	if l.pos+n >= len(l.input) {
		return 0
	}
	return l.input[l.pos+n]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		} else {
			l.column++
		}
		l.readChar()
	}

	if l.ch == '-' && l.peekChar() == '-' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
		l.skipWhitespace()
	}
}

func (l *Lexer) readIdentifier() Token {
	start := l.pos
	startColumn := l.column
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	lit := l.input[start:l.pos]
	if typ, ok := keywords[strings.ToLower(lit)]; ok {
		return Token{Type: typ, Literal: lit, Line: l.line, Column: startColumn}
	}
	return Token{Type: IDENTIFIER, Literal: lit, Line: l.line, Column: startColumn}
}

func (l *Lexer) readNumber() Token {
	start := l.pos
	startColumn := l.column
	if l.ch == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X') {
		l.readChar()
		l.readChar()
		for isHexDigit(l.ch) {
			l.readChar()
		}
	} else if isDigit(l.ch) {
		for isDigit(l.ch) {
			l.readChar()
		}
		if l.ch == '.' && isDigit(l.peekChar()) {
			l.readChar()
			for isDigit(l.ch) {
				l.readChar()
			}
		}
	}
	return Token{Type: NUMBER, Literal: l.input[start:l.pos], Line: l.line, Column: startColumn}
}

func (l *Lexer) readString(delim byte) Token {
	startColumn := l.column
	l.readChar()
	var strBuilder strings.Builder
	for l.ch != delim && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				strBuilder.WriteByte('\n')
				l.line++
				l.column = 0
			case 't':
				strBuilder.WriteByte('\t')
				l.column++
			case 'r':
				strBuilder.WriteByte('\r')
				l.column++
			case '\\':
				strBuilder.WriteByte('\\')
				l.column++
			case '"', '\'':
				strBuilder.WriteByte(l.ch)
				l.column++
			default:
				strBuilder.WriteByte(l.ch)
				l.column++
			}
		} else {
			strBuilder.WriteByte(l.ch)
			l.column++
		}
		l.readChar()
	}
	if l.ch == 0 {
		return Token{Type: ILLEGAL, Literal: "unfinished string", Line: l.line, Column: startColumn}
	}
	l.readChar()
	return Token{Type: STRING, Literal: strBuilder.String(), Line: l.line, Column: startColumn}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}
	tokens = append(tokens, Token{Type: EOF, Literal: "", Line: l.line, Column: l.column})
	return tokens
}

func isHexDigit(ch byte) bool {
    return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}