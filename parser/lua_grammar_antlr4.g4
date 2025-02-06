// should be used for operators and new syntax instead of statements now.

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
    : assignStatement
    | controlFlowStatement
    | functionDeclaration
    | returnStatement
    | localDeclaration
    | labelStatement
    | selectStatement
    | 'spawn' expression
    | expression (';')?
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
    : primaryExpression (operatorGroup expression)*
    | expression operatorGroup expression
    |<assoc=right> expression operatorGroup expression
    | unaryOp expression
    | 'match' expression 'with' matchArm+ 'end'
    | 'async' block 'end'
    | 'await' expression
    | expression '!!'       
    ;

primaryExpression
    : (literal
    | functionCall
    | variable
    | unaryOperation
    | tableConstructor
    | functionExpression
    | '(' expression ')'
    | lambdaExpression
    ( callChain )*)
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
    : IDENTIFIER
    | variable safeAccess
    | variable '[' expression ']'
    | variable '.' IDENTIFIER
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
    // | coroutineStatement internal function call too
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

//coroutineStatement
//    : 'coroutine' '.' ( 'create' | 'resume' | 'yield' | 'status' | 'running' | 'wrap' | 'isyieldable' )
//    ;

protectedCallStatement
    : (('pcall' | 'xpcall') ('.' | ':')? IDENTIFIER?)
      '(' (expressionList | namedArgs) ')'
    ;

namedArgs
    : IDENTIFIER '=' expression (',' IDENTIFIER '=' expression)*
    ;

block
    : statement*
    ;

// ---------------------------
/* Declarations (Local Variables, Functions) */
localDeclaration
    : 'local' typeAnnotation? IDENTIFIER ('=' expression)?
    | 'local' IDENTIFIER (',' IDENTIFIER)* '=' expressionList
    | 'local' 'function' IDENTIFIER '(' (parameterList)? ')' block 'end'
    ;

functionDeclaration
    : 'function' (IDENTIFIER '.' | IDENTIFIER ':')? IDENTIFIER
      '(' (parameterList)? ')' block 'end'
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
    | safeAccess
    | nonNullAssertOp
    ;

logicalOp:     'and'|'or';
comparisonOp:  '=='|'>='|'<='|'~='|'>'|'<';
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
    : 'function' (IDENTIFIER ':')? '(' (parameterList)? ')' 
      (block 'end' | '=>' expression)
    ;

parameterList
    : IDENTIFIER (',' IDENTIFIER)* (',' varargOp)?
    | varargOp
    ;

selectStatement
    : 'select' '(' ('#' | expression) ',' expression ')' 
    ;

lambdaExpression
    : '(' (IDENTIFIER (',' IDENTIFIER)*)? ')' '=>' expression
    ;

typeAnnotation
    : ':' typeSpec
    ;

typeSpec
    : 'number' | 'string' | 'boolean' | 'table' | 'function'
    | 'any' | typeSpec '[]' | 'table<' typeSpec ',' typeSpec '>'
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