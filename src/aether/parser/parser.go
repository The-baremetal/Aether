package parser

// The parser of AetherC

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
	"FLUX/src/aether/lexer"
    "fmt"
    "strconv"
)

const (
    LOWEST = iota
    ASSIGN      // =
    OR          // or
    AND         // and
    EQUALS      // ==, !=
    COMPARISON  // <, >, <=, >=
    SUM         // +, -
    PRODUCT     // *, /, %
    PREFIX      // -X, !X
    CALL        // function()
    INDEX       // []
)

type Parser struct { // defining parser as a structure that can be created
	lexer         *lexer.Lexer
	currentToken  lexer.Token
	peekToken     lexer.Token
	errors        []string
	prefixParseFns map[lexer.TokenType]func() Expression
	infixParseFns  map[lexer.TokenType]func(Expression) Expression
	precedences    map[lexer.TokenType]int
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[lexer.TokenType]func() Expression)
	p.registerPrefix(lexer.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(lexer.NUMBER, p.parseNumberLiteral)
	p.registerPrefix(lexer.STRING, p.parseStringLiteral)
	p.registerPrefix(lexer.TRUE, p.parseBoolean)
	p.registerPrefix(lexer.FALSE, p.parseBoolean)
	p.registerPrefix(lexer.NIL, p.parseNilLiteral)
	p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)
	p.registerPrefix(lexer.NOT, p.parsePrefixExpression)
	p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(lexer.LBRACE, p.parseTableConstructor)
	p.registerPrefix(lexer.LBRACKET, p.parseArrayConstructor)
	p.registerPrefix(lexer.FUNCTION, p.parseFunctionLiteral)
    p.registerPrefix(lexer.VAARG, p.parseVarArgs)

	p.infixParseFns = make(map[lexer.TokenType]func(Expression) Expression)
	p.registerInfix(lexer.PLUS, p.parseInfixExpression)
	p.registerInfix(lexer.MINUS, p.parseInfixExpression)
	p.registerInfix(lexer.MULTIPLY, p.parseInfixExpression)
	p.registerInfix(lexer.DIVIDE, p.parseInfixExpression)
	p.registerInfix(lexer.MODULO, p.parseInfixExpression)
	p.registerInfix(lexer.EXPONENT, p.parseInfixExpression)
	p.registerInfix(lexer.EQ, p.parseInfixExpression)
	p.registerInfix(lexer.NE, p.parseInfixExpression)
	p.registerInfix(lexer.LT, p.parseInfixExpression)
	p.registerInfix(lexer.LE, p.parseInfixExpression)
	p.registerInfix(lexer.GT, p.parseInfixExpression)
	p.registerInfix(lexer.GE, p.parseInfixExpression)
	p.registerInfix(lexer.AND, p.parseInfixExpression)
	p.registerInfix(lexer.OR, p.parseInfixExpression)
	p.registerInfix(lexer.CONCAT, p.parseInfixExpression)
	p.registerInfix(lexer.LPAREN, p.parseCallExpression)
	p.registerInfix(lexer.LBRACKET, p.parseIndexExpression)
	p.registerInfix(lexer.DOT, p.parseDotExpression)
	p.registerInfix(lexer.ASSIGN, p.parseAssignExpression)
    p.registerInfix(lexer.COLON, p.parseMethodCall)

	p.registerPrefix(lexer.ASYNC, p.parseAsyncExpression)
	p.registerPrefix(lexer.AWAIT, p.parseAwaitExpression)

	p.precedences = map[lexer.TokenType]int{
    lexer.ASSIGN:   10,
    lexer.OR:       20,
    lexer.AND:      30,
    lexer.IN:       40,
    lexer.EQ:       40, lexer.NE: 40, lexer.LT: 40, lexer.LE: 40, lexer.GT: 40, lexer.GE: 40,
    lexer.CONCAT:   50,
    lexer.PLUS:     60, lexer.MINUS: 60,
    lexer.MULTIPLY: 70, lexer.DIVIDE: 70, lexer.MODULO: 70, lexer.EXPONENT: 70,
    lexer.LPAREN:   80, lexer.LBRACKET: 80, lexer.DOT: 80, lexer.COLON: 80,
}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *Program {
    program := &Program{}
    for p.currentToken.Type != lexer.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case lexer.LOCAL:
		return p.parseLocalStatement()
	case lexer.FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.CONST:
		return p.parseConstDeclaration()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.WHILE:
		return p.parseWhileStatement()
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.REPEAT:
		return p.parseRepeatStatement()
	case lexer.BREAK:
		return p.parseBreakStatement()
	case lexer.TRY:
		return p.parseTryCatchStatement()
	case lexer.COROUTINE:
		return p.parseCoroutineStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseConstDeclaration() Statement {
    stmt := &ConstDeclaration{Token: p.currentToken}
    p.nextToken()
    stmt.Names = p.parseIdentifierList()
    
    if !p.expectPeek(lexer.ASSIGN) {
        return nil
    }
    
    p.nextToken()
    stmt.Values = p.parseExpressionList(lexer.ASSIGN)
    return stmt
}

func (p *Parser) parseMethodCall(left Expression) Expression {
    expr := &MethodCallExpression{
        Token:   p.currentToken,
        Object:  left,
    }
    p.nextToken()

    if !p.expectPeek(lexer.IDENTIFIER) {
        return nil
    }
    expr.Method = p.parseIdentifier().(*Identifier)

    if !p.expectPeek(lexer.LPAREN) {
        return nil
    }
    expr.Arguments = p.parseArgumentList()

    return expr
}

func (p *Parser) parseReturnStatement() Statement {
    stmt := &ReturnStatement{Token: p.currentToken}
    p.nextToken()
    stmt.Value = p.parseExpression(LOWEST)
    return stmt
}

func (p *Parser) parseRepeatStatement() Statement {
    stmt := &RepeatStatement{Token: p.currentToken}
    stmt.Body = p.parseBlockStatement(lexer.UNTIL)
    if !p.expectPeek(lexer.UNTIL) {
        return nil
    }
    
    p.nextToken()
    stmt.Condition = p.parseExpression(LOWEST)
    return stmt
}

func (p *Parser) parseBreakStatement() Statement {
    return &BreakStatement{Token: p.currentToken}
}

func (p *Parser) parseExpressionStatement() Statement {
    expr := p.parseExpression(LOWEST)
    return &ExpressionStatement{Expression: expr}
}

func (p *Parser) parseLocalStatement() *LocalDeclaration {
    stmt := &LocalDeclaration{Token: p.currentToken}
    p.nextToken()
    
    if p.currentToken.Type == lexer.CONST {
        stmt.Constant = true
        p.nextToken()
    }
    
    stmt.Names = p.parseIdentifierList()
    
    if !p.expectPeek(lexer.ASSIGN) {
        return nil
    }
    
    p.nextToken()

    var values []Expression
    values = append(values, p.parseExpression(LOWEST))
    
    for p.peekTokenIs(lexer.COMMA) {
        p.nextToken()
        p.nextToken()
        values = append(values, p.parseExpression(LOWEST))
    }
    
    stmt.Values = values
    
    return stmt
}
func (p *Parser) parseCallExpression(function Expression) Expression {
    expr := &CallExpression{
        Token:    p.currentToken,
        Function: function,
    }
    p.nextToken()
    expr.Arguments = p.parseArgumentList()
    return expr
}

func (p *Parser) parseArgumentList() []Expression {
    args := []Expression{}
    if p.peekTokenIs(lexer.RPAREN) {
        p.nextToken()
        return args
    }
    p.nextToken()
    args = append(args, p.parseArgument())
    for p.peekTokenIs(lexer.COMMA) {
        p.nextToken()
        p.nextToken()
        args = append(args, p.parseArgument())
    }
    if !p.expectPeek(lexer.RPAREN) {
        return nil
    }
    return args
}

func (p *Parser) parseArgument() Expression {
    if p.currentToken.Type == lexer.IDENTIFIER && p.peekTokenIs(lexer.ASSIGN) {
        nameExpr := p.parseIdentifier()
        name, ok := nameExpr.(*Identifier)
        if !ok {
            p.errors = append(p.errors, "expected identifier")
            return nil
        }
        p.nextToken()
        p.nextToken()
        value := p.parseExpression(LOWEST)
        return &NamedArgument{Name: name, Value: value}
    }
    return p.parseExpression(LOWEST)
}

func (p *Parser) parseAsyncExpression() Expression {
    expr := &AsyncExpression{Token: p.currentToken}
    p.nextToken()
    expr.Block = p.parseBlockStatement(lexer.END)
    if !p.expectPeek(lexer.END) {
        return nil
    }
    return expr
}

func (p *Parser) parseAwaitExpression() Expression {
    expr := &AwaitExpression{Token: p.currentToken}
    p.nextToken()
    expr.Expression = p.parseExpression(PREFIX)
    return expr
}

func (p *Parser) parseNumberLiteral() Expression {
    value, err := strconv.ParseFloat(p.currentToken.Literal, 64)
    if err != nil {
        p.errors = append(p.errors, fmt.Sprintf("could not parse %q as number", p.currentToken.Literal))
        return nil
    }
    return &NumberLiteral{
        Token: p.currentToken,
        Value: value,
    }
}

func (p *Parser) parseStringLiteral() Expression {
    return &StringLiteral{
        Token: p.currentToken,
        Value: p.currentToken.Literal,
    }
}

func (p *Parser) parseBoolean() Expression {
    value := p.currentToken.Type == lexer.TRUE
    return &BooleanLiteral{
        Token: p.currentToken,
        Value: value,
    }
}

func (p *Parser) parseNilLiteral() Expression {
    return &NilLiteral{Token: p.currentToken}
}

func (p *Parser) parseGroupedExpression() Expression {
    p.nextToken()
    exp := p.parseExpression(LOWEST)
    if !p.expectPeek(lexer.RPAREN) {
        return nil
    }
    return exp
}

func (p *Parser) parseIndexExpression(left Expression) Expression {
    exp := &IndexExpression{
        Token: p.currentToken,
        Left:  left,
    }
    p.nextToken()
    exp.Index = p.parseExpression(LOWEST)
    if !p.expectPeek(lexer.RBRACKET) {
        return nil
    }
    return exp
}

func (p *Parser) parseDotExpression(left Expression) Expression {
    exp := &DotExpression{
        Token:    p.currentToken,
        Object:   left,
    }
    if !p.expectPeek(lexer.IDENTIFIER) {
        return nil
    }
    ident, ok := p.parseIdentifier().(*Identifier)
	if !ok {
		p.errors = append(p.errors, "expected identifier")
		return nil
	}
    exp.Property = ident
    return exp
}

func (p *Parser) parseAssignExpression(left Expression) Expression {
    exp := &AssignExpression{
        Token: p.currentToken,
        Left:  left,
    }
    precedence := p.curPrecedence()
    p.nextToken()
    exp.Value = p.parseExpression(precedence)
    return exp
}

func (p *Parser) parseFunctionDeclaration() *FunctionDeclaration {
    stmt := &FunctionDeclaration{Token: p.currentToken}
    p.nextToken()
    var hasColon bool
    if p.peekToken.Type == lexer.COLON {
        hasColon = true
        p.nextToken()
        ident, ok := p.parseIdentifier().(*Identifier)
        if !ok {
            p.errors = append(p.errors, "expected identifier")
            return nil
        }
        stmt.Receiver = ident
        p.nextToken()
    }
    stmt.Name = p.parseIdentifier().(*Identifier)
    if !p.expectPeek(lexer.LPAREN) {
        p.errors = append(p.errors, fmt.Sprintf("expected '(' after function name at line %d, column %d, got %s (%q) instead", p.peekToken.Line, p.peekToken.Column, p.peekToken.Type, p.peekToken.Literal))
        return nil
    }
    p.nextToken()
    if hasColon {
        selfParam := &Identifier{
            Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: "self"},
            Value: "self",
        }
        stmt.Parameters = append([]*Identifier{selfParam}, p.parseFunctionParameters()...)
    } else {
        stmt.Parameters = p.parseFunctionParameters()
    }
    stmt.Body = p.parseBlockStatement(lexer.END)
    if p.currentToken.Type != lexer.END {
        p.errors = append(p.errors, fmt.Sprintf("expected 'end' after function body at line %d, column %d, got %s (%q) instead", p.currentToken.Line, p.currentToken.Column, p.currentToken.Type, p.currentToken.Literal))
        return nil
    }
    p.nextToken()
    return stmt
}

