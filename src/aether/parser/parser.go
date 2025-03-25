package parser

import (
	"fmt"
	"strconv"
	"strings"

	"FLUX/src/lib/lexer"
	"FLUX/src/aether/ast"
)

/*
replace FLUX/src/aether/ast => ./src/aether/ast

replace FLUX/src/aether/visitor => ./src/aether/visitor

replace FLUX/src/aether/codegen => ./src/aether/codegen

replace FLUX/src/aether/parser => ./src/aether/parser

replace FLUX/src/lib/lexer => ./src/lib/lexer

replace FLUX/src/aether/utils => ./src/aether/utils
*/


const (
	_ int = iota
	LOWEST
	EQUALS  // ==
	LESSGREATER // > or <
	SUM   // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL  // myFunction(X)
)

//types

type Node interface {
	String(indent int) string
}

type Statement interface {
	Node
	statementNode() // marker method
}

type Expression interface {
	Node
	expressionNode() // marker method
}

type ProgramNode struct {
	Statements []Statement
}

type Parser struct {
	lexer        *lexer.Lexer
	current_token  lexer.Token
	peek_token     lexer.Token
	errors       []string
	error_handler *utils.ErrorHandler
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:        l,
		errors:       []string{},
		error_handler: utils.NewErrorHandler(),
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) parseAssignStatement() Statement {
	// parse variable on left-hand side
	varExpr := p.parseVariable()
	op := p.current_token.Literal
	p.nextToken()
	value := p.parseExpression(LOWEST)
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &AssignStmt{Variable: varExpr, Op: op, Value: value}
}

func (p *Parser) parseIfStatement() Statement {
	// current token is IF; consume it and parse condition.
	p.nextToken()
	cond := p.parseExpression(LOWEST)
	if !p.expectPeek(lexer.THEN) {
		return nil
	}
	thenBlock := p.parseBlock()
	var elseIfs []*ast.IfStatementNode
	var elseBlock *ast.BlockNode
	// Parse zero or more elseif clauses.
	for p.current_token.Type == lexer.ELSEIF {
		ii := &ast.IfStatementNode{}
		p.nextToken()
		ii.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.THEN) {
			return nil
		}
		ii.Then = p.parseBlock()
		elseIfs = append(elseIfs, ii)
	}
	// Optionally, parse an else clause.
	if p.current_token.Type == lexer.ELSE {
		p.nextToken()
		elseBlock = p.parseBlock()
	}
	if !p.expectPeek(lexer.END) {
		return nil
	}
	return &ast.IfStatementNode{Condition: cond, Then: thenBlock, Else: elseBlock}
}

func (p *Parser) parseWhileStatement() Statement {
	p.nextToken() // consume 'while'
	cond := p.parseExpression(LOWEST)
	if !p.expectPeek(lexer.DO) {
		return nil
	}
	body := p.parseBlock()
	if !p.expectPeek(lexer.END) {
		return nil
	}
	return &WhileStmt{Cond: cond, Body: body}
}

func (p *Parser) parseRepeatStatement() Statement {
	p.nextToken() // consume 'repeat'
	body := p.parseBlock()
	if !p.expectPeek(lexer.UNTIL) {
		return nil
	}
	cond := p.parseExpression(LOWEST)
	return &RepeatStmt{Body: body, Cond: cond}
}

func (p *Parser) parseForStatement() Statement {
	fs := &ForStmt{}
	p.nextToken() // consume 'for'
	// Numeric for: IDENT '=' expression ',' expression (',' expression)? 'do' block 'end'
	if p.current_token.Type == lexer.IDENT && p.peekTokenIs(lexer.ASSIGN) {
		fs.Identifier = p.current_token.Literal
		p.nextToken() // consume identifier
		p.nextToken() // consume '='
		fs.Init = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.COMMA) {
			return nil
		}
		p.nextToken()
		fs.End = p.parseExpression(LOWEST)
		// Optional step expression.
		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken() // consume comma
			p.nextToken()
			fs.Step = p.parseExpression(LOWEST)
		}
		fs.Body = p.parseBlock()
		if !p.expectPeek(lexer.END) {
			return nil
		}
		fs.Generic = false
	} else {
		// Generic for: for identifierList 'in' expressionList 'do' block 'end'
		fs.Generic = true
		// For simplicity, parse one identifier.
		fs.Identifier = p.current_token.Literal
		if !p.expectPeek(lexer.IN) {
			return nil
		}
		p.nextToken()
		// Parse one expression as iterator list.
		fs.IterExprs = []Expression{p.parseExpression(LOWEST)}
		fs.Body = p.parseBlock()
		if !p.expectPeek(lexer.END) {
			return nil
		}
	}
	return fs
}

