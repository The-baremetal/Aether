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
		"'>>'", "'..'", "':'", "'table.insert'", "','", "'{'", "'}'", "'['",
		"']'", "'.'", "'--'", "'_'", "'::'", "'__'", "'__add'", "'__sub'", "'__mul'",
		"'__div'", "'__mod'", "'__pow'", "'__unm'", "'__concat'", "'__len'",
		"'__eq'", "'__lt'", "'__le'", "", "'in'", "'print'", "'and'", "'break'",
		"'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'", "'goto'",
		"'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'",
		"'true'", "'until'", "'while'", "'local'", "'function'", "'index'",
		"'newindex'", "'mode'", "'pcall'", "'xpcall'", "'coroutine'", "'create'",
		"'resume'", "'yield'", "'status'", "'nan'", "'inf'", "", "", "", "",
		"", "", "", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "SEPARATOR", "KW_IN",
		"KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE", "KW_ELSEIF", "KW_END",
		"KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL", "KW_NOT", "KW_OR",
		"KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL", "KW_WHILE",
		"KW_LOCAL", "KW_FUNCTION", "KW_INDEX", "KW_NEWINDEX", "KW_MODE", "KW_PCALL",
		"KW_XPCALL", "KW_COROUTINE", "KW_CREATE", "KW_RESUME", "KW_YIELD", "KW_STATUS",
		"KW_NAN", "KW_INF", "NUMBER", "STRING", "LETTER", "DIGIT", "WS", "SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT", "VARARG",
	}
	staticData.RuleNames = []string{
		"program", "statement", "control_statement", "statement_terminator",
		"assignment", "expression", "prefix_expression", "primary_expression",
		"operators", "comparison_operator", "arith_operator", "logical_operator",
		"bitwise_operator", "concat_operator", "literal", "function_call", "table_insert",
		"function_declaration", "block", "if_statement", "for_statement", "while_statement",
		"table", "field", "table_access", "single_line_comment", "print_statement",
		"identifier", "repeat_statement", "identifier_list", "expression_list",
		"return_statement", "break_statement", "goto_statement", "label_statement",
		"function_expression", "method_call", "metatable_field", "metamethod",
		"coroutine_statement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 89, 442, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 5, 0, 82, 8, 0, 10, 0,
		12, 0, 85, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		95, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 3, 1,
		3, 1, 4, 3, 4, 108, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 126, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 132, 8, 5, 10, 5, 12, 5, 135, 9, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 142, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 153, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 160, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 185, 8, 15, 1, 15, 1, 15, 3, 15, 189,
		8, 15, 1, 15, 1, 15, 3, 15, 193, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 204, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 3, 17, 214, 8, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 5, 17, 222, 8, 17, 10, 17, 12, 17, 225, 9, 17,
		1, 17, 3, 17, 228, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 5, 18, 235,
		8, 18, 10, 18, 12, 18, 238, 9, 18, 1, 18, 1, 18, 3, 18, 242, 8, 18, 4,
		18, 244, 8, 18, 11, 18, 12, 18, 245, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 257, 8, 19, 10, 19, 12, 19, 260, 9,
		19, 1, 19, 1, 19, 3, 19, 264, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 284, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20,
		290, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 5, 22, 302, 8, 22, 10, 22, 12, 22, 305, 9, 22, 1, 22, 1, 22,
		3, 22, 309, 8, 22, 1, 22, 3, 22, 312, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 321, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 332, 8, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 343, 8, 27, 10, 27, 12,
		27, 346, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 354, 8,
		28, 1, 29, 1, 29, 1, 29, 5, 29, 359, 8, 29, 10, 29, 12, 29, 362, 9, 29,
		1, 30, 1, 30, 1, 30, 5, 30, 367, 8, 30, 10, 30, 12, 30, 370, 9, 30, 1,
		30, 3, 30, 373, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 379, 8, 31, 10,
		31, 12, 31, 382, 9, 31, 3, 31, 384, 8, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 5, 35, 404, 8, 35, 10, 35, 12, 35, 407, 9, 35,
		3, 35, 409, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 5, 36, 422, 8, 36, 10, 36, 12, 36, 425, 9, 36,
		3, 36, 427, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 0, 1, 10, 40, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		0, 10, 1, 0, 5, 10, 1, 0, 11, 15, 2, 0, 49, 49, 61, 61, 1, 0, 16, 20, 4,
		0, 55, 55, 59, 59, 65, 65, 80, 83, 1, 0, 73, 74, 2, 0, 31, 31, 84, 85,
		1, 0, 70, 72, 1, 0, 34, 45, 1, 0, 76, 79, 468, 0, 83, 1, 0, 0, 0, 2, 94,
		1, 0, 0, 0, 4, 102, 1, 0, 0, 0, 6, 104, 1, 0, 0, 0, 8, 107, 1, 0, 0, 0,
		10, 125, 1, 0, 0, 0, 12, 141, 1, 0, 0, 0, 14, 152, 1, 0, 0, 0, 16, 159,
		1, 0, 0, 0, 18, 161, 1, 0, 0, 0, 20, 163, 1, 0, 0, 0, 22, 165, 1, 0, 0,
		0, 24, 167, 1, 0, 0, 0, 26, 169, 1, 0, 0, 0, 28, 171, 1, 0, 0, 0, 30, 203,
		1, 0, 0, 0, 32, 205, 1, 0, 0, 0, 34, 213, 1, 0, 0, 0, 36, 243, 1, 0, 0,
		0, 38, 247, 1, 0, 0, 0, 40, 289, 1, 0, 0, 0, 42, 291, 1, 0, 0, 0, 44, 297,
		1, 0, 0, 0, 46, 320, 1, 0, 0, 0, 48, 331, 1, 0, 0, 0, 50, 333, 1, 0, 0,
		0, 52, 335, 1, 0, 0, 0, 54, 340, 1, 0, 0, 0, 56, 347, 1, 0, 0, 0, 58, 355,
		1, 0, 0, 0, 60, 372, 1, 0, 0, 0, 62, 374, 1, 0, 0, 0, 64, 387, 1, 0, 0,
		0, 66, 390, 1, 0, 0, 0, 68, 394, 1, 0, 0, 0, 70, 398, 1, 0, 0, 0, 72, 414,
		1, 0, 0, 0, 74, 430, 1, 0, 0, 0, 76, 435, 1, 0, 0, 0, 78, 437, 1, 0, 0,
		0, 80, 82, 3, 2, 1, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 1, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		86, 95, 3, 8, 4, 0, 87, 95, 3, 10, 5, 0, 88, 95, 3, 4, 2, 0, 89, 95, 3,
		34, 17, 0, 90, 95, 3, 30, 15, 0, 91, 95, 3, 62, 31, 0, 92, 95, 3, 64, 32,
		0, 93, 95, 3, 68, 34, 0, 94, 86, 1, 0, 0, 0, 94, 87, 1, 0, 0, 0, 94, 88,
		1, 0, 0, 0, 94, 89, 1, 0, 0, 0, 94, 90, 1, 0, 0, 0, 94, 91, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 3, 1, 0, 0, 0, 96, 103, 3,
		38, 19, 0, 97, 103, 3, 40, 20, 0, 98, 103, 3, 42, 21, 0, 99, 103, 3, 56,
		28, 0, 100, 103, 3, 66, 33, 0, 101, 103, 3, 78, 39, 0, 102, 96, 1, 0, 0,
		0, 102, 97, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 102,
		100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 5, 1, 0, 0, 0, 104, 105, 5,
		46, 0, 0, 105, 7, 1, 0, 0, 0, 106, 108, 5, 68, 0, 0, 107, 106, 1, 0, 0,
		0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 3, 58, 29, 0,
		110, 111, 5, 1, 0, 0, 111, 112, 3, 60, 30, 0, 112, 113, 3, 6, 3, 0, 113,
		9, 1, 0, 0, 0, 114, 115, 6, 5, -1, 0, 115, 126, 3, 28, 14, 0, 116, 126,
		3, 54, 27, 0, 117, 118, 5, 2, 0, 0, 118, 119, 3, 10, 5, 0, 119, 120, 5,
		3, 0, 0, 120, 126, 1, 0, 0, 0, 121, 126, 3, 30, 15, 0, 122, 126, 3, 44,
		22, 0, 123, 126, 3, 48, 24, 0, 124, 126, 3, 70, 35, 0, 125, 114, 1, 0,
		0, 0, 125, 116, 1, 0, 0, 0, 125, 117, 1, 0, 0, 0, 125, 121, 1, 0, 0, 0,
		125, 122, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126,
		133, 1, 0, 0, 0, 127, 128, 10, 5, 0, 0, 128, 129, 3, 16, 8, 0, 129, 130,
		3, 10, 5, 6, 130, 132, 1, 0, 0, 0, 131, 127, 1, 0, 0, 0, 132, 135, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 11, 1, 0, 0, 0,
		135, 133, 1, 0, 0, 0, 136, 142, 3, 14, 7, 0, 137, 138, 5, 60, 0, 0, 138,
		142, 3, 12, 6, 0, 139, 140, 5, 4, 0, 0, 140, 142, 3, 12, 6, 0, 141, 136,
		1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 13, 1, 0,
		0, 0, 143, 153, 3, 28, 14, 0, 144, 153, 3, 54, 27, 0, 145, 146, 5, 2, 0,
		0, 146, 147, 3, 10, 5, 0, 147, 148, 5, 3, 0, 0, 148, 153, 1, 0, 0, 0, 149,
		153, 3, 30, 15, 0, 150, 153, 3, 44, 22, 0, 151, 153, 3, 48, 24, 0, 152,
		143, 1, 0, 0, 0, 152, 144, 1, 0, 0, 0, 152, 145, 1, 0, 0, 0, 152, 149,
		1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 15, 1, 0,
		0, 0, 154, 160, 3, 18, 9, 0, 155, 160, 3, 20, 10, 0, 156, 160, 3, 22, 11,
		0, 157, 160, 3, 24, 12, 0, 158, 160, 3, 26, 13, 0, 159, 154, 1, 0, 0, 0,
		159, 155, 1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		158, 1, 0, 0, 0, 160, 17, 1, 0, 0, 0, 161, 162, 7, 0, 0, 0, 162, 19, 1,
		0, 0, 0, 163, 164, 7, 1, 0, 0, 164, 21, 1, 0, 0, 0, 165, 166, 7, 2, 0,
		0, 166, 23, 1, 0, 0, 0, 167, 168, 7, 3, 0, 0, 168, 25, 1, 0, 0, 0, 169,
		170, 5, 21, 0, 0, 170, 27, 1, 0, 0, 0, 171, 172, 7, 4, 0, 0, 172, 29, 1,
		0, 0, 0, 173, 174, 7, 5, 0, 0, 174, 175, 5, 2, 0, 0, 175, 176, 3, 60, 30,
		0, 176, 177, 5, 3, 0, 0, 177, 204, 1, 0, 0, 0, 178, 185, 3, 54, 27, 0,
		179, 185, 3, 48, 24, 0, 180, 181, 5, 2, 0, 0, 181, 182, 3, 10, 5, 0, 182,
		183, 5, 3, 0, 0, 183, 185, 1, 0, 0, 0, 184, 178, 1, 0, 0, 0, 184, 179,
		1, 0, 0, 0, 184, 180, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 187, 5, 22,
		0, 0, 187, 189, 3, 54, 27, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0,
		0, 189, 190, 1, 0, 0, 0, 190, 192, 5, 2, 0, 0, 191, 193, 3, 60, 30, 0,
		192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194,
		195, 5, 3, 0, 0, 195, 204, 1, 0, 0, 0, 196, 204, 3, 32, 16, 0, 197, 198,
		5, 48, 0, 0, 198, 199, 5, 2, 0, 0, 199, 200, 3, 60, 30, 0, 200, 201, 5,
		3, 0, 0, 201, 204, 1, 0, 0, 0, 202, 204, 3, 72, 36, 0, 203, 173, 1, 0,
		0, 0, 203, 184, 1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 203, 197, 1, 0, 0, 0,
		203, 202, 1, 0, 0, 0, 204, 31, 1, 0, 0, 0, 205, 206, 5, 23, 0, 0, 206,
		207, 5, 2, 0, 0, 207, 208, 3, 54, 27, 0, 208, 209, 5, 24, 0, 0, 209, 210,
		3, 10, 5, 0, 210, 211, 5, 3, 0, 0, 211, 33, 1, 0, 0, 0, 212, 214, 5, 68,
		0, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 216, 5, 69, 0, 0, 216, 217, 3, 54, 27, 0, 217, 227, 5, 2, 0, 0, 218,
		223, 3, 54, 27, 0, 219, 220, 5, 24, 0, 0, 220, 222, 3, 54, 27, 0, 221,
		219, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 228, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 228, 5, 89,
		0, 0, 227, 218, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0,
		228, 229, 1, 0, 0, 0, 229, 230, 5, 3, 0, 0, 230, 231, 3, 36, 18, 0, 231,
		232, 5, 54, 0, 0, 232, 35, 1, 0, 0, 0, 233, 235, 3, 6, 3, 0, 234, 233,
		1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0,
		0, 0, 237, 239, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 241, 3, 2, 1, 0,
		240, 242, 3, 6, 3, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		244, 1, 0, 0, 0, 243, 236, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 37, 1, 0, 0, 0, 247, 248, 5, 58,
		0, 0, 248, 249, 3, 10, 5, 0, 249, 250, 5, 64, 0, 0, 250, 258, 3, 36, 18,
		0, 251, 252, 5, 53, 0, 0, 252, 253, 3, 10, 5, 0, 253, 254, 5, 64, 0, 0,
		254, 255, 3, 36, 18, 0, 255, 257, 1, 0, 0, 0, 256, 251, 1, 0, 0, 0, 257,
		260, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 263,
		1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 262, 5, 52, 0, 0, 262, 264, 3, 36,
		18, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0,
		265, 266, 5, 54, 0, 0, 266, 39, 1, 0, 0, 0, 267, 268, 5, 56, 0, 0, 268,
		269, 3, 54, 27, 0, 269, 270, 5, 47, 0, 0, 270, 271, 3, 10, 5, 0, 271, 272,
		5, 51, 0, 0, 272, 273, 3, 36, 18, 0, 273, 274, 5, 54, 0, 0, 274, 290, 1,
		0, 0, 0, 275, 276, 5, 56, 0, 0, 276, 277, 3, 54, 27, 0, 277, 278, 5, 1,
		0, 0, 278, 279, 3, 10, 5, 0, 279, 280, 5, 24, 0, 0, 280, 283, 3, 10, 5,
		0, 281, 282, 5, 24, 0, 0, 282, 284, 3, 10, 5, 0, 283, 281, 1, 0, 0, 0,
		283, 284, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 5, 51, 0, 0, 286,
		287, 3, 36, 18, 0, 287, 288, 5, 54, 0, 0, 288, 290, 1, 0, 0, 0, 289, 267,
		1, 0, 0, 0, 289, 275, 1, 0, 0, 0, 290, 41, 1, 0, 0, 0, 291, 292, 5, 67,
		0, 0, 292, 293, 3, 10, 5, 0, 293, 294, 5, 51, 0, 0, 294, 295, 3, 36, 18,
		0, 295, 296, 5, 54, 0, 0, 296, 43, 1, 0, 0, 0, 297, 311, 5, 25, 0, 0, 298,
		303, 3, 46, 23, 0, 299, 300, 5, 24, 0, 0, 300, 302, 3, 46, 23, 0, 301,
		299, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304,
		1, 0, 0, 0, 304, 308, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 307, 5, 24,
		0, 0, 307, 309, 3, 74, 37, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0,
		0, 309, 312, 1, 0, 0, 0, 310, 312, 3, 74, 37, 0, 311, 298, 1, 0, 0, 0,
		311, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 5, 26, 0, 0, 314,
		45, 1, 0, 0, 0, 315, 316, 3, 54, 27, 0, 316, 317, 5, 1, 0, 0, 317, 318,
		3, 10, 5, 0, 318, 321, 1, 0, 0, 0, 319, 321, 3, 10, 5, 0, 320, 315, 1,
		0, 0, 0, 320, 319, 1, 0, 0, 0, 321, 47, 1, 0, 0, 0, 322, 323, 3, 54, 27,
		0, 323, 324, 5, 27, 0, 0, 324, 325, 3, 10, 5, 0, 325, 326, 5, 28, 0, 0,
		326, 332, 1, 0, 0, 0, 327, 328, 3, 54, 27, 0, 328, 329, 5, 29, 0, 0, 329,
		330, 3, 54, 27, 0, 330, 332, 1, 0, 0, 0, 331, 322, 1, 0, 0, 0, 331, 327,
		1, 0, 0, 0, 332, 49, 1, 0, 0, 0, 333, 334, 5, 30, 0, 0, 334, 51, 1, 0,
		0, 0, 335, 336, 5, 48, 0, 0, 336, 337, 5, 2, 0, 0, 337, 338, 3, 10, 5,
		0, 338, 339, 5, 3, 0, 0, 339, 53, 1, 0, 0, 0, 340, 344, 5, 84, 0, 0, 341,
		343, 7, 6, 0, 0, 342, 341, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 55, 1, 0, 0, 0, 346, 344, 1, 0,
		0, 0, 347, 348, 5, 62, 0, 0, 348, 349, 3, 36, 18, 0, 349, 350, 5, 66, 0,
		0, 350, 353, 3, 10, 5, 0, 351, 354, 3, 6, 3, 0, 352, 354, 5, 0, 0, 1, 353,
		351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0, 354, 57, 1, 0, 0, 0, 355, 360, 3,
		54, 27, 0, 356, 357, 5, 24, 0, 0, 357, 359, 3, 54, 27, 0, 358, 356, 1,
		0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0,
		0, 361, 59, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 368, 3, 10, 5, 0, 364,
		365, 5, 24, 0, 0, 365, 367, 3, 10, 5, 0, 366, 364, 1, 0, 0, 0, 367, 370,
		1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 373, 1, 0,
		0, 0, 370, 368, 1, 0, 0, 0, 371, 373, 5, 89, 0, 0, 372, 363, 1, 0, 0, 0,
		372, 371, 1, 0, 0, 0, 373, 61, 1, 0, 0, 0, 374, 383, 5, 63, 0, 0, 375,
		380, 3, 10, 5, 0, 376, 377, 5, 24, 0, 0, 377, 379, 3, 10, 5, 0, 378, 376,
		1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0,
		0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 375, 1, 0, 0, 0,
		383, 384, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386, 3, 6, 3, 0, 386,
		63, 1, 0, 0, 0, 387, 388, 5, 50, 0, 0, 388, 389, 3, 6, 3, 0, 389, 65, 1,
		0, 0, 0, 390, 391, 5, 57, 0, 0, 391, 392, 3, 54, 27, 0, 392, 393, 3, 6,
		3, 0, 393, 67, 1, 0, 0, 0, 394, 395, 5, 32, 0, 0, 395, 396, 3, 54, 27,
		0, 396, 397, 5, 32, 0, 0, 397, 69, 1, 0, 0, 0, 398, 399, 5, 69, 0, 0, 399,
		408, 5, 2, 0, 0, 400, 405, 3, 54, 27, 0, 401, 402, 5, 24, 0, 0, 402, 404,
		3, 54, 27, 0, 403, 401, 1, 0, 0, 0, 404, 407, 1, 0, 0, 0, 405, 403, 1,
		0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0,
		0, 408, 400, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410,
		411, 5, 3, 0, 0, 411, 412, 3, 36, 18, 0, 412, 413, 5, 54, 0, 0, 413, 71,
		1, 0, 0, 0, 414, 415, 3, 54, 27, 0, 415, 416, 5, 22, 0, 0, 416, 417, 3,
		54, 27, 0, 417, 426, 5, 2, 0, 0, 418, 423, 3, 10, 5, 0, 419, 420, 5, 24,
		0, 0, 420, 422, 3, 10, 5, 0, 421, 419, 1, 0, 0, 0, 422, 425, 1, 0, 0, 0,
		423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425,
		423, 1, 0, 0, 0, 426, 418, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428,
		1, 0, 0, 0, 428, 429, 5, 3, 0, 0, 429, 73, 1, 0, 0, 0, 430, 431, 5, 33,
		0, 0, 431, 432, 7, 7, 0, 0, 432, 433, 5, 1, 0, 0, 433, 434, 3, 10, 5, 0,
		434, 75, 1, 0, 0, 0, 435, 436, 7, 8, 0, 0, 436, 77, 1, 0, 0, 0, 437, 438,
		5, 75, 0, 0, 438, 439, 5, 29, 0, 0, 439, 440, 7, 9, 0, 0, 440, 79, 1, 0,
		0, 0, 39, 83, 94, 102, 107, 125, 133, 141, 152, 159, 184, 188, 192, 203,
		213, 223, 227, 236, 241, 245, 258, 263, 283, 289, 303, 308, 311, 320, 331,
		344, 353, 360, 368, 372, 380, 383, 405, 408, 423, 426,
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
	lua_grammar_antlr4ParserSEPARATOR           = 46
	lua_grammar_antlr4ParserKW_IN               = 47
	lua_grammar_antlr4ParserKW_PRINT            = 48
	lua_grammar_antlr4ParserKW_AND              = 49
	lua_grammar_antlr4ParserKW_BREAK            = 50
	lua_grammar_antlr4ParserKW_DO               = 51
	lua_grammar_antlr4ParserKW_ELSE             = 52
	lua_grammar_antlr4ParserKW_ELSEIF           = 53
	lua_grammar_antlr4ParserKW_END              = 54
	lua_grammar_antlr4ParserKW_FALSE            = 55
	lua_grammar_antlr4ParserKW_FOR              = 56
	lua_grammar_antlr4ParserKW_GOTO             = 57
	lua_grammar_antlr4ParserKW_IF               = 58
	lua_grammar_antlr4ParserKW_NIL              = 59
	lua_grammar_antlr4ParserKW_NOT              = 60
	lua_grammar_antlr4ParserKW_OR               = 61
	lua_grammar_antlr4ParserKW_REPEAT           = 62
	lua_grammar_antlr4ParserKW_RETURN           = 63
	lua_grammar_antlr4ParserKW_THEN             = 64
	lua_grammar_antlr4ParserKW_TRUE             = 65
	lua_grammar_antlr4ParserKW_UNTIL            = 66
	lua_grammar_antlr4ParserKW_WHILE            = 67
	lua_grammar_antlr4ParserKW_LOCAL            = 68
	lua_grammar_antlr4ParserKW_FUNCTION         = 69
	lua_grammar_antlr4ParserKW_INDEX            = 70
	lua_grammar_antlr4ParserKW_NEWINDEX         = 71
	lua_grammar_antlr4ParserKW_MODE             = 72
	lua_grammar_antlr4ParserKW_PCALL            = 73
	lua_grammar_antlr4ParserKW_XPCALL           = 74
	lua_grammar_antlr4ParserKW_COROUTINE        = 75
	lua_grammar_antlr4ParserKW_CREATE           = 76
	lua_grammar_antlr4ParserKW_RESUME           = 77
	lua_grammar_antlr4ParserKW_YIELD            = 78
	lua_grammar_antlr4ParserKW_STATUS           = 79
	lua_grammar_antlr4ParserKW_NAN              = 80
	lua_grammar_antlr4ParserKW_INF              = 81
	lua_grammar_antlr4ParserNUMBER              = 82
	lua_grammar_antlr4ParserSTRING              = 83
	lua_grammar_antlr4ParserLETTER              = 84
	lua_grammar_antlr4ParserDIGIT               = 85
	lua_grammar_antlr4ParserWS                  = 86
	lua_grammar_antlr4ParserSINGLE_LINE_COMMENT = 87
	lua_grammar_antlr4ParserMULTI_LINE_COMMENT  = 88
	lua_grammar_antlr4ParserVARARG              = 89
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
	lua_grammar_antlr4ParserRULE_block                = 18
	lua_grammar_antlr4ParserRULE_if_statement         = 19
	lua_grammar_antlr4ParserRULE_for_statement        = 20
	lua_grammar_antlr4ParserRULE_while_statement      = 21
	lua_grammar_antlr4ParserRULE_table                = 22
	lua_grammar_antlr4ParserRULE_field                = 23
	lua_grammar_antlr4ParserRULE_table_access         = 24
	lua_grammar_antlr4ParserRULE_single_line_comment  = 25
	lua_grammar_antlr4ParserRULE_print_statement      = 26
	lua_grammar_antlr4ParserRULE_identifier           = 27
	lua_grammar_antlr4ParserRULE_repeat_statement     = 28
	lua_grammar_antlr4ParserRULE_identifier_list      = 29
	lua_grammar_antlr4ParserRULE_expression_list      = 30
	lua_grammar_antlr4ParserRULE_return_statement     = 31
	lua_grammar_antlr4ParserRULE_break_statement      = 32
	lua_grammar_antlr4ParserRULE_goto_statement       = 33
	lua_grammar_antlr4ParserRULE_label_statement      = 34
	lua_grammar_antlr4ParserRULE_function_expression  = 35
	lua_grammar_antlr4ParserRULE_method_call          = 36
	lua_grammar_antlr4ParserRULE_metatable_field      = 37
	lua_grammar_antlr4ParserRULE_metamethod           = 38
	lua_grammar_antlr4ParserRULE_coroutine_statement  = 39
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3493385931619041276) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&1017629) != 0) {
		{
			p.SetState(80)
			p.Statement()
		}

		p.SetState(85)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Function_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.Function_call()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.Return_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.Break_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(93)
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
	Coroutine_statement() ICoroutine_statementContext

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

func (s *Control_statementContext) Coroutine_statement() ICoroutine_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoroutine_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoroutine_statementContext)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserKW_IF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.If_statement()
		}

	case lua_grammar_antlr4ParserKW_FOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.For_statement()
		}

	case lua_grammar_antlr4ParserKW_WHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.While_statement()
		}

	case lua_grammar_antlr4ParserKW_REPEAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.Repeat_statement()
		}

	case lua_grammar_antlr4ParserKW_GOTO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.Goto_statement()
		}

	case lua_grammar_antlr4ParserKW_COROUTINE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.Coroutine_statement()
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
		p.SetState(104)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(106)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(109)
		p.Identifier_list()
	}
	{
		p.SetState(110)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Expression_list()
	}
	{
		p.SetState(112)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(115)
			p.Literal()
		}

	case 2:
		{
			p.SetState(116)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(117)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expression(0)
		}
		{
			p.SetState(119)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(121)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(122)
			p.Table()
		}

	case 6:
		{
			p.SetState(123)
			p.Table_access()
		}

	case 7:
		{
			p.SetState(124)
			p.Function_expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(133)
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
			p.SetState(127)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(128)
				p.Operators()
			}
			{
				p.SetState(129)
				p.expression(6)
			}

		}
		p.SetState(135)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__22, lua_grammar_antlr4ParserT__24, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_PCALL, lua_grammar_antlr4ParserKW_XPCALL, lua_grammar_antlr4ParserKW_NAN, lua_grammar_antlr4ParserKW_INF, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Primary_expression()
		}

	case lua_grammar_antlr4ParserKW_NOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(lua_grammar_antlr4ParserKW_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Prefix_expression()
		}

	case lua_grammar_antlr4ParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(lua_grammar_antlr4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.expression(0)
		}
		{
			p.SetState(147)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Function_call()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(150)
			p.Table()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
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
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__4, lua_grammar_antlr4ParserT__5, lua_grammar_antlr4ParserT__6, lua_grammar_antlr4ParserT__7, lua_grammar_antlr4ParserT__8, lua_grammar_antlr4ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Comparison_operator()
		}

	case lua_grammar_antlr4ParserT__10, lua_grammar_antlr4ParserT__11, lua_grammar_antlr4ParserT__12, lua_grammar_antlr4ParserT__13, lua_grammar_antlr4ParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Arith_operator()
		}

	case lua_grammar_antlr4ParserKW_AND, lua_grammar_antlr4ParserKW_OR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.Logical_operator()
		}

	case lua_grammar_antlr4ParserT__15, lua_grammar_antlr4ParserT__16, lua_grammar_antlr4ParserT__17, lua_grammar_antlr4ParserT__18, lua_grammar_antlr4ParserT__19:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(157)
			p.Bitwise_operator()
		}

	case lua_grammar_antlr4ParserT__20:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)
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
		p.SetState(161)
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
		p.SetState(163)
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
		p.SetState(165)
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
		p.SetState(167)
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
		p.SetState(169)
		p.Match(lua_grammar_antlr4ParserT__20)
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
		p.SetState(171)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&503317521) != 0) {
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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			_la = p.GetTokenStream().LA(1)

			if !(_la == lua_grammar_antlr4ParserKW_PCALL || _la == lua_grammar_antlr4ParserKW_XPCALL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(174)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Expression_list()
		}
		{
			p.SetState(176)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(178)
				p.Identifier()
			}

		case 2:
			{
				p.SetState(179)
				p.Table_access()
			}

		case 3:
			{
				p.SetState(180)
				p.Match(lua_grammar_antlr4ParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(181)
				p.expression(0)
			}
			{
				p.SetState(182)
				p.Match(lua_grammar_antlr4ParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__21 {
			{
				p.SetState(186)
				p.Match(lua_grammar_antlr4ParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(187)
				p.Identifier()
			}

		}
		{
			p.SetState(190)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&612771024341041156) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&17793809) != 0) {
			{
				p.SetState(191)
				p.Expression_list()
			}

		}
		{
			p.SetState(194)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.Table_insert()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)
			p.Match(lua_grammar_antlr4ParserKW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(lua_grammar_antlr4ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Expression_list()
		}
		{
			p.SetState(200)
			p.Match(lua_grammar_antlr4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
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
		p.SetState(205)
		p.Match(lua_grammar_antlr4ParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Identifier()
	}
	{
		p.SetState(208)
		p.Match(lua_grammar_antlr4ParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.expression(0)
	}
	{
		p.SetState(210)
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
	KW_LOCAL() antlr.TerminalNode
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
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_LOCAL {
		{
			p.SetState(212)
			p.Match(lua_grammar_antlr4ParserKW_LOCAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(215)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Identifier()
	}
	{
		p.SetState(217)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserLETTER:
		{
			p.SetState(218)
			p.Identifier()
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(219)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(220)
				p.Identifier()
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case lua_grammar_antlr4ParserVARARG:
		{
			p.SetState(226)
			p.Match(lua_grammar_antlr4ParserVARARG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case lua_grammar_antlr4ParserT__2:

	default:
	}
	{
		p.SetState(229)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Block()
	}
	{
		p.SetState(231)
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
	p.EnterRule(localctx, 36, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3493315562874863612) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&1017629) != 0) {
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserSEPARATOR {
			{
				p.SetState(233)
				p.Statement_terminator()
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(239)
			p.Statement()
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(240)
				p.Statement_terminator()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(245)
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
	p.EnterRule(localctx, 38, lua_grammar_antlr4ParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(lua_grammar_antlr4ParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.expression(0)
	}
	{
		p.SetState(249)
		p.Match(lua_grammar_antlr4ParserKW_THEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Block()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserKW_ELSEIF {
		{
			p.SetState(251)
			p.Match(lua_grammar_antlr4ParserKW_ELSEIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.expression(0)
		}
		{
			p.SetState(253)
			p.Match(lua_grammar_antlr4ParserKW_THEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Block()
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserKW_ELSE {
		{
			p.SetState(261)
			p.Match(lua_grammar_antlr4ParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Block()
		}

	}
	{
		p.SetState(265)
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
	p.EnterRule(localctx, 40, lua_grammar_antlr4ParserRULE_for_statement)
	var _la int

	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Identifier()
		}
		{
			p.SetState(269)
			p.Match(lua_grammar_antlr4ParserKW_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.expression(0)
		}
		{
			p.SetState(271)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Block()
		}
		{
			p.SetState(273)
			p.Match(lua_grammar_antlr4ParserKW_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(lua_grammar_antlr4ParserKW_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Identifier()
		}
		{
			p.SetState(277)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.expression(0)
		}
		{
			p.SetState(279)
			p.Match(lua_grammar_antlr4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.expression(0)
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(281)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(282)
				p.expression(0)
			}

		}
		{
			p.SetState(285)
			p.Match(lua_grammar_antlr4ParserKW_DO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Block()
		}
		{
			p.SetState(287)
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
	p.EnterRule(localctx, 42, lua_grammar_antlr4ParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(lua_grammar_antlr4ParserKW_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.expression(0)
	}
	{
		p.SetState(293)
		p.Match(lua_grammar_antlr4ParserKW_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Block()
	}
	{
		p.SetState(295)
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
	p.EnterRule(localctx, 44, lua_grammar_antlr4ParserRULE_table)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(lua_grammar_antlr4ParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__22, lua_grammar_antlr4ParserT__24, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_FUNCTION, lua_grammar_antlr4ParserKW_PCALL, lua_grammar_antlr4ParserKW_XPCALL, lua_grammar_antlr4ParserKW_NAN, lua_grammar_antlr4ParserKW_INF, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		{
			p.SetState(298)
			p.Field()
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(299)
					p.Match(lua_grammar_antlr4ParserT__23)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(300)
					p.Field()
				}

			}
			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(306)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(307)
				p.Metatable_field()
			}

		}

	case lua_grammar_antlr4ParserT__32:
		{
			p.SetState(310)
			p.Metatable_field()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(313)
		p.Match(lua_grammar_antlr4ParserT__25)
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
	p.EnterRule(localctx, 46, lua_grammar_antlr4ParserRULE_field)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Identifier()
		}
		{
			p.SetState(316)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
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
	p.EnterRule(localctx, 48, lua_grammar_antlr4ParserRULE_table_access)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Identifier()
		}
		{
			p.SetState(323)
			p.Match(lua_grammar_antlr4ParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.expression(0)
		}
		{
			p.SetState(325)
			p.Match(lua_grammar_antlr4ParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Identifier()
		}
		{
			p.SetState(328)
			p.Match(lua_grammar_antlr4ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
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
	p.EnterRule(localctx, 50, lua_grammar_antlr4ParserRULE_single_line_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(lua_grammar_antlr4ParserT__29)
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
	p.EnterRule(localctx, 52, lua_grammar_antlr4ParserRULE_print_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(lua_grammar_antlr4ParserKW_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
		p.expression(0)
	}
	{
		p.SetState(338)
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
	p.EnterRule(localctx, 54, lua_grammar_antlr4ParserRULE_identifier)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(lua_grammar_antlr4ParserLETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(344)
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
				p.SetState(341)
				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&27021597764222977) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 56, lua_grammar_antlr4ParserRULE_repeat_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(lua_grammar_antlr4ParserKW_REPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Block()
	}
	{
		p.SetState(349)
		p.Match(lua_grammar_antlr4ParserKW_UNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.expression(0)
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserSEPARATOR:
		{
			p.SetState(351)
			p.Statement_terminator()
		}

	case lua_grammar_antlr4ParserEOF:
		{
			p.SetState(352)
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
	p.EnterRule(localctx, 58, lua_grammar_antlr4ParserRULE_identifier_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Identifier()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__23 {
		{
			p.SetState(356)
			p.Match(lua_grammar_antlr4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Identifier()
		}

		p.SetState(362)
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
	p.EnterRule(localctx, 60, lua_grammar_antlr4ParserRULE_expression_list)
	var _la int

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__1, lua_grammar_antlr4ParserT__22, lua_grammar_antlr4ParserT__24, lua_grammar_antlr4ParserKW_PRINT, lua_grammar_antlr4ParserKW_FALSE, lua_grammar_antlr4ParserKW_NIL, lua_grammar_antlr4ParserKW_TRUE, lua_grammar_antlr4ParserKW_FUNCTION, lua_grammar_antlr4ParserKW_PCALL, lua_grammar_antlr4ParserKW_XPCALL, lua_grammar_antlr4ParserKW_NAN, lua_grammar_antlr4ParserKW_INF, lua_grammar_antlr4ParserNUMBER, lua_grammar_antlr4ParserSTRING, lua_grammar_antlr4ParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.expression(0)
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(364)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(365)
				p.expression(0)
			}

			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case lua_grammar_antlr4ParserVARARG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)
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
	p.EnterRule(localctx, 62, lua_grammar_antlr4ParserRULE_return_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(lua_grammar_antlr4ParserKW_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&612771024341041156) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&1016593) != 0) {
		{
			p.SetState(375)
			p.expression(0)
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(376)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(377)
				p.expression(0)
			}

			p.SetState(382)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(385)
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
	p.EnterRule(localctx, 64, lua_grammar_antlr4ParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(lua_grammar_antlr4ParserKW_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
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
	p.EnterRule(localctx, 66, lua_grammar_antlr4ParserRULE_goto_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(lua_grammar_antlr4ParserKW_GOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Identifier()
	}
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 68, lua_grammar_antlr4ParserRULE_label_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(lua_grammar_antlr4ParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Identifier()
	}
	{
		p.SetState(396)
		p.Match(lua_grammar_antlr4ParserT__31)
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
	p.EnterRule(localctx, 70, lua_grammar_antlr4ParserRULE_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(lua_grammar_antlr4ParserKW_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserLETTER {
		{
			p.SetState(400)
			p.Identifier()
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(401)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(402)
				p.Identifier()
			}

			p.SetState(407)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(410)
		p.Match(lua_grammar_antlr4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Block()
	}
	{
		p.SetState(412)
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
	p.EnterRule(localctx, 72, lua_grammar_antlr4ParserRULE_method_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Identifier()
	}
	{
		p.SetState(415)
		p.Match(lua_grammar_antlr4ParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Identifier()
	}
	{
		p.SetState(417)
		p.Match(lua_grammar_antlr4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&612771024341041156) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&1016593) != 0) {
		{
			p.SetState(418)
			p.expression(0)
		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__23 {
			{
				p.SetState(419)
				p.Match(lua_grammar_antlr4ParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(420)
				p.expression(0)
			}

			p.SetState(425)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(428)
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
	p.EnterRule(localctx, 74, lua_grammar_antlr4ParserRULE_metatable_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(lua_grammar_antlr4ParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-70)) & ^0x3f) == 0 && ((int64(1)<<(_la-70))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(432)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
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
	p.EnterRule(localctx, 76, lua_grammar_antlr4ParserRULE_metamethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70351564308480) != 0) {
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
	p.EnterRule(localctx, 78, lua_grammar_antlr4ParserRULE_coroutine_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(lua_grammar_antlr4ParserKW_COROUTINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.Match(lua_grammar_antlr4ParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-76)) & ^0x3f) == 0 && ((int64(1)<<(_la-76))&15) != 0) {
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
