package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// token.go
const (
    DOT     TokenType = "DOT"
    // Special tokens
    EOF     TokenType = "EOF"
    ILLEGAL TokenType = "ILLEGAL"

    // Literals
    IDENTIFIER TokenType = "IDENTIFIER"
    NUMBER     TokenType = "NUMBER"
    STRING     TokenType = "STRING"
    BOOL       TokenType = "BOOL"
    NIL        TokenType = "NIL"

    // Operators
    PLUS      TokenType = "PLUS"
    MINUS     TokenType = "MINUS"
    MULTIPLY  TokenType = "MULTIPLY"
    DIVIDE    TokenType = "DIVIDE"
    MODULO    TokenType = "MODULO"
    EXPONENT  TokenType = "EXPONENT"
    LEN       TokenType = "LEN"
    EQ        TokenType = "EQ"
    GE        TokenType = "GE"
    LE        TokenType = "LE"
    NE        TokenType = "NE"
    GT        TokenType = "GT"
    LT        TokenType = "LT"
    CONCAT    TokenType = "CONCAT"
    VAARG     TokenType = "VAARG"

    // Booleans
    TRUE  TokenType = "TRUE"
    FALSE TokenType = "FALSE"

    // Delimiters
    COMMA     TokenType = "COMMA" // ,
    COLON     TokenType = "COLON" // : & op "call"
    SEMICOLON TokenType = "SEMICOLON" // ;
    LPAREN    TokenType = "LPAREN" // (
    RPAREN    TokenType = "RPAREN" // )
    LBRACE    TokenType = "LBRACE" // {
    RBRACE    TokenType = "RBRACE" // }
    LBRACKET  TokenType = "LBRACKET" // [
    RBRACKET  TokenType = "RBRACKET" // ]
	ASSIGN    TokenType = "ASSIGN" // =

    // Keywords
    FUNCTION TokenType = "FUNCTION"
    LOCAL    TokenType = "LOCAL"
    CONST    TokenType = "CONST"
    RETURN   TokenType = "RETURN"
    IF       TokenType = "IF"
    THEN     TokenType = "THEN"
    ELSE     TokenType = "ELSE"
    ELSEIF   TokenType = "ELSEIF"
    END      TokenType = "END"
    WHILE    TokenType = "WHILE"
    DO       TokenType = "DO"
    REPEAT   TokenType = "REPEAT"
    UNTIL    TokenType = "UNTIL"
    FOR      TokenType = "FOR"
    IN       TokenType = "IN"
    BREAK    TokenType = "BREAK"
    GOTO     TokenType = "GOTO"
    ASYNC    TokenType = "ASYNC"
    AWAIT    TokenType = "AWAIT"
    TRY      TokenType = "TRY"
    CATCH    TokenType = "CATCH"
    FINALLY  TokenType = "FINALLY"
    IMPORT   TokenType = "IMPORT"
    NOT      TokenType = "NOT"
    AND      TokenType = "AND"
    OR       TokenType = "OR"
    COROUTINE TokenType = "COROUTINE"
    //PRINT    TokenType = "PRINT"
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
	"async":    ASYNC,
	"await":    AWAIT,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"nil":      NIL,
	"true":     BOOL,
	"false":    BOOL,
	"not":      NOT,
	"or":       OR,
	"and":      AND,
	"import":   IMPORT,
    //"print":    PRINT,
    "coroutine":COROUTINE,
}

const (
	LOWEST  = 1
	EQUALS  = 2
	LESSGREATER = 3
	SUM     = 4
	PRODUCT = 5
	PREFIX  = 6
	CALL    = 7
)

var Precedences = map[TokenType]int{
    EQ:       EQUALS,
    NE:       EQUALS,
    LT:       LESSGREATER,
    GT:       LESSGREATER,
    LE:       LESSGREATER,
    GE:       LESSGREATER,
    PLUS:     SUM,
    MINUS:    SUM,
    DIVIDE:   PRODUCT,
    MULTIPLY: PRODUCT,
    MODULO:   PRODUCT,
    EXPONENT: PRODUCT,
    CONCAT:   CALL,
}
