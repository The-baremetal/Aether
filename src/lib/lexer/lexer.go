package lexer

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	DOT     TokenType = "."
	// Special tokens
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	// Literals
	IDENT   TokenType = "IDENT"
	NUMBER  TokenType = "NUMBER"
	STRING  TokenType = "STRING"
	BOOL    TokenType = "BOOL"
	NIL     TokenType = "NIL"

	// Operators
	ASSIGN         TokenType = "="
	PLUS           TokenType = "+"
	MINUS          TokenType = "-"
	ASTERISK       TokenType = "*"
	SLASH          TokenType = "/"
	FLOORDIV       TokenType = "//"
	MODULO         TokenType = "%"
	CARET          TokenType = "^"
	BIT_AND        TokenType = "&"
	BIT_OR         TokenType = "|"
	BIT_XOR        TokenType = "~"
	LSHIFT         TokenType = "<<"
	RSHIFT         TokenType = ">>"
	EQ             TokenType = "=="
	GE             TokenType = ">="
	LE             TokenType = "<="
	NE             TokenType = "~="
	GT             TokenType = ">"
	LT             TokenType = "<"
	CONCAT         TokenType = ".."
	NOT            TokenType = "not"
	NEG            TokenType = "-"
	TYPEOF         TokenType = "typeof"

	// Compound operators
	PLUS_ASSIGN        TokenType = "+="
	MINUS_ASSIGN       TokenType = "-="
	ASTERISK_ASSIGN    TokenType = "*="
	SLASH_ASSIGN       TokenType = "/="
	FLOORDIV_ASSIGN    TokenType = "//="
	CARET_ASSIGN       TokenType = "^="
	BIT_AND_ASSIGN     TokenType = "&="
	BIT_OR_ASSIGN      TokenType = "|="
	CONCAT_ASSIGN      TokenType = "..="
	COMPOUND_ASSIGN    TokenType = "??="
	INCR               TokenType = "+_"
	DECR               TokenType = "-_"
	SHIFT_ASSIGN_L     TokenType = "<<="
	SHIFT_ASSIGN_R     TokenType = ">>="
	COALESCE         TokenType = "??"
	NON_NULL_ASSERT  TokenType = "!!"

	// Delimiters
	COMMA     TokenType = ","
	COLON     TokenType = ":"
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	LBRACKET  TokenType = "["
	RBRACKET  TokenType = "]"

	// Keywords
	FUNCTION TokenType = "function"
	LOCAL    TokenType = "local"
	CONST    TokenType = "const"
	RETURN   TokenType = "return"
	IF       TokenType = "if"
	THEN     TokenType = "then"
	ELSE     TokenType = "else"
	ELSEIF   TokenType = "elseif"
	END      TokenType = "end"
	WHILE    TokenType = "while"
	DO       TokenType = "do"
	REPEAT   TokenType = "repeat"
	UNTIL    TokenType = "until"
	FOR      TokenType = "for"
	IN       TokenType = "in"
	BREAK    TokenType = "break"
	GOTO     TokenType = "goto"
	ASYNC    TokenType = "async"
	AWAIT    TokenType = "await"
	MATCH    TokenType = "match"
	WITH     TokenType = "with"
	SELECT   TokenType = "select"
	TRY      TokenType = "try"
	CATCH    TokenType = "catch"
	FINALLY  TokenType = "finally"
	TYPE     TokenType = "type"
	SPAWN    TokenType = "spawn"
	VOID     TokenType = "void"

	// Experimental / special operators
	SAFE_NAV TokenType = "?."
	PIPE     TokenType = "|>"
	ARROW    TokenType = "=>"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"local":    LOCAL,
	"const":    CONST,
	"return":   RETURN,
	"if":       IF,
	"then":     THEN,
	"else":     ELSE,
	"elseif":   ELSEIF,
	"end":      END,
	"while":    WHILE,
	"do":       DO,
	"repeat":   REPEAT,
	"until":    UNTIL,
	"for":      FOR,
	"in":       IN,
	"break":    BREAK,
	"goto":     GOTO,
	"async":    ASYNC,
	"await":    AWAIT,
	"match":    MATCH,
	"with":     WITH,
	"select":   SELECT,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"type":     TYPE,
	"spawn":    SPAWN,
	"void":     VOID,
	"nil":      NIL,
	"true":     BOOL,
	"false":    BOOL,
	"not":      NOT,
	"typeof":   TYPEOF,
}

