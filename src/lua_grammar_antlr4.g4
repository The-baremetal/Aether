// should be used for operators and new syntax instead of statements now.
// 1 based indexing is disabled by default
// immutable internal module allows immutable objects or it can be built into the compiler which is better

grammar Lua_grammar_antlr4;

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
    | functionDeclaration         #functionDecl
    | localDeclaration            #localDecl
    | labelStatement              #labelStmt
    | selectStatement             #selectStmt
    | SPAWN expression ';'?       #spawnStatement
    | expression (';')?           #expressionStmt
    | constDeclaration            #constDecl
    | returnStatement             #returnStmt
    | typeDeclaration             #typeDecl
    ;

// **Adding a new statement**: 
// 1. Add a new rule in the 'statement' section. 
// 2. It should follow the pattern: newStatementName
// Example:
// newStatement
//     : 'newKeyword' expression ';'
//     ;
assignStatement
    : variable (assignOp | incrOp | shiftAssignOp | coalesceOp) expression
    ;

// ---------------------------
/* Expressions (Literals, Variables, Operations) */
expression
    : unaryOp expression
    | 'match' expression 'with' matchArm+ 'end'
    | 'async' block 'end'
    | 'await' expression
    | primaryExpression (operatorGroup expression)*
    | expression operatorGroup expression
    |<assoc=right> expression operatorGroup expression
    | expression '|>' expression
    | expression '?.' IDENTIFIER
    | expression '?[' expression ']'
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
    ) ( callChain | typeAssertion )*
    ;

callChain
    : ':' IDENTIFIER '(' expressionList? ')'  #methodChain
    | '.' IDENTIFIER                           #propertyChain
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

safeAccess
    : '?.' IDENTIFIER
    | '?[' expression ']'
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

labelStatement
    : '::' IDENTIFIER '::'
    ;

matchArm
    : pattern ('when' expression)? '=>' expression
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

gotoStatement
    : 'goto' IDENTIFIER
    ;

// TODO: Should i add the throw statement or should it be a function call? well i guess the compiler will handle it, its not the parser's problem
// anyways, throw should be like this
// throw("Something went wrong!", aether.ErrorTypes.RuntimeError)

//coroutineStatement
//    : 'coroutine' '.' ( 'create' | 'resume' | 'yield' | 'status' | 'running' | 'wrap' | 'isyieldable' )
//    ;

// corountine can be autodetected because it is essentially a function call

/* pcall is trash, proper error handling
protectedCallStatement
    : (('pcall' | 'xpcall') ('.' | ':')? IDENTIFIER?)
      '(' (expressionList | namedArgs) ')'
    ;
*/



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
    : 'local' 'const'? (typeAnnotation IDENTIFIER (',' typeAnnotation IDENTIFIER)* '=' expressionList)  #typedLocalDecl
    | 'local' 'const'? IDENTIFIER (',' IDENTIFIER)* '=' expressionList  #untypedLocalDecl
    | 'local' 'function' IDENTIFIER '(' (parameterList)? ')' (typeAnnotation)? block 'end'  #localFunctionDecl
    ;

constDeclaration
    : 'const' typeAnnotation IDENTIFIER '=' expression  #typedConstDecl
    | 'const' IDENTIFIER (',' IDENTIFIER)* '=' expressionList  #untypedConstDecl
    ;

functionDeclaration
    : (DOC_COMMENT* 'function') 
      (IDENTIFIER '.' | IDENTIFIER ':')? IDENTIFIER
      '(' (parameterList)? ')' 
      (typeAnnotation)? 
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
    | bitwiseOp
    | assignOp
    | unaryOp
    | concatOp
    | compoundAssignOp
    | incrOp
    | coalesceOp
    | shiftAssignOp
    | safeAccess
    | nonNullAssertOp
    ;

