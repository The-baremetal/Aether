# Parser Specification

This is the parser specification for the Aether language.

## Structure

The parser is a `Parser` struct, which has these fields:

- `lexer`: Pointer to parser
- `errors`: Strings that has errors which may have been unresolved by the parser.
- `curToken`: The token that is being parsed
- `peekToken`: The next token in stream

## Operator Precedence

The parser precedence is this with new ones being added.

- `LOWEST`: 1
- `COMPARISON`: 2 (`==`, `<`, `>`)
- `CONCAT`: 3 (`.`)
- `SUM`: 4 (`+`, `-`)
- `PRODUCT`: 5 (`*`, `/`, `%`)
- `EXPONENT`: 6 (`^`)

The `precedences` map in the `parser.go` file maintains priority.