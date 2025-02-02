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
		"", "'='", "'+='", "'-='", "'*='", "'/='", "'('", "')'", "','", "'{'",
		"'}'", "'__metatable'", "'::'", "'__add'", "'__sub'", "'__mul'", "'__div'",
		"'__mod'", "'__pow'", "'__unm'", "'__concat'", "'__len'", "'__eq'",
		"'__lt'", "'__le'", "'__tostring'", "'__pairs'", "'__ipairs'", "'__call'",
		"'#'", "'if'", "'then'", "'elseif'", "'else'", "'end'", "'while'", "'do'",
		"'repeat'", "'until'", "'for'", "'break'", "'goto'", "'coroutine'",
		"'.'", "'create'", "'resume'", "'yield'", "'status'", "'local'", "'function'",
		"'return'", "'and'", "'or'", "'=='", "'>='", "'<='", "'~='", "'>'",
		"'<'", "'+'", "'-'", "'*'", "'/'", "'//'", "'%'", "'^'", "'&'", "'|'",
		"'~'", "'<<'", "'>>'", "'//='", "'^='", "'&='", "'|='", "'not'", "'..'",
		"'...'", "'..='", "'??='", "'+_'", "'-_'", "'??'", "'<<='", "'>>='",
		"", "", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"IDENTIFIER", "BOOL", "NIL", "NUMBER", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "assignStatement", "expression", "primaryExpression",
		"literal", "variable", "functionCall", "tableConstructor", "metatable",
		"metamethods", "labelStatement", "metamethod", "field", "binaryOperation",
		"unaryOperation", "controlFlowStatement", "ifStatement", "whileStatement",
		"repeatStatement", "forStatement", "breakStatement", "gotoStatement",
		"coroutineStatement", "block", "localDeclaration", "functionDeclaration",
		"returnStatement", "operatorGroup", "logicalOp", "comparisonOp", "arithOp",
		"bitwiseOp", "assignOp", "unaryOp", "concatOp", "varargOp", "compoundAssignOp",
		"incrOp", "coalesceOp", "shiftAssignOp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 91, 378, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 5, 0, 84,
		8, 0, 10, 0, 12, 0, 87, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 97, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 118, 8, 3, 10, 3, 12, 3, 121, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 143, 8, 7, 10, 7, 12, 7, 146, 9, 7, 3, 7, 148,
		8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 156, 8, 8, 10, 8, 12, 8,
		159, 9, 8, 3, 8, 161, 8, 8, 1, 8, 1, 8, 3, 8, 165, 8, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 176, 8, 10, 10, 10, 12,
		10, 179, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 191, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		229, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 236, 8, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 245, 8, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 256, 8, 17,
		10, 17, 12, 17, 259, 9, 17, 1, 17, 1, 17, 3, 17, 263, 8, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 286,
		8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 5, 24, 302, 8, 24, 10, 24, 12, 24, 305,
		9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 311, 8, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 5, 26, 319, 8, 26, 10, 26, 12, 26, 322, 9, 26,
		3, 26, 324, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 5, 27, 334, 8, 27, 10, 27, 12, 27, 337, 9, 27, 3, 27, 339, 8, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		3, 28, 352, 8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 0, 1, 6, 41, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 0, 13, 1, 0, 1, 5, 1, 0, 86, 89, 1, 0, 13, 28, 1, 0, 44, 47, 1, 0,
		51, 52, 1, 0, 53, 58, 1, 0, 59, 65, 1, 0, 66, 70, 2, 0, 1, 5, 71, 74, 3,
		0, 29, 29, 60, 60, 75, 75, 3, 0, 2, 5, 71, 72, 78, 79, 1, 0, 80, 81, 1,
		0, 83, 84, 391, 0, 85, 1, 0, 0, 0, 2, 96, 1, 0, 0, 0, 4, 98, 1, 0, 0, 0,
		6, 107, 1, 0, 0, 0, 8, 131, 1, 0, 0, 0, 10, 133, 1, 0, 0, 0, 12, 135, 1,
		0, 0, 0, 14, 137, 1, 0, 0, 0, 16, 151, 1, 0, 0, 0, 18, 168, 1, 0, 0, 0,
		20, 172, 1, 0, 0, 0, 22, 180, 1, 0, 0, 0, 24, 184, 1, 0, 0, 0, 26, 190,
		1, 0, 0, 0, 28, 228, 1, 0, 0, 0, 30, 235, 1, 0, 0, 0, 32, 244, 1, 0, 0,
		0, 34, 246, 1, 0, 0, 0, 36, 266, 1, 0, 0, 0, 38, 272, 1, 0, 0, 0, 40, 277,
		1, 0, 0, 0, 42, 291, 1, 0, 0, 0, 44, 293, 1, 0, 0, 0, 46, 296, 1, 0, 0,
		0, 48, 303, 1, 0, 0, 0, 50, 306, 1, 0, 0, 0, 52, 312, 1, 0, 0, 0, 54, 329,
		1, 0, 0, 0, 56, 351, 1, 0, 0, 0, 58, 353, 1, 0, 0, 0, 60, 355, 1, 0, 0,
		0, 62, 357, 1, 0, 0, 0, 64, 359, 1, 0, 0, 0, 66, 361, 1, 0, 0, 0, 68, 363,
		1, 0, 0, 0, 70, 365, 1, 0, 0, 0, 72, 367, 1, 0, 0, 0, 74, 369, 1, 0, 0,
		0, 76, 371, 1, 0, 0, 0, 78, 373, 1, 0, 0, 0, 80, 375, 1, 0, 0, 0, 82, 84,
		3, 2, 1, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5,
		0, 0, 1, 89, 1, 1, 0, 0, 0, 90, 97, 3, 4, 2, 0, 91, 97, 3, 32, 16, 0, 92,
		97, 3, 52, 26, 0, 93, 97, 3, 54, 27, 0, 94, 97, 3, 50, 25, 0, 95, 97, 3,
		22, 11, 0, 96, 90, 1, 0, 0, 0, 96, 91, 1, 0, 0, 0, 96, 92, 1, 0, 0, 0,
		96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 3, 1, 0,
		0, 0, 98, 99, 3, 12, 6, 0, 99, 100, 7, 0, 0, 0, 100, 101, 3, 6, 3, 0, 101,
		5, 1, 0, 0, 0, 102, 103, 6, 3, -1, 0, 103, 108, 3, 8, 4, 0, 104, 105, 3,
		68, 34, 0, 105, 106, 3, 6, 3, 1, 106, 108, 1, 0, 0, 0, 107, 102, 1, 0,
		0, 0, 107, 104, 1, 0, 0, 0, 108, 119, 1, 0, 0, 0, 109, 110, 10, 3, 0, 0,
		110, 111, 3, 56, 28, 0, 111, 112, 3, 6, 3, 4, 112, 118, 1, 0, 0, 0, 113,
		114, 10, 2, 0, 0, 114, 115, 3, 56, 28, 0, 115, 116, 3, 6, 3, 2, 116, 118,
		1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 118, 121, 1, 0,
		0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 7, 1, 0, 0, 0, 121,
		119, 1, 0, 0, 0, 122, 132, 3, 10, 5, 0, 123, 132, 3, 12, 6, 0, 124, 132,
		3, 14, 7, 0, 125, 132, 3, 30, 15, 0, 126, 132, 3, 16, 8, 0, 127, 128, 5,
		6, 0, 0, 128, 129, 3, 6, 3, 0, 129, 130, 5, 7, 0, 0, 130, 132, 1, 0, 0,
		0, 131, 122, 1, 0, 0, 0, 131, 123, 1, 0, 0, 0, 131, 124, 1, 0, 0, 0, 131,
		125, 1, 0, 0, 0, 131, 126, 1, 0, 0, 0, 131, 127, 1, 0, 0, 0, 132, 9, 1,
		0, 0, 0, 133, 134, 7, 1, 0, 0, 134, 11, 1, 0, 0, 0, 135, 136, 5, 85, 0,
		0, 136, 13, 1, 0, 0, 0, 137, 138, 5, 85, 0, 0, 138, 147, 5, 6, 0, 0, 139,
		144, 3, 6, 3, 0, 140, 141, 5, 8, 0, 0, 141, 143, 3, 6, 3, 0, 142, 140,
		1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 139, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 5, 7, 0, 0, 150,
		15, 1, 0, 0, 0, 151, 160, 5, 9, 0, 0, 152, 157, 3, 26, 13, 0, 153, 154,
		5, 8, 0, 0, 154, 156, 3, 26, 13, 0, 155, 153, 1, 0, 0, 0, 156, 159, 1,
		0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 161, 1, 0, 0,
		0, 159, 157, 1, 0, 0, 0, 160, 152, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		164, 1, 0, 0, 0, 162, 163, 5, 8, 0, 0, 163, 165, 3, 18, 9, 0, 164, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 10,
		0, 0, 167, 17, 1, 0, 0, 0, 168, 169, 5, 11, 0, 0, 169, 170, 5, 1, 0, 0,
		170, 171, 3, 6, 3, 0, 171, 19, 1, 0, 0, 0, 172, 177, 3, 24, 12, 0, 173,
		174, 5, 8, 0, 0, 174, 176, 3, 24, 12, 0, 175, 173, 1, 0, 0, 0, 176, 179,
		1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 21, 1, 0,
		0, 0, 179, 177, 1, 0, 0, 0, 180, 181, 5, 12, 0, 0, 181, 182, 5, 85, 0,
		0, 182, 183, 5, 12, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 7, 2, 0, 0, 185,
		25, 1, 0, 0, 0, 186, 187, 5, 85, 0, 0, 187, 188, 5, 1, 0, 0, 188, 191,
		3, 6, 3, 0, 189, 191, 3, 6, 3, 0, 190, 186, 1, 0, 0, 0, 190, 189, 1, 0,
		0, 0, 191, 27, 1, 0, 0, 0, 192, 193, 3, 6, 3, 0, 193, 194, 3, 62, 31, 0,
		194, 195, 3, 6, 3, 0, 195, 229, 1, 0, 0, 0, 196, 197, 3, 6, 3, 0, 197,
		198, 3, 64, 32, 0, 198, 199, 3, 6, 3, 0, 199, 229, 1, 0, 0, 0, 200, 201,
		3, 6, 3, 0, 201, 202, 3, 60, 30, 0, 202, 203, 3, 6, 3, 0, 203, 229, 1,
		0, 0, 0, 204, 205, 3, 6, 3, 0, 205, 206, 3, 58, 29, 0, 206, 207, 3, 6,
		3, 0, 207, 229, 1, 0, 0, 0, 208, 209, 3, 6, 3, 0, 209, 210, 3, 70, 35,
		0, 210, 211, 3, 6, 3, 0, 211, 229, 1, 0, 0, 0, 212, 213, 3, 6, 3, 0, 213,
		214, 3, 74, 37, 0, 214, 215, 3, 6, 3, 0, 215, 229, 1, 0, 0, 0, 216, 217,
		3, 6, 3, 0, 217, 218, 3, 78, 39, 0, 218, 219, 3, 6, 3, 0, 219, 229, 1,
		0, 0, 0, 220, 221, 3, 6, 3, 0, 221, 222, 3, 80, 40, 0, 222, 223, 3, 6,
		3, 0, 223, 229, 1, 0, 0, 0, 224, 225, 3, 6, 3, 0, 225, 226, 3, 76, 38,
		0, 226, 227, 3, 6, 3, 0, 227, 229, 1, 0, 0, 0, 228, 192, 1, 0, 0, 0, 228,
		196, 1, 0, 0, 0, 228, 200, 1, 0, 0, 0, 228, 204, 1, 0, 0, 0, 228, 208,
		1, 0, 0, 0, 228, 212, 1, 0, 0, 0, 228, 216, 1, 0, 0, 0, 228, 220, 1, 0,
		0, 0, 228, 224, 1, 0, 0, 0, 229, 29, 1, 0, 0, 0, 230, 231, 3, 68, 34, 0,
		231, 232, 3, 6, 3, 0, 232, 236, 1, 0, 0, 0, 233, 234, 5, 29, 0, 0, 234,
		236, 3, 6, 3, 0, 235, 230, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 31, 1,
		0, 0, 0, 237, 245, 3, 34, 17, 0, 238, 245, 3, 36, 18, 0, 239, 245, 3, 38,
		19, 0, 240, 245, 3, 40, 20, 0, 241, 245, 3, 42, 21, 0, 242, 245, 3, 44,
		22, 0, 243, 245, 3, 46, 23, 0, 244, 237, 1, 0, 0, 0, 244, 238, 1, 0, 0,
		0, 244, 239, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 33, 1, 0, 0, 0, 246, 247, 5,
		30, 0, 0, 247, 248, 3, 6, 3, 0, 248, 249, 5, 31, 0, 0, 249, 257, 3, 48,
		24, 0, 250, 251, 5, 32, 0, 0, 251, 252, 3, 6, 3, 0, 252, 253, 5, 31, 0,
		0, 253, 254, 3, 48, 24, 0, 254, 256, 1, 0, 0, 0, 255, 250, 1, 0, 0, 0,
		256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258,
		262, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 261, 5, 33, 0, 0, 261, 263,
		3, 48, 24, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1,
		0, 0, 0, 264, 265, 5, 34, 0, 0, 265, 35, 1, 0, 0, 0, 266, 267, 5, 35, 0,
		0, 267, 268, 3, 6, 3, 0, 268, 269, 5, 36, 0, 0, 269, 270, 3, 48, 24, 0,
		270, 271, 5, 34, 0, 0, 271, 37, 1, 0, 0, 0, 272, 273, 5, 37, 0, 0, 273,
		274, 3, 48, 24, 0, 274, 275, 5, 38, 0, 0, 275, 276, 3, 6, 3, 0, 276, 39,
		1, 0, 0, 0, 277, 278, 5, 39, 0, 0, 278, 279, 5, 85, 0, 0, 279, 280, 5,
		1, 0, 0, 280, 281, 3, 6, 3, 0, 281, 282, 5, 8, 0, 0, 282, 285, 3, 6, 3,
		0, 283, 284, 5, 8, 0, 0, 284, 286, 3, 6, 3, 0, 285, 283, 1, 0, 0, 0, 285,
		286, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 5, 36, 0, 0, 288, 289,
		3, 48, 24, 0, 289, 290, 5, 34, 0, 0, 290, 41, 1, 0, 0, 0, 291, 292, 5,
		40, 0, 0, 292, 43, 1, 0, 0, 0, 293, 294, 5, 41, 0, 0, 294, 295, 5, 85,
		0, 0, 295, 45, 1, 0, 0, 0, 296, 297, 5, 42, 0, 0, 297, 298, 5, 43, 0, 0,
		298, 299, 7, 3, 0, 0, 299, 47, 1, 0, 0, 0, 300, 302, 3, 2, 1, 0, 301, 300,
		1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0,
		0, 0, 304, 49, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 307, 5, 48, 0, 0,
		307, 310, 5, 85, 0, 0, 308, 309, 5, 1, 0, 0, 309, 311, 3, 6, 3, 0, 310,
		308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 51, 1, 0, 0, 0, 312, 313, 5,
		49, 0, 0, 313, 314, 5, 85, 0, 0, 314, 323, 5, 6, 0, 0, 315, 320, 5, 85,
		0, 0, 316, 317, 5, 8, 0, 0, 317, 319, 5, 85, 0, 0, 318, 316, 1, 0, 0, 0,
		319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 315, 1, 0, 0, 0, 323, 324,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 5, 7, 0, 0, 326, 327, 3, 48,
		24, 0, 327, 328, 5, 34, 0, 0, 328, 53, 1, 0, 0, 0, 329, 338, 5, 50, 0,
		0, 330, 335, 3, 6, 3, 0, 331, 332, 5, 8, 0, 0, 332, 334, 3, 6, 3, 0, 333,
		331, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336,
		1, 0, 0, 0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 330, 1, 0,
		0, 0, 338, 339, 1, 0, 0, 0, 339, 55, 1, 0, 0, 0, 340, 352, 3, 58, 29, 0,
		341, 352, 3, 60, 30, 0, 342, 352, 3, 62, 31, 0, 343, 352, 3, 64, 32, 0,
		344, 352, 3, 66, 33, 0, 345, 352, 3, 68, 34, 0, 346, 352, 3, 70, 35, 0,
		347, 352, 3, 74, 37, 0, 348, 352, 3, 76, 38, 0, 349, 352, 3, 78, 39, 0,
		350, 352, 3, 80, 40, 0, 351, 340, 1, 0, 0, 0, 351, 341, 1, 0, 0, 0, 351,
		342, 1, 0, 0, 0, 351, 343, 1, 0, 0, 0, 351, 344, 1, 0, 0, 0, 351, 345,
		1, 0, 0, 0, 351, 346, 1, 0, 0, 0, 351, 347, 1, 0, 0, 0, 351, 348, 1, 0,
		0, 0, 351, 349, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 57, 1, 0, 0, 0,
		353, 354, 7, 4, 0, 0, 354, 59, 1, 0, 0, 0, 355, 356, 7, 5, 0, 0, 356, 61,
		1, 0, 0, 0, 357, 358, 7, 6, 0, 0, 358, 63, 1, 0, 0, 0, 359, 360, 7, 7,
		0, 0, 360, 65, 1, 0, 0, 0, 361, 362, 7, 8, 0, 0, 362, 67, 1, 0, 0, 0, 363,
		364, 7, 9, 0, 0, 364, 69, 1, 0, 0, 0, 365, 366, 5, 76, 0, 0, 366, 71, 1,
		0, 0, 0, 367, 368, 5, 77, 0, 0, 368, 73, 1, 0, 0, 0, 369, 370, 7, 10, 0,
		0, 370, 75, 1, 0, 0, 0, 371, 372, 7, 11, 0, 0, 372, 77, 1, 0, 0, 0, 373,
		374, 5, 82, 0, 0, 374, 79, 1, 0, 0, 0, 375, 376, 7, 12, 0, 0, 376, 81,
		1, 0, 0, 0, 26, 85, 96, 107, 117, 119, 131, 144, 147, 157, 160, 164, 177,
		190, 228, 235, 244, 257, 262, 285, 303, 310, 320, 323, 335, 338, 351,
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
	lua_grammar_antlr4ParserEOF        = antlr.TokenEOF
	lua_grammar_antlr4ParserT__0       = 1
	lua_grammar_antlr4ParserT__1       = 2
	lua_grammar_antlr4ParserT__2       = 3
	lua_grammar_antlr4ParserT__3       = 4
	lua_grammar_antlr4ParserT__4       = 5
	lua_grammar_antlr4ParserT__5       = 6
	lua_grammar_antlr4ParserT__6       = 7
	lua_grammar_antlr4ParserT__7       = 8
	lua_grammar_antlr4ParserT__8       = 9
	lua_grammar_antlr4ParserT__9       = 10
	lua_grammar_antlr4ParserT__10      = 11
	lua_grammar_antlr4ParserT__11      = 12
	lua_grammar_antlr4ParserT__12      = 13
	lua_grammar_antlr4ParserT__13      = 14
	lua_grammar_antlr4ParserT__14      = 15
	lua_grammar_antlr4ParserT__15      = 16
	lua_grammar_antlr4ParserT__16      = 17
	lua_grammar_antlr4ParserT__17      = 18
	lua_grammar_antlr4ParserT__18      = 19
	lua_grammar_antlr4ParserT__19      = 20
	lua_grammar_antlr4ParserT__20      = 21
	lua_grammar_antlr4ParserT__21      = 22
	lua_grammar_antlr4ParserT__22      = 23
	lua_grammar_antlr4ParserT__23      = 24
	lua_grammar_antlr4ParserT__24      = 25
	lua_grammar_antlr4ParserT__25      = 26
	lua_grammar_antlr4ParserT__26      = 27
	lua_grammar_antlr4ParserT__27      = 28
	lua_grammar_antlr4ParserT__28      = 29
	lua_grammar_antlr4ParserT__29      = 30
	lua_grammar_antlr4ParserT__30      = 31
	lua_grammar_antlr4ParserT__31      = 32
	lua_grammar_antlr4ParserT__32      = 33
	lua_grammar_antlr4ParserT__33      = 34
	lua_grammar_antlr4ParserT__34      = 35
	lua_grammar_antlr4ParserT__35      = 36
	lua_grammar_antlr4ParserT__36      = 37
	lua_grammar_antlr4ParserT__37      = 38
	lua_grammar_antlr4ParserT__38      = 39
	lua_grammar_antlr4ParserT__39      = 40
	lua_grammar_antlr4ParserT__40      = 41
	lua_grammar_antlr4ParserT__41      = 42
	lua_grammar_antlr4ParserT__42      = 43
	lua_grammar_antlr4ParserT__43      = 44
	lua_grammar_antlr4ParserT__44      = 45
	lua_grammar_antlr4ParserT__45      = 46
	lua_grammar_antlr4ParserT__46      = 47
	lua_grammar_antlr4ParserT__47      = 48
	lua_grammar_antlr4ParserT__48      = 49
	lua_grammar_antlr4ParserT__49      = 50
	lua_grammar_antlr4ParserT__50      = 51
	lua_grammar_antlr4ParserT__51      = 52
	lua_grammar_antlr4ParserT__52      = 53
	lua_grammar_antlr4ParserT__53      = 54
	lua_grammar_antlr4ParserT__54      = 55
	lua_grammar_antlr4ParserT__55      = 56
	lua_grammar_antlr4ParserT__56      = 57
	lua_grammar_antlr4ParserT__57      = 58
	lua_grammar_antlr4ParserT__58      = 59
	lua_grammar_antlr4ParserT__59      = 60
	lua_grammar_antlr4ParserT__60      = 61
	lua_grammar_antlr4ParserT__61      = 62
	lua_grammar_antlr4ParserT__62      = 63
	lua_grammar_antlr4ParserT__63      = 64
	lua_grammar_antlr4ParserT__64      = 65
	lua_grammar_antlr4ParserT__65      = 66
	lua_grammar_antlr4ParserT__66      = 67
	lua_grammar_antlr4ParserT__67      = 68
	lua_grammar_antlr4ParserT__68      = 69
	lua_grammar_antlr4ParserT__69      = 70
	lua_grammar_antlr4ParserT__70      = 71
	lua_grammar_antlr4ParserT__71      = 72
	lua_grammar_antlr4ParserT__72      = 73
	lua_grammar_antlr4ParserT__73      = 74
	lua_grammar_antlr4ParserT__74      = 75
	lua_grammar_antlr4ParserT__75      = 76
	lua_grammar_antlr4ParserT__76      = 77
	lua_grammar_antlr4ParserT__77      = 78
	lua_grammar_antlr4ParserT__78      = 79
	lua_grammar_antlr4ParserT__79      = 80
	lua_grammar_antlr4ParserT__80      = 81
	lua_grammar_antlr4ParserT__81      = 82
	lua_grammar_antlr4ParserT__82      = 83
	lua_grammar_antlr4ParserT__83      = 84
	lua_grammar_antlr4ParserIDENTIFIER = 85
	lua_grammar_antlr4ParserBOOL       = 86
	lua_grammar_antlr4ParserNIL        = 87
	lua_grammar_antlr4ParserNUMBER     = 88
	lua_grammar_antlr4ParserSTRING     = 89
	lua_grammar_antlr4ParserWS         = 90
	lua_grammar_antlr4ParserCOMMENT    = 91
)

