// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type lua_grammar_antlr4Parser struct {
	*antlr.BaseParser
}

var Lua_grammar_antlr4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lua_grammar_antlr4ParserInit() {
	staticData := &Lua_grammar_antlr4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'^'", "'('", "')'", "'#'", "'-'", "'~'", "'>'", "'<'", "'>='",
		"'=='", "'<='", "'~='", "'*'", "'/'", "'+'", "'//'", "'&'", "'|'", "'<<'",
		"'>>'", "'>>>'", "'..'", "':'", "'table.insert'", "','", "'.'", "'{'",
		"'}'", "';'", "'['", "']'", "'_'", "'::'", "'__'", "'__add'", "'__sub'",
		"'__mul'", "'__div'", "'__mod'", "'__pow'", "'__unm'", "'__concat'",
		"'__len'", "'__eq'", "'__lt'", "'__le'", "'__tostring'", "'__pairs'",
		"'__ipairs'", "'__call'", "", "'in'", "'print'", "'and'", "'break'",
		"'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'", "'goto'",
		"'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'",
		"'true'", "'until'", "'while'", "'local'", "'function'", "'index'",
		"'newindex'", "'mode'", "'pcall'", "'xpcall'", "'coroutine'", "'create'",
		"'resume'", "'yield'", "'status'", "'nan'", "'inf'", "'error'", "'assert'",
		"'pairs'", "'ipairs'", "", "", "", "", "", "", "", "'...'", "'++'",
		"'--'", "'+='", "'-='", "'*='", "'/='", "'//='", "'^='", "'..='", "'??'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "SEPARATOR", "KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO",
		"KW_ELSE", "KW_ELSEIF", "KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF",
		"KW_NIL", "KW_NOT", "KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE",
		"KW_UNTIL", "KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "KW_INDEX", "KW_NEWINDEX",
		"KW_MODE", "KW_PCALL", "KW_XPCALL", "KW_COROUTINE", "KW_CREATE", "KW_RESUME",
		"KW_YIELD", "KW_STATUS", "KW_NAN", "KW_INF", "KW_ERROR", "KW_ASSERT",
		"KW_PAIRS", "KW_IPAIRS", "NUMBER", "STRING", "LETTER", "DIGIT", "WS",
		"SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT", "VARARG", "INCREMENT",
		"DECREMENT", "PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN",
		"FLOOR_DIV_ASSIGN", "EXP_ASSIGN", "CONCAT_ASSIGN", "NULL_COALESCE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "control_statement", "statement_terminator",
		"assignment", "expression", "prefix_expression", "primary_expression",
		"operators", "comparison_operator", "arith_operator", "logical_operator",
		"bitwise_operator", "concat_operator", "literal", "function_call", "table_insert",
		"function_declaration", "parameter", "block", "if_statement", "for_statement",
		"while_statement", "do_statement", "table", "field_separator", "field",
		"table_access", "single_line_comment", "print_statement", "identifier",
		"repeat_statement", "identifier_list", "expression_list", "return_statement",
		"break_statement", "goto_statement", "label_statement", "function_expression",
		"method_call", "metatable_field", "metamethod", "coroutine_statement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 109, 511, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 101, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 4, 3, 4, 114, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 137, 8, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 5, 5, 153, 8, 5, 10, 5, 12, 5, 156, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 171, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 183,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 189, 8, 7, 10, 7, 12, 7, 192, 9, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 199, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 229, 8, 15, 1, 15, 1, 15, 3, 15, 233, 8, 15,
		1, 15, 1, 15, 3, 15, 237, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 248, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 3, 17, 258, 8, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 265, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17,
		272, 8, 17, 10, 17, 12, 17, 275, 9, 17, 1, 17, 3, 17, 278, 8, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 287, 8, 18, 1, 19, 5,
		19, 290, 8, 19, 10, 19, 12, 19, 293, 9, 19, 1, 19, 1, 19, 3, 19, 297, 8,
		19, 4, 19, 299, 8, 19, 11, 19, 12, 19, 300, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 312, 8, 20, 10, 20, 12, 20, 315,
		9, 20, 1, 20, 1, 20, 3, 20, 319, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 339, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 345, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 362, 8, 24, 10,
		24, 12, 24, 365, 9, 24, 1, 24, 3, 24, 368, 8, 24, 3, 24, 370, 8, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 387, 8, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 398, 8, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 409, 8, 30, 10,
		30, 12, 30, 412, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		420, 8, 31, 1, 32, 1, 32, 1, 32, 5, 32, 425, 8, 32, 10, 32, 12, 32, 428,
		9, 32, 1, 33, 1, 33, 1, 33, 5, 33, 433, 8, 33, 10, 33, 12, 33, 436, 9,
		33, 1, 33, 3, 33, 439, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 445, 8,
		34, 10, 34, 12, 34, 448, 9, 34, 3, 34, 450, 8, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 470, 8, 38, 10, 38, 12, 38, 473,
		9, 38, 3, 38, 475, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 488, 8, 39, 10, 39, 12, 39, 491,
		9, 39, 3, 39, 493, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 3, 40, 500,
		8, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 0, 2, 10, 14, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 12, 2, 0, 1, 1, 102, 107,
		1, 0, 8, 13, 2, 0, 6, 6, 14, 17, 2, 0, 55, 55, 67, 67, 2, 0, 7, 7, 18,
		22, 5, 0, 61, 61, 65, 65, 71, 71, 86, 87, 92, 93, 1, 0, 79, 80, 1, 0, 88,
		89, 2, 0, 26, 26, 30, 30, 2, 0, 33, 33, 94, 95, 1, 0, 36, 51, 1, 0, 82,
		85, 552, 0, 89, 1, 0, 0, 0, 2, 100, 1, 0, 0, 0, 4, 108, 1, 0, 0, 0, 6,
		110, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 136, 1, 0, 0, 0, 12, 170, 1, 0,
		0, 0, 14, 182, 1, 0, 0, 0, 16, 198, 1, 0, 0, 0, 18, 200, 1, 0, 0, 0, 20,
		202, 1, 0, 0, 0, 22, 204, 1, 0, 0, 0, 24, 206, 1, 0, 0, 0, 26, 208, 1,
		0, 0, 0, 28, 210, 1, 0, 0, 0, 30, 247, 1, 0, 0, 0, 32, 249, 1, 0, 0, 0,
		34, 257, 1, 0, 0, 0, 36, 283, 1, 0, 0, 0, 38, 298, 1, 0, 0, 0, 40, 302,
		1, 0, 0, 0, 42, 344, 1, 0, 0, 0, 44, 346, 1, 0, 0, 0, 46, 352, 1, 0, 0,
		0, 48, 356, 1, 0, 0, 0, 50, 373, 1, 0, 0, 0, 52, 386, 1, 0, 0, 0, 54, 397,
		1, 0, 0, 0, 56, 399, 1, 0, 0, 0, 58, 401, 1, 0, 0, 0, 60, 406, 1, 0, 0,
		0, 62, 413, 1, 0, 0, 0, 64, 421, 1, 0, 0, 0, 66, 438, 1, 0, 0, 0, 68, 440,
		1, 0, 0, 0, 70, 453, 1, 0, 0, 0, 72, 456, 1, 0, 0, 0, 74, 460, 1, 0, 0,
		0, 76, 464, 1, 0, 0, 0, 78, 480, 1, 0, 0, 0, 80, 496, 1, 0, 0, 0, 82, 504,
		1, 0, 0, 0, 84, 506, 1, 0, 0, 0, 86, 88, 3, 2, 1, 0, 87, 86, 1, 0, 0, 0,
		88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 1, 1, 0,
		0, 0, 91, 89, 1, 0, 0, 0, 92, 101, 3, 8, 4, 0, 93, 101, 3, 10, 5, 0, 94,
		101, 3, 4, 2, 0, 95, 101, 3, 34, 17, 0, 96, 101, 3, 30, 15, 0, 97, 101,
		3, 68, 34, 0, 98, 101, 3, 70, 35, 0, 99, 101, 3, 74, 37, 0, 100, 92, 1,
		0, 0, 0, 100, 93, 1, 0, 0, 0, 100, 94, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0,
		100, 96, 1, 0, 0, 0, 100, 97, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 99,
		1, 0, 0, 0, 101, 3, 1, 0, 0, 0, 102, 109, 3, 40, 20, 0, 103, 109, 3, 42,
		21, 0, 104, 109, 3, 44, 22, 0, 105, 109, 3, 62, 31, 0, 106, 109, 3, 72,
		36, 0, 107, 109, 3, 46, 23, 0, 108, 102, 1, 0, 0, 0, 108, 103, 1, 0, 0,
		0, 108, 104, 1, 0, 0, 0, 108, 105, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		107, 1, 0, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111, 5, 52, 0, 0, 111, 7, 1,
		0, 0, 0, 112, 114, 5, 74, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 116, 3, 64, 32, 0, 116, 117, 7, 0, 0, 0,
		117, 118, 3, 66, 33, 0, 118, 119, 3, 6, 3, 0, 119, 9, 1, 0, 0, 0, 120,
		121, 6, 5, -1, 0, 121, 137, 3, 28, 14, 0, 122, 137, 3, 60, 30, 0, 123,
		137, 3, 12, 6, 0, 124, 125, 5, 3, 0, 0, 125, 126, 3, 10, 5, 0, 126, 127,
		5, 4, 0, 0, 127, 137, 1, 0, 0, 0, 128, 137, 3, 30, 15, 0, 129, 137, 3,
		48, 24, 0, 130, 137, 3, 54, 27, 0, 131, 137, 3, 76, 38, 0, 132, 133, 5,
		100, 0, 0, 133, 137, 3, 10, 5, 3, 134, 135, 5, 101, 0, 0, 135, 137, 3,
		10, 5, 1, 136, 120, 1, 0, 0, 0, 136, 122, 1, 0, 0, 0, 136, 123, 1, 0, 0,
		0, 136, 124, 1, 0, 0, 0, 136, 128, 1, 0, 0, 0, 136, 129, 1, 0, 0, 0, 136,
		130, 1, 0, 0, 0, 136, 131, 1, 0, 0, 0, 136, 132, 1, 0, 0, 0, 136, 134,
		1, 0, 0, 0, 137, 154, 1, 0, 0, 0, 138, 139, 10, 12, 0, 0, 139, 140, 5,
		2, 0, 0, 140, 153, 3, 10, 5, 13, 141, 142, 10, 10, 0, 0, 142, 143, 3, 16,
		8, 0, 143, 144, 3, 10, 5, 11, 144, 153, 1, 0, 0, 0, 145, 146, 10, 5, 0,
		0, 146, 147, 5, 109, 0, 0, 147, 153, 3, 10, 5, 6, 148, 149, 10, 4, 0, 0,
		149, 153, 5, 100, 0, 0, 150, 151, 10, 2, 0, 0, 151, 153, 5, 101, 0, 0,
		152, 138, 1, 0, 0, 0, 152, 141, 1, 0, 0, 0, 152, 145, 1, 0, 0, 0, 152,
		148, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 11, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 157, 171, 3, 14, 7, 0, 158, 159, 5, 66, 0, 0, 159, 171, 3, 12, 6,
		0, 160, 161, 5, 5, 0, 0, 161, 171, 3, 12, 6, 0, 162, 163, 5, 6, 0, 0, 163,
		171, 3, 12, 6, 0, 164, 165, 5, 7, 0, 0, 165, 171, 3, 12, 6, 0, 166, 167,
		5, 100, 0, 0, 167, 171, 3, 12, 6, 0, 168, 169, 5, 101, 0, 0, 169, 171,
		3, 12, 6, 0, 170, 157, 1, 0, 0, 0, 170, 158, 1, 0, 0, 0, 170, 160, 1, 0,
		0, 0, 170, 162, 1, 0, 0, 0, 170, 164, 1, 0, 0, 0, 170, 166, 1, 0, 0, 0,
		170, 168, 1, 0, 0, 0, 171, 13, 1, 0, 0, 0, 172, 173, 6, 7, -1, 0, 173,
		183, 3, 28, 14, 0, 174, 183, 3, 60, 30, 0, 175, 176, 5, 3, 0, 0, 176, 177,
		3, 10, 5, 0, 177, 178, 5, 4, 0, 0, 178, 183, 1, 0, 0, 0, 179, 183, 3, 30,
		15, 0, 180, 183, 3, 48, 24, 0, 181, 183, 3, 54, 27, 0, 182, 172, 1, 0,
		0, 0, 182, 174, 1, 0, 0, 0, 182, 175, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0,
		182, 180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 190, 1, 0, 0, 0, 184,
		185, 10, 2, 0, 0, 185, 189, 5, 100, 0, 0, 186, 187, 10, 1, 0, 0, 187, 189,
		5, 101, 0, 0, 188, 184, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 192, 1,
		0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 15, 1, 0, 0,
		0, 192, 190, 1, 0, 0, 0, 193, 199, 3, 18, 9, 0, 194, 199, 3, 20, 10, 0,
		195, 199, 3, 22, 11, 0, 196, 199, 3, 24, 12, 0, 197, 199, 3, 26, 13, 0,
		198, 193, 1, 0, 0, 0, 198, 194, 1, 0, 0, 0, 198, 195, 1, 0, 0, 0, 198,
		196, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 17, 1, 0, 0, 0, 200, 201, 7,
		1, 0, 0, 201, 19, 1, 0, 0, 0, 202, 203, 7, 2, 0, 0, 203, 21, 1, 0, 0, 0,
		204, 205, 7, 3, 0, 0, 205, 23, 1, 0, 0, 0, 206, 207, 7, 4, 0, 0, 207, 25,
		1, 0, 0, 0, 208, 209, 5, 23, 0, 0, 209, 27, 1, 0, 0, 0, 210, 211, 7, 5,
		0, 0, 211, 29, 1, 0, 0, 0, 212, 213, 7, 6, 0, 0, 213, 214, 5, 3, 0, 0,
		214, 215, 3, 66, 33, 0, 215, 216, 5, 4, 0, 0, 216, 248, 1, 0, 0, 0, 217,
		218, 7, 7, 0, 0, 218, 219, 5, 3, 0, 0, 219, 220, 3, 66, 33, 0, 220, 221,
		5, 4, 0, 0, 221, 248, 1, 0, 0, 0, 222, 229, 3, 60, 30, 0, 223, 229, 3,
		54, 27, 0, 224, 225, 5, 3, 0, 0, 225, 226, 3, 10, 5, 0, 226, 227, 5, 4,
		0, 0, 227, 229, 1, 0, 0, 0, 228, 222, 1, 0, 0, 0, 228, 223, 1, 0, 0, 0,
		228, 224, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 231, 5, 24, 0, 0, 231,
		233, 3, 60, 30, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 236, 5, 3, 0, 0, 235, 237, 3, 66, 33, 0, 236, 235, 1,
		0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 5, 4, 0,
		0, 239, 248, 1, 0, 0, 0, 240, 248, 3, 32, 16, 0, 241, 242, 5, 54, 0, 0,
		242, 243, 5, 3, 0, 0, 243, 244, 3, 66, 33, 0, 244, 245, 5, 4, 0, 0, 245,
		248, 1, 0, 0, 0, 246, 248, 3, 78, 39, 0, 247, 212, 1, 0, 0, 0, 247, 217,
		1, 0, 0, 0, 247, 228, 1, 0, 0, 0, 247, 240, 1, 0, 0, 0, 247, 241, 1, 0,
		0, 0, 247, 246, 1, 0, 0, 0, 248, 31, 1, 0, 0, 0, 249, 250, 5, 25, 0, 0,
		250, 251, 5, 3, 0, 0, 251, 252, 3, 60, 30, 0, 252, 253, 5, 26, 0, 0, 253,
		254, 3, 10, 5, 0, 254, 255, 5, 4, 0, 0, 255, 33, 1, 0, 0, 0, 256, 258,
		5, 74, 0, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 1, 0,
		0, 0, 259, 264, 5, 75, 0, 0, 260, 261, 3, 60, 30, 0, 261, 262, 5, 27, 0,
		0, 262, 263, 3, 60, 30, 0, 263, 265, 1, 0, 0, 0, 264, 260, 1, 0, 0, 0,
		264, 265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 3, 60, 30, 0, 267,
		277, 5, 3, 0, 0, 268, 273, 3, 36, 18, 0, 269, 270, 5, 26, 0, 0, 270, 272,
		3, 36, 18, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1,
		0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 278, 1, 0, 0, 0, 275, 273, 1, 0, 0,
		0, 276, 278, 5, 99, 0, 0, 277, 268, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 277,
		278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 5, 4, 0, 0, 280, 281,
		3, 38, 19, 0, 281, 282, 5, 60, 0, 0, 282, 35, 1, 0, 0, 0, 283, 286, 3,
		60, 30, 0, 284, 285, 5, 1, 0, 0, 285, 287, 3, 10, 5, 0, 286, 284, 1, 0,
		0, 0, 286, 287, 1, 0, 0, 0, 287, 37, 1, 0, 0, 0, 288, 290, 3, 6, 3, 0,
		289, 288, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 296,
		3, 2, 1, 0, 295, 297, 3, 6, 3, 0, 296, 295, 1, 0, 0, 0, 296, 297, 1, 0,
		0, 0, 297, 299, 1, 0, 0, 0, 298, 291, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0,
		300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 39, 1, 0, 0, 0, 302, 303,
		5, 64, 0, 0, 303, 304, 3, 10, 5, 0, 304, 305, 5, 70, 0, 0, 305, 313, 3,
		38, 19, 0, 306, 307, 5, 59, 0, 0, 307, 308, 3, 10, 5, 0, 308, 309, 5, 70,
		0, 0, 309, 310, 3, 38, 19, 0, 310, 312, 1, 0, 0, 0, 311, 306, 1, 0, 0,
		0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314,
		318, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 5, 58, 0, 0, 317, 319,
		3, 38, 19, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 1,
		0, 0, 0, 320, 321, 5, 60, 0, 0, 321, 41, 1, 0, 0, 0, 322, 323, 5, 62, 0,
		0, 323, 324, 3, 60, 30, 0, 324, 325, 5, 53, 0, 0, 325, 326, 3, 10, 5, 0,
		326, 327, 5, 57, 0, 0, 327, 328, 3, 38, 19, 0, 328, 329, 5, 60, 0, 0, 329,
		345, 1, 0, 0, 0, 330, 331, 5, 62, 0, 0, 331, 332, 3, 60, 30, 0, 332, 333,
		5, 1, 0, 0, 333, 334, 3, 10, 5, 0, 334, 335, 5, 26, 0, 0, 335, 338, 3,
		10, 5, 0, 336, 337, 5, 26, 0, 0, 337, 339, 3, 10, 5, 0, 338, 336, 1, 0,
		0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 5, 57, 0, 0,
		341, 342, 3, 38, 19, 0, 342, 343, 5, 60, 0, 0, 343, 345, 1, 0, 0, 0, 344,
		322, 1, 0, 0, 0, 344, 330, 1, 0, 0, 0, 345, 43, 1, 0, 0, 0, 346, 347, 5,
		73, 0, 0, 347, 348, 3, 10, 5, 0, 348, 349, 5, 57, 0, 0, 349, 350, 3, 38,
		19, 0, 350, 351, 5, 60, 0, 0, 351, 45, 1, 0, 0, 0, 352, 353, 5, 57, 0,
		0, 353, 354, 3, 38, 19, 0, 354, 355, 5, 60, 0, 0, 355, 47, 1, 0, 0, 0,
		356, 369, 5, 28, 0, 0, 357, 363, 3, 52, 26, 0, 358, 359, 3, 50, 25, 0,
		359, 360, 3, 52, 26, 0, 360, 362, 1, 0, 0, 0, 361, 358, 1, 0, 0, 0, 362,
		365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 367,
		1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 368, 3, 50, 25, 0, 367, 366, 1,
		0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369, 357, 1, 0, 0,
		0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 5, 29, 0, 0, 372,
		49, 1, 0, 0, 0, 373, 374, 7, 8, 0, 0, 374, 51, 1, 0, 0, 0, 375, 376, 3,
		60, 30, 0, 376, 377, 5, 1, 0, 0, 377, 378, 3, 10, 5, 0, 378, 387, 1, 0,
		0, 0, 379, 380, 5, 31, 0, 0, 380, 381, 3, 10, 5, 0, 381, 382, 5, 32, 0,
		0, 382, 383, 5, 1, 0, 0, 383, 384, 3, 10, 5, 0, 384, 387, 1, 0, 0, 0, 385,
		387, 3, 10, 5, 0, 386, 375, 1, 0, 0, 0, 386, 379, 1, 0, 0, 0, 386, 385,
		1, 0, 0, 0, 387, 53, 1, 0, 0, 0, 388, 389, 3, 60, 30, 0, 389, 390, 5, 31,
		0, 0, 390, 391, 3, 10, 5, 0, 391, 392, 5, 32, 0, 0, 392, 398, 1, 0, 0,
		0, 393, 394, 3, 60, 30, 0, 394, 395, 5, 27, 0, 0, 395, 396, 3, 60, 30,
		0, 396, 398, 1, 0, 0, 0, 397, 388, 1, 0, 0, 0, 397, 393, 1, 0, 0, 0, 398,
		55, 1, 0, 0, 0, 399, 400, 5, 101, 0, 0, 400, 57, 1, 0, 0, 0, 401, 402,
		5, 54, 0, 0, 402, 403, 5, 3, 0, 0, 403, 404, 3, 10, 5, 0, 404, 405, 5,
		4, 0, 0, 405, 59, 1, 0, 0, 0, 406, 410, 5, 94, 0, 0, 407, 409, 7, 9, 0,
		0, 408, 407, 1, 0, 0, 0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410,
		411, 1, 0, 0, 0, 411, 61, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 414, 5,
		68, 0, 0, 414, 415, 3, 38, 19, 0, 415, 416, 5, 72, 0, 0, 416, 419, 3, 10,
		5, 0, 417, 420, 3, 6, 3, 0, 418, 420, 5, 0, 0, 1, 419, 417, 1, 0, 0, 0,
		419, 418, 1, 0, 0, 0, 420, 63, 1, 0, 0, 0, 421, 426, 3, 60, 30, 0, 422,
		423, 5, 26, 0, 0, 423, 425, 3, 60, 30, 0, 424, 422, 1, 0, 0, 0, 425, 428,
		1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 65, 1, 0,
		0, 0, 428, 426, 1, 0, 0, 0, 429, 434, 3, 10, 5, 0, 430, 431, 5, 26, 0,
		0, 431, 433, 3, 10, 5, 0, 432, 430, 1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434,
		432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 439, 1, 0, 0, 0, 436, 434,
		1, 0, 0, 0, 437, 439, 5, 99, 0, 0, 438, 429, 1, 0, 0, 0, 438, 437, 1, 0,
		0, 0, 439, 67, 1, 0, 0, 0, 440, 449, 5, 69, 0, 0, 441, 446, 3, 10, 5, 0,
		442, 443, 5, 26, 0, 0, 443, 445, 3, 10, 5, 0, 444, 442, 1, 0, 0, 0, 445,
		448, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 450,
		1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 449, 441, 1, 0, 0, 0, 449, 450, 1, 0,
		0, 0, 450, 451, 1, 0, 0, 0, 451, 452, 3, 6, 3, 0, 452, 69, 1, 0, 0, 0,
		453, 454, 5, 56, 0, 0, 454, 455, 3, 6, 3, 0, 455, 71, 1, 0, 0, 0, 456,
		457, 5, 63, 0, 0, 457, 458, 3, 60, 30, 0, 458, 459, 3, 6, 3, 0, 459, 73,
		1, 0, 0, 0, 460, 461, 5, 34, 0, 0, 461, 462, 3, 60, 30, 0, 462, 463, 5,
		34, 0, 0, 463, 75, 1, 0, 0, 0, 464, 465, 5, 75, 0, 0, 465, 474, 5, 3, 0,
		0, 466, 471, 3, 60, 30, 0, 467, 468, 5, 26, 0, 0, 468, 470, 3, 60, 30,
		0, 469, 467, 1, 0, 0, 0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471,
		472, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 474, 466,
		1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 477, 5, 4,
		0, 0, 477, 478, 3, 38, 19, 0, 478, 479, 5, 60, 0, 0, 479, 77, 1, 0, 0,
		0, 480, 481, 3, 60, 30, 0, 481, 482, 5, 24, 0, 0, 482, 483, 3, 60, 30,
		0, 483, 492, 5, 3, 0, 0, 484, 489, 3, 10, 5, 0, 485, 486, 5, 26, 0, 0,
		486, 488, 3, 10, 5, 0, 487, 485, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489,
		487, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 493, 1, 0, 0, 0, 491, 489,
		1, 0, 0, 0, 492, 484, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 494, 1, 0,
		0, 0, 494, 495, 5, 4, 0, 0, 495, 79, 1, 0, 0, 0, 496, 499, 5, 35, 0, 0,
		497, 500, 3, 82, 41, 0, 498, 500, 3, 60, 30, 0, 499, 497, 1, 0, 0, 0, 499,
		498, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 502, 5, 1, 0, 0, 502, 503,
		3, 10, 5, 0, 503, 81, 1, 0, 0, 0, 504, 505, 7, 10, 0, 0, 505, 83, 1, 0,
		0, 0, 506, 507, 5, 81, 0, 0, 507, 508, 5, 27, 0, 0, 508, 509, 7, 11, 0,
		0, 509, 85, 1, 0, 0, 0, 45, 89, 100, 108, 113, 136, 152, 154, 170, 182,
		188, 190, 198, 228, 232, 236, 247, 257, 264, 273, 277, 286, 291, 296, 300,
		313, 318, 338, 344, 363, 367, 369, 386, 397, 410, 419, 426, 434, 438, 446,
		449, 471, 474, 489, 492, 499,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// lua_grammar_antlr4ParserInit initializes any static state used to implement lua_grammar_antlr4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newlua_grammar_antlr4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lua_grammar_antlr4ParserInit() {
	staticData := &Lua_grammar_antlr4ParserStaticData
	staticData.once.Do(lua_grammar_antlr4ParserInit)
}

