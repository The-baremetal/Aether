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
    | selectStatement
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
    |<assoc=right> expression operatorGroup expression
    | unaryOp expression
    ;

primaryExpression
    : literal
    | variable
    | functionCall
    | unaryOperation
    | tableConstructor
    | functionExpression
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
    | variable '[' expression ']'
    | variable '.' IDENTIFIER
    ;

functionCall
    : variable '(' (expression (',' expression)*)? ')'
    | variable ':' IDENTIFIER '(' (expression (',' expression)*)? ')'
    ;

tableConstructor
    : '{' (field (',' field)*)? (',' metatable)? '}'
    ;

metatable
    : '__metatable' '=' expression
    | '{' metamethods '}'
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
    | '[' expression ']' '=' expression
    | expression
    | IDENTIFIER ':' functionExpression
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
    | protectedCallStatement
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
    : 'for' IDENTIFIER '=' expression ',' expression (',' expression)? 'do' block 'end'  #numericFor
    | 'for' identifierList 'in' expressionList 'do' block 'end'                            #genericFor
    ;

breakStatement
    : 'break'
    ;

gotoStatement
    : 'goto' IDENTIFIER
    ;

coroutineStatement
    : 'coroutine' '.' ( 'create' | 'resume' | 'yield' | 'status' | 'running' | 'wrap' | 'isyieldable' )
    ;

protectedCallStatement
    : (('pcall' | 'xpcall') ('.' | ':')? IDENTIFIER?)
      '(' expression (',' expression)? ')'
    ;

block
    : statement*
    ;

// ---------------------------
/* Declarations (Local Variables, Functions) */
localDeclaration
    : 'local' IDENTIFIER ('=' expression)?
    | 'local' IDENTIFIER (',' IDENTIFIER)* '=' expressionList
    | 'local' 'function' IDENTIFIER '(' (IDENTIFIER (',' IDENTIFIER)*)? ')' block 'end'
    ;

functionDeclaration
    : 'function' (IDENTIFIER '.' | IDENTIFIER ':')? IDENTIFIER
      '(' (IDENTIFIER (',' IDENTIFIER)* (',' varargOp)? | varargOp)? ')' block 'end'
    ;

returnStatement
    : 'return' (expression (',' expression)* | functionCall)?
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
    | '=>'
    ;

logicalOp:     'and'|'or';
comparisonOp:  '=='|'>='|'<='|'~='|'>'|'<';
arithOp:       '+'|'-'|'*'|'/'|'//'|'%'|'^';
bitwiseOp:     '&'|'|'|'~'|'<<'|'>>';
assignOp:      '='|'+='|'-='|'*='|'/='|'//='|'^='|'&='|'|=';
unaryOp:       'not'|'#'|'-'|'~'|'typeof';
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

identifierList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

expressionList
    : expression (',' expression)*
    | varargOp
    ;

functionExpression
    : 'function' (IDENTIFIER ':')? '(' (IDENTIFIER (',' IDENTIFIER)*)? ')' block 'end'
    ;

selectStatement
    : 'select' '(' ('#' | expression) ',' expression ')' 
    ;