func (p *Parser) parseVarArgs() Expression {
    return &VarArgLiteral{
        Token: p.currentToken,
    }
}
func (p *Parser) parseFunctionParameters() []*Identifier { // issue #5 made by luohoa97: when you parse more than 1 argument, it gives expected ')' after function parameters at line 0, column 0, got COMMA (",")
    var params []*Identifier
    if p.peekTokenIs(lexer.RPAREN) {
        p.nextToken()
        return params
    }
    p.nextToken()
    if p.currentToken.Type == lexer.VAARG {
        params = append(params, &Identifier{Token: p.currentToken, Value: "..."})
        p.nextToken()
    } else {
        params = append(params, p.parseIdentifier().(*Identifier))
        for p.peekTokenIs(lexer.COMMA) {
            p.nextToken()
            p.nextToken()
            if p.currentToken.Type == lexer.VAARG {
                params = append(params, &Identifier{Token: p.currentToken, Value: "..."})
                p.nextToken()
                break
            }
            params = append(params, p.parseIdentifier().(*Identifier))
        }
    }
    if p.currentToken.Type != lexer.RPAREN {
        p.errors = append(p.errors, fmt.Sprintf(
            "expected ')' after function parameters at line %d, column %d, got %s (%q)",
            p.currentToken.Line, p.currentToken.Column, p.currentToken.Type, p.currentToken.Literal))
        return nil
    }

    return params
}