// Newlua_grammar_antlr4Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newlua_grammar_antlr4Parser(input antlr.TokenStream) *lua_grammar_antlr4Parser {
	Lua_grammar_antlr4ParserInit()
	this := new(lua_grammar_antlr4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Lua_grammar_antlr4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "lua_grammar_antlr4.g4"

	return this
}

// lua_grammar_antlr4Parser tokens.
const (
	lua_grammar_antlr4ParserEOF                 = antlr.TokenEOF
	lua_grammar_antlr4ParserT__0                = 1
	lua_grammar_antlr4ParserT__1                = 2
	lua_grammar_antlr4ParserT__2                = 3
	lua_grammar_antlr4ParserT__3                = 4
	lua_grammar_antlr4ParserT__4                = 5
	lua_grammar_antlr4ParserT__5                = 6
	lua_grammar_antlr4ParserT__6                = 7
	lua_grammar_antlr4ParserT__7                = 8
	lua_grammar_antlr4ParserT__8                = 9
	lua_grammar_antlr4ParserT__9                = 10
	lua_grammar_antlr4ParserT__10               = 11
	lua_grammar_antlr4ParserT__11               = 12
	lua_grammar_antlr4ParserT__12               = 13
	lua_grammar_antlr4ParserT__13               = 14
	lua_grammar_antlr4ParserT__14               = 15
	lua_grammar_antlr4ParserT__15               = 16
	lua_grammar_antlr4ParserT__16               = 17
	lua_grammar_antlr4ParserT__17               = 18
	lua_grammar_antlr4ParserT__18               = 19
	lua_grammar_antlr4ParserT__19               = 20
	lua_grammar_antlr4ParserT__20               = 21
	lua_grammar_antlr4ParserT__21               = 22
	lua_grammar_antlr4ParserT__22               = 23
	lua_grammar_antlr4ParserT__23               = 24
	lua_grammar_antlr4ParserT__24               = 25
	lua_grammar_antlr4ParserT__25               = 26
	lua_grammar_antlr4ParserT__26               = 27
	lua_grammar_antlr4ParserT__27               = 28
	lua_grammar_antlr4ParserT__28               = 29
	lua_grammar_antlr4ParserT__29               = 30
	lua_grammar_antlr4ParserT__30               = 31
	lua_grammar_antlr4ParserT__31               = 32
	lua_grammar_antlr4ParserT__32               = 33
	lua_grammar_antlr4ParserT__33               = 34
	lua_grammar_antlr4ParserT__34               = 35
	lua_grammar_antlr4ParserT__35               = 36
	lua_grammar_antlr4ParserT__36               = 37
	lua_grammar_antlr4ParserT__37               = 38
	lua_grammar_antlr4ParserT__38               = 39
	lua_grammar_antlr4ParserT__39               = 40
	lua_grammar_antlr4ParserT__40               = 41
	lua_grammar_antlr4ParserT__41               = 42
	lua_grammar_antlr4ParserT__42               = 43
	lua_grammar_antlr4ParserT__43               = 44
	lua_grammar_antlr4ParserT__44               = 45
	lua_grammar_antlr4ParserT__45               = 46
	lua_grammar_antlr4ParserT__46               = 47
	lua_grammar_antlr4ParserT__47               = 48
	lua_grammar_antlr4ParserT__48               = 49
	lua_grammar_antlr4ParserT__49               = 50
	lua_grammar_antlr4ParserT__50               = 51
	lua_grammar_antlr4ParserSEPARATOR           = 52
	lua_grammar_antlr4ParserKW_IN               = 53
	lua_grammar_antlr4ParserKW_PRINT            = 54
	lua_grammar_antlr4ParserKW_AND              = 55
	lua_grammar_antlr4ParserKW_BREAK            = 56
	lua_grammar_antlr4ParserKW_DO               = 57
	lua_grammar_antlr4ParserKW_ELSE             = 58
	lua_grammar_antlr4ParserKW_ELSEIF           = 59
	lua_grammar_antlr4ParserKW_END              = 60
	lua_grammar_antlr4ParserKW_FALSE            = 61
	lua_grammar_antlr4ParserKW_FOR              = 62
	lua_grammar_antlr4ParserKW_GOTO             = 63
	lua_grammar_antlr4ParserKW_IF               = 64
	lua_grammar_antlr4ParserKW_NIL              = 65
	lua_grammar_antlr4ParserKW_NOT              = 66
	lua_grammar_antlr4ParserKW_OR               = 67
	lua_grammar_antlr4ParserKW_REPEAT           = 68
	lua_grammar_antlr4ParserKW_RETURN           = 69
	lua_grammar_antlr4ParserKW_THEN             = 70
	lua_grammar_antlr4ParserKW_TRUE             = 71
	lua_grammar_antlr4ParserKW_UNTIL            = 72
	lua_grammar_antlr4ParserKW_WHILE            = 73
	lua_grammar_antlr4ParserKW_LOCAL            = 74
	lua_grammar_antlr4ParserKW_FUNCTION         = 75
	lua_grammar_antlr4ParserKW_INDEX            = 76
	lua_grammar_antlr4ParserKW_NEWINDEX         = 77
	lua_grammar_antlr4ParserKW_MODE             = 78
	lua_grammar_antlr4ParserKW_PCALL            = 79
	lua_grammar_antlr4ParserKW_XPCALL           = 80
	lua_grammar_antlr4ParserKW_COROUTINE        = 81
	lua_grammar_antlr4ParserKW_CREATE           = 82
	lua_grammar_antlr4ParserKW_RESUME           = 83
	lua_grammar_antlr4ParserKW_YIELD            = 84
	lua_grammar_antlr4ParserKW_STATUS           = 85
	lua_grammar_antlr4ParserKW_NAN              = 86
	lua_grammar_antlr4ParserKW_INF              = 87
	lua_grammar_antlr4ParserKW_ERROR            = 88
	lua_grammar_antlr4ParserKW_ASSERT           = 89
	lua_grammar_antlr4ParserKW_PAIRS            = 90
	lua_grammar_antlr4ParserKW_IPAIRS           = 91
	lua_grammar_antlr4ParserNUMBER              = 92
	lua_grammar_antlr4ParserSTRING              = 93
	lua_grammar_antlr4ParserLETTER              = 94
	lua_grammar_antlr4ParserDIGIT               = 95
	lua_grammar_antlr4ParserWS                  = 96
	lua_grammar_antlr4ParserSINGLE_LINE_COMMENT = 97
	lua_grammar_antlr4ParserMULTI_LINE_COMMENT  = 98
	lua_grammar_antlr4ParserVARARG              = 99
	lua_grammar_antlr4ParserINCREMENT           = 100
	lua_grammar_antlr4ParserDECREMENT           = 101
	lua_grammar_antlr4ParserPLUS_ASSIGN         = 102
	lua_grammar_antlr4ParserMINUS_ASSIGN        = 103
	lua_grammar_antlr4ParserMULT_ASSIGN         = 104
	lua_grammar_antlr4ParserDIV_ASSIGN          = 105
	lua_grammar_antlr4ParserFLOOR_DIV_ASSIGN    = 106
	lua_grammar_antlr4ParserEXP_ASSIGN          = 107
	lua_grammar_antlr4ParserCONCAT_ASSIGN       = 108
	lua_grammar_antlr4ParserNULL_COALESCE       = 109
)