func (p *Parser) parseTryCatchStatement() Statement {
	tcs := &TryCatchStmt{}
	p.nextToken() // consume 'try'
	tcs.TryBlock = p.parseBlock()
	if p.current_token.Type == lexer.CATCH {
		p.nextToken() // consume 'catch'
		tcs.CatchIdent = p.current_token.Literal
		p.nextToken()
		tcs.CatchBlock = p.parseBlock()
	}
	if p.current_token.Type == lexer.FINALLY {
		p.nextToken() // consume 'finally'
		tcs.Finally = p.parseBlock()
	}
	if !p.expectPeek(lexer.END) {
		return nil
	}
	return tcs
}

func (p *Parser) parseFunctionDeclaration() Statement {
	fd := &FunctionDecl{}
	p.nextToken() // consume 'function'
	// Optionally, parse receiver if token is '.' or ':'
	if p.peekTokenIs(lexer.DOT) || p.peekTokenIs(lexer.COLON) {
		p.nextToken() // consume dot or colon
		fd.Receiver = p.current_token.Literal
	}
	if p.current_token.Type != lexer.IDENT {
		p.errors = append(p.errors, "expected function name")
		return nil
	}
	fd.Name = p.current_token.Literal
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	fd.Params = p.parseParameterList()
	// Optionally parse type annotation for return types.
	if p.peekTokenIs(lexer.TYPEOF) {
		p.nextToken() // consume typeof indicator
		fd.RetTypes = []string{p.current_token.Literal} // simplified; real parser would parse complex type annotations.
	}
	fd.Body = p.parseBlock()
	if !p.expectPeek(lexer.END) {
		return nil
	}
	return fd
}

func (p *Parser) parseLocalDeclaration() Statement {
	ld := &LocalDecl{}
	// Check for optional "const"
	if p.peekTokenIs(lexer.CONST) {
		ld.IsConst = true
		p.nextToken()
	}
	// Parse identifier list (for simplicity, one identifier)
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	ld.Names = []string{p.current_token.Literal}
	// Optionally parse type annotations; skipped here.
	if !p.expectPeek(lexer.ASSIGN) {
		return ld
	}
	p.nextToken()
	ld.Values = []Expression{p.parseExpression(LOWEST)}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return ld
}

func (p *Parser) parseConstDeclaration() Statement {
	cd := &ConstDecl{}
	// For simplicity, assume one identifier.
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	cd.Names = []string{p.current_token.Literal}
	if !p.expectPeek(lexer.ASSIGN) {
		return cd
	}
	p.nextToken()
	cd.Values = []Expression{p.parseExpression(LOWEST)}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return cd
}

func (p *Parser) parseLabelStatement() Statement {
	ls := &LabelStmt{Label: p.current_token.Literal}
	p.nextToken()
	return ls
}

func (p *Parser) parseSelectStatement() Statement {
	ss := &SelectStmt{}
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	p.nextToken()
	ss.Args = append(ss.Args, p.parseExpression(LOWEST))
	if !p.expectPeek(lexer.COMMA) {
		return nil
	}
	p.nextToken()
	ss.Args = append(ss.Args, p.parseExpression(LOWEST))
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	return ss
}

func (p *Parser) parseSpawnStatement() Statement {
	p.nextToken() // consume 'spawn'
	expr := p.parseExpression(LOWEST)
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &SpawnStmt{Expr: expr}
}

func (p *Parser) parseReturnStatement() Statement {
	rs := &ReturnStmt{}
	p.nextToken() // consume 'return'
	// For simplicity, parse one expression.
	rs.Exprs = append(rs.Exprs, p.parseExpression(LOWEST))
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return rs
}

func (p *Parser) parseTypeDeclaration() Statement {
	p.nextToken() // consume 'type'
	name := p.current_token.Literal
	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}
	p.nextToken() // skip '='
	ts := p.parseTypeSpec()
	return &TypeDecl{Name: name, Spec: ts}
}

func (p *Parser) parseTypeSpec() *TypeSpec {
	// Simplified: if current token is IDENT, return that as a type spec.
	if p.current_token.Type == lexer.IDENT {
		return &TypeSpec{Name: p.current_token.Literal}
	}
	return &TypeSpec{Name: "unknown"}
}

func (p *Parser) parseExprStatement() Statement {
	expr := p.parseExpression(LOWEST)
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.ExpressionStatementNode{Expr: expr}
}

type Ident struct {
	Value string
}

func (i *Ident) expressionNode() {}
func (i *Ident) String() string   { return i.Value }

type IntLiteral struct {
	Value int64
}

func (il *IntLiteral) expressionNode() {}
func (il *IntLiteral) String() string   { return fmt.Sprintf("%d", il.Value) }

type StrLiteral struct {
	Value string
}

func (sl *StrLiteral) expressionNode() {}
func (sl *StrLiteral) String() string   { return sl.Value }

type BoolLiteral struct {
	Value bool
}