func (p *Parser) parseTryCatchStatement() *TryCatchStatement {
    stmt := &TryCatchStatement{Token: p.currentToken}
    p.nextToken()

    stmt.TryBlock = p.parseBlockStatement(lexer.END)

    if p.peekTokenIs(lexer.CATCH) {
        p.nextToken()
        if p.peekTokenIs(lexer.IDENTIFIER) {
            p.nextToken()
            stmt.CatchVar = p.parseIdentifier().(*Identifier)
        }
        stmt.CatchBlock = p.parseBlockStatement(lexer.END)
    }

    if p.peekTokenIs(lexer.FINALLY) {
        p.nextToken()
        stmt.FinallyBlock = p.parseBlockStatement(lexer.END)
    }

    if !p.expectPeek(lexer.END) {
        return nil
    }

    return stmt
}

func (p *Parser) parseExpression(precedence int) Expression {
    prefix := p.prefixParseFns[p.currentToken.Type]
    if prefix == nil {
        p.noPrefixParseFnError(p.currentToken.Type)
        return nil
    }
    leftExp := prefix()
    
    for precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        p.nextToken()
        leftExp = infix(leftExp)
    }
    
    return leftExp
}

func (p *Parser) parsePrefixExpression() Expression {
    expr := &PrefixExpression{
        Token:    p.currentToken,
        Operator: p.currentToken.Literal,
    }
    p.nextToken()
    expr.Right = p.parseExpression(PREFIX)
    return expr
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
    expr := &InfixExpression{
        Token:    p.currentToken,
        Left:     left,
        Operator: p.currentToken.Literal,
    }
    
    precedence := p.curPrecedence()
    p.nextToken()
    expr.Right = p.parseExpression(precedence)
    return expr
}