// lua_grammar_antlr4Parser rules.
const (
	lua_grammar_antlr4ParserRULE_program              = 0
	lua_grammar_antlr4ParserRULE_statement            = 1
	lua_grammar_antlr4ParserRULE_assignStatement      = 2
	lua_grammar_antlr4ParserRULE_expression           = 3
	lua_grammar_antlr4ParserRULE_primaryExpression    = 4
	lua_grammar_antlr4ParserRULE_literal              = 5
	lua_grammar_antlr4ParserRULE_variable             = 6
	lua_grammar_antlr4ParserRULE_functionCall         = 7
	lua_grammar_antlr4ParserRULE_tableConstructor     = 8
	lua_grammar_antlr4ParserRULE_metatable            = 9
	lua_grammar_antlr4ParserRULE_metamethods          = 10
	lua_grammar_antlr4ParserRULE_labelStatement       = 11
	lua_grammar_antlr4ParserRULE_metamethod           = 12
	lua_grammar_antlr4ParserRULE_field                = 13
	lua_grammar_antlr4ParserRULE_binaryOperation      = 14
	lua_grammar_antlr4ParserRULE_unaryOperation       = 15
	lua_grammar_antlr4ParserRULE_controlFlowStatement = 16
	lua_grammar_antlr4ParserRULE_ifStatement          = 17
	lua_grammar_antlr4ParserRULE_whileStatement       = 18
	lua_grammar_antlr4ParserRULE_repeatStatement      = 19
	lua_grammar_antlr4ParserRULE_forStatement         = 20
	lua_grammar_antlr4ParserRULE_breakStatement       = 21
	lua_grammar_antlr4ParserRULE_gotoStatement        = 22
	lua_grammar_antlr4ParserRULE_coroutineStatement   = 23
	lua_grammar_antlr4ParserRULE_block                = 24
	lua_grammar_antlr4ParserRULE_localDeclaration     = 25
	lua_grammar_antlr4ParserRULE_functionDeclaration  = 26
	lua_grammar_antlr4ParserRULE_returnStatement      = 27
	lua_grammar_antlr4ParserRULE_operatorGroup        = 28
	lua_grammar_antlr4ParserRULE_logicalOp            = 29
	lua_grammar_antlr4ParserRULE_comparisonOp         = 30
	lua_grammar_antlr4ParserRULE_arithOp              = 31
	lua_grammar_antlr4ParserRULE_bitwiseOp            = 32
	lua_grammar_antlr4ParserRULE_assignOp             = 33
	lua_grammar_antlr4ParserRULE_unaryOp              = 34
	lua_grammar_antlr4ParserRULE_concatOp             = 35
	lua_grammar_antlr4ParserRULE_varargOp             = 36
	lua_grammar_antlr4ParserRULE_compoundAssignOp     = 37
	lua_grammar_antlr4ParserRULE_incrOp               = 38
	lua_grammar_antlr4ParserRULE_coalesceOp           = 39
	lua_grammar_antlr4ParserRULE_shiftAssignOp        = 40
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
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

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserEOF, 0)
}

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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1978744046620672) != 0) || _la == lua_grammar_antlr4ParserIDENTIFIER {
		{
			p.SetState(82)
			p.Statement()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(lua_grammar_antlr4ParserEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignStatement() IAssignStatementContext
	ControlFlowStatement() IControlFlowStatementContext
	FunctionDeclaration() IFunctionDeclarationContext
	ReturnStatement() IReturnStatementContext
	LocalDeclaration() ILocalDeclarationContext
	LabelStatement() ILabelStatementContext

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

func (s *StatementContext) AssignStatement() IAssignStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *StatementContext) ControlFlowStatement() IControlFlowStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlFlowStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlFlowStatementContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) LocalDeclaration() ILocalDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalDeclarationContext)
}