// lua_grammar_antlr4Parser rules.
const (
	lua_grammar_antlr4ParserRULE_program              = 0
	lua_grammar_antlr4ParserRULE_statement            = 1
	lua_grammar_antlr4ParserRULE_control_statement    = 2
	lua_grammar_antlr4ParserRULE_statement_terminator = 3
	lua_grammar_antlr4ParserRULE_assignment           = 4
	lua_grammar_antlr4ParserRULE_expression           = 5
	lua_grammar_antlr4ParserRULE_prefix_expression    = 6
	lua_grammar_antlr4ParserRULE_primary_expression   = 7
	lua_grammar_antlr4ParserRULE_operators            = 8
	lua_grammar_antlr4ParserRULE_comparison_operator  = 9
	lua_grammar_antlr4ParserRULE_arith_operator       = 10
	lua_grammar_antlr4ParserRULE_logical_operator     = 11
	lua_grammar_antlr4ParserRULE_bitwise_operator     = 12
	lua_grammar_antlr4ParserRULE_concat_operator      = 13
	lua_grammar_antlr4ParserRULE_literal              = 14
	lua_grammar_antlr4ParserRULE_function_call        = 15
	lua_grammar_antlr4ParserRULE_table_insert         = 16
	lua_grammar_antlr4ParserRULE_function_declaration = 17
	lua_grammar_antlr4ParserRULE_parameter            = 18
	lua_grammar_antlr4ParserRULE_block                = 19
	lua_grammar_antlr4ParserRULE_if_statement         = 20
	lua_grammar_antlr4ParserRULE_for_statement        = 21
	lua_grammar_antlr4ParserRULE_while_statement      = 22
	lua_grammar_antlr4ParserRULE_do_statement         = 23
	lua_grammar_antlr4ParserRULE_table                = 24
	lua_grammar_antlr4ParserRULE_field_separator      = 25
	lua_grammar_antlr4ParserRULE_field                = 26
	lua_grammar_antlr4ParserRULE_table_access         = 27
	lua_grammar_antlr4ParserRULE_single_line_comment  = 28
	lua_grammar_antlr4ParserRULE_print_statement      = 29
	lua_grammar_antlr4ParserRULE_identifier           = 30
	lua_grammar_antlr4ParserRULE_repeat_statement     = 31
	lua_grammar_antlr4ParserRULE_identifier_list      = 32
	lua_grammar_antlr4ParserRULE_expression_list      = 33
	lua_grammar_antlr4ParserRULE_return_statement     = 34
	lua_grammar_antlr4ParserRULE_break_statement      = 35
	lua_grammar_antlr4ParserRULE_goto_statement       = 36
	lua_grammar_antlr4ParserRULE_label_statement      = 37
	lua_grammar_antlr4ParserRULE_function_expression  = 38
	lua_grammar_antlr4ParserRULE_method_call          = 39
	lua_grammar_antlr4ParserRULE_metatable_field      = 40
	lua_grammar_antlr4ParserRULE_metamethod           = 41
	lua_grammar_antlr4ParserRULE_coroutine_statement  = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lua_grammar_antlr4ParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2071655811108568856) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&208100495031) != 0) {
		{
			p.SetState(86)
			p.Statement()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Expression() IExpressionContext
	Control_statement() IControl_statementContext
	Function_declaration() IFunction_declarationContext
	Function_call() IFunction_callContext
	Return_statement() IReturn_statementContext
	Break_statement() IBreak_statementContext
	Label_statement() ILabel_statementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Control_statement() IControl_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControl_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControl_statementContext)
}

