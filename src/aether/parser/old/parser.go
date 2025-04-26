package parser

import (
	"fmt"
	"strconv"

	"FLUX/src/aether/ast"
	"FLUX/src/aether/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
	errors    []string
	tokens    []lexer.Token
	tokenIndex int
	currentRecursionDepth int
}

func (p *Parser) nextToken() {
    if p.tokenIndex < len(p.tokens)-1 {
        p.tokenIndex++
    }
    p.curToken = p.tokens[p.tokenIndex]
    if p.tokenIndex+1 < len(p.tokens) {
        p.peekToken = p.tokens[p.tokenIndex+1]
    } else {
        p.peekToken = lexer.Token{Type: lexer.EOF}
    }
    fmt.Printf("Advanced token: curToken=%v, peekToken=%v\n", p.curToken, p.peekToken)
}

func (p *Parser) currentToken() lexer.Token {
    return p.tokens[p.tokenIndex]
}

func (p *Parser) raiseError(message string) error {
    return fmt.Errorf("Error: %s", message)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
    program := ast.NewProgram()
    for p.curToken.Type != lexer.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}
func (p *Parser) parseStatement() ast.Statement {
	fmt.Println("breakpoint reached")
	switch p.curToken.Type {
	case lexer.LOCAL:
		fmt.Println("local reached")
		return p.parseLocalStatement()
	case lexer.CONST:
		return p.parseConstStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.WHILE:
		return p.parseWhileStatement()
	case lexer.REPEAT:
		return p.parseRepeatStatement()
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.TRY:
		return p.parseTryCatchStatement()
	case lexer.FUNCTION:
		return p.parseFunctionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() ast.Statement {
    expr := p.parseExpression(lexer.LOWEST)
    if expr == nil {
        return nil
    }
    // Check if it's a call expression
    if p.peekTokenIs(lexer.LPAREN) {
        expr = p.parseCallExpression(expr)
    }
    return &ast.ExpressionStatement{Expr: expr}
}

func (p *Parser) parseRepeatStatement() *ast.RepeatStatement {
	stmt := &ast.RepeatStatement{Token: p.curToken}
	p.nextToken()

	stmt.Block = []ast.Statement{}
	for !p.curTokenIs(lexer.UNTIL) && !p.curTokenIs(lexer.EOF) {
		stmt.Block = append(stmt.Block, p.parseStatement())
		p.nextToken()
	}

	p.expectPeek(lexer.UNTIL)
	p.nextToken()
	stmt.Condition = p.parseExpression(lexer.LOWEST)

	return stmt
}
func (p *Parser) parseForInitialization() ast.Statement {
    if p.peekTokenIs(lexer.IDENTIFIER) {
        p.nextToken()
        nameExpr := p.parseIdentifier()
        name, ok := nameExpr.(*ast.Identifier)
        if !ok {
            p.errors = append(p.errors, "expected identifier")
            return nil
        }
        p.nextToken()
        value := p.parseExpression(lexer.LOWEST)
        return ast.NewLocalStatement(p.curToken, name, value)
    }
    return nil
}
func (p *Parser) parseForStatement() *ast.ForStatement {
	p.consumeExpectedToken(lexer.FOR)
	p.consumeExpectedToken(lexer.LPAREN)
	initStmt := p.parseForInitialization()
	conditionExpr := p.parseExpression(lexer.LOWEST)
	incrementExpr := p.parseExpression(lexer.LOWEST)
	p.consumeExpectedToken(lexer.DO)
	block := p.parseBlockStatement()
	p.consumeExpectedToken(lexer.END)
	return ast.NewForStatement(initStmt, conditionExpr, incrementExpr, block)
}

func (p *Parser) consumeToken() lexer.Token {
	p.advance()
	return p.currentToken()
}

func (p *Parser) consumeExpectedToken(expected lexer.TokenType) {
    if p.curToken.Type != expected {
        p.errors = append(p.errors, fmt.Sprintf("expected token %s, got %s", expected, p.curToken.Type))
        return
    }
    p.nextToken()
}

func (p *Parser) advance() {
	if p.tokenIndex < len(p.tokens)-1 {
		p.tokenIndex++
	}
}

func (p *Parser) parseFunctionStatement() ast.Statement {
	p.consumeExpectedToken(lexer.FUNCTION)

	if p.peekToken.Type != lexer.IDENTIFIER {
		p.peekError(lexer.IDENTIFIER)
		return nil
	}
	p.nextToken()
	name := ast.NewIdentifier(p.curToken)

	if p.peekToken.Type != lexer.LPAREN {
		p.peekError(lexer.LPAREN)
		return nil
	}
	p.nextToken()

	parameters := []*ast.Identifier{}
	for p.peekToken.Type == lexer.IDENTIFIER {
		p.nextToken()
		parameters = append(parameters, ast.NewIdentifier(p.curToken))
		if p.peekToken.Type == lexer.COMMA {
			p.nextToken()
		}
	}

	if p.peekToken.Type != lexer.RPAREN {
		p.peekError(lexer.RPAREN)
		return nil
	}
	p.nextToken()

	// Parse body
	body := []ast.Statement{}
	for p.curToken.Type != lexer.END && p.curToken.Type != lexer.EOF {
		p.nextToken() // Consume the token
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
		}
	}

	if p.curToken.Type != lexer.END {
		p.peekError(lexer.END)
		return nil
	}
	p.nextToken() // Consume the END token

	return ast.NewFunctionStatement(name, parameters, body)
}
func (p *Parser) parseIfStatement() ast.Statement {
	fmt.Printf("parseIfStatement: curToken=%v\n", p.curToken)
    p.consumeExpectedToken(lexer.IF)
    condition := p.parseExpression(lexer.LOWEST)
	if condition == nil {
		return nil
	}
    p.consumeExpectedToken(lexer.THEN)
    thenBlock := p.wrapInExpressionStatements(p.parseBlockStatement())
	fmt.Printf("parseIfStatement: thenBlock parsed, curToken=%v\n", p.curToken)

    elseIfs := []*ast.ElseIfClause{}
    for p.curTokenIs(lexer.ELSEIF) {
        p.nextToken()
        elseifCond := p.parseExpression(lexer.LOWEST)
		if elseifCond == nil {
			return nil
		}
		p.consumeExpectedToken(lexer.THEN)
        elseifBlock := p.wrapInExpressionStatements(p.parseBlockStatement())
		fmt.Printf("parseIfStatement: elseifBlock parsed, curToken=%v\n", p.curToken)
        elseIfs = append(elseIfs, &ast.ElseIfClause{
            Condition: elseifCond,
            Block:     elseifBlock,
        })
    }

    var elseBlock []ast.Statement
    if p.curTokenIs(lexer.ELSE) {
        p.nextToken()
        elseBlock = p.wrapInExpressionStatements(p.parseBlockStatement())
    }

    p.consumeExpectedToken(lexer.END)
    return ast.NewIfStatement(condition, thenBlock, elseIfs, elseBlock)
}

func (p *Parser) wrapInExpressionStatements(block []ast.Statement) []ast.Statement {
	wrappedBlock := []ast.Statement{}
	for _, stmt := range block {
		_, alreadyExpressionStatement := stmt.(*ast.ExpressionStatement)
		if expr, ok := stmt.(ast.Expression); ok && !alreadyExpressionStatement {
			wrappedBlock = append(wrappedBlock, &ast.ExpressionStatement{Expr: expr, Type: "ExpressionStatement"})
		} else {
			wrappedBlock = append(wrappedBlock, stmt)
		}
	}
	return wrappedBlock
}
func (p *Parser) parseLocalStatement() ast.Statement {
    localToken := p.curToken
    p.nextToken()
    if p.peekTokenIs(lexer.FUNCTION) {
        p.nextToken()
        return p.parseLocalFunctionStatement(localToken)
    }
    if !p.curTokenIs(lexer.IDENTIFIER) {
        p.errors = append(p.errors, "expected identifier")
        return nil
    }
    name := ast.NewIdentifier(p.curToken)
    if !p.expectPeek(lexer.ASSIGN) {
        return nil
    }
    p.nextToken()
    value := p.parseExpression(lexer.LOWEST)
    return ast.NewLocalStatement(localToken, name, value)
}
func (p *Parser) parseLocalFunctionStatement(localToken lexer.Token) ast.Statement {
    p.consumeExpectedToken(lexer.FUNCTION)
    name := ast.NewIdentifier(p.curToken)
    p.nextToken()
    p.consumeExpectedToken(lexer.LPAREN)
    params := []*ast.Identifier{}
    for !p.peekTokenIs(lexer.RPAREN) {
        p.nextToken()
        if p.curTokenIs(lexer.IDENTIFIER) {
            params = append(params, ast.NewIdentifier(p.curToken))
        }
        if p.peekTokenIs(lexer.COMMA) {
            p.nextToken()
        }
    }
    p.consumeExpectedToken(lexer.RPAREN)
    p.nextToken()
    body := p.parseBlockStatement()
    return ast.NewLocalFunctionStatement(name, params, body)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{}
	p.consumeExpectedToken(lexer.RETURN)
	returnValue := p.parseExpression(lexer.LOWEST)
	stmt = ast.NewReturnStatement(returnValue)
	return stmt
}
func (p *Parser) parseExpression(precedence int) ast.Expression {
    leftExpr := p.parsePrimary()
    if leftExpr == nil {
        return nil
    }
    for precedence < p.peekTokenPrecedence() {
        operator := p.peekToken.Literal
        p.nextToken()
        rightExpr := p.parseExpression(precedence)
        if rightExpr == nil {
            return nil
        }
        leftExpr = ast.NewInfixExpression(leftExpr, operator, rightExpr)
    }
    return leftExpr
}
func (p *Parser) parseTryCatchStatement() ast.Statement {
    stmt := &ast.TryCatchStatement{Token: p.curToken}
    p.nextToken() // Consume TRY
    stmt.TryBlock = p.parseBlockStatement()
    if !p.expectPeek(lexer.CATCH) {
        return nil
    }
    p.nextToken() // Consume CATCH
    stmt.CatchBlock = p.parseBlockStatement()
    return stmt
}
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	callExp := &ast.CallExpression{
		Token:    p.curToken,
		Function: function,
		Type:     "CallExpression",
	}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	callExp.Arguments = p.parseExpressionList(lexer.RPAREN)

	return callExp
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return args
	}

	p.nextToken()
	arg := p.parseExpression(lexer.LOWEST)
	if arg != nil {
		args = append(args, arg)
	} else {
		return nil // Error while parsing expression
	}

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken() // Consume comma
		p.nextToken() // Advance to next argument

		arg := p.parseExpression(lexer.LOWEST)
		if arg != nil {
			args = append(args, arg)
		} else {
			return nil // Error while parsing expression
		}
	}

	if !p.expectPeek(end) {
		return nil
	}

	return args
}
func (p *Parser) parseInfixExpression(precedence int) ast.Expression {
	leftExp := p.parsePrimary()
	if leftExp == nil {
		return nil
	}

	for p.peekToken.Type != lexer.EOF && precedence < p.peekTokenPrecedence() {
		operator := p.peekToken.Literal
		p.nextToken() // Move to the operator

		rightExp := p.parseExpression(precedence)
		if rightExp == nil {
			return nil
		}

		leftExp = &ast.InfixExpression{
			LeftNode:     leftExp,
			Operator: operator,
			RightNode:    rightExp,
		}
	}

	return leftExp
}