logicalOp:     'and'|'or';
comparisonOp:  '=='|'>='|'<='|'~='|'>'|'<'|'in';
arithOp:       '+'|'-'|'*'|'/'|'//'|'%'|'^';
bitwiseOp:     '&'|'|'|'~'|'<<'|'>>';
assignOp:      '='|'+='|'-='|'*='|'/='|'//='|'^='|'&='|'|=';
unaryOp:       'not'|'#'|'-'|'~'|'typeof';
concatOp:      '..';
varargOp: '...';
compoundAssignOp: 
    '+=' | '-=' | '*=' | '/=' | '//=' | '^=' | '..=' | '??=' |
    '+_=' | '-_='
    ;
incrOp: '+_' | '-_';
coalesceOp: '??';
shiftAssignOp: '<<=' | '>>=';
nonNullAssertOp: '!!';

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
GOTO: 'goto';
AND: 'and';
OR: 'or';
NOT: 'not';
ASYNC: 'async';
AWAIT: 'await';
MATCH: 'match';
WITH: 'with';
SELECT: 'select';
TRY: 'try';
CATCH: 'catch';
FINALLY: 'finally';
TYPEOF: 'typeof';

SPAWN: 'spawn';
VOID: 'void';
TYPE: 'type';
SAFE_NAV: '?.';
PIPE: '|>';
ARROW: '=>';

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
    : 'function' (IDENTIFIER ':')? '(' (parameterList)? ')' (typeAnnotation)?
      block 'end'
    ;

parameterList
    : (IDENTIFIER typeAnnotation? (',' IDENTIFIER typeAnnotation?)*) 
      (',' varargOp typeAnnotation?)?
    | varargOp typeAnnotation?
    ;

selectStatement
    : 'select' '(' ('#' | expression) ',' expression ')' 
    ;

typeAnnotation
    : ':' typeSpec ('?' | '|' typeSpec)*
    ;

typeDeclaration
    : 'type' IDENTIFIER '=' typeSpec
    ;

typeSpec
    : ( 'number' | 'string' | 'boolean' | 'table' | 'function' | 'void'
      | 'any' | 'nil' | 'unknown' | IDENTIFIER )
      ('[]' | '<' typeSpec (',' typeSpec)* '>')*
    | 'function' '(' (typeSpec (',' typeSpec)*)? ')' '=>' typeSpec
    | '(' typeSpec ')'
    ;

typeAssertion
    : primaryExpression 'as' typeSpec
    ;