func (s *StatementContext) Function_declaration() IFunction_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *StatementContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) Break_statement() IBreak_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_statementContext)
}

func (s *StatementContext) Label_statement() ILabel_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabel_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabel_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lua_grammar_antlr4ParserRULE_statement)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.Function_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.Function_call()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(97)
			p.Return_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.Break_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.Label_statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControl_statementContext is an interface to support dynamic dispatch.
type IControl_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_statement() IIf_statementContext
	For_statement() IFor_statementContext
	While_statement() IWhile_statementContext
	Repeat_statement() IRepeat_statementContext
	Goto_statement() IGoto_statementContext
	Do_statement() IDo_statementContext

	// IsControl_statementContext differentiates from other interfaces.
	IsControl_statementContext()
}

type Control_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControl_statementContext() *Control_statementContext {
	var p = new(Control_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement
	return p
}

func InitEmptyControl_statementContext(p *Control_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement
}

func (*Control_statementContext) IsControl_statementContext() {}

func NewControl_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_statementContext {
	var p = new(Control_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_control_statement

	return p
}

func (s *Control_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_statementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Control_statementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *Control_statementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *Control_statementContext) Repeat_statement() IRepeat_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeat_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeat_statementContext)
}

func (s *Control_statementContext) Goto_statement() IGoto_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoto_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoto_statementContext)
}

func (s *Control_statementContext) Do_statement() IDo_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDo_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDo_statementContext)
}

func (s *Control_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterControl_statement(s)
	}
}

func (s *Control_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitControl_statement(s)
	}
}

func (s *Control_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitControl_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Control_statement() (localctx IControl_statementContext) {
	localctx = NewControl_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lua_grammar_antlr4ParserRULE_control_statement)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserKW_IF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.If_statement()
		}

	case lua_grammar_antlr4ParserKW_FOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.For_statement()
		}

	case lua_grammar_antlr4ParserKW_WHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.While_statement()
		}

	case lua_grammar_antlr4ParserKW_REPEAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Repeat_statement()
		}

	case lua_grammar_antlr4ParserKW_GOTO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.Goto_statement()
		}

	case lua_grammar_antlr4ParserKW_DO:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Do_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatement_terminatorContext is an interface to support dynamic dispatch.
type IStatement_terminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEPARATOR() antlr.TerminalNode

	// IsStatement_terminatorContext differentiates from other interfaces.
	IsStatement_terminatorContext()
}

type Statement_terminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_terminatorContext() *Statement_terminatorContext {
	var p = new(Statement_terminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator
	return p
}

func InitEmptyStatement_terminatorContext(p *Statement_terminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator
}

func (*Statement_terminatorContext) IsStatement_terminatorContext() {}

func NewStatement_terminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_terminatorContext {
	var p = new(Statement_terminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_statement_terminator

	return p
}

func (s *Statement_terminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_terminatorContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserSEPARATOR, 0)
}

func (s *Statement_terminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_terminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_terminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterStatement_terminator(s)
	}
}

func (s *Statement_terminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitStatement_terminator(s)
	}
}

func (s *Statement_terminatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitStatement_terminator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Statement_terminator() (localctx IStatement_terminatorContext) {
	localctx = NewStatement_terminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lua_grammar_antlr4ParserRULE_statement_terminator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(lua_grammar_antlr4ParserSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier_list() IIdentifier_listContext
	Expression_list() IExpression_listContext
	Statement_terminator() IStatement_terminatorContext
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	MULT_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	FLOOR_DIV_ASSIGN() antlr.TerminalNode
	EXP_ASSIGN() antlr.TerminalNode
	KW_LOCAL() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier_list() IIdentifier_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifier_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifier_listContext)
}

func (s *AssignmentContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *AssignmentContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *AssignmentContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserPLUS_ASSIGN, 0)
}

func (s *AssignmentContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserMINUS_ASSIGN, 0)
}

func (s *AssignmentContext) MULT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserMULT_ASSIGN, 0)
}

func (s *AssignmentContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDIV_ASSIGN, 0)
}

func (s *AssignmentContext) FLOOR_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserFLOOR_DIV_ASSIGN, 0)
}

func (s *AssignmentContext) EXP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserEXP_ASSIGN, 0)
}

func (s *AssignmentContext) KW_LOCAL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_LOCAL, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lua_grammar_antlr4ParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(112)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(115)
		p.Identifier_list()
	}
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserT__0 || ((int64((_la-102)) & ^0x3f) == 0 && ((int64(1)<<(_la-102))&63) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(117)
		p.Expression_list()
	}
	{
		p.SetState(118)
		p.Statement_terminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	Identifier() IIdentifierContext
	Prefix_expression() IPrefix_expressionContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Function_call() IFunction_callContext
	Table() ITableContext
	Table_access() ITable_accessContext
	Function_expression() IFunction_expressionContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	Operators() IOperatorsContext
	NULL_COALESCE() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Prefix_expression() IPrefix_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_expressionContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *ExpressionContext) Table_access() ITable_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_accessContext)
}

func (s *ExpressionContext) Function_expression() IFunction_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_expressionContext)
}

func (s *ExpressionContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserINCREMENT, 0)
}

func (s *ExpressionContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDECREMENT, 0)
}

func (s *ExpressionContext) Operators() IOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorsContext)
}

