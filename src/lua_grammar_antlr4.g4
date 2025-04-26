// should be used for operators and new syntax instead of statements now.
// 1 based indexing is disabled by default
// immutable internal module allows immutable objects or it can be built into the compiler which is better

grammar lua_grammar_antlr4;

options {
    language = Go;
}

// **Main entry point**
program
    : statement* EOF
    ;

// ---------------------------
/* Statements (Assignment, Control Flow, etc.) */
statement
    : assignStatement             #assignStmt
    | controlFlowStatement        #controlFlowStmt
    | coroutineStatement          #coroutineStmt
    | functionDeclaration         #functionDecl
    | localDeclaration            #localDecl
    | expression (';')?           #expressionStmt
    | constDeclaration            #constDecl
    | returnStatement             #returnStmt
    ;

// **Adding a new statement**: 
// 1. Add a new rule in the 'statement' section. 
// 2. It should follow the pattearn: newStatementName
// Example:
// newStatement
//     : 'newKeyword' expression ';'
//     ;
assignStatement
    : variable (assignOp) expression
    ;

// ---------------------------
/* Expressions (Literals, Variables, Operations) */
expression
    : unaryOp expression
    | 'async' block 'end'
    | 'await' expression
    | primaryExpression (operatorGroup expression)*
    | expression operatorGroup expression
    |<assoc=right> expression operatorGroup expression
    ;

primaryExpression
    : (literal
    | functionCall
    | variable
    | unaryOperation
    | tableConstructor
    | arrayConstructor
    | functionExpression
    | '(' expression ')'
    ) ( callChain )*
    ;

callChain
    : ':' IDENTIFIER '(' expressionList? ')'  #methodChain
    | '[' expression ']'                       #indexChain
    ;

literal
    : BOOL
    | NIL
    | NUMBER
    | STRING
    ;

variable
    : ( 'const' )? IDENTIFIER
    | variable ( '.' IDENTIFIER | '[' expression ']' | '?.' IDENTIFIER | '?[' expression ']' )
    ;

functionCall
    : variable '(' (expression (',' expression)*)? ')'
    | variable ':' IDENTIFIER '(' (expression (',' expression)*)? ')'
    ;

tableConstructor
    : '{' (field (',' field)*)? (',' metatable)? '}'
    ;

arrayConstructor
    : '[' expressionList? ']'
    ;

metatable
    : '__metatable' '=' expression
    | '{' metamethods '}'
    ;

metamethods
    : metamethod (',' metamethod)*
    ;

pattern
    : literal
    | tableConstructor
    | IDENTIFIER
    | pattern '|' pattern
    | '{' fieldPattern (',' fieldPattern)* '}'
    ;

fieldPattern
    : IDENTIFIER '=' pattern
    | '[' expression ']' '=' pattern
    | pattern
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
    | expression comparisonOp expression
    | expression logicalOp expression
    | expression concatOp expression
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
    | tryCatchStatement
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

coroutineStatement
    : 'coroutine' '.' ( 'create' | 'resume' | 'wrap' ) '(' expressionList? ')'        #coroutineFuncWithArgs
    | 'coroutine' '.' ( 'yield' | 'status' | 'running' | 'isyieldable' )              #coroutinePropAccess
    ;

namedArgs
    : IDENTIFIER '=' expression (',' IDENTIFIER '=' expression)*
    ;

block
    : statement*
    ;

// ---------------------------
/* Declarations (Local Variables, Functions) */

// Static typed is recommended for readability
localDeclaration
    : 'local' 'const'? ( IDENTIFIER (','  IDENTIFIER)* '=' expressionList)  #typedLocalDecl
    | 'local' 'const'? IDENTIFIER (',' IDENTIFIER)* '=' expressionList  #untypedLocalDecl
    | 'local' 'function' IDENTIFIER '(' (parameterList)? ')' ? block 'end'  #localFunctionDecl
    ;

constDeclaration
    : 'const' IDENTIFIER '=' expression  #typedConstDecl
    | 'const' IDENTIFIER (',' IDENTIFIER)* '=' expressionList  #untypedConstDecl
    ;

functionDeclaration
    : (DOC_COMMENT* 'function') 
      (IDENTIFIER '.' | IDENTIFIER ':')? IDENTIFIER
      '(' (parameterList)? ')' 
      block 
      'end'
    ;

returnStatement
    : RETURN (expression (',' expression)* | functionCall)?
    ;

// ---------------------------
/* Central Operator Hub */
operatorGroup
    : logicalOp
    | comparisonOp
    | arithOp
    | assignOp
    | unaryOp
    | concatOp
    ;

logicalOp:     'and'|'or';
comparisonOp:  '=='|'>='|'<='|'~='|'>'|'<'|'in';
arithOp:       '+'|'-'|'*'|'/'|'%'|'^';
assignOp:      '=';
unaryOp:       'not'|'#'|'-'|'~';
concatOp:      '..';
varargOp: '...';

// ---------------------------
/* Keywords and Identifiers */
CONST: 'const';
NIL: 'nil';
BOOL: 'true'|'false';
FUNCTION: 'function';
LOCAL: 'local';
RETURN: 'return';
IF: 'if';
THEN: 'then';
ELSE: 'else';
ELSEIF: 'elseif';
END: 'end';
WHILE: 'while';
DO: 'do';
REPEAT: 'repeat';
UNTIL: 'until';
FOR: 'for';
IN: 'in';
BREAK: 'break';
AND: 'and';
OR: 'or';
NOT: 'not';
ASYNC: 'async';
AWAIT: 'await';
WITH: 'with';
TRY: 'try';
CATCH: 'catch';
FINALLY: 'finally';

IDENTIFIER
    : [a-zA-Z_][a-zA-Z_0-9]*
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

DOC_COMMENT
    : '---' ~[\r\n]*
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
    : 'function' (IDENTIFIER ':')? '(' (parameterList)? ')'?
      block 'end'
    ;

parameterList
    : (IDENTIFIER  (',' IDENTIFIER )*) 
      (',' varargOp)
    | varargOp
    ;

// Add require as a special expression form
//requireExpression
//    : 'require' '(' expression ')'
//    ;

// ---------------------------
/* Error Handling */
tryCatchStatement
    : 'try' block 
      ('catch' (IDENTIFIER)? block)?
      ('finally' block)?
      'end'
    ;