func (p *Parser) parseCoroutineStatement() Statement {
    stmt := &CoroutineStatement{Token: p.currentToken}
    p.nextToken()

    if !p.expectPeek(lexer.DOT) {
        return nil
    }
    p.nextToken()

    method := p.currentToken.Literal
    stmt.Method = method
    p.nextToken()

    switch method {
    case "create", "resume", "wrap":
        if !p.expectPeek(lexer.LPAREN) {
            return nil
        }
        stmt.Args = p.parseExpressionList(lexer.RPAREN)
    case "yield", "status", "running", "isyieldable":
    default:
        p.errors = append(p.errors, "unknown coroutine method: "+method)
        return nil
    }

    return stmt
}

func (p *Parser) nextToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curPrecedence() int {
    if p, ok := p.precedences[p.currentToken.Type]; ok {
        return p
    }
    return LOWEST
}

func (p *Parser) peekPrecedence() int {
    if p, ok := p.precedences[p.peekToken.Type]; ok {
        return p
    }
    return LOWEST
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    }
    p.peekError(t)
    return false
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn func() Expression) {
    p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn func(Expression) Expression) {
    p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
    p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s (%q) found at line %d, column %d", t, p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column))
}

func (p *Parser) peekError(t lexer.TokenType) {
    p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, got %s (%q) instead at line %d, column %d",
        t, p.peekToken.Type, p.peekToken.Literal, p.peekToken.Line, p.peekToken.Column))
}