func (s *ExpressionContext) NULL_COALESCE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserNULL_COALESCE, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *lua_grammar_antlr4Parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, lua_grammar_antlr4ParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(121)
			p.Literal()
		}

	case 2:
		{
			p.SetState(122)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(123)
			p.Prefix_expression()
		}

	case 4:
		{
			p.SetState(124)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expression(0)
		}
		{
			p.SetState(126)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(128)
			p.Function_call()
		}

	case 6:
		{
			p.SetState(129)
			p.Table()
		}

	case 7:
		{
			p.SetState(130)
			p.Table_access()
		}

	case 8:
		{
			p.SetState(131)
			p.Function_expression()
		}

	case 9:
		{
			p.SetState(132)
			p.Match(lua_grammar_antlr4ParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expression(3)
		}

	case 10:
		{
			p.SetState(134)
			p.Match(lua_grammar_antlr4ParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(139)
					p.Match(lua_grammar_antlr4ParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(140)
					p.expression(13)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(142)
					p.Operators()
				}
				{
					p.SetState(143)
					p.expression(11)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(146)
					p.Match(lua_grammar_antlr4ParserNULL_COALESCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(147)
					p.expression(6)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(149)
					p.Match(lua_grammar_antlr4ParserINCREMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(151)
					p.Match(lua_grammar_antlr4ParserDECREMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefix_expressionContext is an interface to support dynamic dispatch.
type IPrefix_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary_expression() IPrimary_expressionContext
	KW_NOT() antlr.TerminalNode
	Prefix_expression() IPrefix_expressionContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode

	// IsPrefix_expressionContext differentiates from other interfaces.
	IsPrefix_expressionContext()
}

type Prefix_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_expressionContext() *Prefix_expressionContext {
	var p = new(Prefix_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_prefix_expression
	return p
}

func InitEmptyPrefix_expressionContext(p *Prefix_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_prefix_expression
}

func (*Prefix_expressionContext) IsPrefix_expressionContext() {}

func NewPrefix_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_expressionContext {
	var p = new(Prefix_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_prefix_expression

	return p
}

func (s *Prefix_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Prefix_expressionContext) Primary_expression() IPrimary_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *Prefix_expressionContext) KW_NOT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_NOT, 0)
}

func (s *Prefix_expressionContext) Prefix_expression() IPrefix_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_expressionContext)
}

func (s *Prefix_expressionContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserINCREMENT, 0)
}

func (s *Prefix_expressionContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDECREMENT, 0)
}

func (s *Prefix_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterPrefix_expression(s)
	}
}

func (s *Prefix_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitPrefix_expression(s)
	}
}

func (s *Prefix_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitPrefix_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Prefix_expression() (localctx IPrefix_expressionContext) {
	localctx = NewPrefix_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lua_grammar_antlr4ParserRULE_prefix_expression)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__2, lua_grammar_antlr4ParserT__24, lua_grammar_antlr4ParserT__27, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_PCALL, lua_grammar_antlr4ParserKW_XPCALL, lua_grammar_antlr4ParserKW_NAN, lua_grammar_antlr4ParserKW_INF, lua_grammar_antlr4ParserKW_ERROR, lua_grammar_antlr4ParserKW_ASSERT, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.primary_expression(0)
		}

	case lua_grammar_antlr4ParserKW_NOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(lua_grammar_antlr4ParserKW_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(lua_grammar_antlr4ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Match(lua_grammar_antlr4ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.Match(lua_grammar_antlr4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserINCREMENT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.Match(lua_grammar_antlr4ParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserDECREMENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(168)
			p.Match(lua_grammar_antlr4ParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Prefix_expression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimary_expressionContext is an interface to support dynamic dispatch.
type IPrimary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	Identifier() IIdentifierContext
	Expression() IExpressionContext
	Function_call() IFunction_callContext
	Table() ITableContext
	Table_access() ITable_accessContext
	Primary_expression() IPrimary_expressionContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode

	// IsPrimary_expressionContext differentiates from other interfaces.
	IsPrimary_expressionContext()
}

type Primary_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_expressionContext() *Primary_expressionContext {
	var p = new(Primary_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primary_expression
	return p
}

func InitEmptyPrimary_expressionContext(p *Primary_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primary_expression
}

func (*Primary_expressionContext) IsPrimary_expressionContext() {}

func NewPrimary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_expressionContext {
	var p = new(Primary_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primary_expression

	return p
}

func (s *Primary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_expressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Primary_expressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Primary_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Primary_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Primary_expressionContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *Primary_expressionContext) Table_access() ITable_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_accessContext)
}

func (s *Primary_expressionContext) Primary_expression() IPrimary_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *Primary_expressionContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserINCREMENT, 0)
}

func (s *Primary_expressionContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDECREMENT, 0)
}

func (s *Primary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterPrimary_expression(s)
	}
}

func (s *Primary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitPrimary_expression(s)
	}
}

func (s *Primary_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitPrimary_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Primary_expression() (localctx IPrimary_expressionContext) {
	return p.primary_expression(0)
}

func (p *lua_grammar_antlr4Parser) primary_expression(_p int) (localctx IPrimary_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimary_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, lua_grammar_antlr4ParserRULE_primary_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(173)
			p.Literal()
		}

	case 2:
		{
			p.SetState(174)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(175)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.expression(0)
		}
		{
			p.SetState(177)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(179)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(180)
			p.Table()
		}

	case 6:
		{
			p.SetState(181)
			p.Table_access()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_primary_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(185)
					p.Match(lua_grammar_antlr4ParserINCREMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_primary_expression)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(187)
					p.Match(lua_grammar_antlr4ParserDECREMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorsContext is an interface to support dynamic dispatch.
type IOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comparison_operator() IComparison_operatorContext
	Arith_operator() IArith_operatorContext
	Logical_operator() ILogical_operatorContext
	Bitwise_operator() IBitwise_operatorContext
	Concat_operator() IConcat_operatorContext

	// IsOperatorsContext differentiates from other interfaces.
	IsOperatorsContext()
}

type OperatorsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorsContext() *OperatorsContext {
	var p = new(OperatorsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators
	return p
}

func InitEmptyOperatorsContext(p *OperatorsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators
}

func (*OperatorsContext) IsOperatorsContext() {}

func NewOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorsContext {
	var p = new(OperatorsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operators

	return p
}

func (s *OperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorsContext) Comparison_operator() IComparison_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_operatorContext)
}

func (s *OperatorsContext) Arith_operator() IArith_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArith_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArith_operatorContext)
}

func (s *OperatorsContext) Logical_operator() ILogical_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_operatorContext)
}

func (s *OperatorsContext) Bitwise_operator() IBitwise_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitwise_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitwise_operatorContext)
}

func (s *OperatorsContext) Concat_operator() IConcat_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcat_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcat_operatorContext)
}

func (s *OperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterOperators(s)
	}
}

func (s *OperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitOperators(s)
	}
}

func (s *OperatorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitOperators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Operators() (localctx IOperatorsContext) {
	localctx = NewOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lua_grammar_antlr4ParserRULE_operators)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__7, lua_grammar_antlr4ParserT__8, lua_grammar_antlr4ParserT__9, lua_grammar_antlr4ParserT__10, lua_grammar_antlr4ParserT__11, lua_grammar_antlr4ParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Comparison_operator()
		}

	case lua_grammar_antlr4ParserT__5, lua_grammar_antlr4ParserT__13, lua_grammar_antlr4ParserT__14, lua_grammar_antlr4ParserT__15, lua_grammar_antlr4ParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Arith_operator()
		}

	case lua_grammar_antlr4ParserKW_AND, lua_grammar_antlr4ParserKW_OR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Logical_operator()
		}

	case lua_grammar_antlr4ParserT__6, lua_grammar_antlr4ParserT__17, lua_grammar_antlr4ParserT__18, lua_grammar_antlr4ParserT__19, lua_grammar_antlr4ParserT__20, lua_grammar_antlr4ParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(196)
			p.Bitwise_operator()
		}

	case lua_grammar_antlr4ParserT__22:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(197)
			p.Concat_operator()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparison_operatorContext is an interface to support dynamic dispatch.
type IComparison_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparison_operatorContext differentiates from other interfaces.
	IsComparison_operatorContext()
}

type Comparison_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_operatorContext() *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator
	return p
}

func InitEmptyComparison_operatorContext(p *Comparison_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator
}

func (*Comparison_operatorContext) IsComparison_operatorContext() {}

func NewComparison_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparison_operator

	return p
}

func (s *Comparison_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Comparison_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitComparison_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Comparison_operator() (localctx IComparison_operatorContext) {
	localctx = NewComparison_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lua_grammar_antlr4ParserRULE_comparison_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArith_operatorContext is an interface to support dynamic dispatch.
type IArith_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArith_operatorContext differentiates from other interfaces.
	IsArith_operatorContext()
}

type Arith_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArith_operatorContext() *Arith_operatorContext {
	var p = new(Arith_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator
	return p
}

func InitEmptyArith_operatorContext(p *Arith_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator
}

func (*Arith_operatorContext) IsArith_operatorContext() {}

func NewArith_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arith_operatorContext {
	var p = new(Arith_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arith_operator

	return p
}

func (s *Arith_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Arith_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arith_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arith_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterArith_operator(s)
	}
}

func (s *Arith_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitArith_operator(s)
	}
}

func (s *Arith_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitArith_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Arith_operator() (localctx IArith_operatorContext) {
	localctx = NewArith_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lua_grammar_antlr4ParserRULE_arith_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245824) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_operatorContext is an interface to support dynamic dispatch.
type ILogical_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_AND() antlr.TerminalNode
	KW_OR() antlr.TerminalNode

	// IsLogical_operatorContext differentiates from other interfaces.
	IsLogical_operatorContext()
}

type Logical_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_operatorContext() *Logical_operatorContext {
	var p = new(Logical_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator
	return p
}

func InitEmptyLogical_operatorContext(p *Logical_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator
}

func (*Logical_operatorContext) IsLogical_operatorContext() {}

func NewLogical_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_operatorContext {
	var p = new(Logical_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logical_operator

	return p
}

func (s *Logical_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_operatorContext) KW_AND() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_AND, 0)
}

func (s *Logical_operatorContext) KW_OR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_OR, 0)
}

func (s *Logical_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLogical_operator(s)
	}
}

func (s *Logical_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLogical_operator(s)
	}
}

func (s *Logical_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLogical_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Logical_operator() (localctx ILogical_operatorContext) {
	localctx = NewLogical_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, lua_grammar_antlr4ParserRULE_logical_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserKW_AND || _la == lua_grammar_antlr4ParserKW_OR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBitwise_operatorContext is an interface to support dynamic dispatch.
type IBitwise_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBitwise_operatorContext differentiates from other interfaces.
	IsBitwise_operatorContext()
}

type Bitwise_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitwise_operatorContext() *Bitwise_operatorContext {
	var p = new(Bitwise_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwise_operator
	return p
}

func InitEmptyBitwise_operatorContext(p *Bitwise_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwise_operator
}

func (*Bitwise_operatorContext) IsBitwise_operatorContext() {}

func NewBitwise_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bitwise_operatorContext {
	var p = new(Bitwise_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwise_operator

	return p
}

func (s *Bitwise_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Bitwise_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bitwise_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bitwise_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBitwise_operator(s)
	}
}

func (s *Bitwise_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBitwise_operator(s)
	}
}

func (s *Bitwise_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBitwise_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Bitwise_operator() (localctx IBitwise_operatorContext) {
	localctx = NewBitwise_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, lua_grammar_antlr4ParserRULE_bitwise_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8126592) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConcat_operatorContext is an interface to support dynamic dispatch.
type IConcat_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConcat_operatorContext differentiates from other interfaces.
	IsConcat_operatorContext()
}

type Concat_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcat_operatorContext() *Concat_operatorContext {
	var p = new(Concat_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concat_operator
	return p
}

func InitEmptyConcat_operatorContext(p *Concat_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concat_operator
}

func (*Concat_operatorContext) IsConcat_operatorContext() {}

func NewConcat_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Concat_operatorContext {
	var p = new(Concat_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concat_operator

	return p
}

func (s *Concat_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Concat_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Concat_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Concat_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterConcat_operator(s)
	}
}

func (s *Concat_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitConcat_operator(s)
	}
}

func (s *Concat_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitConcat_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Concat_operator() (localctx IConcat_operatorContext) {
	localctx = NewConcat_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, lua_grammar_antlr4ParserRULE_concat_operator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(lua_grammar_antlr4ParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	KW_TRUE() antlr.TerminalNode
	KW_FALSE() antlr.TerminalNode
	KW_NIL() antlr.TerminalNode
	KW_NAN() antlr.TerminalNode
	KW_INF() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserSTRING, 0)
}

func (s *LiteralContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_TRUE, 0)
}

func (s *LiteralContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FALSE, 0)
}