func (s *StatementContext) LabelStatement() ILabelStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelStatementContext)
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
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.AssignStatement()
		}

	case lua_grammar_antlr4ParserT__29, lua_grammar_antlr4ParserT__34, lua_grammar_antlr4ParserT__36, lua_grammar_antlr4ParserT__38, lua_grammar_antlr4ParserT__39, lua_grammar_antlr4ParserT__40, lua_grammar_antlr4ParserT__41:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.ControlFlowStatement()
		}

	case lua_grammar_antlr4ParserT__48:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.FunctionDeclaration()
		}

	case lua_grammar_antlr4ParserT__49:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.ReturnStatement()
		}

	case lua_grammar_antlr4ParserT__47:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.LocalDeclaration()
		}

	case lua_grammar_antlr4ParserT__11:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.LabelStatement()
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

// IAssignStatementContext is an interface to support dynamic dispatch.
type IAssignStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Expression() IExpressionContext

	// IsAssignStatementContext differentiates from other interfaces.
	IsAssignStatementContext()
}

type AssignStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStatementContext() *AssignStatementContext {
	var p = new(AssignStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignStatement
	return p
}

func InitEmptyAssignStatementContext(p *AssignStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignStatement
}

func (*AssignStatementContext) IsAssignStatementContext() {}

func NewAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStatementContext {
	var p = new(AssignStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignStatement

	return p
}

func (s *AssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStatementContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignStatementContext) Expression() IExpressionContext {
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

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) AssignStatement() (localctx IAssignStatementContext) {
	localctx = NewAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lua_grammar_antlr4ParserRULE_assignStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Variable()
	}
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(100)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	UnaryOp() IUnaryOpContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	OperatorGroup() IOperatorGroupContext

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

func (s *ExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
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

func (s *ExpressionContext) OperatorGroup() IOperatorGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorGroupContext)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, lua_grammar_antlr4ParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(103)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(104)
			p.UnaryOp()
		}
		{
			p.SetState(105)
			p.expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(110)
					p.OperatorGroup()
				}
				{
					p.SetState(111)
					p.expression(4)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, lua_grammar_antlr4ParserRULE_expression)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					p.OperatorGroup()
				}
				{
					p.SetState(115)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	Variable() IVariableContext
	FunctionCall() IFunctionCallContext
	UnaryOperation() IUnaryOperationContext
	TableConstructor() ITableConstructorContext
	Expression() IExpressionContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Literal() ILiteralContext {
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

func (s *PrimaryExpressionContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PrimaryExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PrimaryExpressionContext) UnaryOperation() IUnaryOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperationContext)
}

