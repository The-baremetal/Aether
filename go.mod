module FLUX

go 1.22

toolchain go1.22.11

require github.com/antlr4-go/antlr/v4 v4.13.1

require (
	FLUX/src/aether/ast v0.0.0-00010101000000-000000000000 // indirect
	FLUX/src/aether/parser v0.0.0-00010101000000-000000000000 // indirect
	FLUX/src/aether/utils v0.0.0-00010101000000-000000000000 // indirect
	FLUX/src/lib/lexer v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
)

replace FLUX/src/aether/ast => ./src/aether/ast

replace FLUX/src/aether/visitor => ./src/aether/visitor

replace FLUX/src/aether/codegen => ./src/aether/codegen

replace FLUX/src/aether/parser => ./src/aether/parser

replace FLUX/src/lib/lexer => ./src/lib/lexer

replace FLUX/src/aether/utils => ./src/aether/utils