func (s *LiteralContext) KW_NIL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_NIL, 0)
}

func (s *LiteralContext) KW_NAN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_NAN, 0)
}

func (s *LiteralContext) KW_INF() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_INF, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, lua_grammar_antlr4ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-61)) & ^0x3f) == 0 && ((int64(1)<<(_la-61))&6543115281) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_list() IExpression_listContext
	KW_PCALL() antlr.TerminalNode
	KW_XPCALL() antlr.TerminalNode
	KW_ERROR() antlr.TerminalNode
	KW_ASSERT() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Table_access() ITable_accessContext
	Expression() IExpressionContext
	Table_insert() ITable_insertContext
	KW_PRINT() antlr.TerminalNode
	Method_call() IMethod_callContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Function_callContext) KW_PCALL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_PCALL, 0)
}

func (s *Function_callContext) KW_XPCALL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_XPCALL, 0)
}

func (s *Function_callContext) KW_ERROR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ERROR, 0)
}

func (s *Function_callContext) KW_ASSERT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ASSERT, 0)
}

func (s *Function_callContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_callContext) Table_access() ITable_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_accessContext)
}

func (s *Function_callContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_callContext) Table_insert() ITable_insertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_insertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_insertContext)
}

func (s *Function_callContext) KW_PRINT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_PRINT, 0)
}

func (s *Function_callContext) Method_call() IMethod_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_callContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, lua_grammar_antlr4ParserRULE_function_call)
	var _la int

	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			_la = p.GetTokenStream().LA(1)

			if !(_la == lua_grammar_antlr4ParserKW_PCALL || _la == lua_grammar_antlr4ParserKW_XPCALL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(213)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Expression_list()
		}
		{
			p.SetState(215)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			_la = p.GetTokenStream().LA(1)

			if !(_la == lua_grammar_antlr4ParserKW_ERROR || _la == lua_grammar_antlr4ParserKW_ASSERT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(218)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Expression_list()
		}
		{
			p.SetState(220)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(222)
				p.Identifier()
			}

		case 2:
			{
				p.SetState(223)
				p.Table_access()
			}

		case 3:
			{
				p.SetState(224)
				p.Match(lua_grammar_antlr4ParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(225)
				p.expression(0)
			}
			{
				p.SetState(226)
				p.Match(lua_grammar_antlr4ParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(230)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.Identifier()
			}

		}
		{
			p.SetState(234)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2323857408025166056) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&121230115907) != 0) {
			{
				p.SetState(235)
				p.Expression_list()
			}

		}
		{
			p.SetState(238)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
			p.Table_insert()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.Match(lua_grammar_antlr4ParserKW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Expression_list()
		}
		{
			p.SetState(244)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(246)
			p.Method_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_insertContext is an interface to support dynamic dispatch.
type ITable_insertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsTable_insertContext differentiates from other interfaces.
	IsTable_insertContext()
}

type Table_insertContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_insertContext() *Table_insertContext {
	var p = new(Table_insertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert
	return p
}

func InitEmptyTable_insertContext(p *Table_insertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert
}

func (*Table_insertContext) IsTable_insertContext() {}

func NewTable_insertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_insertContext {
	var p = new(Table_insertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_insert

	return p
}

func (s *Table_insertContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_insertContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_insertContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Table_insertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_insertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_insertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable_insert(s)
	}
}

func (s *Table_insertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable_insert(s)
	}
}

func (s *Table_insertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable_insert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table_insert() (localctx ITable_insertContext) {
	localctx = NewTable_insertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, lua_grammar_antlr4ParserRULE_table_insert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(lua_grammar_antlr4ParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Identifier()
	}
	{
		p.SetState(252)
		p.Match(lua_grammar_antlr4ParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.expression(0)
	}
	{
		p.SetState(254)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FUNCTION() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Block() IBlockContext
	KW_END() antlr.TerminalNode
	KW_LOCAL() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	VARARG() antlr.TerminalNode

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration
	return p
}

func InitEmptyFunction_declarationContext(p *Function_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) KW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FUNCTION, 0)
}

func (s *Function_declarationContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Function_declarationContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_declarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Function_declarationContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *Function_declarationContext) KW_LOCAL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_LOCAL, 0)
}

func (s *Function_declarationContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *Function_declarationContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Function_declarationContext) VARARG() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserVARARG, 0)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (s *Function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunction_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Function_declaration() (localctx IFunction_declarationContext) {
	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, lua_grammar_antlr4ParserRULE_function_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(256)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(259)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(260)
			p.Identifier()
		}
		{
			p.SetState(261)
			p.Match(lua_grammar_antlr4ParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Identifier()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(266)
		p.Identifier()
	}
	{
		p.SetState(267)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserLETTER:
		{
			p.SetState(268)
			p.Parameter()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(269)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(270)
				p.Parameter()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case lua_grammar_antlr4ParserVARARG:
		{
			p.SetState(276)
			p.Match(lua_grammar_antlr4ParserVARARG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case lua_grammar_antlr4ParserT__3:

	default:
	}
	{
		p.SetState(279)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Block()
	}
	{
		p.SetState(281)
		p.Match(lua_grammar_antlr4ParserKW_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParameterContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, lua_grammar_antlr4ParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Identifier()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__0 {
		{
			p.SetState(284)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.expression(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllStatement_terminator() []IStatement_terminatorContext
	Statement_terminator(i int) IStatement_terminatorContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllStatement_terminator() []IStatement_terminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			len++
		}
	}

	tst := make([]IStatement_terminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatement_terminatorContext); ok {
			tst[i] = t.(IStatement_terminatorContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement_terminator(i int) IStatement_terminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2067152211481198360) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&208100495031) != 0) {
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserSEPARATOR {
			{
				p.SetState(288)
				p.Statement_terminator()
			}

			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(294)
			p.Statement()
		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(295)
				p.Statement_terminator()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_IF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllKW_THEN() []antlr.TerminalNode
	KW_THEN(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	KW_END() antlr.TerminalNode
	AllKW_ELSEIF() []antlr.TerminalNode
	KW_ELSEIF(i int) antlr.TerminalNode
	KW_ELSE() antlr.TerminalNode

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_IF, 0)
}

func (s *If_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_statementContext) AllKW_THEN() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserKW_THEN)
}

func (s *If_statementContext) KW_THEN(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_THEN, i)
}

func (s *If_statementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *If_statementContext) AllKW_ELSEIF() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserKW_ELSEIF)
}

func (s *If_statementContext) KW_ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ELSEIF, i)
}

func (s *If_statementContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_ELSE, 0)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, lua_grammar_antlr4ParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(lua_grammar_antlr4ParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.expression(0)
	}
	{
		p.SetState(304)
		p.Match(lua_grammar_antlr4ParserKW_THEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Block()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserKW_ELSEIF {
		{
			p.SetState(306)
			p.Match(lua_grammar_antlr4ParserKW_ELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.expression(0)
		}
		{
			p.SetState(308)
			p.Match(lua_grammar_antlr4ParserKW_THEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Block()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_ELSE {
		{
			p.SetState(316)
			p.Match(lua_grammar_antlr4ParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.Block()
		}

	}
	{
		p.SetState(320)
		p.Match(lua_grammar_antlr4ParserKW_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FOR() antlr.TerminalNode
	Identifier() IIdentifierContext
	KW_IN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	KW_DO() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode

	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement
	return p
}

func InitEmptyFor_statementContext(p *For_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) KW_FOR() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FOR, 0)
}

func (s *For_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *For_statementContext) KW_IN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_IN, 0)
}

func (s *For_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *For_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_statementContext) KW_DO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_DO, 0)
}

func (s *For_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFor_statement(s)
	}
}

func (s *For_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFor_statement(s)
	}
}

func (s *For_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFor_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, lua_grammar_antlr4ParserRULE_for_statement)
	var _la int

	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Identifier()
		}
		{
			p.SetState(324)
			p.Match(lua_grammar_antlr4ParserKW_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.expression(0)
		}
		{
			p.SetState(326)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Block()
		}
		{
			p.SetState(328)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Identifier()
		}
		{
			p.SetState(332)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.expression(0)
		}
		{
			p.SetState(334)
			p.Match(lua_grammar_antlr4ParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.expression(0)
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(336)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(337)
				p.expression(0)
			}

		}
		{
			p.SetState(340)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Block()
		}
		{
			p.SetState(342)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	KW_DO() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) KW_WHILE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_WHILE, 0)
}

func (s *While_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_statementContext) KW_DO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_DO, 0)
}

func (s *While_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *While_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterWhile_statement(s)
	}
}

func (s *While_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitWhile_statement(s)
	}
}

func (s *While_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitWhile_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, lua_grammar_antlr4ParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(lua_grammar_antlr4ParserKW_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.expression(0)
	}
	{
		p.SetState(348)
		p.Match(lua_grammar_antlr4ParserKW_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Block()
	}
	{
		p.SetState(350)
		p.Match(lua_grammar_antlr4ParserKW_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDo_statementContext is an interface to support dynamic dispatch.
type IDo_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_DO() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode

	// IsDo_statementContext differentiates from other interfaces.
	IsDo_statementContext()
}

type Do_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDo_statementContext() *Do_statementContext {
	var p = new(Do_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_do_statement
	return p
}

func InitEmptyDo_statementContext(p *Do_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_do_statement
}

func (*Do_statementContext) IsDo_statementContext() {}

func NewDo_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Do_statementContext {
	var p = new(Do_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_do_statement

	return p
}

func (s *Do_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Do_statementContext) KW_DO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_DO, 0)
}

func (s *Do_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Do_statementContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *Do_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Do_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Do_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterDo_statement(s)
	}
}

func (s *Do_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitDo_statement(s)
	}
}

func (s *Do_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitDo_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Do_statement() (localctx IDo_statementContext) {
	localctx = NewDo_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, lua_grammar_antlr4ParserRULE_do_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(lua_grammar_antlr4ParserKW_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Block()
	}
	{
		p.SetState(354)
		p.Match(lua_grammar_antlr4ParserKW_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllField_separator() []IField_separatorContext
	Field_separator(i int) IField_separatorContext

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table
	return p
}

func InitEmptyTableContext(p *TableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TableContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TableContext) AllField_separator() []IField_separatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_separatorContext); ok {
			len++
		}
	}

	tst := make([]IField_separatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_separatorContext); ok {
			tst[i] = t.(IField_separatorContext)
			i++
		}
	}

	return tst
}