func (s *PrimaryExpressionContext) TableConstructor() ITableConstructorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableConstructorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableConstructorContext)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
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

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lua_grammar_antlr4ParserRULE_primaryExpression)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.FunctionCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.UnaryOperation()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.TableConstructor()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.Match(lua_grammar_antlr4ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.expression(0)
		}
		{
			p.SetState(129)
			p.Match(lua_grammar_antlr4ParserT__6)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL() antlr.TerminalNode
	NIL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

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

func (s *LiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserBOOL, 0)
}

func (s *LiteralContext) NIL() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserNIL, 0)
}

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserSTRING, 0)
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
	p.EnterRule(localctx, 10, lua_grammar_antlr4ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-86)) & ^0x3f) == 0 && ((int64(1)<<(_la-86))&15) != 0) {
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lua_grammar_antlr4ParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) AllExpression() []IExpressionContext {
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

func (s *FunctionCallContext) Expression(i int) IExpressionContext {
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

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lua_grammar_antlr4ParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(lua_grammar_antlr4ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505143718464) != 0) || ((int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&31745) != 0) {
		{
			p.SetState(139)
			p.expression(0)
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__7 {
			{
				p.SetState(140)
				p.Match(lua_grammar_antlr4ParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(141)
				p.expression(0)
			}

			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(149)
		p.Match(lua_grammar_antlr4ParserT__6)
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

// ITableConstructorContext is an interface to support dynamic dispatch.
type ITableConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Metatable() IMetatableContext

	// IsTableConstructorContext differentiates from other interfaces.
	IsTableConstructorContext()
}

type TableConstructorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableConstructorContext() *TableConstructorContext {
	var p = new(TableConstructorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_tableConstructor
	return p
}

func InitEmptyTableConstructorContext(p *TableConstructorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_tableConstructor
}

func (*TableConstructorContext) IsTableConstructorContext() {}

func NewTableConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableConstructorContext {
	var p = new(TableConstructorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_tableConstructor

	return p
}

func (s *TableConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableConstructorContext) AllField() []IFieldContext {
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

func (s *TableConstructorContext) Field(i int) IFieldContext {
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

func (s *TableConstructorContext) Metatable() IMetatableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetatableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetatableContext)
}

func (s *TableConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterTableConstructor(s)
	}
}

func (s *TableConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitTableConstructor(s)
	}
}

func (s *TableConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitTableConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) TableConstructor() (localctx ITableConstructorContext) {
	localctx = NewTableConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lua_grammar_antlr4ParserRULE_tableConstructor)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(lua_grammar_antlr4ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921505143718464) != 0) || ((int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&31745) != 0) {
		{
			p.SetState(152)
			p.Field()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(153)
					p.Match(lua_grammar_antlr4ParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(154)
					p.Field()
				}

			}
			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__7 {
		{
			p.SetState(162)
			p.Match(lua_grammar_antlr4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Metatable()
		}

	}
	{
		p.SetState(166)
		p.Match(lua_grammar_antlr4ParserT__9)
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

// IMetatableContext is an interface to support dynamic dispatch.
type IMetatableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMetatableContext differentiates from other interfaces.
	IsMetatableContext()
}

type MetatableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetatableContext() *MetatableContext {
	var p = new(MetatableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable
	return p
}

func InitEmptyMetatableContext(p *MetatableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable
}

func (*MetatableContext) IsMetatableContext() {}

func NewMetatableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetatableContext {
	var p = new(MetatableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metatable

	return p
}

func (s *MetatableContext) GetParser() antlr.Parser { return s.parser }

func (s *MetatableContext) Expression() IExpressionContext {
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

func (s *MetatableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetatableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetatableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMetatable(s)
	}
}

func (s *MetatableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMetatable(s)
	}
}

func (s *MetatableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMetatable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Metatable() (localctx IMetatableContext) {
	localctx = NewMetatableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lua_grammar_antlr4ParserRULE_metatable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(lua_grammar_antlr4ParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
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

// IMetamethodsContext is an interface to support dynamic dispatch.
type IMetamethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMetamethod() []IMetamethodContext
	Metamethod(i int) IMetamethodContext

	// IsMetamethodsContext differentiates from other interfaces.
	IsMetamethodsContext()
}

type MetamethodsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetamethodsContext() *MetamethodsContext {
	var p = new(MetamethodsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethods
	return p
}

func InitEmptyMetamethodsContext(p *MetamethodsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethods
}

func (*MetamethodsContext) IsMetamethodsContext() {}

func NewMetamethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetamethodsContext {
	var p = new(MetamethodsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_metamethods

	return p
}

func (s *MetamethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *MetamethodsContext) AllMetamethod() []IMetamethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetamethodContext); ok {
			len++
		}
	}

	tst := make([]IMetamethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetamethodContext); ok {
			tst[i] = t.(IMetamethodContext)
			i++
		}
	}

	return tst
}

func (s *MetamethodsContext) Metamethod(i int) IMetamethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetamethodContext); ok {
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

	return t.(IMetamethodContext)
}

func (s *MetamethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetamethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetamethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterMetamethods(s)
	}
}

func (s *MetamethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitMetamethods(s)
	}
}

func (s *MetamethodsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitMetamethods(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) Metamethods() (localctx IMetamethodsContext) {
	localctx = NewMetamethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lua_grammar_antlr4ParserRULE_metamethods)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Metamethod()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__7 {
		{
			p.SetState(173)
			p.Match(lua_grammar_antlr4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Metamethod()
		}

		p.SetState(179)
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

// ILabelStatementContext is an interface to support dynamic dispatch.
type ILabelStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLabelStatementContext differentiates from other interfaces.
	IsLabelStatementContext()
}

type LabelStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelStatementContext() *LabelStatementContext {
	var p = new(LabelStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_labelStatement
	return p
}

func InitEmptyLabelStatementContext(p *LabelStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_labelStatement
}

func (*LabelStatementContext) IsLabelStatementContext() {}

func NewLabelStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelStatementContext {
	var p = new(LabelStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_labelStatement

	return p
}

func (s *LabelStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *LabelStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLabelStatement(s)
	}
}

func (s *LabelStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLabelStatement(s)
	}
}

func (s *LabelStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLabelStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) LabelStatement() (localctx ILabelStatementContext) {
	localctx = NewLabelStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, lua_grammar_antlr4ParserRULE_labelStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(lua_grammar_antlr4ParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(lua_grammar_antlr4ParserT__11)
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
	p.EnterRule(localctx, 24, lua_grammar_antlr4ParserRULE_metamethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536862720) != 0) {
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
	IDENTIFIER() antlr.TerminalNode
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

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 26, lua_grammar_antlr4ParserRULE_field)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(lua_grammar_antlr4ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
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

// IBinaryOperationContext is an interface to support dynamic dispatch.
type IBinaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ArithOp() IArithOpContext
	BitwiseOp() IBitwiseOpContext
	ComparisonOp() IComparisonOpContext
	LogicalOp() ILogicalOpContext
	ConcatOp() IConcatOpContext
	CompoundAssignOp() ICompoundAssignOpContext
	CoalesceOp() ICoalesceOpContext
	ShiftAssignOp() IShiftAssignOpContext
	IncrOp() IIncrOpContext

	// IsBinaryOperationContext differentiates from other interfaces.
	IsBinaryOperationContext()
}

type BinaryOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperationContext() *BinaryOperationContext {
	var p = new(BinaryOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_binaryOperation
	return p
}

func InitEmptyBinaryOperationContext(p *BinaryOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_binaryOperation
}

func (*BinaryOperationContext) IsBinaryOperationContext() {}

func NewBinaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperationContext {
	var p = new(BinaryOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_binaryOperation

	return p
}

func (s *BinaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperationContext) AllExpression() []IExpressionContext {
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

func (s *BinaryOperationContext) Expression(i int) IExpressionContext {
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

func (s *BinaryOperationContext) ArithOp() IArithOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithOpContext)
}

func (s *BinaryOperationContext) BitwiseOp() IBitwiseOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitwiseOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitwiseOpContext)
}

func (s *BinaryOperationContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *BinaryOperationContext) LogicalOp() ILogicalOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOpContext)
}

func (s *BinaryOperationContext) ConcatOp() IConcatOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcatOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcatOpContext)
}

func (s *BinaryOperationContext) CompoundAssignOp() ICompoundAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundAssignOpContext)
}

