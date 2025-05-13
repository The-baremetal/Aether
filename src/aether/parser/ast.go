package parser

import (
    "FLUX/src/aether/lexer"
)

// Node represents a node in the AST
type Node interface {
    node()
}

// Statement represents a statement node
type Statement interface {
    Node
    statementNode()
}

// Expression represents an expression node
type Expression interface {
    Node
    expressionNode()
}

// Program root node
type Program struct {
    Statements []Statement
}

func (p *Program) node() {}

// ========== Statements ==========
type LocalDeclaration struct {
    Token    lexer.Token
    Constant bool
    Names    []*Identifier
    Values   []Expression
}

func (ld *LocalDeclaration) statementNode() {}
func (ld *LocalDeclaration) node() {}

type ImportStatement struct {
    Token  lexer.Token
    Path   *StringLiteral
    AsName *Identifier
}

func (is *ImportStatement) statementNode() {}
func (is *ImportStatement) node() {}

type FunctionDeclaration struct {
    Token      lexer.Token
    Receiver   *Identifier
    Name       *Identifier
    Parameters []*Identifier
    Body       *BlockStatement
}

func (fd *FunctionDeclaration) statementNode() {}
func (fd *FunctionDeclaration) node() {}

type ConstDeclaration struct {
    Token  lexer.Token
    Names   []*Identifier
    Values  []Expression
}

func (cd *ConstDeclaration) statementNode() {}
func (cd *ConstDeclaration) node() {}

type ReturnStatement struct {
    Token       lexer.Token
    Value Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) node() {}

type IfStatement struct {
    Token        lexer.Token
    Condition    Expression
    Consequence  *BlockStatement
    Alternatives []*IfStatement
    Alternative  *BlockStatement
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) node() {}

type RepeatStatement struct {
    Token     lexer.Token
    Body      *BlockStatement
    Condition Expression
}

func (rps *RepeatStatement) statementNode() {}
func (rps *RepeatStatement) node() {}

type WhileStatement struct {
    Token     lexer.Token
    Condition Expression
    Body      *BlockStatement
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) node() {}

type ForStatement struct {
    Token       lexer.Token
    Identifier  *Identifier
    Identifiers []*Identifier
    InitValue   Expression
    EndValue    Expression
    StepValue   Expression
    Iterators   []Expression
    Body        *BlockStatement
}

func (fs *ForStatement) statementNode() {}
func (fs *ForStatement) node() {}

type BreakStatement struct {
    Token lexer.Token
}

func (bs *BreakStatement) statementNode() {}
func (bs *BreakStatement) node() {}

type TryCatchStatement struct {
    Token        lexer.Token
    TryBlock     *BlockStatement
    CatchVar     *Identifier
    CatchBlock   *BlockStatement
    FinallyBlock *BlockStatement
}

func (tcs *TryCatchStatement) statementNode() {}
func (tcs *TryCatchStatement) node() {}

type CoroutineStatement struct {
    Token  lexer.Token
    Method string
    Args   []Expression
}

func (cs *CoroutineStatement) statementNode() {}
func (cs *CoroutineStatement) node() {}

// ========== Expressions ==========
type Identifier struct {
    Token lexer.Token
    Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) node() {}

type NumberLiteral struct {
    Token lexer.Token
    Value float64
}

func (nl *NumberLiteral) expressionNode() {}
func (nl *NumberLiteral) node() {}

type StringLiteral struct {
    Token lexer.Token
    Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) node() {}

type Boolean struct {
    Token lexer.Token
    Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) node() {}

type NilLiteral struct {
    Token lexer.Token
}

func (nl *NilLiteral) expressionNode() {}
func (nl *NilLiteral) node() {}

type PrefixExpression struct {
    Token    lexer.Token
    Operator string
    Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) node() {}

type InfixExpression struct {
    Token    lexer.Token
    Left     Expression
    Operator string
    Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) node() {}

type TableLiteral struct {
    Token     lexer.Token
    Fields    []*TableField
    Metatable Expression
}

func (tl *TableLiteral) expressionNode() {}
func (tl *TableLiteral) node() {}

type TableField struct {
    Key   Expression
    Value Expression
}

type ArrayLiteral struct {
    Token    lexer.Token
    Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) node() {}

type FunctionLiteral struct {
    Token      lexer.Token
    Parameters []*Identifier
    Body       *BlockStatement
    Receiver   *Identifier // For method definitions
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) node() {}

type CallExpression struct {
    Token     lexer.Token
    Function  Expression
    Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) node() {}

type IndexExpression struct {
    Token   lexer.Token
    Left    Expression
    Index   Expression
    Optional bool // For optional chaining
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) node() {}

type ExpressionStatement struct {
    Expression Expression
}

func (es *ExpressionStatement) expressionNode() {}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) node() {}

type BooleanLiteral struct {
    Token lexer.Token
    Value bool
}

func (bl *BooleanLiteral) expressionNode() {}
func (bl *BooleanLiteral) node() {}

type DoStatement struct {
    Token lexer.Token
    Body  *BlockStatement
}

func (ds *DoStatement) statementNode() {}
func (ds *DoStatement) node() {}

type DotExpression struct {
    Token    lexer.Token
    Object   Expression
    Property *Identifier
    Left  Expression
    Right Expression
    Optional bool // For optional chaining
}

func (de *DotExpression) expressionNode() {}
func (de *DotExpression) node() {}

type AssignExpression struct {
    Token lexer.Token
    Left  Expression
    Value Expression
}

func (ae *AssignExpression) expressionNode() {}
func (ae *AssignExpression) node() {}

type AsyncExpression struct {
    Token lexer.Token
    Block *BlockStatement
}

func (ae *AsyncExpression) expressionNode() {}
func (ae *AsyncExpression) node() {}

type AwaitExpression struct {
    Token      lexer.Token
    Expression Expression
}

func (ae *AwaitExpression) expressionNode() {}
func (ae *AwaitExpression) node() {}

type NamedArgument struct {
    Token lexer.Token
    Name  *Identifier
    Value Expression
}

func (na *NamedArgument) expressionNode() {}
func (na *NamedArgument) node() {}

// ========== Helper Structures ==========
type BlockStatement struct {
    Token      lexer.Token
    Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) node() {}

type FunctionParameter struct {
    Token lexer.Token
    Name  *Identifier
}

type Metamethod struct {
    Token lexer.Token
    Name  string
    Value Expression
}

type MethodCallExpression struct {
    Token     lexer.Token
    Object    Expression
    Method    *Identifier
    Arguments []Expression
}

func (mce *MethodCallExpression) expressionNode() {}
func (mce *MethodCallExpression) node() {}

type VarArgLiteral struct {
    Token lexer.Token
}

func (v *VarArgLiteral) expressionNode() {}
func (v *VarArgLiteral) TokenLiteral() string { return v.Token.Literal }
func (v *VarArgLiteral) String() string       { return v.Token.Literal }
func (v *VarArgLiteral) node() {}