func (s *TableContext) Field_separator(i int) IField_separatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_separatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_separatorContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable(s)
	}
}

func (s *TableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, lua_grammar_antlr4ParserRULE_table)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(lua_grammar_antlr4ParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2323857410172649704) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&104050246723) != 0) {
		{
			p.SetState(357)
			p.Field()
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(358)
					p.Field_separator()
				}
				{
					p.SetState(359)
					p.Field()
				}

			}
			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__25 || _la == lua_grammar_antlr4ParserT__29 {
			{
				p.SetState(366)
				p.Field_separator()
			}

		}

	}
	{
		p.SetState(371)
		p.Match(lua_grammar_antlr4ParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_separatorContext is an interface to support dynamic dispatch.
type IField_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsField_separatorContext differentiates from other interfaces.
	IsField_separatorContext()
}

type Field_separatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_separatorContext() *Field_separatorContext {
	var p = new(Field_separatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field_separator
	return p
}

func InitEmptyField_separatorContext(p *Field_separatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field_separator
}

func (*Field_separatorContext) IsField_separatorContext() {}

func NewField_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_separatorContext {
	var p = new(Field_separatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field_separator

	return p
}

func (s *Field_separatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Field_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterField_separator(s)
	}
}

func (s *Field_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitField_separator(s)
	}
}

func (s *Field_separatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitField_separator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Field_separator() (localctx IField_separatorContext) {
	localctx = NewField_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, lua_grammar_antlr4ParserRULE_field_separator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserT__25 || _la == lua_grammar_antlr4ParserT__29) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, lua_grammar_antlr4ParserRULE_field)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Identifier()
		}
		{
			p.SetState(376)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.Match(lua_grammar_antlr4ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.expression(0)
		}
		{
			p.SetState(381)
			p.Match(lua_grammar_antlr4ParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(385)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_accessContext is an interface to support dynamic dispatch.
type ITable_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Expression() IExpressionContext

	// IsTable_accessContext differentiates from other interfaces.
	IsTable_accessContext()
}

type Table_accessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_accessContext() *Table_accessContext {
	var p = new(Table_accessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access
	return p
}

func InitEmptyTable_accessContext(p *Table_accessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access
}

func (*Table_accessContext) IsTable_accessContext() {}

func NewTable_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_accessContext {
	var p = new(Table_accessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_table_access

	return p
}

func (s *Table_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_accessContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Table_accessContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_accessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Table_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_accessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTable_access(s)
	}
}

func (s *Table_accessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTable_access(s)
	}
}

func (s *Table_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTable_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Table_access() (localctx ITable_accessContext) {
	localctx = NewTable_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, lua_grammar_antlr4ParserRULE_table_access)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Identifier()
		}
		{
			p.SetState(389)
			p.Match(lua_grammar_antlr4ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.expression(0)
		}
		{
			p.SetState(391)
			p.Match(lua_grammar_antlr4ParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Identifier()
		}
		{
			p.SetState(394)
			p.Match(lua_grammar_antlr4ParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Identifier()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_line_commentContext is an interface to support dynamic dispatch.
type ISingle_line_commentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECREMENT() antlr.TerminalNode

	// IsSingle_line_commentContext differentiates from other interfaces.
	IsSingle_line_commentContext()
}

type Single_line_commentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_line_commentContext() *Single_line_commentContext {
	var p = new(Single_line_commentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment
	return p
}

func InitEmptySingle_line_commentContext(p *Single_line_commentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment
}

func (*Single_line_commentContext) IsSingle_line_commentContext() {}

func NewSingle_line_commentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_line_commentContext {
	var p = new(Single_line_commentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_single_line_comment

	return p
}

func (s *Single_line_commentContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_line_commentContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDECREMENT, 0)
}

func (s *Single_line_commentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_line_commentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_line_commentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterSingle_line_comment(s)
	}
}

func (s *Single_line_commentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitSingle_line_comment(s)
	}
}

func (s *Single_line_commentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitSingle_line_comment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Single_line_comment() (localctx ISingle_line_commentContext) {
	localctx = NewSingle_line_commentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, lua_grammar_antlr4ParserRULE_single_line_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(lua_grammar_antlr4ParserDECREMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrint_statementContext is an interface to support dynamic dispatch.
type IPrint_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_PRINT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrint_statementContext differentiates from other interfaces.
	IsPrint_statementContext()
}

type Print_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_statementContext() *Print_statementContext {
	var p = new(Print_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement
	return p
}

func InitEmptyPrint_statementContext(p *Print_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement
}

func (*Print_statementContext) IsPrint_statementContext() {}

func NewPrint_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_statementContext {
	var p = new(Print_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_print_statement

	return p
}

func (s *Print_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_statementContext) KW_PRINT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_PRINT, 0)
}

func (s *Print_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Print_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterPrint_statement(s)
	}
}

func (s *Print_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitPrint_statement(s)
	}
}

func (s *Print_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitPrint_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Print_statement() (localctx IPrint_statementContext) {
	localctx = NewPrint_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, lua_grammar_antlr4ParserRULE_print_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(lua_grammar_antlr4ParserKW_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.expression(0)
	}
	{
		p.SetState(404)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLETTER() []antlr.TerminalNode
	LETTER(i int) antlr.TerminalNode
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserLETTER)
}

func (s *IdentifierContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserLETTER, i)
}

func (s *IdentifierContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserDIGIT)
}

func (s *IdentifierContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserDIGIT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, lua_grammar_antlr4ParserRULE_identifier)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(lua_grammar_antlr4ParserLETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(407)
				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&6917529027641081857) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeat_statementContext is an interface to support dynamic dispatch.
type IRepeat_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_REPEAT() antlr.TerminalNode
	Block() IBlockContext
	KW_UNTIL() antlr.TerminalNode
	Expression() IExpressionContext
	Statement_terminator() IStatement_terminatorContext
	EOF() antlr.TerminalNode

	// IsRepeat_statementContext differentiates from other interfaces.
	IsRepeat_statementContext()
}

type Repeat_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeat_statementContext() *Repeat_statementContext {
	var p = new(Repeat_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeat_statement
	return p
}

func InitEmptyRepeat_statementContext(p *Repeat_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeat_statement
}

func (*Repeat_statementContext) IsRepeat_statementContext() {}

func NewRepeat_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Repeat_statementContext {
	var p = new(Repeat_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeat_statement

	return p
}

func (s *Repeat_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Repeat_statementContext) KW_REPEAT() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_REPEAT, 0)
}

func (s *Repeat_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Repeat_statementContext) KW_UNTIL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_UNTIL, 0)
}

func (s *Repeat_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Repeat_statementContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *Repeat_statementContext) EOF() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserEOF, 0)
}

func (s *Repeat_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Repeat_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Repeat_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterRepeat_statement(s)
	}
}

func (s *Repeat_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitRepeat_statement(s)
	}
}

func (s *Repeat_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitRepeat_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Repeat_statement() (localctx IRepeat_statementContext) {
	localctx = NewRepeat_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, lua_grammar_antlr4ParserRULE_repeat_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(lua_grammar_antlr4ParserKW_REPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Block()
	}
	{
		p.SetState(415)
		p.Match(lua_grammar_antlr4ParserKW_UNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.expression(0)
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserSEPARATOR:
		{
			p.SetState(417)
			p.Statement_terminator()
		}

	case lua_grammar_antlr4ParserEOF:
		{
			p.SetState(418)
			p.Match(lua_grammar_antlr4ParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifier_listContext is an interface to support dynamic dispatch.
type IIdentifier_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsIdentifier_listContext differentiates from other interfaces.
	IsIdentifier_listContext()
}

type Identifier_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_listContext() *Identifier_listContext {
	var p = new(Identifier_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier_list
	return p
}

func InitEmptyIdentifier_listContext(p *Identifier_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier_list
}

func (*Identifier_listContext) IsIdentifier_listContext() {}

func NewIdentifier_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_listContext {
	var p = new(Identifier_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_identifier_list

	return p
}

func (s *Identifier_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_listContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Identifier_listContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Identifier_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIdentifier_list(s)
	}
}

func (s *Identifier_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIdentifier_list(s)
	}
}

func (s *Identifier_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIdentifier_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Identifier_list() (localctx IIdentifier_listContext) {
	localctx = NewIdentifier_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, lua_grammar_antlr4ParserRULE_identifier_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Identifier()
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__25 {
		{
			p.SetState(422)
			p.Match(lua_grammar_antlr4ParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Identifier()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	VARARG() antlr.TerminalNode

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression_list
	return p
}

func InitEmptyExpression_listContext(p *Expression_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression_list
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_listContext) VARARG() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserVARARG, 0)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (s *Expression_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitExpression_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Expression_list() (localctx IExpression_listContext) {
	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, lua_grammar_antlr4ParserRULE_expression_list)
	var _la int

	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__2, lua_grammar_antlr4ParserT__4, lua_grammar_antlr4ParserT__5, lua_grammar_antlr4ParserT__6, lua_grammar_antlr4ParserT__24, lua_grammar_antlr4ParserT__27, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_NOT, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_FUNCTION, lua_grammar_antlr4ParserKW_PCALL, lua_grammar_antlr4ParserKW_XPCALL, lua_grammar_antlr4ParserKW_NAN, lua_grammar_antlr4ParserKW_INF, lua_grammar_antlr4ParserKW_ERROR, lua_grammar_antlr4ParserKW_ASSERT, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER, lua_grammar_antlr4ParserINCREMENT, lua_grammar_antlr4ParserDECREMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.expression(0)
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(430)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(431)
				p.expression(0)
			}

			p.SetState(436)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case lua_grammar_antlr4ParserVARARG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(lua_grammar_antlr4ParserVARARG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_RETURN() antlr.TerminalNode
	Statement_terminator() IStatement_terminatorContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) KW_RETURN() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_RETURN, 0)
}

func (s *Return_statementContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *Return_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Return_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, lua_grammar_antlr4ParserRULE_return_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Match(lua_grammar_antlr4ParserKW_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2323857408025166056) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&104050246723) != 0) {
		{
			p.SetState(441)
			p.expression(0)
		}
		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(442)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(443)
				p.expression(0)
			}

			p.SetState(448)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(451)
		p.Statement_terminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_statementContext is an interface to support dynamic dispatch.
type IBreak_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_BREAK() antlr.TerminalNode
	Statement_terminator() IStatement_terminatorContext

	// IsBreak_statementContext differentiates from other interfaces.
	IsBreak_statementContext()
}

type Break_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_statementContext() *Break_statementContext {
	var p = new(Break_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_break_statement
	return p
}

func InitEmptyBreak_statementContext(p *Break_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_break_statement
}

func (*Break_statementContext) IsBreak_statementContext() {}

func NewBreak_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_statementContext {
	var p = new(Break_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_break_statement

	return p
}

func (s *Break_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_statementContext) KW_BREAK() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_BREAK, 0)
}

func (s *Break_statementContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBreak_statement(s)
	}
}

func (s *Break_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBreak_statement(s)
	}
}

func (s *Break_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBreak_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, lua_grammar_antlr4ParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Match(lua_grammar_antlr4ParserKW_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Statement_terminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoto_statementContext is an interface to support dynamic dispatch.
type IGoto_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_GOTO() antlr.TerminalNode
	Identifier() IIdentifierContext
	Statement_terminator() IStatement_terminatorContext

	// IsGoto_statementContext differentiates from other interfaces.
	IsGoto_statementContext()
}

type Goto_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoto_statementContext() *Goto_statementContext {
	var p = new(Goto_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_goto_statement
	return p
}

func InitEmptyGoto_statementContext(p *Goto_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_goto_statement
}

func (*Goto_statementContext) IsGoto_statementContext() {}

func NewGoto_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Goto_statementContext {
	var p = new(Goto_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_goto_statement

	return p
}

func (s *Goto_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Goto_statementContext) KW_GOTO() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_GOTO, 0)
}

func (s *Goto_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Goto_statementContext) Statement_terminator() IStatement_terminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_terminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_terminatorContext)
}

func (s *Goto_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Goto_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Goto_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterGoto_statement(s)
	}
}

