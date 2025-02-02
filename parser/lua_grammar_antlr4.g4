grammar lua_grammar_antlr4;

// **Main entry point**
program
    : statement* EOF
    ;

// ---------------------------
/* Statements (Assignment, Control Flow, etc.) */
statement
    : assignStatement
    | controlFlowStatement
    | functionDeclaration
    | returnStatement
    | localDeclaration
    | labelStatement
    ;

// **Adding a new statement**: 
// 1. Add a new rule in the 'statement' section. 
// 2. It should follow the pattern: newStatementName
// Example:
// newStatement
//     : 'newKeyword' expression ';'
//     ;
assignStatement
    : variable ('=' | '+=' | '-=' | '*=' | '/=') expression
    ;

// ---------------------------
/* Expressions (Literals, Variables, Operations) */
expression
    : primaryExpression
    | expression operatorGroup expression
    |<assoc=right> expression operatorGroup expression  // For right-associative
    | unaryOp expression
    ;

primaryExpression
    : literal
    | variable
    | functionCall
    | unaryOperation
    | tableConstructor
    | '(' expression ')'
    ;

literal
    : BOOL
    | NIL
    | NUMBER
    | STRING
    ;

variable
    : IDENTIFIER
    ;

functionCall
    : IDENTIFIER '(' (expression (',' expression)*)? ')'
    ;

tableConstructor
    : '{' (field (',' field)*)? (',' metatable)? '}'
    ;

metatable
    : '__metatable' '=' expression
    ;

metamethods
    : metamethod (',' metamethod)*
    ;

labelStatement
    : '::' IDENTIFIER '::'
    ;

metamethod: '__add' | '__sub' | '__mul' | '__div' | '__mod' | '__pow' 
          | '__unm' | '__concat' | '__len' | '__eq' | '__lt' | '__le'
          | '__tostring' | '__pairs' | '__ipairs' | '__call';

field
    : IDENTIFIER '=' expression
    | expression
    ;

// ---------------------------
/* Binary and Unary Operations */
binaryOperation
    : expression arithOp expression
    | expression bitwiseOp expression
    | expression comparisonOp expression
    | expression logicalOp expression
    | expression concatOp expression
    | expression compoundAssignOp expression
    | expression coalesceOp expression
    | expression shiftAssignOp expression
    | expression incrOp expression
    ;

unaryOperation
    : unaryOp expression
    | '#' expression
    ;
// ---------------------------
/* Control Flow Statements (If, While, For, etc.) */
controlFlowStatement
    : ifStatement
    | whileStatement
    | repeatStatement
    | forStatement
    | breakStatement
    | gotoStatement
    | coroutineStatement
    ;

ifStatement
    : 'if' expression 'then' block ('elseif' expression 'then' block)* ('else' block)? 'end'
    ;

whileStatement
    : 'while' expression 'do' block 'end'
    ;

repeatStatement
    : 'repeat' block 'until' expression
    ;

forStatement
    : 'for' IDENTIFIER '=' expression ',' expression (',' expression)? 'do' block 'end'
    ;

breakStatement
    : 'break'
    ;

gotoStatement
    : 'goto' IDENTIFIER
    ;

coroutineStatement
    : 'coroutine' '.' ( 'create' | 'resume' | 'yield' | 'status' )
    ;

block
    : statement*
    ;

// ---------------------------
/* Declarations (Local Variables, Functions) */
localDeclaration
    : 'local' IDENTIFIER ('=' expression)?
    ;

functionDeclaration
    : 'function' IDENTIFIER '(' (IDENTIFIER (',' IDENTIFIER)*)? ')' block 'end'
    ;

returnStatement
    : 'return' (expression (',' expression)*)?
    ;

// ---------------------------
/* Central Operator Hub */
operatorGroup
    : logicalOp
    | comparisonOp
    | arithOp
    | bitwiseOp
    | assignOp
    | unaryOp
    | concatOp
    | compoundAssignOp
    | incrOp
    | coalesceOp
    | shiftAssignOp
    ;

logicalOp:     'and'|'or';
comparisonOp:  '=='|'>='|'<='|'~='|'>'|'<';
arithOp:       '+'|'-'|'*'|'/'|'//'|'%'|'^';
bitwiseOp:     '&'|'|'|'~'|'<<'|'>>';
assignOp:      '='|'+='|'-='|'*='|'/='|'//='|'^='|'&='|'|=';
unaryOp:       'not'|'#'|'-';
concatOp:      '..';
varargOp: '...';
compoundAssignOp: '+=' | '-=' | '*=' | '/=' | '//=' | '^=' | '..=' | '??=';
incrOp: '+_' | '-_';
coalesceOp: '??';
shiftAssignOp: '<<=' | '>>=';


// ---------------------------
/* Keywords and Identifiers */
IDENTIFIER
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;

BOOL
    : 'true' | 'false'
    ;

NIL
    : 'nil'
    ;

NUMBER
    : [0-9]+ ('.' [0-9]+)?
    ;

STRING
    : '"' (~["\\\r\n])* '"'
    | '\'' (~['\\\r\n])* '\''
    ;

// ---------------------------
/* Whitespaces and Comments */
WS
    : [ \t\r\n]+ -> skip
    ;

COMMENT
    : '--' ~[\r\n]* -> skip
    ;