// ---------------------------
/* 
$$\      $$\  $$$$$$\  $$$$$$\ $$$$$$$$\ $$\       $$$$$$$\  $$$$$$$$\ $$$$$$$$\  $$$$$$\  $$$$$$$\  $$$$$$$$\       $$\      $$\  $$$$$$\  $$\   $$\ $$$$$$\ $$\   $$\  $$$$$$\         $$$$$$\  $$\   $$\  $$$$$$\  $$\   $$\  $$$$$$\  $$$$$$$$\  $$$$$$\            $$$$$$$\  $$\       $$$$$$$$\  $$$$$$\   $$$$$$\  $$$$$$$$\       $$$$$$$$\ $$$$$$\ $$$$$$$\   $$$$$$\ $$$$$$$$\       $$\   $$\  $$$$$$\  $$$$$$$$\       $$$$$$$$\ $$\   $$\ $$$$$$$$\       $$$$$$$$\ $$\   $$\ $$$$$$$\  $$$$$$$$\ $$$$$$$\  $$$$$$\ $$\      $$\ $$$$$$$$\ $$\   $$\ $$$$$$$$\  $$$$$$\         $$$$$$\  $$$$$$$$\  $$$$$$\ $$$$$$$$\ $$$$$$\  $$$$$$\  $$\   $$\       $$$$$$$$\  $$$$$$\         $$$$$$\  $$\    $$\  $$$$$$\  $$$$$$\ $$$$$$$\        $$$$$$$\  $$$$$$$\  $$$$$$$$\  $$$$$$\  $$\   $$\ $$$$$$\ $$\   $$\  $$$$$$\        $$\   $$\  $$$$$$\  $$$$$$$$\ $$$$$$$\         $$$$$$\   $$$$$$\  $$$$$$$\  $$$$$$$$\ $$\ 
$$ | $\  $$ |$$  __$$\ \_$$  _|\__$$  __|$$ |      $$  __$$\ $$  _____|$$  _____|$$  __$$\ $$  __$$\ $$  _____|      $$$\    $$$ |$$  __$$\ $$ | $$  |\_$$  _|$$$\  $$ |$$  __$$\       $$  __$$\ $$ |  $$ |$$  __$$\ $$$\  $$ |$$  __$$\ $$  _____|$$  __$$\           $$  __$$\ $$ |      $$  _____|$$  __$$\ $$  __$$\ $$  _____|      $$  _____|\_$$  _|$$  __$$\ $$  __$$\\__$$  __|      $$ |  $$ |$$  __$$\ $$  _____|      \__$$  __|$$ |  $$ |$$  _____|      $$  _____|$$ |  $$ |$$  __$$\ $$  _____|$$  __$$\ \_$$  _|$$$\    $$$ |$$  _____|$$$\  $$ |\__$$  __|$$  __$$\       $$  __$$\ $$  _____|$$  __$$\\__$$  __|\_$$  _|$$  __$$\ $$$\  $$ |      \__$$  __|$$  __$$\       $$  __$$\ $$ |   $$ |$$  __$$\ \_$$  _|$$  __$$\       $$  __$$\ $$  __$$\ $$  _____|$$  __$$\ $$ | $$  |\_$$  _|$$$\  $$ |$$  __$$\       $$ |  $$ |$$  __$$\ $$  _____|$$  __$$\       $$  __$$\ $$  __$$\ $$  __$$\ $$  _____|$$ |
$$ |$$$\ $$ |$$ /  $$ |  $$ |     $$ |   $$ |      $$ |  $$ |$$ |      $$ |      $$ /  $$ |$$ |  $$ |$$ |            $$$$\  $$$$ |$$ /  $$ |$$ |$$  /   $$ |  $$$$\ $$ |$$ /  \__|      $$ /  \__|$$ |  $$ |$$ /  $$ |$$$$\ $$ |$$ /  \__|$$ |      $$ /  \__|          $$ |  $$ |$$ |      $$ |      $$ /  $$ |$$ /  \__|$$ |            $$ |        $$ |  $$ |  $$ |$$ /  \__|  $$ |         $$ |  $$ |$$ /  \__|$$ |               $$ |   $$ |  $$ |$$ |            $$ |      \$$\ $$  |$$ |  $$ |$$ |      $$ |  $$ |  $$ |  $$$$\  $$$$ |$$ |      $$$$\ $$ |   $$ |   $$ /  \__|      $$ /  \__|$$ |      $$ /  \__|  $$ |     $$ |  $$ /  $$ |$$$$\ $$ |         $$ |   $$ /  $$ |      $$ /  $$ |$$ |   $$ |$$ /  $$ |  $$ |  $$ |  $$ |      $$ |  $$ |$$ |  $$ |$$ |      $$ /  $$ |$$ |$$  /   $$ |  $$$$\ $$ |$$ /  \__|      $$ |  $$ |$$ /  \__|$$ |      $$ |  $$ |      $$ /  \__|$$ /  $$ |$$ |  $$ |$$ |      $$ |
$$ $$ $$\$$ |$$$$$$$$ |  $$ |     $$ |   $$ |      $$$$$$$\ |$$$$$\    $$$$$\    $$ |  $$ |$$$$$$$  |$$$$$\          $$\$$\$$ $$ |$$$$$$$$ |$$$$$  /    $$ |  $$ $$\$$ |$$ |\_$$ |      $$ |      $$$$$$$$ |$$$$$$$$ |$$ $$\$$ |$$ |$$$$\ $$$$$\    \$$$$$$\            $$$$$$$  |$$ |      $$$$$\    $$$$$$$$ |\$$$$$$\  $$$$$\          $$$$$\      $$ |  $$$$$$$  |\$$$$$$\    $$ |         $$ |  $$ |\$$$$$$\  $$$$$\             $$ |   $$$$$$$$ |$$$$$\          $$$$$\     \$$$$  / $$$$$$$  |$$$$$\    $$$$$$$  |  $$ |  $$\$$\$$ $$ |$$$$$\    $$ $$\$$ |   $$ |   \$$$$$$\        \$$$$$$\  $$$$$\    $$ |        $$ |     $$ |  $$ |  $$ |$$ $$\$$ |         $$ |   $$ |  $$ |      $$$$$$$$ |\$$\  $$  |$$ |  $$ |  $$ |  $$ |  $$ |      $$$$$$$\ |$$$$$$$  |$$$$$\    $$$$$$$$ |$$$$$  /    $$ |  $$ $$\$$ |$$ |$$$$\       $$ |  $$ |\$$$$$$\  $$$$$\    $$$$$$$  |      $$ |      $$ |  $$ |$$ |  $$ |$$$$$\    $$ |
$$$$  _$$$$ |$$  __$$ |  $$ |     $$ |   \__|      $$  __$$\ $$  __|   $$  __|   $$ |  $$ |$$  __$$< $$  __|         $$ \$$$  $$ |$$  __$$ |$$  $$<     $$ |  $$ \$$$$ |$$ |\_$$ |      $$ |      $$  __$$ |$$  __$$ |$$ \$$$$ |$$ |\_$$ |$$  __|    \____$$\           $$  ____/ $$ |      $$  __|   $$  __$$ | \____$$\ $$  __|         $$  __|     $$ |  $$  __$$<  \____$$\   $$ |         $$ |  $$ | \____$$\ $$  __|            $$ |   $$  __$$ |$$  __|         $$  __|    $$  $$<  $$  ____/ $$  __|   $$  __$$<   $$ |  $$ \$$$  $$ |$$  __|   $$ \$$$$ |   $$ |    \____$$\        \____$$\ $$  __|   $$ |        $$ |     $$ |  $$ |  $$ |$$ \$$$$ |         $$ |   $$ |  $$ |      $$  __$$ | \$$\$$  / $$ |  $$ |  $$ |  $$ |  $$ |      $$  __$$\ $$  __$$< $$  __|   $$  __$$ |$$  $$<     $$ |  $$ \$$$$ |$$ |\_$$ |      $$ |  $$ | \____$$\ $$  __|   $$  __$$<       $$ |      $$ |  $$ |$$ |  $$ |$$  __|   \__|
$$$  / \$$$ |$$ |  $$ |  $$ |     $$ |             $$ |  $$ |$$ |      $$ |      $$ |  $$ |$$ |  $$ |$$ |            $$ |\$  /$$ |$$ |  $$ |$$ |\$$\    $$ |  $$ |\$$$ |$$ |  $$ |      $$ |  $$\ $$ |  $$ |$$ |  $$ |$$ |\$$$ |$$ |  $$ |$$ |      $$\   $$ |          $$ |      $$ |      $$ |      $$ |  $$ |$$\   $$ |$$ |            $$ |        $$ |  $$ |  $$ |$$\   $$ |  $$ |         $$ |  $$ |$$\   $$ |$$ |               $$ |   $$ |  $$ |$$ |            $$ |      $$  /\$$\ $$ |      $$ |      $$ |  $$ |  $$ |  $$ |\$  /$$ |$$ |      $$ |\$$$ |   $$ |   $$\   $$ |      $$\   $$ |$$ |      $$ |  $$\   $$ |     $$ |  $$ |  $$ |$$ |\$$$ |         $$ |   $$ |  $$ |      $$ |  $$ |  \$$$  /  $$ |  $$ |  $$ |  $$ |  $$ |      $$ |  $$ |$$ |  $$ |$$ |      $$ |  $$ |$$ |\$$\    $$ |  $$ |\$$$ |$$ |  $$ |      $$ |  $$ |$$\   $$ |$$ |      $$ |  $$ |      $$ |  $$\ $$ |  $$ |$$ |  $$ |$$ |          
$$  /   \$$ |$$ |  $$ |$$$$$$\    $$ |   $$\       $$$$$$$  |$$$$$$$$\ $$ |       $$$$$$  |$$ |  $$ |$$$$$$$$\       $$ | \_/ $$ |$$ |  $$ |$$ | \$$\ $$$$$$\ $$ | \$$ |\$$$$$$  |      \$$$$$$  |$$ |  $$ |$$ |  $$ |$$ | \$$ |\$$$$$$  |$$$$$$$$\ \$$$$$$  |$$\       $$ |      $$$$$$$$\ $$$$$$$$\ $$ |  $$ |\$$$$$$  |$$$$$$$$\       $$ |      $$$$$$\ $$ |  $$ |\$$$$$$  |  $$ |         \$$$$$$  |\$$$$$$  |$$$$$$$$\          $$ |   $$ |  $$ |$$$$$$$$\       $$$$$$$$\ $$ /  $$ |$$ |      $$$$$$$$\ $$ |  $$ |$$$$$$\ $$ | \_/ $$ |$$$$$$$$\ $$ | \$$ |   $$ |   \$$$$$$  |      \$$$$$$  |$$$$$$$$\ \$$$$$$  |  $$ |   $$$$$$\  $$$$$$  |$$ | \$$ |         $$ |    $$$$$$  |      $$ |  $$ |   \$  /    $$$$$$  |$$$$$$\ $$$$$$$  |      $$$$$$$  |$$ |  $$ |$$$$$$$$\ $$ |  $$ |$$ | \$$\ $$$$$$\ $$ | \$$ |\$$$$$$  |      \$$$$$$  |\$$$$$$  |$$$$$$$$\ $$ |  $$ |      \$$$$$$  | $$$$$$  |$$$$$$$  |$$$$$$$$\ $$\ 
\__/     \__|\__|  \__|\______|   \__|   \__|      \_______/ \________|\__|       \______/ \__|  \__|\________|      \__|     \__|\__|  \__|\__|  \__|\______|\__|  \__| \______/        \______/ \__|  \__|\__|  \__|\__|  \__| \______/ \________| \______/ $  |      \__|      \________|\________|\__|  \__| \______/ \________|      \__|      \______|\__|  \__| \______/   \__|          \______/  \______/ \________|         \__|   \__|  \__|\________|      \________|\__|  \__|\__|      \________|\__|  \__|\______|\__|     \__|\________|\__|  \__|   \__|    \______/        \______/ \________| \______/   \__|   \______| \______/ \__|  \__|         \__|    \______/       \__|  \__|    \_/     \______/ \______|\_______/       \_______/ \__|  \__|\________|\__|  \__|\__|  \__|\______|\__|  \__| \______/        \______/  \______/ \________|\__|  \__|       \______/  \______/ \_______/ \________|\__|

(very cool right? the script to handle this is is in the location /INTERNALSCRIPTS/EXPERIMENTS/init.lua) NOTE THIS SCRIPT ISNT IN STANDARD LUA, IT IS IN THE LANGUAGE THIS PROJECT IS ITSELF. This project is coming together and I'm very excited to have the build ready in the next 8 months.
okay fine, I'm sorry for making you read that. For now, the internal scripts wont work because the parser is not finished and the compiler doesn't exist. after that, you can stop moving parts of your code from experimentalExpression to the Expression rule. Thanks for reading this pointless comment. These experiment rules are for multi merge collaboration so when you work on it, please move the code back...when you commit unfinshed code, please move this back to the experiment rule.
*/

experimentalExpression
    : safeNavigation // {experiments.isEnabled("safe_nav")}? (removed for: IMPLEMENTATION.NOTDONE)
    | pipeOperator // {experiments.isEnabled("pipe")}? (removed for: IMPLEMENTATION.NOTDONE)
    | decoratorSyntax // {experiments.isEnabled("decorators")}? (removed for: IMPLEMENTATION.NOTDONE)
    ;

safeNavigation
    : expression '?.' IDENTIFIER
    | expression '?[' expression ']'
    ;

pipeOperator
    : expression '|>' expression
    ;

decoratorSyntax
    : '@' IDENTIFIER ( '(' expressionList? ')' )?
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