func (s *Goto_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitGoto_statement(s)
	}
}

func (s *Goto_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitGoto_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Goto_statement() (localctx IGoto_statementContext) {
	localctx = NewGoto_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, lua_grammar_antlr4ParserRULE_goto_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(lua_grammar_antlr4ParserKW_GOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Identifier()
	}
	{
		p.SetState(458)
		p.Statement_terminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabel_statementContext is an interface to support dynamic dispatch.
type ILabel_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsLabel_statementContext differentiates from other interfaces.
	IsLabel_statementContext()
}

type Label_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabel_statementContext() *Label_statementContext {
	var p = new(Label_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_label_statement
	return p
}

func InitEmptyLabel_statementContext(p *Label_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_label_statement
}

func (*Label_statementContext) IsLabel_statementContext() {}

func NewLabel_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Label_statementContext {
	var p = new(Label_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_label_statement

	return p
}

func (s *Label_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Label_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Label_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Label_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Label_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLabel_statement(s)
	}
}

func (s *Label_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLabel_statement(s)
	}
}

func (s *Label_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLabel_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Label_statement() (localctx ILabel_statementContext) {
	localctx = NewLabel_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, lua_grammar_antlr4ParserRULE_label_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(lua_grammar_antlr4ParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.Identifier()
	}
	{
		p.SetState(462)
		p.Match(lua_grammar_antlr4ParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_expressionContext is an interface to support dynamic dispatch.
type IFunction_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FUNCTION() antlr.TerminalNode
	Block() IBlockContext
	KW_END() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsFunction_expressionContext differentiates from other interfaces.
	IsFunction_expressionContext()
}

type Function_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_expressionContext() *Function_expressionContext {
	var p = new(Function_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_expression
	return p
}

func InitEmptyFunction_expressionContext(p *Function_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_expression
}

func (*Function_expressionContext) IsFunction_expressionContext() {}

func NewFunction_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_expressionContext {
	var p = new(Function_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_function_expression

	return p
}

func (s *Function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_expressionContext) KW_FUNCTION() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_FUNCTION, 0)
}

func (s *Function_expressionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Function_expressionContext) KW_END() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_END, 0)
}

func (s *Function_expressionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Function_expressionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunction_expression(s)
	}
}

func (s *Function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunction_expression(s)
	}
}

func (s *Function_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunction_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Function_expression() (localctx IFunction_expressionContext) {
	localctx = NewFunction_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, lua_grammar_antlr4ParserRULE_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(466)
			p.Identifier()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(467)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(468)
				p.Identifier()
			}

			p.SetState(473)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(476)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
		p.Block()
	}
	{
		p.SetState(478)
		p.Match(lua_grammar_antlr4ParserKW_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_callContext is an interface to support dynamic dispatch.
type IMethod_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsMethod_callContext differentiates from other interfaces.
	IsMethod_callContext()
}

type Method_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_callContext() *Method_callContext {
	var p = new(Method_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_method_call
	return p
}

func InitEmptyMethod_callContext(p *Method_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_method_call
}

func (*Method_callContext) IsMethod_callContext() {}

func NewMethod_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_callContext {
	var p = new(Method_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_method_call

	return p
}

func (s *Method_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_callContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Method_callContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Method_callContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Method_callContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Method_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMethod_call(s)
	}
}

func (s *Method_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMethod_call(s)
	}
}

func (s *Method_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMethod_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Method_call() (localctx IMethod_callContext) {
	localctx = NewMethod_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, lua_grammar_antlr4ParserRULE_method_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Identifier()
	}
	{
		p.SetState(481)
		p.Match(lua_grammar_antlr4ParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.Identifier()
	}
	{
		p.SetState(483)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2323857408025166056) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&104050246723) != 0) {
		{
			p.SetState(484)
			p.expression(0)
		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__25 {
			{
				p.SetState(485)
				p.Match(lua_grammar_antlr4ParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(486)
				p.expression(0)
			}

			p.SetState(491)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(494)
		p.Match(lua_grammar_antlr4ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetatable_fieldContext is an interface to support dynamic dispatch.
type IMetatable_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Metamethod() IMetamethodContext
	Identifier() IIdentifierContext

	// IsMetatable_fieldContext differentiates from other interfaces.
	IsMetatable_fieldContext()
}

type Metatable_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetatable_fieldContext() *Metatable_fieldContext {
	var p = new(Metatable_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable_field
	return p
}

func InitEmptyMetatable_fieldContext(p *Metatable_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable_field
}

func (*Metatable_fieldContext) IsMetatable_fieldContext() {}

func NewMetatable_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Metatable_fieldContext {
	var p = new(Metatable_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable_field

	return p
}

func (s *Metatable_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Metatable_fieldContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Metatable_fieldContext) Metamethod() IMetamethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetamethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetamethodContext)
}

func (s *Metatable_fieldContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Metatable_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Metatable_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Metatable_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMetatable_field(s)
	}
}

func (s *Metatable_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMetatable_field(s)
	}
}

func (s *Metatable_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMetatable_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Metatable_field() (localctx IMetatable_fieldContext) {
	localctx = NewMetatable_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, lua_grammar_antlr4ParserRULE_metatable_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Match(lua_grammar_antlr4ParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__35, lua_grammar_antlr4ParserT__36, lua_grammar_antlr4ParserT__37, lua_grammar_antlr4ParserT__38, lua_grammar_antlr4ParserT__39, lua_grammar_antlr4ParserT__40, lua_grammar_antlr4ParserT__41, lua_grammar_antlr4ParserT__42, lua_grammar_antlr4ParserT__43, lua_grammar_antlr4ParserT__44, lua_grammar_antlr4ParserT__45, lua_grammar_antlr4ParserT__46, lua_grammar_antlr4ParserT__47, lua_grammar_antlr4ParserT__48, lua_grammar_antlr4ParserT__49, lua_grammar_antlr4ParserT__50:
		{
			p.SetState(497)
			p.Metamethod()
		}

	case lua_grammar_antlr4ParserLETTER:
		{
			p.SetState(498)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(501)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetamethodContext is an interface to support dynamic dispatch.
type IMetamethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMetamethodContext differentiates from other interfaces.
	IsMetamethodContext()
}

type MetamethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetamethodContext() *MetamethodContext {
	var p = new(MetamethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethod
	return p
}

func InitEmptyMetamethodContext(p *MetamethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethod
}

func (*MetamethodContext) IsMetamethodContext() {}

func NewMetamethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetamethodContext {
	var p = new(MetamethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethod

	return p
}

func (s *MetamethodContext) GetParser() antlr.Parser { return s.parser }
func (s *MetamethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetamethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetamethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMetamethod(s)
	}
}

func (s *MetamethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMetamethod(s)
	}
}

func (s *MetamethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMetamethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Metamethod() (localctx IMetamethodContext) {
	localctx = NewMetamethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, lua_grammar_antlr4ParserRULE_metamethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503530907893760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICoroutine_statementContext is an interface to support dynamic dispatch.
type ICoroutine_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_COROUTINE() antlr.TerminalNode
	KW_CREATE() antlr.TerminalNode
	KW_RESUME() antlr.TerminalNode
	KW_YIELD() antlr.TerminalNode
	KW_STATUS() antlr.TerminalNode

	// IsCoroutine_statementContext differentiates from other interfaces.
	IsCoroutine_statementContext()
}

type Coroutine_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoroutine_statementContext() *Coroutine_statementContext {
	var p = new(Coroutine_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutine_statement
	return p
}

func InitEmptyCoroutine_statementContext(p *Coroutine_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutine_statement
}

func (*Coroutine_statementContext) IsCoroutine_statementContext() {}

func NewCoroutine_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Coroutine_statementContext {
	var p = new(Coroutine_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutine_statement

	return p
}

func (s *Coroutine_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Coroutine_statementContext) KW_COROUTINE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_COROUTINE, 0)
}

func (s *Coroutine_statementContext) KW_CREATE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_CREATE, 0)
}

func (s *Coroutine_statementContext) KW_RESUME() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_RESUME, 0)
}

func (s *Coroutine_statementContext) KW_YIELD() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_YIELD, 0)
}

func (s *Coroutine_statementContext) KW_STATUS() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_STATUS, 0)
}

func (s *Coroutine_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Coroutine_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Coroutine_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterCoroutine_statement(s)
	}
}

func (s *Coroutine_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitCoroutine_statement(s)
	}
}

func (s *Coroutine_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitCoroutine_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Coroutine_statement() (localctx ICoroutine_statementContext) {
	localctx = NewCoroutine_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, lua_grammar_antlr4ParserRULE_coroutine_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Match(lua_grammar_antlr4ParserKW_COROUTINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(507)
		p.Match(lua_grammar_antlr4ParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-82)) & ^0x3f) == 0 && ((int64(1)<<(_la-82))&15) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *lua_grammar_antlr4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 7:
		var t *Primary_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Primary_expressionContext)
		}
		return p.Primary_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *lua_grammar_antlr4Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *lua_grammar_antlr4Parser) Primary_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