func (p *Parser) parseIfStatement() *IfStatement {
    stmt := &IfStatement{Token: p.currentToken}
    p.nextToken()

    stmt.Condition = p.parseExpression(LOWEST)

    if !p.expectPeek(lexer.THEN) {
        p.errors = append(p.errors, fmt.Sprintf("expected 'then' after 'if' condition at line %d, column %d, got %s (%q) instead", p.peekToken.Line, p.peekToken.Column, p.peekToken.Type, p.peekToken.Literal))
        return nil
    }

    stmt.Consequence = p.parseBlockStatement(lexer.END)

    for p.peekTokenIs(lexer.ELSEIF) {
        p.nextToken()
        elseif := &IfStatement{
            Token:       p.currentToken,
            Condition:   p.parseExpression(LOWEST),
            Consequence: p.parseBlockStatement(lexer.END),
        }
        stmt.Alternatives = append(stmt.Alternatives, elseif)
    }

    if p.peekTokenIs(lexer.ELSE) {
        p.nextToken()
        stmt.Alternative = p.parseBlockStatement(lexer.END)
    }

    if !p.expectPeek(lexer.END) {
        return nil
    }

    return stmt
}

func (p *Parser) parseWhileStatement() *WhileStatement {
    stmt := &WhileStatement{Token: p.currentToken}
    p.nextToken()

    stmt.Condition = p.parseExpression(LOWEST)

    if !p.expectPeek(lexer.DO) {
        return nil
    }

    stmt.Body = p.parseBlockStatement(lexer.END)

    if !p.expectPeek(lexer.END) {
        return nil
    }

    return stmt
}

func (p *Parser) parseForStatement() *ForStatement {
    stmt := &ForStatement{Token: p.currentToken}
    p.nextToken()

    if p.peekTokenIs(lexer.ASSIGN) {
        stmt.Identifier = p.parseIdentifier().(*Identifier)
        p.nextToken()
        stmt.InitValue = p.parseExpression(LOWEST)
        if !p.expectPeek(lexer.COMMA) {
            return nil
        }
        p.nextToken()
        stmt.EndValue = p.parseExpression(LOWEST)
        if p.peekTokenIs(lexer.COMMA) {
            p.nextToken()
            p.nextToken()
            stmt.StepValue = p.parseExpression(LOWEST)
        }
    } else {
        stmt.Identifiers = p.parseIdentifierList()
        if !p.expectPeek(lexer.IN) {
            return nil
        }
        p.nextToken()
        stmt.Iterators = p.parseExpressionList(lexer.DO)
    }
    if !p.expectPeek(lexer.DO) {
        return nil
    }
    stmt.Body = p.parseBlockStatement(lexer.END)

    if !p.expectPeek(lexer.END) {
        return nil
    }
    return stmt
}

