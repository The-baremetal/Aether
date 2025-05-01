// this code is pretty self explainatory.

package lexer

import (
	"strings"
	"unicode"
)
type Lexer struct {
	input        string
	pos, readPos int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
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

	// ADDED
	if l.handleImportKeyword() {
		return tok
	}

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
	case '<':
		tok = l.newSingleToken(LT)
	case '>':
		tok = l.newSingleToken(GT)
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch)}
		l.readChar()
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch)}
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
	tok := Token{Type: t, Literal: string(l.ch)}
	l.readChar()
	return tok
}

func (l *Lexer) newDoubleToken(t TokenType) Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	l.readChar()
	return Token{Type: t, Literal: literal}
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
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	lit := l.input[start:l.pos]
	if typ, ok := keywords[strings.ToLower(lit)]; ok {
		return Token{Type: typ, Literal: lit}
	}
	return Token{Type: IDENTIFIER, Literal: lit}
}

func (l *Lexer) readNumber() Token {
    start := l.pos
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
    return Token{Type: NUMBER, Literal: l.input[start:l.pos]}
}

func (l *Lexer) readString(delim byte) Token { // rich strings
	start := l.pos + 1
	_ = start
	l.readChar()

	var strBuilder strings.Builder

	for l.ch != delim && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				strBuilder.WriteByte('\n')
			case 't':
				strBuilder.WriteByte('\t')
			case 'r':
				strBuilder.WriteByte('\r')
			case '\\':
				strBuilder.WriteByte('\\')
			case '"', '\'':
				strBuilder.WriteByte(l.ch)
			default:
				strBuilder.WriteByte(l.ch)
			}
		} else {
			strBuilder.WriteByte(l.ch)
		}
		l.readChar()
	}
	if l.ch == 0 {
		return Token{Type: ILLEGAL, Literal: "unfinished string"}
	}
	l.readChar()

	return Token{Type: STRING, Literal: strBuilder.String()}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) handleImportKeyword() bool {
	if strings.HasPrefix(l.input[l.pos:], "import") {
		if l.pos+len("import") < len(l.input) {
			nextChar := l.input[l.pos+len("import")]
			if nextChar == ' ' || nextChar == '\n' || nextChar == '\t' || nextChar == '(' {
				l.pos += len("import")
				l.readChar()
				l.skipWhitespace()
				return true
			}
		} else {
			l.pos += len("import")
			l.readChar()
			l.skipWhitespace()
			return true
		}
	}
	return false
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}
	tokens = append(tokens, Token{Type: EOF, Literal: ""})
	return tokens
}

func isHexDigit(ch byte) bool {
    return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}