func (s *BinaryOperationContext) CoalesceOp() ICoalesceOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoalesceOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoalesceOpContext)
}

func (s *BinaryOperationContext) ShiftAssignOp() IShiftAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftAssignOpContext)
}

func (s *BinaryOperationContext) IncrOp() IIncrOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrOpContext)
}

func (s *BinaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBinaryOperation(s)
	}
}

func (s *BinaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBinaryOperation(s)
	}
}

func (s *BinaryOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBinaryOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) BinaryOperation() (localctx IBinaryOperationContext) {
	localctx = NewBinaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, lua_grammar_antlr4ParserRULE_binaryOperation)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.expression(0)
		}
		{
			p.SetState(193)
			p.ArithOp()
		}
		{
			p.SetState(194)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.expression(0)
		}
		{
			p.SetState(197)
			p.BitwiseOp()
		}
		{
			p.SetState(198)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.expression(0)
		}
		{
			p.SetState(201)
			p.ComparisonOp()
		}
		{
			p.SetState(202)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.expression(0)
		}
		{
			p.SetState(205)
			p.LogicalOp()
		}
		{
			p.SetState(206)
			p.expression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(208)
			p.expression(0)
		}
		{
			p.SetState(209)
			p.ConcatOp()
		}
		{
			p.SetState(210)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(212)
			p.expression(0)
		}
		{
			p.SetState(213)
			p.CompoundAssignOp()
		}
		{
			p.SetState(214)
			p.expression(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(216)
			p.expression(0)
		}
		{
			p.SetState(217)
			p.CoalesceOp()
		}
		{
			p.SetState(218)
			p.expression(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(220)
			p.expression(0)
		}
		{
			p.SetState(221)
			p.ShiftAssignOp()
		}
		{
			p.SetState(222)
			p.expression(0)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(224)
			p.expression(0)
		}
		{
			p.SetState(225)
			p.IncrOp()
		}
		{
			p.SetState(226)
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

// IUnaryOperationContext is an interface to support dynamic dispatch.
type IUnaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOp() IUnaryOpContext
	Expression() IExpressionContext

	// IsUnaryOperationContext differentiates from other interfaces.
	IsUnaryOperationContext()
}

type UnaryOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperationContext() *UnaryOperationContext {
	var p = new(UnaryOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOperation
	return p
}

func InitEmptyUnaryOperationContext(p *UnaryOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOperation
}

func (*UnaryOperationContext) IsUnaryOperationContext() {}

func NewUnaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperationContext {
	var p = new(UnaryOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOperation

	return p
}

func (s *UnaryOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperationContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryOperationContext) Expression() IExpressionContext {
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

func (s *UnaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitUnaryOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) UnaryOperation() (localctx IUnaryOperationContext) {
	localctx = NewUnaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, lua_grammar_antlr4ParserRULE_unaryOperation)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.UnaryOp()
		}
		{
			p.SetState(231)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(lua_grammar_antlr4ParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
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

// IControlFlowStatementContext is an interface to support dynamic dispatch.
type IControlFlowStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	WhileStatement() IWhileStatementContext
	RepeatStatement() IRepeatStatementContext
	ForStatement() IForStatementContext
	BreakStatement() IBreakStatementContext
	GotoStatement() IGotoStatementContext
	CoroutineStatement() ICoroutineStatementContext

	// IsControlFlowStatementContext differentiates from other interfaces.
	IsControlFlowStatementContext()
}

type ControlFlowStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlFlowStatementContext() *ControlFlowStatementContext {
	var p = new(ControlFlowStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_controlFlowStatement
	return p
}

func InitEmptyControlFlowStatementContext(p *ControlFlowStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_controlFlowStatement
}

func (*ControlFlowStatementContext) IsControlFlowStatementContext() {}

func NewControlFlowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlFlowStatementContext {
	var p = new(ControlFlowStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_controlFlowStatement

	return p
}

func (s *ControlFlowStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlFlowStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *ControlFlowStatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *ControlFlowStatementContext) RepeatStatement() IRepeatStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatStatementContext)
}

func (s *ControlFlowStatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *ControlFlowStatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *ControlFlowStatementContext) GotoStatement() IGotoStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGotoStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGotoStatementContext)
}

func (s *ControlFlowStatementContext) CoroutineStatement() ICoroutineStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoroutineStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoroutineStatementContext)
}

func (s *ControlFlowStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlFlowStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlFlowStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterControlFlowStatement(s)
	}
}