type Token struct {
	Type    TokenType
	Literal string
}

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
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(PLUS_ASSIGN)
		} else if l.peekChar() == '_' {
			tok = l.newDoubleToken(INCR)
		} else {
			tok = l.newSingleToken(PLUS)
		}
	case '-':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(MINUS_ASSIGN)
		} else if l.peekChar() == '_' {
			tok = l.newDoubleToken(DECR)
		} else {
			tok = l.newSingleToken(MINUS)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(ASTERISK_ASSIGN)
		} else {
			tok = l.newSingleToken(ASTERISK)
		}
	case '/':
		if l.peekChar() == '/' {
			if l.peekAhead(2) == '=' {
				tok = l.newMultiCharToken(SLASH_ASSIGN, 3)
			} else {
				tok = l.newMultiCharToken(FLOORDIV, 2)
			}
		} else if l.peekChar() == '=' {
			tok = l.newDoubleToken(SLASH_ASSIGN)
		} else {
			tok = l.newSingleToken(SLASH)
		}
	case '%':
		tok = l.newSingleToken(MODULO)
	case '^':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(CARET_ASSIGN)
		} else {
			tok = l.newSingleToken(CARET)
		}
	case '&':
		if l.peekChar() == '=' {
			// similar handling could be added here
			tok = l.newDoubleToken(BIT_AND_ASSIGN)
		} else {
			tok = l.newSingleToken(BIT_AND)
		}
	case '|':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(BIT_OR_ASSIGN)
		} else if l.peekChar() == '>' {
			tok = l.newDoubleToken(PIPE)
		} else {
			tok = l.newSingleToken(BIT_OR)
		}
	case '~':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(NE)
		} else {
			tok = l.newSingleToken(BIT_XOR)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(LE)
		} else if l.peekChar() == '<' {
			if l.peekAhead(2) == '=' {
				tok = l.newMultiCharToken(SHIFT_ASSIGN_L, 3)
			} else {
				tok = l.newMultiCharToken(LSHIFT, 2)
			}
		} else {
			tok = l.newSingleToken(LT)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = l.newDoubleToken(GE)
		} else if l.peekChar() == '>' {
			if l.peekAhead(2) == '=' {
				tok = l.newMultiCharToken(SHIFT_ASSIGN_R, 3)
			} else {
				tok = l.newMultiCharToken(RSHIFT, 2)
			}
		} else {
			tok = l.newSingleToken(GT)
		}
	case '.':
		if l.peekChar() == '.' {
			tok = l.newDoubleToken(CONCAT)
		} else {
			tok = l.newSingleToken(DOT)
		}
	case '?':
		if l.peekChar() == '.' {
			tok = l.newDoubleToken(SAFE_NAV)
		} else if l.peekChar() == '[' {
			tok = l.newMultiCharToken(SAFE_NAV, 2) // reuse token type for safe access with '['
		} else if l.peekChar() == '?' {
			tok = l.newDoubleToken(COALESCE)
		} else {
			tok = l.newSingleToken(ILLEGAL)
		}
	case '!':
		if l.peekChar() == '!' {
			tok = l.newDoubleToken(NON_NULL_ASSERT)
		} else {
			tok = l.newSingleToken(ILLEGAL)
		}
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
	return Token{Type: IDENT, Literal: lit}
}

func (l *Lexer) readNumber() Token {
	start := l.pos
	for isDigit(l.ch) {
		l.readChar()
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
