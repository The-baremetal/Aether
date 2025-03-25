
package utils

import "fmt"

type ErrorHandler struct {
	Errors []string
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		Errors: []string{},
	}
}

func (h *ErrorHandler) AddError(line int, message string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: %s", line, message))
}

func (h *ErrorHandler) AddExpectedTokenError(line int, expected string, got string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Expected token '%s', got '%s'", line, expected, got))
}

func (h *ErrorHandler) AddUnexpectedTokenError(line int, got string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Unexpected token '%s'", line, got))
}

func (h *ErrorHandler) AddInvalidIdentifierError(line int, identifier string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid identifier '%s'", line, identifier))
}

func (h *ErrorHandler) AddInvalidOperatorError(line int, operator string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid operator '%s'", line, operator))
}

func (h *ErrorHandler) AddInvalidTypeError(line int, expected string, got string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid type: expected '%s', got '%s'", line, expected, got))
}

func (h *ErrorHandler) AddNotEnoughArgumentsError(line int, expected int, got int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Expected %d arguments, got %d", line, expected, got))
}

func (h *ErrorHandler) AddTooManyArgumentsError(line int, expected int, got int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Expected %d arguments, got %d", line, expected, got))
}

func (h *ErrorHandler) AddUndefinedVariableError(line int, variable string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Undefined variable '%s'", line, variable))
}

func (h *ErrorHandler) AddRedefinedVariableError(line int, variable string) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Redefined variable '%s'", line, variable))
}

func (h *ErrorHandler) AddMissingReturnStatementError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing return statement", line))
}

func (h *ErrorHandler) AddInvalidReturnStatementError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid return statement", line))
}

func (h *ErrorHandler) AddInvalidAssignmentError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid assignment", line))
}

func (h *ErrorHandler) AddInvalidExpressionError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid expression", line))
}

func (h *ErrorHandler) AddInvalidStatementError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid statement", line))
}

func (h *ErrorHandler) AddMissingEndError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'end'", line))
}

func (h *ErrorHandler) AddMissingThenError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'then'", line))
}

func (h *ErrorHandler) AddMissingDoError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'do'", line))
}

func (h *ErrorHandler) AddMissingUntilError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'until'", line))
}

func (h *ErrorHandler) AddMissingInError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'in'", line))
}

func (h *ErrorHandler) AddMissingWithError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'with'", line))
}

func (h *ErrorHandler) AddMissingCatchError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'catch'", line))
}

func (h *ErrorHandler) AddMissingFinallyError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Missing 'finally'", line))
}

func (h *ErrorHandler) AddInvalidTableConstructorError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid table constructor", line))
}

func (h *ErrorHandler) AddInvalidArrayConstructorError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid array constructor", line))
}

func (h *ErrorHandler) AddInvalidSelectStatementError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid select statement", line))
}

func (h *ErrorHandler) AddInvalidMatchArmError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid match arm", line))
}

func (h *ErrorHandler) AddInvalidFieldPatternError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid field pattern", line))
}

func (h *ErrorHandler) AddInvalidTypeAnnotationError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid type annotation", line))
}

func (h *ErrorHandler) AddInvalidTypeDeclarationError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid type declaration", line))
}

func (h *ErrorHandler) AddInvalidVarargError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid vararg", line))
}

func (h *ErrorHandler) AddInvalidSafeAccessError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid safe access", line))
}

func (h *ErrorHandler) AddInvalidPipeOperatorError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid pipe operator", line))
}

func (h *ErrorHandler) AddInvalidDecoratorError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid decorator", line))
}

func (h *ErrorHandler) AddInvalidGotoError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid goto", line))
}

func (h *ErrorHandler) AddInvalidLabelError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid label", line))
}

func (h *ErrorHandler) AddInvalidBreakError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid break", line))
}

func (h *ErrorHandler) AddInvalidContinueError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid continue", line))
}

func (h *ErrorHandler) AddInvalidSpawnError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid spawn", line))
}

func (h *ErrorHandler) AddInvalidAwaitError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid await", line))
}

func (h *ErrorHandler) AddInvalidAsyncError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid async", line))
}

func (h *ErrorHandler) AddInvalidCoalesceError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid coalesce", line))
}

func (h *ErrorHandler) AddInvalidShiftAssignError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid shift assign", line))
}

func (h *ErrorHandler) AddInvalidIncrError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid incr", line))
}

func (h *ErrorHandler) AddInvalidNonNullAssertError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid non-null assert", line))
}

func (h *ErrorHandler) AddInvalidCompoundAssignError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid compound assign", line))
}

func (h *ErrorHandler) AddInvalidConcatError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid concat", line))
}

func (h *ErrorHandler) AddInvalidBitwiseError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid bitwise", line))
}

func (h *ErrorHandler) AddInvalidArithError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid arith", line))
}

func (h *ErrorHandler) AddInvalidComparisonError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid comparison", line))
}

func (h *ErrorHandler) AddInvalidLogicalError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid logical", line))
}

func (h *ErrorHandler) AddInvalidUnaryError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid unary", line))
}

func (h *ErrorHandler) AddInvalidAssignError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid assign", line))
}

func (h *ErrorHandler) AddInvalidLiteralError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid literal", line))
}

func (h *ErrorHandler) AddInvalidVariableError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid variable", line))
}

func (h *ErrorHandler) AddInvalidFunctionCallError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid function call", line))
}

func (h *ErrorHandler) AddInvalidTableError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid table", line))
}

func (h *ErrorHandler) AddInvalidArrayError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid array", line))
}

func (h *ErrorHandler) AddInvalidMetatableError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid metatable", line))
}

func (h *ErrorHandler) AddInvalidMetamethodError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid metamethod", line))
}

func (h *ErrorHandler) AddInvalidFieldError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid field", line))
}

func (h *ErrorHandler) AddInvalidBlockError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid block", line))
}

func (h *ErrorHandler) AddInvalidParameterListError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid parameter list", line))
}

func (h *ErrorHandler) AddInvalidFunctionExpressionError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid function expression", line))
}

func (h *ErrorHandler) AddInvalidNamedArgsError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid named args", line))
}

func (h *ErrorHandler) AddAlreadyDefinedError(redefinedline int, existingline int, variablename str) {
	h.Errors = append(h.Errors, fmt.Sprintf("Variable: (variablename) was redefined at (redefinedline) while being defined at (existingline)")) // this error only happens when the user tries to redefine an immutable variable they have manually made immutable, with the automatic borrowing and mutability, the compiler just defines it as atomic mutable instead.
}

func (h *ErrorHandler) AddInvalidDocCommentError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid doc comment", line))
}

func (h *ErrorHandler) AddInvalidCommentError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid comment", line))
}

func (h *ErrorHandler) AddInvalidIdentifierListError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid identifier list", line))
}

func (h *ErrorHandler) AddInvalidExpressionListError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid expression list", line))
}

func (h *ErrorHandler) AddInvalidProgramError(line int) {
	h.Errors = append(h.Errors, fmt.Sprintf("Error on line %d: Invalid program", line))
}

func (h *ErrorHandler) HasErrors() bool {
	return len(h.Errors) > 0
}

func (h *ErrorHandler) GetErrors() []string {
	return h.Errors
}

func (h *ErrorHandler) PrintErrors() {
	for _, err := range h.Errors {
		fmt.Println(err)
	}
}