func (s *ControlFlowStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitControlFlowStatement(s)
	}
}

func (s *ControlFlowStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitControlFlowStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ControlFlowStatement() (localctx IControlFlowStatementContext) {
	localctx = NewControlFlowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, lua_grammar_antlr4ParserRULE_controlFlowStatement)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case lua_grammar_antlr4ParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.IfStatement()
		}

	case lua_grammar_antlr4ParserT__34:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.WhileStatement()
		}

	case lua_grammar_antlr4ParserT__36:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.RepeatStatement()
		}

	case lua_grammar_antlr4ParserT__38:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
			p.ForStatement()
		}

	case lua_grammar_antlr4ParserT__39:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.BreakStatement()
		}

	case lua_grammar_antlr4ParserT__40:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(242)
			p.GotoStatement()
		}

	case lua_grammar_antlr4ParserT__41:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(243)
			p.CoroutineStatement()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllExpression() []IExpressionContext {
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

func (s *IfStatementContext) Expression(i int) IExpressionContext {
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

func (s *IfStatementContext) AllBlock() []IBlockContext {
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

func (s *IfStatementContext) Block(i int) IBlockContext {
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

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, lua_grammar_antlr4ParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(lua_grammar_antlr4ParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.expression(0)
	}
	{
		p.SetState(248)
		p.Match(lua_grammar_antlr4ParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Block()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == lua_grammar_antlr4ParserT__31 {
		{
			p.SetState(250)
			p.Match(lua_grammar_antlr4ParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.expression(0)
		}
		{
			p.SetState(252)
			p.Match(lua_grammar_antlr4ParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Block()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__32 {
		{
			p.SetState(260)
			p.Match(lua_grammar_antlr4ParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Block()
		}

	}
	{
		p.SetState(264)
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

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) Expression() IExpressionContext {
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

func (s *WhileStatementContext) Block() IBlockContext {
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

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, lua_grammar_antlr4ParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(lua_grammar_antlr4ParserT__34)
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
		p.Match(lua_grammar_antlr4ParserT__35)
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

// IRepeatStatementContext is an interface to support dynamic dispatch.
type IRepeatStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	Expression() IExpressionContext

	// IsRepeatStatementContext differentiates from other interfaces.
	IsRepeatStatementContext()
}

type RepeatStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatStatementContext() *RepeatStatementContext {
	var p = new(RepeatStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeatStatement
	return p
}

func InitEmptyRepeatStatementContext(p *RepeatStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeatStatement
}

func (*RepeatStatementContext) IsRepeatStatementContext() {}

func NewRepeatStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatStatementContext {
	var p = new(RepeatStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_repeatStatement

	return p
}

func (s *RepeatStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatStatementContext) Block() IBlockContext {
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

func (s *RepeatStatementContext) Expression() IExpressionContext {
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

func (s *RepeatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterRepeatStatement(s)
	}
}

func (s *RepeatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitRepeatStatement(s)
	}
}

func (s *RepeatStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitRepeatStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) RepeatStatement() (localctx IRepeatStatementContext) {
	localctx = NewRepeatStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, lua_grammar_antlr4ParserRULE_repeatStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(lua_grammar_antlr4ParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Block()
	}
	{
		p.SetState(274)
		p.Match(lua_grammar_antlr4ParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Block() IBlockContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
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

func (s *ForStatementContext) Expression(i int) IExpressionContext {
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

func (s *ForStatementContext) Block() IBlockContext {
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

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, lua_grammar_antlr4ParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(lua_grammar_antlr4ParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(lua_grammar_antlr4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.expression(0)
	}
	{
		p.SetState(281)
		p.Match(lua_grammar_antlr4ParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.expression(0)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__7 {
		{
			p.SetState(283)
			p.Match(lua_grammar_antlr4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.expression(0)
		}

	}
	{
		p.SetState(287)
		p.Match(lua_grammar_antlr4ParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Block()
	}
	{
		p.SetState(289)
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

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, lua_grammar_antlr4ParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(lua_grammar_antlr4ParserT__39)
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

// IGotoStatementContext is an interface to support dynamic dispatch.
type IGotoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsGotoStatementContext differentiates from other interfaces.
	IsGotoStatementContext()
}

type GotoStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoStatementContext() *GotoStatementContext {
	var p = new(GotoStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_gotoStatement
	return p
}

func InitEmptyGotoStatementContext(p *GotoStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_gotoStatement
}

func (*GotoStatementContext) IsGotoStatementContext() {}

func NewGotoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoStatementContext {
	var p = new(GotoStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_gotoStatement

	return p
}

func (s *GotoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *GotoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterGotoStatement(s)
	}
}

func (s *GotoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitGotoStatement(s)
	}
}

func (s *GotoStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitGotoStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) GotoStatement() (localctx IGotoStatementContext) {
	localctx = NewGotoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, lua_grammar_antlr4ParserRULE_gotoStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(lua_grammar_antlr4ParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
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

// ICoroutineStatementContext is an interface to support dynamic dispatch.
type ICoroutineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCoroutineStatementContext differentiates from other interfaces.
	IsCoroutineStatementContext()
}

type CoroutineStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoroutineStatementContext() *CoroutineStatementContext {
	var p = new(CoroutineStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutineStatement
	return p
}

func InitEmptyCoroutineStatementContext(p *CoroutineStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutineStatement
}

func (*CoroutineStatementContext) IsCoroutineStatementContext() {}

func NewCoroutineStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoroutineStatementContext {
	var p = new(CoroutineStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coroutineStatement

	return p
}

func (s *CoroutineStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *CoroutineStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoroutineStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoroutineStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterCoroutineStatement(s)
	}
}

func (s *CoroutineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitCoroutineStatement(s)
	}
}

func (s *CoroutineStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitCoroutineStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) CoroutineStatement() (localctx ICoroutineStatementContext) {
	localctx = NewCoroutineStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, lua_grammar_antlr4ParserRULE_coroutineStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(lua_grammar_antlr4ParserT__41)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(lua_grammar_antlr4ParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0) {
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.EnterRule(localctx, 48, lua_grammar_antlr4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1978744046620672) != 0) || _la == lua_grammar_antlr4ParserIDENTIFIER {
		{
			p.SetState(300)
			p.Statement()
		}

		p.SetState(305)
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

// ILocalDeclarationContext is an interface to support dynamic dispatch.
type ILocalDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsLocalDeclarationContext differentiates from other interfaces.
	IsLocalDeclarationContext()
}

type LocalDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalDeclarationContext() *LocalDeclarationContext {
	var p = new(LocalDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_localDeclaration
	return p
}

func InitEmptyLocalDeclarationContext(p *LocalDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_localDeclaration
}

func (*LocalDeclarationContext) IsLocalDeclarationContext() {}

func NewLocalDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalDeclarationContext {
	var p = new(LocalDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_localDeclaration

	return p
}

func (s *LocalDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, 0)
}

func (s *LocalDeclarationContext) Expression() IExpressionContext {
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

func (s *LocalDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLocalDeclaration(s)
	}
}

func (s *LocalDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLocalDeclaration(s)
	}
}

func (s *LocalDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLocalDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) LocalDeclaration() (localctx ILocalDeclarationContext) {
	localctx = NewLocalDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, lua_grammar_antlr4ParserRULE_localDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(lua_grammar_antlr4ParserT__47)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserT__0 {
		{
			p.SetState(308)
			p.Match(lua_grammar_antlr4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
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

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	Block() IBlockContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(lua_grammar_antlr4ParserIDENTIFIER)
}

func (s *FunctionDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(lua_grammar_antlr4ParserIDENTIFIER, i)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
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

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, lua_grammar_antlr4ParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(lua_grammar_antlr4ParserT__48)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(lua_grammar_antlr4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(lua_grammar_antlr4ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == lua_grammar_antlr4ParserIDENTIFIER {
		{
			p.SetState(315)
			p.Match(lua_grammar_antlr4ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__7 {
			{
				p.SetState(316)
				p.Match(lua_grammar_antlr4ParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(317)
				p.Match(lua_grammar_antlr4ParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(325)
		p.Match(lua_grammar_antlr4ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.Block()
	}
	{
		p.SetState(327)
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) AllExpression() []IExpressionContext {
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

func (s *ReturnStatementContext) Expression(i int) IExpressionContext {
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

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, lua_grammar_antlr4ParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(lua_grammar_antlr4ParserT__49)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(330)
			p.expression(0)
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == lua_grammar_antlr4ParserT__7 {
			{
				p.SetState(331)
				p.Match(lua_grammar_antlr4ParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(332)
				p.expression(0)
			}

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IOperatorGroupContext is an interface to support dynamic dispatch.
type IOperatorGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOp() ILogicalOpContext
	ComparisonOp() IComparisonOpContext
	ArithOp() IArithOpContext
	BitwiseOp() IBitwiseOpContext
	AssignOp() IAssignOpContext
	UnaryOp() IUnaryOpContext
	ConcatOp() IConcatOpContext
	CompoundAssignOp() ICompoundAssignOpContext
	IncrOp() IIncrOpContext
	CoalesceOp() ICoalesceOpContext
	ShiftAssignOp() IShiftAssignOpContext

	// IsOperatorGroupContext differentiates from other interfaces.
	IsOperatorGroupContext()
}

type OperatorGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorGroupContext() *OperatorGroupContext {
	var p = new(OperatorGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operatorGroup
	return p
}

func InitEmptyOperatorGroupContext(p *OperatorGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operatorGroup
}

func (*OperatorGroupContext) IsOperatorGroupContext() {}

func NewOperatorGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorGroupContext {
	var p = new(OperatorGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_operatorGroup

	return p
}

func (s *OperatorGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorGroupContext) LogicalOp() ILogicalOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOpContext)
}

func (s *OperatorGroupContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *OperatorGroupContext) ArithOp() IArithOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithOpContext)
}

func (s *OperatorGroupContext) BitwiseOp() IBitwiseOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitwiseOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitwiseOpContext)
}

func (s *OperatorGroupContext) AssignOp() IAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignOpContext)
}

func (s *OperatorGroupContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *OperatorGroupContext) ConcatOp() IConcatOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcatOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcatOpContext)
}

func (s *OperatorGroupContext) CompoundAssignOp() ICompoundAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundAssignOpContext)
}

func (s *OperatorGroupContext) IncrOp() IIncrOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrOpContext)
}

func (s *OperatorGroupContext) CoalesceOp() ICoalesceOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoalesceOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoalesceOpContext)
}

func (s *OperatorGroupContext) ShiftAssignOp() IShiftAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftAssignOpContext)
}

func (s *OperatorGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterOperatorGroup(s)
	}
}

func (s *OperatorGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitOperatorGroup(s)
	}
}

func (s *OperatorGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitOperatorGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) OperatorGroup() (localctx IOperatorGroupContext) {
	localctx = NewOperatorGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, lua_grammar_antlr4ParserRULE_operatorGroup)
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.LogicalOp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(341)
			p.ComparisonOp()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(342)
			p.ArithOp()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(343)
			p.BitwiseOp()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(344)
			p.AssignOp()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(345)
			p.UnaryOp()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(346)
			p.ConcatOp()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(347)
			p.CompoundAssignOp()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(348)
			p.IncrOp()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(349)
			p.CoalesceOp()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(350)
			p.ShiftAssignOp()
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

// ILogicalOpContext is an interface to support dynamic dispatch.
type ILogicalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogicalOpContext differentiates from other interfaces.
	IsLogicalOpContext()
}

type LogicalOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOpContext() *LogicalOpContext {
	var p = new(LogicalOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logicalOp
	return p
}

func InitEmptyLogicalOpContext(p *LogicalOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logicalOp
}

func (*LogicalOpContext) IsLogicalOpContext() {}

func NewLogicalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOpContext {
	var p = new(LogicalOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_logicalOp

	return p
}

func (s *LogicalOpContext) GetParser() antlr.Parser { return s.parser }
func (s *LogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterLogicalOp(s)
	}
}

func (s *LogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitLogicalOp(s)
	}
}

func (s *LogicalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitLogicalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) LogicalOp() (localctx ILogicalOpContext) {
	localctx = NewLogicalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, lua_grammar_antlr4ParserRULE_logicalOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserT__50 || _la == lua_grammar_antlr4ParserT__51) {
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

// IComparisonOpContext is an interface to support dynamic dispatch.
type IComparisonOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparisonOpContext differentiates from other interfaces.
	IsComparisonOpContext()
}

type ComparisonOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOpContext() *ComparisonOpContext {
	var p = new(ComparisonOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparisonOp
	return p
}

func InitEmptyComparisonOpContext(p *ComparisonOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparisonOp
}

func (*ComparisonOpContext) IsComparisonOpContext() {}

func NewComparisonOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOpContext {
	var p = new(ComparisonOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_comparisonOp

	return p
}

func (s *ComparisonOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterComparisonOp(s)
	}
}

func (s *ComparisonOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitComparisonOp(s)
	}
}

func (s *ComparisonOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitComparisonOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ComparisonOp() (localctx IComparisonOpContext) {
	localctx = NewComparisonOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, lua_grammar_antlr4ParserRULE_comparisonOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&567453553048682496) != 0) {
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

// IArithOpContext is an interface to support dynamic dispatch.
type IArithOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArithOpContext differentiates from other interfaces.
	IsArithOpContext()
}

type ArithOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithOpContext() *ArithOpContext {
	var p = new(ArithOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arithOp
	return p
}

func InitEmptyArithOpContext(p *ArithOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arithOp
}

func (*ArithOpContext) IsArithOpContext() {}

func NewArithOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithOpContext {
	var p = new(ArithOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_arithOp

	return p
}

func (s *ArithOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ArithOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterArithOp(s)
	}
}

func (s *ArithOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitArithOp(s)
	}
}

func (s *ArithOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitArithOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ArithOp() (localctx IArithOpContext) {
	localctx = NewArithOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, lua_grammar_antlr4ParserRULE_arithOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-59)) & ^0x3f) == 0 && ((int64(1)<<(_la-59))&127) != 0) {
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

// IBitwiseOpContext is an interface to support dynamic dispatch.
type IBitwiseOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBitwiseOpContext differentiates from other interfaces.
	IsBitwiseOpContext()
}

type BitwiseOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitwiseOpContext() *BitwiseOpContext {
	var p = new(BitwiseOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwiseOp
	return p
}

func InitEmptyBitwiseOpContext(p *BitwiseOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwiseOp
}

func (*BitwiseOpContext) IsBitwiseOpContext() {}

func NewBitwiseOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitwiseOpContext {
	var p = new(BitwiseOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_bitwiseOp

	return p
}

func (s *BitwiseOpContext) GetParser() antlr.Parser { return s.parser }
func (s *BitwiseOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitwiseOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterBitwiseOp(s)
	}
}

func (s *BitwiseOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitBitwiseOp(s)
	}
}

func (s *BitwiseOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitBitwiseOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) BitwiseOp() (localctx IBitwiseOpContext) {
	localctx = NewBitwiseOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, lua_grammar_antlr4ParserRULE_bitwiseOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&31) != 0) {
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

// IAssignOpContext is an interface to support dynamic dispatch.
type IAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignOpContext differentiates from other interfaces.
	IsAssignOpContext()
}

type AssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOpContext() *AssignOpContext {
	var p = new(AssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignOp
	return p
}

func InitEmptyAssignOpContext(p *AssignOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignOp
}

func (*AssignOpContext) IsAssignOpContext() {}

func NewAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOpContext {
	var p = new(AssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_assignOp

	return p
}

func (s *AssignOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterAssignOp(s)
	}
}

func (s *AssignOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitAssignOp(s)
	}
}

func (s *AssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) AssignOp() (localctx IAssignOpContext) {
	localctx = NewAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, lua_grammar_antlr4ParserRULE_assignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) || ((int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&15) != 0)) {
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

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOp
	return p
}

func InitEmptyUnaryOpContext(p *UnaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOp
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, lua_grammar_antlr4ParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-29)) & ^0x3f) == 0 && ((int64(1)<<(_la-29))&70370891661313) != 0) {
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

// IConcatOpContext is an interface to support dynamic dispatch.
type IConcatOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConcatOpContext differentiates from other interfaces.
	IsConcatOpContext()
}

type ConcatOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatOpContext() *ConcatOpContext {
	var p = new(ConcatOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concatOp
	return p
}

func InitEmptyConcatOpContext(p *ConcatOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concatOp
}

func (*ConcatOpContext) IsConcatOpContext() {}

func NewConcatOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatOpContext {
	var p = new(ConcatOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_concatOp

	return p
}

func (s *ConcatOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ConcatOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterConcatOp(s)
	}
}

func (s *ConcatOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitConcatOp(s)
	}
}

func (s *ConcatOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitConcatOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ConcatOp() (localctx IConcatOpContext) {
	localctx = NewConcatOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, lua_grammar_antlr4ParserRULE_concatOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(lua_grammar_antlr4ParserT__75)
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

// IVarargOpContext is an interface to support dynamic dispatch.
type IVarargOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarargOpContext differentiates from other interfaces.
	IsVarargOpContext()
}

type VarargOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarargOpContext() *VarargOpContext {
	var p = new(VarargOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_varargOp
	return p
}

func InitEmptyVarargOpContext(p *VarargOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_varargOp
}

func (*VarargOpContext) IsVarargOpContext() {}

func NewVarargOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarargOpContext {
	var p = new(VarargOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_varargOp

	return p
}

func (s *VarargOpContext) GetParser() antlr.Parser { return s.parser }
func (s *VarargOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarargOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarargOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterVarargOp(s)
	}
}

func (s *VarargOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitVarargOp(s)
	}
}

func (s *VarargOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitVarargOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) VarargOp() (localctx IVarargOpContext) {
	localctx = NewVarargOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, lua_grammar_antlr4ParserRULE_varargOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(lua_grammar_antlr4ParserT__76)
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

// ICompoundAssignOpContext is an interface to support dynamic dispatch.
type ICompoundAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCompoundAssignOpContext differentiates from other interfaces.
	IsCompoundAssignOpContext()
}

type CompoundAssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundAssignOpContext() *CompoundAssignOpContext {
	var p = new(CompoundAssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_compoundAssignOp
	return p
}

func InitEmptyCompoundAssignOpContext(p *CompoundAssignOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_compoundAssignOp
}

func (*CompoundAssignOpContext) IsCompoundAssignOpContext() {}

func NewCompoundAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundAssignOpContext {
	var p = new(CompoundAssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_compoundAssignOp

	return p
}

func (s *CompoundAssignOpContext) GetParser() antlr.Parser { return s.parser }
func (s *CompoundAssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundAssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundAssignOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterCompoundAssignOp(s)
	}
}

func (s *CompoundAssignOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitCompoundAssignOp(s)
	}
}

func (s *CompoundAssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitCompoundAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) CompoundAssignOp() (localctx ICompoundAssignOpContext) {
	localctx = NewCompoundAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, lua_grammar_antlr4ParserRULE_compoundAssignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60) != 0) || ((int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&387) != 0)) {
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

// IIncrOpContext is an interface to support dynamic dispatch.
type IIncrOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncrOpContext differentiates from other interfaces.
	IsIncrOpContext()
}

type IncrOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrOpContext() *IncrOpContext {
	var p = new(IncrOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_incrOp
	return p
}

func InitEmptyIncrOpContext(p *IncrOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_incrOp
}

func (*IncrOpContext) IsIncrOpContext() {}

func NewIncrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrOpContext {
	var p = new(IncrOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_incrOp

	return p
}

func (s *IncrOpContext) GetParser() antlr.Parser { return s.parser }
func (s *IncrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterIncrOp(s)
	}
}

func (s *IncrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitIncrOp(s)
	}
}

func (s *IncrOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitIncrOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) IncrOp() (localctx IIncrOpContext) {
	localctx = NewIncrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, lua_grammar_antlr4ParserRULE_incrOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserT__79 || _la == lua_grammar_antlr4ParserT__80) {
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

// ICoalesceOpContext is an interface to support dynamic dispatch.
type ICoalesceOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCoalesceOpContext differentiates from other interfaces.
	IsCoalesceOpContext()
}

type CoalesceOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoalesceOpContext() *CoalesceOpContext {
	var p = new(CoalesceOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coalesceOp
	return p
}

func InitEmptyCoalesceOpContext(p *CoalesceOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coalesceOp
}

func (*CoalesceOpContext) IsCoalesceOpContext() {}

func NewCoalesceOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoalesceOpContext {
	var p = new(CoalesceOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_coalesceOp

	return p
}

func (s *CoalesceOpContext) GetParser() antlr.Parser { return s.parser }
func (s *CoalesceOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoalesceOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterCoalesceOp(s)
	}
}

func (s *CoalesceOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitCoalesceOp(s)
	}
}

func (s *CoalesceOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitCoalesceOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) CoalesceOp() (localctx ICoalesceOpContext) {
	localctx = NewCoalesceOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, lua_grammar_antlr4ParserRULE_coalesceOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(lua_grammar_antlr4ParserT__81)
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

// IShiftAssignOpContext is an interface to support dynamic dispatch.
type IShiftAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsShiftAssignOpContext differentiates from other interfaces.
	IsShiftAssignOpContext()
}

type ShiftAssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftAssignOpContext() *ShiftAssignOpContext {
	var p = new(ShiftAssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_shiftAssignOp
	return p
}

func InitEmptyShiftAssignOpContext(p *ShiftAssignOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = lua_grammar_antlr4ParserRULE_shiftAssignOp
}

func (*ShiftAssignOpContext) IsShiftAssignOpContext() {}

func NewShiftAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftAssignOpContext {
	var p = new(ShiftAssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = lua_grammar_antlr4ParserRULE_shiftAssignOp

	return p
}

func (s *ShiftAssignOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ShiftAssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftAssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftAssignOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.EnterShiftAssignOp(s)
	}
}

func (s *ShiftAssignOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lua_grammar_antlr4Listener); ok {
		listenerT.ExitShiftAssignOp(s)
	}
}

func (s *ShiftAssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case lua_grammar_antlr4Visitor:
		return t.VisitShiftAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *lua_grammar_antlr4Parser) ShiftAssignOp() (localctx IShiftAssignOpContext) {
	localctx = NewShiftAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, lua_grammar_antlr4ParserRULE_shiftAssignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lua_grammar_antlr4ParserT__82 || _la == lua_grammar_antlr4ParserT__83) {
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
	case 3:
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