func (p *Parser) parsePrimary() ast.Expression {
	fmt.Errorf("breakpoint reached")
	fmt.Errorf("breakpoint reached")
	fmt.Errorf("breakpoint reached")
	fmt.Errorf("breakpoint reached")
	switch p.curToken.Type {
	case lexer.IDENTIFIER:
		return p.parseIdentifier()
	case lexer.NUMBER:
		return p.parseIntegerLiteral()
	case lexer.STRING:
		return p.parseStringLiteral()
	case lexer.TRUE, lexer.FALSE:
		return p.parseBooleanLiteral()
	case lexer.NIL:
		return p.parseNil()
	case lexer.MINUS, lexer.NOT:
		return p.parsePrefixExpression()
	case lexer.LPAREN:
		p.nextToken() // Consume the LPAREN
		exp := p.parseExpression(lexer.LOWEST)
		if p.peekToken.Type != lexer.RPAREN {
			p.peekError(lexer.RPAREN)
		}

		p.nextToken() // Consume the RPAREN
		return exp
    default:
		msg := fmt.Sprintf("unexpected token %s with literal '%s'", p.curToken.Type, p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return ast.NewIdentifier(p.curToken)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
		return nil
	}
	return ast.NewIntegerLiteral(value)
}

func (p *Parser) parseConstStatement() ast.Statement {
	stmt := &ast.ConstStatement{
		Token: p.curToken,
		Type: "ConstStatement",
	}
	p.consumeExpectedToken(lexer.CONST)
	p.consumeExpectedToken(lexer.IDENTIFIER)
	stmt.Name = p.curToken.Literal
	stmt.Value = p.parseExpression(lexer.LOWEST)
	return stmt
}

func (p *Parser) parseWhileStatement() ast.Statement {
    stmt := &ast.WhileStatement{Token: p.curToken}
    p.consumeExpectedToken(lexer.WHILE)
    stmt.Condition = p.parseExpression(lexer.LOWEST)
    stmt.Body = p.parseBlockStatement()
    return stmt
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return ast.NewStringLiteral(p.curToken.Literal)
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	value := p.curToken.Type == lexer.TRUE
	return ast.NewBooleanLiteral(value)
}

func (p *Parser) parseNil() ast.Expression {
	return ast.NewNil()
}

func (p *Parser) curTokenPrecedence() int {
	return p.tokenPrecedence(p.curToken)
}

func (p *Parser) peekTokenPrecedence() int {
	if p, ok := lexer.Precedences[p.peekToken.Type]; ok {
		return p
	}
	return 0
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type))
	return false
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expr.Right = p.parseExpression(lexer.PREFIX)

	return expr
}
func (p *Parser) tokenPrecedence(token lexer.Token) int {
    switch token.Type {
    case lexer.OR:
        return 1
    case lexer.AND:
        return 2
    case lexer.EQ, lexer.NE:
        return 3
    case lexer.LT, lexer.GT, lexer.LE, lexer.GE:
        return 4
    case lexer.PLUS, lexer.MINUS:
        return 5
    case lexer.MULTIPLY, lexer.DIVIDE, lexer.MODULO:
        return 6
    case lexer.EXPONENT:
        return 7
	case lexer.CONCAT:
        return 8
    default:
        return 0
    }
}
func (p *Parser) parseBlockStatement() []ast.Statement {
    block := []ast.Statement{}
    for !p.curTokenIs(lexer.END) && 
    !p.curTokenIs(lexer.ELSE) && 
    !p.curTokenIs(lexer.ELSEIF) && 
    !p.curTokenIs(lexer.UNTIL) && 
    !p.curTokenIs(lexer.CATCH) && 
    !p.curTokenIs(lexer.EOF) {
        if stmt := p.parseStatement(); stmt != nil {
            block = append(block, stmt)
        }
        p.nextToken()
    }

	// Check if the block contains only one statement and if that statement is an ExpressionStatement
	if len(block) == 1 {
		if exprStmt, ok := block[0].(*ast.ExpressionStatement); ok {
			// If so, return the expression directly
			return []ast.Statement{&ast.ExpressionStatement{Expr: exprStmt.Expr}}
		}
	}

    return block
}

func NewParser(l *lexer.Lexer) *Parser {
	tokens := l.Tokenize()
	p := &Parser{
		l:      l,
		errors: []string{},
		tokens: tokens,
		tokenIndex: 0,
	}
	if len(tokens) > 0 {
		p.curToken = tokens[0]
		if len(tokens) > 1 {
			p.peekToken = tokens[1]
		}
	}
	return p
}
