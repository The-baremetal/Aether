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
		"", "'='", "'('", "')'", "'#'", "'>'", "'<'", "'>='", "'=='", "'<='",
		"'~='", "'*'", "'/'", "'+'", "'-'", "'//'", "'&'", "'|'", "'~'", "'<<'",
		"'>>'", "':'", "'table.insert'", "','", "'{'", "'}'", "'['", "']'",
		"'.'", "'--'", "'_'", "'::'", "'__'", "", "'in'", "'print'", "'and'",
		"'break'", "'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'",
		"'goto'", "'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'",
		"'then'", "'true'", "'until'", "'while'", "'local'", "'function'", "'index'",
		"'newindex'", "'mode'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "SEPARATOR",
		"KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE", "KW_ELSEIF",
		"KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL", "KW_NOT",
		"KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL",
		"KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "KW_INDEX", "KW_NEWINDEX", "KW_MODE",
		"NUMBER", "STRING", "LETTER", "DIGIT", "WS", "SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "control_statement", "statement_terminator",
		"assignment", "expression", "prefix_expression", "primary_expression",
		"operators", "comparison_operator", "arith_operator", "logical_operator",
		"bitwise_operator", "literal", "function_call", "table_insert", "function_declaration",
		"block", "if_statement", "for_statement", "while_statement", "table",
		"field", "table_access", "single_line_comment", "print_statement", "identifier",
		"repeat_statement", "identifier_list", "expression_list", "return_statement",
		"break_statement", "goto_statement", "label_statement", "function_expression",
		"method_call", "metatable_field",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 407, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 89, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		96, 8, 2, 1, 3, 1, 3, 1, 4, 3, 4, 101, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 119, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 125, 8, 5, 10, 5, 12, 5, 128,
		9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 135, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 146, 8, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 152, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 170,
		8, 14, 1, 14, 1, 14, 3, 14, 174, 8, 14, 1, 14, 1, 14, 3, 14, 178, 8, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 189,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 5, 16, 204, 8, 16, 10, 16, 12, 16, 207, 9, 16,
		3, 16, 209, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 217,
		8, 17, 4, 17, 219, 8, 17, 11, 17, 12, 17, 220, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 232, 8, 18, 10, 18, 12, 18,
		235, 9, 18, 1, 18, 1, 18, 3, 18, 239, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 3, 19, 259, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		3, 19, 265, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 5, 21, 277, 8, 21, 10, 21, 12, 21, 280, 9, 21, 1, 21,
		1, 21, 3, 21, 284, 8, 21, 1, 21, 3, 21, 287, 8, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 296, 8, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 307, 8, 23, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 318, 8, 26, 10,
		26, 12, 26, 321, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 328,
		8, 27, 1, 28, 1, 28, 1, 28, 5, 28, 333, 8, 28, 10, 28, 12, 28, 336, 9,
		28, 1, 29, 1, 29, 1, 29, 5, 29, 341, 8, 29, 10, 29, 12, 29, 344, 9, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 350, 8, 30, 10, 30, 12, 30, 353, 9,
		30, 3, 30, 355, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 5, 34, 375, 8, 34, 10, 34, 12, 34, 378, 9, 34, 3, 34, 380, 8, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		5, 35, 393, 8, 35, 10, 35, 12, 35, 396, 9, 35, 3, 35, 398, 8, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 0, 1, 10, 37, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 0, 7, 1, 0,
		5, 10, 1, 0, 11, 15, 2, 0, 36, 36, 48, 48, 1, 0, 16, 20, 4, 0, 42, 42,
		46, 46, 52, 52, 60, 61, 2, 0, 30, 30, 62, 63, 1, 0, 57, 59, 429, 0, 77,
		1, 0, 0, 0, 2, 88, 1, 0, 0, 0, 4, 95, 1, 0, 0, 0, 6, 97, 1, 0, 0, 0, 8,
		100, 1, 0, 0, 0, 10, 118, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 145, 1,
		0, 0, 0, 16, 151, 1, 0, 0, 0, 18, 153, 1, 0, 0, 0, 20, 155, 1, 0, 0, 0,
		22, 157, 1, 0, 0, 0, 24, 159, 1, 0, 0, 0, 26, 161, 1, 0, 0, 0, 28, 188,
		1, 0, 0, 0, 30, 190, 1, 0, 0, 0, 32, 197, 1, 0, 0, 0, 34, 218, 1, 0, 0,
		0, 36, 222, 1, 0, 0, 0, 38, 264, 1, 0, 0, 0, 40, 266, 1, 0, 0, 0, 42, 272,
		1, 0, 0, 0, 44, 295, 1, 0, 0, 0, 46, 306, 1, 0, 0, 0, 48, 308, 1, 0, 0,
		0, 50, 310, 1, 0, 0, 0, 52, 315, 1, 0, 0, 0, 54, 322, 1, 0, 0, 0, 56, 329,
		1, 0, 0, 0, 58, 337, 1, 0, 0, 0, 60, 345, 1, 0, 0, 0, 62, 358, 1, 0, 0,
		0, 64, 361, 1, 0, 0, 0, 66, 365, 1, 0, 0, 0, 68, 369, 1, 0, 0, 0, 70, 385,
		1, 0, 0, 0, 72, 401, 1, 0, 0, 0, 74, 76, 3, 2, 1, 0, 75, 74, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 1, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 80, 89, 3, 8, 4, 0, 81, 89, 3, 10, 5, 0, 82,
		89, 3, 4, 2, 0, 83, 89, 3, 32, 16, 0, 84, 89, 3, 28, 14, 0, 85, 89, 3,
		60, 30, 0, 86, 89, 3, 62, 31, 0, 87, 89, 3, 66, 33, 0, 88, 80, 1, 0, 0,
		0, 88, 81, 1, 0, 0, 0, 88, 82, 1, 0, 0, 0, 88, 83, 1, 0, 0, 0, 88, 84,
		1, 0, 0, 0, 88, 85, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0,
		89, 3, 1, 0, 0, 0, 90, 96, 3, 36, 18, 0, 91, 96, 3, 38, 19, 0, 92, 96,
		3, 40, 20, 0, 93, 96, 3, 54, 27, 0, 94, 96, 3, 64, 32, 0, 95, 90, 1, 0,
		0, 0, 95, 91, 1, 0, 0, 0, 95, 92, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 94,
		1, 0, 0, 0, 96, 5, 1, 0, 0, 0, 97, 98, 5, 33, 0, 0, 98, 7, 1, 0, 0, 0,
		99, 101, 5, 55, 0, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 103, 3, 56, 28, 0, 103, 104, 5, 1, 0, 0, 104, 105, 3,
		58, 29, 0, 105, 106, 3, 6, 3, 0, 106, 9, 1, 0, 0, 0, 107, 108, 6, 5, -1,
		0, 108, 119, 3, 26, 13, 0, 109, 119, 3, 52, 26, 0, 110, 111, 5, 2, 0, 0,
		111, 112, 3, 10, 5, 0, 112, 113, 5, 3, 0, 0, 113, 119, 1, 0, 0, 0, 114,
		119, 3, 28, 14, 0, 115, 119, 3, 42, 21, 0, 116, 119, 3, 46, 23, 0, 117,
		119, 3, 68, 34, 0, 118, 107, 1, 0, 0, 0, 118, 109, 1, 0, 0, 0, 118, 110,
		1, 0, 0, 0, 118, 114, 1, 0, 0, 0, 118, 115, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 118, 117, 1, 0, 0, 0, 119, 126, 1, 0, 0, 0, 120, 121, 10, 5, 0, 0,
		121, 122, 3, 16, 8, 0, 122, 123, 3, 10, 5, 6, 123, 125, 1, 0, 0, 0, 124,
		120, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 11, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 135, 3, 14,
		7, 0, 130, 131, 5, 47, 0, 0, 131, 135, 3, 12, 6, 0, 132, 133, 5, 4, 0,
		0, 133, 135, 3, 12, 6, 0, 134, 129, 1, 0, 0, 0, 134, 130, 1, 0, 0, 0, 134,
		132, 1, 0, 0, 0, 135, 13, 1, 0, 0, 0, 136, 146, 3, 26, 13, 0, 137, 146,
		3, 52, 26, 0, 138, 139, 5, 2, 0, 0, 139, 140, 3, 10, 5, 0, 140, 141, 5,
		3, 0, 0, 141, 146, 1, 0, 0, 0, 142, 146, 3, 28, 14, 0, 143, 146, 3, 42,
		21, 0, 144, 146, 3, 46, 23, 0, 145, 136, 1, 0, 0, 0, 145, 137, 1, 0, 0,
		0, 145, 138, 1, 0, 0, 0, 145, 142, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145,
		144, 1, 0, 0, 0, 146, 15, 1, 0, 0, 0, 147, 152, 3, 18, 9, 0, 148, 152,
		3, 20, 10, 0, 149, 152, 3, 22, 11, 0, 150, 152, 3, 24, 12, 0, 151, 147,
		1, 0, 0, 0, 151, 148, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 150, 1, 0,
		0, 0, 152, 17, 1, 0, 0, 0, 153, 154, 7, 0, 0, 0, 154, 19, 1, 0, 0, 0, 155,
		156, 7, 1, 0, 0, 156, 21, 1, 0, 0, 0, 157, 158, 7, 2, 0, 0, 158, 23, 1,
		0, 0, 0, 159, 160, 7, 3, 0, 0, 160, 25, 1, 0, 0, 0, 161, 162, 7, 4, 0,
		0, 162, 27, 1, 0, 0, 0, 163, 170, 3, 52, 26, 0, 164, 170, 3, 46, 23, 0,
		165, 166, 5, 2, 0, 0, 166, 167, 3, 10, 5, 0, 167, 168, 5, 3, 0, 0, 168,
		170, 1, 0, 0, 0, 169, 163, 1, 0, 0, 0, 169, 164, 1, 0, 0, 0, 169, 165,
		1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 172, 5, 21, 0, 0, 172, 174, 3, 52,
		26, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0,
		175, 177, 5, 2, 0, 0, 176, 178, 3, 58, 29, 0, 177, 176, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 5, 3, 0, 0, 180, 189,
		1, 0, 0, 0, 181, 189, 3, 30, 15, 0, 182, 183, 5, 35, 0, 0, 183, 184, 5,
		2, 0, 0, 184, 185, 3, 58, 29, 0, 185, 186, 5, 3, 0, 0, 186, 189, 1, 0,
		0, 0, 187, 189, 3, 70, 35, 0, 188, 169, 1, 0, 0, 0, 188, 181, 1, 0, 0,
		0, 188, 182, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 29, 1, 0, 0, 0, 190,
		191, 5, 22, 0, 0, 191, 192, 5, 2, 0, 0, 192, 193, 3, 52, 26, 0, 193, 194,
		5, 23, 0, 0, 194, 195, 3, 10, 5, 0, 195, 196, 5, 3, 0, 0, 196, 31, 1, 0,
		0, 0, 197, 198, 5, 56, 0, 0, 198, 199, 3, 52, 26, 0, 199, 208, 5, 2, 0,
		0, 200, 205, 3, 52, 26, 0, 201, 202, 5, 23, 0, 0, 202, 204, 3, 52, 26,
		0, 203, 201, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205,
		206, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 200,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 5, 3,
		0, 0, 211, 212, 3, 34, 17, 0, 212, 213, 5, 41, 0, 0, 213, 33, 1, 0, 0,
		0, 214, 216, 3, 2, 1, 0, 215, 217, 3, 6, 3, 0, 216, 215, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 214, 1, 0, 0, 0, 219, 220,
		1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 35, 1, 0,
		0, 0, 222, 223, 5, 45, 0, 0, 223, 224, 3, 10, 5, 0, 224, 225, 5, 51, 0,
		0, 225, 233, 3, 34, 17, 0, 226, 227, 5, 40, 0, 0, 227, 228, 3, 10, 5, 0,
		228, 229, 5, 51, 0, 0, 229, 230, 3, 34, 17, 0, 230, 232, 1, 0, 0, 0, 231,
		226, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 238, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 237, 5, 39,
		0, 0, 237, 239, 3, 34, 17, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 241, 5, 41, 0, 0, 241, 37, 1, 0, 0, 0, 242,
		243, 5, 43, 0, 0, 243, 244, 3, 52, 26, 0, 244, 245, 5, 34, 0, 0, 245, 246,
		3, 10, 5, 0, 246, 247, 5, 38, 0, 0, 247, 248, 3, 34, 17, 0, 248, 249, 5,
		41, 0, 0, 249, 265, 1, 0, 0, 0, 250, 251, 5, 43, 0, 0, 251, 252, 3, 52,
		26, 0, 252, 253, 5, 1, 0, 0, 253, 254, 3, 10, 5, 0, 254, 255, 5, 23, 0,
		0, 255, 258, 3, 10, 5, 0, 256, 257, 5, 23, 0, 0, 257, 259, 3, 10, 5, 0,
		258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260,
		261, 5, 38, 0, 0, 261, 262, 3, 34, 17, 0, 262, 263, 5, 41, 0, 0, 263, 265,
		1, 0, 0, 0, 264, 242, 1, 0, 0, 0, 264, 250, 1, 0, 0, 0, 265, 39, 1, 0,
		0, 0, 266, 267, 5, 54, 0, 0, 267, 268, 3, 10, 5, 0, 268, 269, 5, 38, 0,
		0, 269, 270, 3, 34, 17, 0, 270, 271, 5, 41, 0, 0, 271, 41, 1, 0, 0, 0,
		272, 286, 5, 24, 0, 0, 273, 278, 3, 44, 22, 0, 274, 275, 5, 23, 0, 0, 275,
		277, 3, 44, 22, 0, 276, 274, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276,
		1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 283, 1, 0, 0, 0, 280, 278, 1, 0,
		0, 0, 281, 282, 5, 23, 0, 0, 282, 284, 3, 72, 36, 0, 283, 281, 1, 0, 0,
		0, 283, 284, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 287, 3, 72, 36, 0,
		286, 273, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288,
		289, 5, 25, 0, 0, 289, 43, 1, 0, 0, 0, 290, 291, 3, 52, 26, 0, 291, 292,
		5, 1, 0, 0, 292, 293, 3, 10, 5, 0, 293, 296, 1, 0, 0, 0, 294, 296, 3, 10,
		5, 0, 295, 290, 1, 0, 0, 0, 295, 294, 1, 0, 0, 0, 296, 45, 1, 0, 0, 0,
		297, 298, 3, 52, 26, 0, 298, 299, 5, 26, 0, 0, 299, 300, 3, 10, 5, 0, 300,
		301, 5, 27, 0, 0, 301, 307, 1, 0, 0, 0, 302, 303, 3, 52, 26, 0, 303, 304,
		5, 28, 0, 0, 304, 305, 3, 52, 26, 0, 305, 307, 1, 0, 0, 0, 306, 297, 1,
		0, 0, 0, 306, 302, 1, 0, 0, 0, 307, 47, 1, 0, 0, 0, 308, 309, 5, 29, 0,
		0, 309, 49, 1, 0, 0, 0, 310, 311, 5, 35, 0, 0, 311, 312, 5, 2, 0, 0, 312,
		313, 3, 10, 5, 0, 313, 314, 5, 3, 0, 0, 314, 51, 1, 0, 0, 0, 315, 319,
		5, 62, 0, 0, 316, 318, 7, 5, 0, 0, 317, 316, 1, 0, 0, 0, 318, 321, 1, 0,
		0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 53, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 322, 323, 5, 49, 0, 0, 323, 324, 3, 34, 17, 0, 324,
		325, 5, 53, 0, 0, 325, 327, 3, 10, 5, 0, 326, 328, 3, 6, 3, 0, 327, 326,
		1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 55, 1, 0, 0, 0, 329, 334, 3, 52,
		26, 0, 330, 331, 5, 23, 0, 0, 331, 333, 3, 52, 26, 0, 332, 330, 1, 0, 0,
		0, 333, 336, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335,
		57, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 342, 3, 10, 5, 0, 338, 339,
		5, 23, 0, 0, 339, 341, 3, 10, 5, 0, 340, 338, 1, 0, 0, 0, 341, 344, 1,
		0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 59, 1, 0, 0,
		0, 344, 342, 1, 0, 0, 0, 345, 354, 5, 50, 0, 0, 346, 351, 3, 10, 5, 0,
		347, 348, 5, 23, 0, 0, 348, 350, 3, 10, 5, 0, 349, 347, 1, 0, 0, 0, 350,
		353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 355,
		1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 354, 346, 1, 0, 0, 0, 354, 355, 1, 0,
		0, 0, 355, 356, 1, 0, 0, 0, 356, 357, 3, 6, 3, 0, 357, 61, 1, 0, 0, 0,
		358, 359, 5, 37, 0, 0, 359, 360, 3, 6, 3, 0, 360, 63, 1, 0, 0, 0, 361,
		362, 5, 44, 0, 0, 362, 363, 3, 52, 26, 0, 363, 364, 3, 6, 3, 0, 364, 65,
		1, 0, 0, 0, 365, 366, 5, 31, 0, 0, 366, 367, 3, 52, 26, 0, 367, 368, 5,
		31, 0, 0, 368, 67, 1, 0, 0, 0, 369, 370, 5, 56, 0, 0, 370, 379, 5, 2, 0,
		0, 371, 376, 3, 52, 26, 0, 372, 373, 5, 23, 0, 0, 373, 375, 3, 52, 26,
		0, 374, 372, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376,
		377, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 379, 371,
		1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 5, 3,
		0, 0, 382, 383, 3, 34, 17, 0, 383, 384, 5, 41, 0, 0, 384, 69, 1, 0, 0,
		0, 385, 386, 3, 52, 26, 0, 386, 387, 5, 21, 0, 0, 387, 388, 3, 52, 26,
		0, 388, 397, 5, 2, 0, 0, 389, 394, 3, 10, 5, 0, 390, 391, 5, 23, 0, 0,
		391, 393, 3, 10, 5, 0, 392, 390, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394,
		392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394,
		1, 0, 0, 0, 397, 389, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 1, 0,
		0, 0, 399, 400, 5, 3, 0, 0, 400, 71, 1, 0, 0, 0, 401, 402, 5, 32, 0, 0,
		402, 403, 7, 6, 0, 0, 403, 404, 5, 1, 0, 0, 404, 405, 3, 10, 5, 0, 405,
		73, 1, 0, 0, 0, 36, 77, 88, 95, 100, 118, 126, 134, 145, 151, 169, 173,
		177, 188, 205, 208, 216, 220, 233, 238, 258, 264, 278, 283, 286, 295, 306,
		319, 327, 334, 342, 351, 354, 376, 379, 394, 397,
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
	lua_grammar_antlr4ParserSEPARATOR           = 33
	lua_grammar_antlr4ParserKW_IN               = 34
	lua_grammar_antlr4ParserKW_PRINT            = 35
	lua_grammar_antlr4ParserKW_AND              = 36
	lua_grammar_antlr4ParserKW_BREAK            = 37
	lua_grammar_antlr4ParserKW_DO               = 38
	lua_grammar_antlr4ParserKW_ELSE             = 39
	lua_grammar_antlr4ParserKW_ELSEIF           = 40
	lua_grammar_antlr4ParserKW_END              = 41
	lua_grammar_antlr4ParserKW_FALSE            = 42
	lua_grammar_antlr4ParserKW_FOR              = 43
	lua_grammar_antlr4ParserKW_GOTO             = 44
	lua_grammar_antlr4ParserKW_IF               = 45
	lua_grammar_antlr4ParserKW_NIL              = 46
	lua_grammar_antlr4ParserKW_NOT              = 47
	lua_grammar_antlr4ParserKW_OR               = 48
	lua_grammar_antlr4ParserKW_REPEAT           = 49
	lua_grammar_antlr4ParserKW_RETURN           = 50
	lua_grammar_antlr4ParserKW_THEN             = 51
	lua_grammar_antlr4ParserKW_TRUE             = 52
	lua_grammar_antlr4ParserKW_UNTIL            = 53
	lua_grammar_antlr4ParserKW_WHILE            = 54
	lua_grammar_antlr4ParserKW_LOCAL            = 55
	lua_grammar_antlr4ParserKW_FUNCTION         = 56
	lua_grammar_antlr4ParserKW_INDEX            = 57
	lua_grammar_antlr4ParserKW_NEWINDEX         = 58
	lua_grammar_antlr4ParserKW_MODE             = 59
	lua_grammar_antlr4ParserNUMBER              = 60
	lua_grammar_antlr4ParserSTRING              = 61
	lua_grammar_antlr4ParserLETTER              = 62
	lua_grammar_antlr4ParserDIGIT               = 63
	lua_grammar_antlr4ParserWS                  = 64
	lua_grammar_antlr4ParserSINGLE_LINE_COMMENT = 65
	lua_grammar_antlr4ParserMULTI_LINE_COMMENT  = 66
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
	lua_grammar_antlr4ParserRULE_literal              = 13
	lua_grammar_antlr4ParserRULE_function_call        = 14
	lua_grammar_antlr4ParserRULE_table_insert         = 15
	lua_grammar_antlr4ParserRULE_function_declaration = 16
	lua_grammar_antlr4ParserRULE_block                = 17
	lua_grammar_antlr4ParserRULE_if_statement         = 18
	lua_grammar_antlr4ParserRULE_for_statement        = 19
	lua_grammar_antlr4ParserRULE_while_statement      = 20
	lua_grammar_antlr4ParserRULE_table                = 21
	lua_grammar_antlr4ParserRULE_field                = 22
	lua_grammar_antlr4ParserRULE_table_access         = 23
	lua_grammar_antlr4ParserRULE_single_line_comment  = 24
	lua_grammar_antlr4ParserRULE_print_statement      = 25
	lua_grammar_antlr4ParserRULE_identifier           = 26
	lua_grammar_antlr4ParserRULE_repeat_statement     = 27
	lua_grammar_antlr4ParserRULE_identifier_list      = 28
	lua_grammar_antlr4ParserRULE_expression_list      = 29
	lua_grammar_antlr4ParserRULE_return_statement     = 30
	lua_grammar_antlr4ParserRULE_break_statement      = 31
	lua_grammar_antlr4ParserRULE_goto_statement       = 32
	lua_grammar_antlr4ParserRULE_label_statement      = 33
	lua_grammar_antlr4ParserRULE_function_expression  = 34
	lua_grammar_antlr4ParserRULE_method_call          = 35
	lua_grammar_antlr4ParserRULE_metatable_field      = 36
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8202880284710928388) != 0 {
		{
			p.SetState(74)
			p.Statement()
		}

		p.SetState(79)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Function_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.Function_call()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
			p.Return_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(86)
			p.Break_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(87)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserKW_IF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.If_statement()
		}

	case lua_grammar_antlr4ParserKW_FOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.For_statement()
		}

	case lua_grammar_antlr4ParserKW_WHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.While_statement()
		}

	case lua_grammar_antlr4ParserKW_REPEAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Repeat_statement()
		}

	case lua_grammar_antlr4ParserKW_GOTO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Goto_statement()
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
		p.SetState(97)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(99)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(102)
		p.Identifier_list()
	}
	{
		p.SetState(103)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Expression_list()
	}
	{
		p.SetState(105)
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
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Function_call() IFunction_callContext
	Table() ITableContext
	Table_access() ITable_accessContext
	Function_expression() IFunction_expressionContext
	Operators() IOperatorsContext

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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(108)
			p.Literal()
		}

	case 2:
		{
			p.SetState(109)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(110)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.expression(0)
		}
		{
			p.SetState(112)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(114)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(115)
			p.Table()
		}

	case 6:
		{
			p.SetState(116)
			p.Table_access()
		}

	case 7:
		{
			p.SetState(117)
			p.Function_expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
			p.SetState(120)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(121)
				p.Operators()
			}
			{
				p.SetState(122)
				p.expression(6)
			}

		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__21, lua_grammar_antlr4ParserT__23, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Primary_expression()
		}

	case lua_grammar_antlr4ParserKW_NOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(lua_grammar_antlr4ParserKW_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
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
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lua_grammar_antlr4ParserRULE_primary_expression)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.expression(0)
		}
		{
			p.SetState(140)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.Function_call()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(143)
			p.Table()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(144)
			p.Table_access()
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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__4, lua_grammar_antlr4ParserT__5, lua_grammar_antlr4ParserT__6, lua_grammar_antlr4ParserT__7, lua_grammar_antlr4ParserT__8, lua_grammar_antlr4ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Comparison_operator()
		}

	case lua_grammar_antlr4ParserT__10, lua_grammar_antlr4ParserT__11, lua_grammar_antlr4ParserT__12, lua_grammar_antlr4ParserT__13, lua_grammar_antlr4ParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Arith_operator()
		}

	case lua_grammar_antlr4ParserKW_AND, lua_grammar_antlr4ParserKW_OR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Logical_operator()
		}

	case lua_grammar_antlr4ParserT__15, lua_grammar_antlr4ParserT__16, lua_grammar_antlr4ParserT__17, lua_grammar_antlr4ParserT__18, lua_grammar_antlr4ParserT__19:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(150)
			p.Bitwise_operator()
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
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2016) != 0) {
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
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63488) != 0) {
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
		p.SetState(157)
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
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2031616) != 0) {
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
	p.EnterRule(localctx, 26, lua_grammar_antlr4ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3463342880238600192) != 0) {
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
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Table_access() ITable_accessContext
	Expression() IExpressionContext
	Expression_list() IExpression_listContext
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
	p.EnterRule(localctx, 28, lua_grammar_antlr4ParserRULE_function_call)
	var _la int

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(163)
				p.Identifier()
			}

		case 2:
			{
				p.SetState(164)
				p.Table_access()
			}

		case 3:
			{
				p.SetState(165)
				p.Match(lua_grammar_antlr4ParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(166)
				p.expression(0)
			}
			{
				p.SetState(167)
				p.Match(lua_grammar_antlr4ParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__20 {
			{
				p.SetState(171)
				p.Match(lua_grammar_antlr4ParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(172)
				p.Identifier()
			}

		}
		{
			p.SetState(175)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8147086527084625924) != 0 {
			{
				p.SetState(176)
				p.Expression_list()
			}

		}
		{
			p.SetState(179)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Table_insert()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(lua_grammar_antlr4ParserKW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Expression_list()
		}
		{
			p.SetState(185)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(187)
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
	p.EnterRule(localctx, 30, lua_grammar_antlr4ParserRULE_table_insert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(lua_grammar_antlr4ParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Identifier()
	}
	{
		p.SetState(193)
		p.Match(lua_grammar_antlr4ParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.expression(0)
	}
	{
		p.SetState(195)
		p.Match(lua_grammar_antlr4ParserT__2)
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
	p.EnterRule(localctx, 32, lua_grammar_antlr4ParserRULE_function_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Identifier()
	}
	{
		p.SetState(199)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(200)
			p.Identifier()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(201)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(202)
				p.Identifier()
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(210)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Block()
	}
	{
		p.SetState(212)
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
	p.EnterRule(localctx, 34, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8202880284710928388) != 0) {
		{
			p.SetState(214)
			p.Statement()
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserSEPARATOR {
			{
				p.SetState(215)
				p.Statement_terminator()
			}

		}

		p.SetState(220)
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
	p.EnterRule(localctx, 36, lua_grammar_antlr4ParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(lua_grammar_antlr4ParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.expression(0)
	}
	{
		p.SetState(224)
		p.Match(lua_grammar_antlr4ParserKW_THEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Block()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserKW_ELSEIF {
		{
			p.SetState(226)
			p.Match(lua_grammar_antlr4ParserKW_ELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.expression(0)
		}
		{
			p.SetState(228)
			p.Match(lua_grammar_antlr4ParserKW_THEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Block()
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_ELSE {
		{
			p.SetState(236)
			p.Match(lua_grammar_antlr4ParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Block()
		}

	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 38, lua_grammar_antlr4ParserRULE_for_statement)
	var _la int

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Identifier()
		}
		{
			p.SetState(244)
			p.Match(lua_grammar_antlr4ParserKW_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.expression(0)
		}
		{
			p.SetState(246)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Block()
		}
		{
			p.SetState(248)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
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
			p.Match(lua_grammar_antlr4ParserT__0)
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
			p.Match(lua_grammar_antlr4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.expression(0)
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(256)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(257)
				p.expression(0)
			}

		}
		{
			p.SetState(260)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Block()
		}
		{
			p.SetState(262)
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
	p.EnterRule(localctx, 40, lua_grammar_antlr4ParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(lua_grammar_antlr4ParserKW_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.expression(0)
	}
	{
		p.SetState(268)
		p.Match(lua_grammar_antlr4ParserKW_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Block()
	}
	{
		p.SetState(270)
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
	Metatable_field() IMetatable_fieldContext

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

func (s *TableContext) Metatable_field() IMetatable_fieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetatable_fieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetatable_fieldContext)
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
	p.EnterRule(localctx, 42, lua_grammar_antlr4ParserRULE_table)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(lua_grammar_antlr4ParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__21, lua_grammar_antlr4ParserT__23, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_FUNCTION, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		{
			p.SetState(273)
			p.Field()
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(274)
					p.Match(lua_grammar_antlr4ParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(275)
					p.Field()
				}

			}
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(281)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(282)
				p.Metatable_field()
			}

		}

	case lua_grammar_antlr4ParserT__31:
		{
			p.SetState(285)
			p.Metatable_field()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(288)
		p.Match(lua_grammar_antlr4ParserT__24)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

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

func (s *FieldContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 44, lua_grammar_antlr4ParserRULE_field)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Identifier()
		}
		{
			p.SetState(291)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
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
	p.EnterRule(localctx, 46, lua_grammar_antlr4ParserRULE_table_access)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Identifier()
		}
		{
			p.SetState(298)
			p.Match(lua_grammar_antlr4ParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.expression(0)
		}
		{
			p.SetState(300)
			p.Match(lua_grammar_antlr4ParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.Identifier()
		}
		{
			p.SetState(303)
			p.Match(lua_grammar_antlr4ParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
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
	p.EnterRule(localctx, 48, lua_grammar_antlr4ParserRULE_single_line_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
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
	p.EnterRule(localctx, 50, lua_grammar_antlr4ParserRULE_print_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(lua_grammar_antlr4ParserKW_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.expression(0)
	}
	{
		p.SetState(313)
		p.Match(lua_grammar_antlr4ParserT__2)
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
	p.EnterRule(localctx, 52, lua_grammar_antlr4ParserRULE_identifier)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(lua_grammar_antlr4ParserLETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(316)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4611686017353646080) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 54, lua_grammar_antlr4ParserRULE_repeat_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(lua_grammar_antlr4ParserKW_REPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Block()
	}
	{
		p.SetState(324)
		p.Match(lua_grammar_antlr4ParserKW_UNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.expression(0)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(326)
			p.Statement_terminator()
		}

	} else if p.HasError() { // JIM
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
	p.EnterRule(localctx, 56, lua_grammar_antlr4ParserRULE_identifier_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Identifier()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__22 {
		{
			p.SetState(330)
			p.Match(lua_grammar_antlr4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Identifier()
		}

		p.SetState(336)
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
	p.EnterRule(localctx, 58, lua_grammar_antlr4ParserRULE_expression_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.expression(0)
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__22 {
		{
			p.SetState(338)
			p.Match(lua_grammar_antlr4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.expression(0)
		}

		p.SetState(344)
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
	p.EnterRule(localctx, 60, lua_grammar_antlr4ParserRULE_return_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(lua_grammar_antlr4ParserKW_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8147086527084625924) != 0 {
		{
			p.SetState(346)
			p.expression(0)
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(347)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(348)
				p.expression(0)
			}

			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(356)
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
	p.EnterRule(localctx, 62, lua_grammar_antlr4ParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(lua_grammar_antlr4ParserKW_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
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
	p.EnterRule(localctx, 64, lua_grammar_antlr4ParserRULE_goto_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(lua_grammar_antlr4ParserKW_GOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.Identifier()
	}
	{
		p.SetState(363)
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
	p.EnterRule(localctx, 66, lua_grammar_antlr4ParserRULE_label_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(lua_grammar_antlr4ParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Identifier()
	}
	{
		p.SetState(367)
		p.Match(lua_grammar_antlr4ParserT__30)
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
	p.EnterRule(localctx, 68, lua_grammar_antlr4ParserRULE_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(371)
			p.Identifier()
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(372)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(373)
				p.Identifier()
			}

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(381)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Block()
	}
	{
		p.SetState(383)
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
	p.EnterRule(localctx, 70, lua_grammar_antlr4ParserRULE_method_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Identifier()
	}
	{
		p.SetState(386)
		p.Match(lua_grammar_antlr4ParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Identifier()
	}
	{
		p.SetState(388)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8147086527084625924) != 0 {
		{
			p.SetState(389)
			p.expression(0)
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__22 {
			{
				p.SetState(390)
				p.Match(lua_grammar_antlr4ParserT__22)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(391)
				p.expression(0)
			}

			p.SetState(396)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(399)
		p.Match(lua_grammar_antlr4ParserT__2)
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
	KW_INDEX() antlr.TerminalNode
	KW_NEWINDEX() antlr.TerminalNode
	KW_MODE() antlr.TerminalNode

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

func (s *Metatable_fieldContext) KW_INDEX() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_INDEX, 0)
}

func (s *Metatable_fieldContext) KW_NEWINDEX() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_NEWINDEX, 0)
}

func (s *Metatable_fieldContext) KW_MODE() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserKW_MODE, 0)
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
	p.EnterRule(localctx, 72, lua_grammar_antlr4ParserRULE_metatable_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(lua_grammar_antlr4ParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008806316530991104) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(403)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
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

func (p *lua_grammar_antlr4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *lua_grammar_antlr4Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