func (p *Parser) parseTableConstructor() Expression {
    table := &TableLiteral{Token: p.currentToken}
    for !p.peekTokenIs(lexer.RBRACE) {
        p.nextToken()
        if p.curTokenIs(lexer.IDENTIFIER) && p.currentToken.Literal == "__metatable" {
            p.nextToken()
            if !p.expectPeek(lexer.ASSIGN) {
                return nil
            }
            p.nextToken()
            table.Metatable = p.parseExpression(LOWEST)
            continue
        }
        key, value := p.parseTableField()
        table.Fields = append(table.Fields, &TableField{Key: key, Value: value})
        if !p.peekTokenIs(lexer.COMMA) && !p.peekTokenIs(lexer.RBRACE) {
            p.peekError(lexer.COMMA)
            return nil
        }
    }
    p.nextToken()
    return table
}

func (p *Parser) parseTableField() (Expression, Expression) {
    var key Expression
    var value Expression

    if p.curTokenIs(lexer.LBRACKET) {
        p.nextToken()
        key = p.parseExpression(LOWEST)
        if !p.expectPeek(lexer.RBRACKET) || !p.expectPeek(lexer.ASSIGN) {
            return nil, nil
        }
        p.nextToken()
        value = p.parseExpression(LOWEST)
        return key, value
    }

    if p.curTokenIs(lexer.IDENTIFIER) {
        if p.peekTokenIs(lexer.ASSIGN) {
            key = &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
            p.nextToken()
            p.nextToken()
            value = p.parseExpression(LOWEST)
            return key, value
        } else if p.peekTokenIs(lexer.COLON) {
            key = &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
            p.nextToken()
            p.nextToken()
            value = p.parseFunctionLiteral()
            return key, value
        } else {
            key = &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
            value = &Identifier{
                Token: p.currentToken,
                Value: p.currentToken.Literal,
            }
            if p.peekTokenIs(lexer.COMMA) || p.peekTokenIs(lexer.RBRACE) {
                p.nextToken()
                return key, value
            }
        }
    }
    key = nil
    value = p.parseExpression(LOWEST)
    return key, value
}

func (p *Parser) parseArrayConstructor() Expression {
    array := &ArrayLiteral{Token: p.currentToken}
    array.Elements = p.parseExpressionList(lexer.RBRACKET)
    return array
}

func (p *Parser) parseFunctionLiteral() Expression {
    lit := &FunctionLiteral{Token: p.currentToken}
    p.nextToken()
    if p.peekTokenIs(lexer.COLON) {
        ident, ok := p.parseIdentifier().(*Identifier)
        if !ok {
            return nil
        }
        lit.Receiver = ident
        p.nextToken()
        p.nextToken()
    }

    if !p.expectPeek(lexer.LPAREN) {
        return nil
    }
    lit.Parameters = p.parseFunctionParameters()
    lit.Body = p.parseBlockStatement(lexer.END)
    return lit
}

func (p *Parser) parseBlockStatement(endToken lexer.TokenType) *BlockStatement {
    block := &BlockStatement{Token: p.currentToken}
    p.nextToken()

    for !p.curTokenIs(endToken) && !p.curTokenIs(lexer.EOF) {
        stmt := p.parseStatement()
        if stmt != nil {
            block.Statements = append(block.Statements, stmt)
        }
        p.nextToken()
    }

    return block
}

func (p *Parser) parseIdentifier() Expression {
    return &Identifier{
        Token: p.currentToken,
        Value: p.currentToken.Literal,
    }
}

func (p *Parser) parseIdentifierList() []*Identifier {
    var identifiers []*Identifier

    identifiers = append(identifiers, &Identifier{
        Token: p.currentToken,
        Value: p.currentToken.Literal,
    })

    for p.peekTokenIs(lexer.COMMA) {
        p.nextToken()
        p.nextToken()
        identifiers = append(identifiers, &Identifier{
            Token: p.currentToken,
            Value: p.currentToken.Literal,
        })
    }

    return identifiers
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []Expression {
    list := []Expression{}

    if p.peekTokenIs(end) {
        p.nextToken()
        return list
    }

    p.nextToken()
    list = append(list, p.parseExpression(LOWEST))

    for p.peekTokenIs(lexer.COMMA) {
        p.nextToken()
        p.nextToken()
        list = append(list, p.parseExpression(LOWEST))
    }

    if !p.expectPeek(end) {
        return nil
    }

    return list
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
    return p.currentToken.Type == t
}
