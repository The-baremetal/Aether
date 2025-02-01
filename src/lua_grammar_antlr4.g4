grammar lua_grammar_antlr4;

program: statement*;

statement
    : assignment
    | expression
    | control_statement
    | function_declaration
    | function_call
    | return_statement
    | break_statement
    | label_statement
    ;

control_statement
    : if_statement
    | for_statement
    | while_statement
    | repeat_statement
    | goto_statement
    | do_statement
    ;

statement_terminator: SEPARATOR;

assignment: (KW_LOCAL)? identifier_list ('=' | '+=' | '-=' | '*=' | '/=' | '//=' | '^=') expression_list statement_terminator;

expression
    : literal
    | identifier
    | prefix_expression
    | expression '^' expression
    | '(' expression ')'
    | expression operators expression
    | function_call
    | table
    | table_access
    | function_expression
    | expression '??' expression
    | expression '++'
    | '++' expression
    | expression '--'
    | '--' expression
    ;

prefix_expression
    : primary_expression
    | KW_NOT prefix_expression
    | '#' prefix_expression
    | '-' prefix_expression
    | '~' prefix_expression
    | INCREMENT prefix_expression
    | DECREMENT prefix_expression
    ;

primary_expression
    : literal
    | identifier
    | '(' expression ')'
    | function_call
    | table
    | table_access
    | primary_expression INCREMENT
    | primary_expression DECREMENT
    ;

operators
    : comparison_operator
    | arith_operator
    | logical_operator
    | bitwise_operator
    | concat_operator
    ;

comparison_operator: '>' | '<' | '>=' | '==' | '<=' | '~=';
arith_operator: '*' | '/' | '+' | '-' | '//';
logical_operator: KW_AND | KW_OR;
bitwise_operator: '&' | '|' | '~' | '<<' | '>>' | '>>>';
concat_operator: '..';

literal
    : NUMBER
    | STRING
    | KW_TRUE
    | KW_FALSE
    | KW_NIL
    | KW_NAN
    | KW_INF
    ;

function_call: 
    (KW_PCALL | KW_XPCALL) '(' expression_list ')'
    | (KW_ERROR | KW_ASSERT) '(' expression_list ')'
    | (identifier | table_access | '(' expression ')') 
    (':' identifier)? '(' expression_list? ')'
    | table_insert
    | KW_PRINT '(' expression_list ')'
    | method_call
    ;

table_insert: 'table.insert' '(' identifier ',' expression ')';

function_declaration: 
    (KW_LOCAL)? KW_FUNCTION (identifier '.' identifier)? 
    identifier '(' (parameter (',' parameter)* | VARARG)? ')' block KW_END;

parameter: identifier ('=' expression)?;

block: (statement_terminator* statement statement_terminator?)+;

if_statement: KW_IF expression KW_THEN block 
    (KW_ELSEIF expression KW_THEN block)* 
    (KW_ELSE block)? 
    KW_END;
for_statement
    : KW_FOR identifier KW_IN expression KW_DO block KW_END
    | KW_FOR identifier '=' expression ',' expression (',' expression)? KW_DO block KW_END
    ;
while_statement: KW_WHILE expression KW_DO block KW_END;
do_statement: KW_DO block KW_END;

table: '{' (field (field_separator field)* field_separator?)? '}';
field_separator: ',' | ';';

field: 
    identifier '=' expression 
    | '[' expression ']' '=' expression
    | expression;

table_access: identifier '[' expression ']' | identifier '.' identifier;

single_line_comment: '--';

SEPARATOR: ';' | ('\r'? '\n' | '\r');

print_statement: KW_PRINT '(' expression ')';

KW_IN: 'in';
KW_PRINT: 'print';
KW_AND: 'and';
KW_BREAK: 'break';
KW_DO: 'do';
KW_ELSE: 'else';
KW_ELSEIF: 'elseif';
KW_END: 'end';
KW_FALSE: 'false';
KW_FOR: 'for';
KW_GOTO: 'goto';
KW_IF: 'if';
KW_NIL: 'nil';
KW_NOT: 'not';
KW_OR: 'or';
KW_REPEAT: 'repeat';
KW_RETURN: 'return';
KW_THEN: 'then';
KW_TRUE: 'true';
KW_UNTIL: 'until';
KW_WHILE: 'while';
KW_LOCAL: 'local';
KW_FUNCTION: 'function';
KW_INDEX: 'index';
KW_NEWINDEX: 'newindex';
KW_MODE: 'mode';
KW_PCALL: 'pcall';
KW_XPCALL: 'xpcall';
KW_COROUTINE: 'coroutine';
KW_CREATE: 'create';
KW_RESUME: 'resume';
KW_YIELD: 'yield';
KW_STATUS: 'status';
KW_NAN: 'nan';
KW_INF: 'inf';
KW_ERROR: 'error';
KW_ASSERT: 'assert';
KW_PAIRS: 'pairs';
KW_IPAIRS: 'ipairs';

identifier: LETTER (LETTER | DIGIT | '_')*;

NUMBER: [0-9]+ ('.' [0-9]+)?;

STRING: 
    '"' (ESC | ~["\\])* '"' 
    | '\'' (ESC | ~['\\])* '\''
    | '[' '='* '[' .*? ']' '='* ']';

fragment ESC: '\\' ["\\/bfnrt];

LETTER: [a-zA-Z_];

DIGIT: [0-9];

WS: [ \t]+ -> skip;

repeat_statement: KW_REPEAT block KW_UNTIL expression (statement_terminator | EOF);

identifier_list: identifier (',' identifier)*;
expression_list: expression (',' expression)* | VARARG;

return_statement: KW_RETURN (expression (',' expression)*)? statement_terminator;

break_statement: KW_BREAK statement_terminator;

goto_statement: KW_GOTO identifier statement_terminator;
label_statement: '::' identifier '::';

function_expression: KW_FUNCTION '(' (identifier (',' identifier)*)? ')' block KW_END;

method_call: identifier ':' identifier '(' (expression (',' expression)*)? ')';

metatable_field: '__' (metamethod | identifier) '=' expression;

metamethod: '__add' | '__sub' | '__mul' | '__div' | '__mod' | '__pow' 
          | '__unm' | '__concat' | '__len' | '__eq' | '__lt' | '__le'
          | '__tostring' | '__pairs' | '__ipairs' | '__call';

coroutine_statement: KW_COROUTINE '.' (KW_CREATE | KW_RESUME | KW_YIELD | KW_STATUS);

SINGLE_LINE_COMMENT: '--' ~[\r\n]* -> skip;
MULTI_LINE_COMMENT: '--[[' .*? ']]' -> skip;

VARARG: '...';

// New operators
INCREMENT: '++';
DECREMENT: '--';
PLUS_ASSIGN: '+=';
MINUS_ASSIGN: '-=';
MULT_ASSIGN: '*=';
DIV_ASSIGN: '/=';
FLOOR_DIV_ASSIGN: '//=';
EXP_ASSIGN: '^=';
CONCAT_ASSIGN: '..=';
NULL_COALESCE: '??';
