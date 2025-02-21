package core

import (
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/value"
)

func RegisterDefaultIRRules(c *Compiler) {
	c.AddIRRule("+", addRule)
	c.AddIRRule("-", subRule)
	c.AddIRRule("*", mulRule)
	c.AddIRRule("/", sdivRule)
	c.AddIRRule("==", eqRule)
	c.AddIRRule("!=", neRule)
	c.AddIRRule("<", sltRule)
	c.AddIRRule(">", sgtRule)
}

func addRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewAdd(left, right), nil
}

func subRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewSub(left, right), nil
}

func mulRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewMul(left, right), nil
}

func sdivRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewSDiv(left, right), nil
}

func eqRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewICmp(enum.IPredEQ, left, right), nil
}

func neRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewICmp(enum.IPredNE, left, right), nil
}

func sltRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewICmp(enum.IPredSLT, left, right), nil
}

func sgtRule(c *Compiler, left, right value.Value) (value.Value, error) {
	return c.Builder.CurrentBlock().NewICmp(enum.IPredSGT, left, right), nil
}