func (bl *BoolLiteral) expressionNode() {}
func (bl *BoolLiteral) String() string   { return fmt.Sprintf("%t", bl.Value) }

type NilLiteral struct{}

func (nl *NilLiteral) expressionNode() {}
func (nl *NilLiteral) String() string   { return "nil" }

type UnaryExpr struct {
	Op    string
	Right Expression
}

func (ue *UnaryExpr) expressionNode() {}
func (ue *UnaryExpr) String() string   { return fmt.Sprintf("(%s%s)", ue.Op, ue.Right.String()) }

type InfixExpr struct {
	Left  Expression
	Op    string
	Right Expression
}

func (ie *InfixExpr) expressionNode() {}
func (ie *InfixExpr) String() string   { return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Op, ie.Right.String()) }

type MatchExpr struct {
	Expr Expression
	Arms []MatchArm
}

func (me *MatchExpr) expressionNode() {}
func (me *MatchExpr) String() string   { return "match expression" }

// --- Expression parsing ---

func (p *Parser) parseExpression(precedence int) Expression {
	var left Expression
	switch p.current_token.Type {
	case lexer.IDENT:
		left = p.parseIdentifier()
	case lexer.NUMBER:
		left = p.parseIntegerLiteral()
	case lexer.STRING:
		left = p.parseStringLiteral()
	case lexer.BOOL:
		left = p.parseBoolLiteral()
	case lexer.NIL:
		left = &NilLiteral{}
	case lexer.MINUS, lexer.NOT:
		op := p.current_token.Literal
		p.nextToken()
		right := p.parseExpression(PREFIX)
		left = &UnaryExpr{Op: op, Right: right}
	case lexer.MATCH:
		left = p.parseMatchExpression()
	case lexer.AWAIT:
		p.nextToken()
		left = p.parseExpression(LOWEST)
	default:
		return nil
	}
	for !p.peekTokenIs(lexer.SEMICOLON) && precedence < p.peekPrecedence() {
		p.nextToken()
		left = p.parseInfixExpression(left)
	}
	return left
}

func (p *Parser) parseIdentifier() Expression {
	return &Ident{Value: p.current_token.Literal}
}

func (p *Parser) parseIntegerLiteral() Expression {
	val, err := strconv.ParseInt(p.current_token.Literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, "could not parse integer")
		return nil
	}
	return &IntLiteral{Value: val}
}

func (p *Parser) parseStringLiteral() Expression {
	return &StrLiteral{Value: p.current_token.Literal}
}

func (p *Parser) parseBoolLiteral() Expression {
	return &BoolLiteral{Value: p.current_token.Literal == "true"}
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expr := &InfixExpr{Left: left, Op: p.current_token.Literal}
	prec := p.curPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(prec)
	return expr
}

func (p *Parser) parseMatchExpression() Expression {
	me := &MatchExpr{}
	p.nextToken() // consume 'match'
	me.Expr = p.parseExpression(LOWEST)
	if !p.expectPeek(lexer.WITH) {
		return nil
	}
	for !p.peekTokenIs(lexer.END) {
		arm := p.parseMatchArm()
		if arm != nil {
			me.Arms = append(me.Arms, arm)
		}
	}
	if !p.expectPeek(lexer.END) {
		return nil
	}
	return me
}

func (p *Parser) parseMatchArm() *MatchArm {
	arm := &MatchArm{}
	arm.Pattern = p.parseExpression(LOWEST)
	if p.peekTokenIs("when") { // for simplicity, we compare literal
		p.nextToken() // consume 'when'
		arm.When = p.parseExpression(LOWEST)
	}
	if !p.expectPeek(lexer.ARROW) {
		return nil
	}
	p.nextToken()
	arm.Result = p.parseExpression(LOWEST)
	return arm
}

func (p *Parser) parseVariable() Expression {
	// For now, assume a simple identifier.
	return &Ident{Value: p.current_token.Literal}
}

func (p *Parser) parseBlock() *ast.BlockNode {
	block := &ast.BlockNode{}
	p.nextToken()
	for p.current_token.Type != lexer.END &&
		p.current_token.Type != lexer.ELSE &&
		p.current_token.Type != lexer.ELSEIF &&
		p.current_token.Type != lexer.UNTIL &&
		p.current_token.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseParameterList() []string {
	var params []string
	if p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		return params
	}
	p.nextToken() // first parameter
	params = append(params, p.current_token.Literal)
	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken() // consume comma
		p.nextToken()
		params = append(params, p.current_token.Literal)
	}
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	return params
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected token %s, got %s", t, p.peekToken.Type))
	return false
}

func (p *Parser) curPrecedence() int {
	if prec, ok := precedences[p.current_token.Type]; ok {
		return prec
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peekToken.Type]; ok {
		return prec
	}
	return LOWEST
}
