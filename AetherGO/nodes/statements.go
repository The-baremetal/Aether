package nodes

import (
	"FLUX/AetherGO/core"
	"github.com/llir/llvm/ir"
)

type IfStatement struct {
	Condition  interface{}
	ThenBlock []interface{}
	ElseBlock []interface{}
}

func (is *IfStatement) GenIR(c *core.Compiler) error {
	condVal, err := core.GenerateExpressionIR(c, is.Condition)
	if err != nil {
		return err
	}
	if !condVal.Type().Equal(types.I1) {
		condVal = c.Builder.CurrentBlock().NewTrunc(condVal, types.I1)
	}
	thenBlock := c.curFunc.NewBlock("if.then")
	elseBlock := c.curFunc.NewBlock("if.else")
	contBlock := c.curFunc.NewBlock("if.end")
	
	c.Builder.CurrentBlock().NewCondBr(condVal, thenBlock, elseBlock)
	c.Builder.blocks = append(c.Builder.blocks, thenBlock)
	for _, stmt := range is.ThenBlock {
		if err := core.GenerateStatementIR(c, stmt); err != nil {
			return err
		}
	}
	thenBlock.NewBr(contBlock)
	c.Builder.blocks = append(c.Builder.blocks, elseBlock)
	for _, stmt := range is.ElseBlock {
		if err := core.GenerateStatementIR(c, stmt); err != nil {
			return err
		}
	}
	elseBlock.NewBr(contBlock)
	c.Builder.blocks = append(c.Builder.blocks, contBlock)
	return nil
}

type Assignment struct {
	Variable string
	Value    interface{}
}

func (is *IfStatement) GenIR(c *core.Compiler) error {
	condVal, err := core.GenerateExpressionIR(c, is.Condition)
	if err != nil {
		return err
	}
	thenBlock := c.curFunc.NewBlock("if.then")
	elseBlock := c.curFunc.NewBlock("if.else")
	contBlock := c.curFunc.NewBlock("if.end")
	c.Builder.CurrentBlock().NewCondBr(condVal, thenBlock, elseBlock)
	c.Builder.blocks = append(c.Builder.blocks, thenBlock)
	for _, stmt := range is.ThenBlock {
		if err := core.GenerateStatementIR(c, stmt); err != nil {
			return err
		}
	}
	thenBlock.NewBr(contBlock)
	c.Builder.blocks = append(c.Builder.blocks, elseBlock)
	for _, stmt := range is.ElseBlock {
		if err := core.GenerateStatementIR(c, stmt); err != nil {
			return err
		}
	}
	elseBlock.NewBr(contBlock)
	c.Builder.blocks = append(c.Builder.blocks, contBlock)
	return nil
}

type PrintStatement struct {
	Expr interface{}
}

func (p *PrintStatement) GenIR(c *core.Compiler) error {
	exprVal, err := core.GenerateExpressionIR(c, p.Expr)
	if err != nil {
		return err
	}
	var formatStr value.Value
	if exprVal.Type().Equal(types.I8Ptr) {
		formatStr = c.Module.Globals["print_str_fmt"]
	} else {
		formatStr = c.Module.Globals["print_int_fmt"]
	}
	
	printf := c.Module.Funcs["printf"]
	c.Builder.CurrentBlock().NewCall(printf, formatStr, exprVal)
	return nil
}