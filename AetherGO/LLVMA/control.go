package LLVMA

import (
    "github.com/llir/llvm/ir"
    "github.com/llir/llvm/ir/value"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
    "FLUX/AetherGO/utils"
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    "fmt"
)

type CompilerInterface interface {
    GetExprGen() *ExprGenerator
    GetMainFn() *ir.Func 
    GetEntryBlock() *ir.Block
    ProcessBlock(antlr.ParseTree)
}

var ControlRegistry = map[string]func(*ControlFlow, CompilerInterface, antlr.ParseTree) error{
    "ifstatement": handleIf,
    "for":         handleFor,
    "while":       handleWhile,
}

type ControlHook interface {
    BeforeControlFlow(compiler interface{}, ctx antlr.ParseTree)
    AfterControlFlow(compiler interface{}, ctx antlr.ParseTree)
}

func (cf *ControlFlow) RegisterControlHook(hook ControlHook) {
    cf.hooks = append(cf.hooks, hook)
}

type ControlFlow struct {
    Module    *ir.Module
    Builder   *ir.Block
    LoopStack []*LoopContext
    hooks     []ControlHook
}

type LoopContext struct {
    CondBlock *ir.Block
    BodyBlock *ir.Block
    ExitBlock *ir.Block
}

func NewControlFlow(mod *ir.Module, entry *ir.Block) *ControlFlow {
    return &ControlFlow{
        Module:    mod,
        Builder:   entry,
        LoopStack: make([]*LoopContext, 0),
    }
}

func (cf *ControlFlow) CreateIf(cond value.Value, thenBlock, elseBlock, mergeBlock *ir.Block) {
    cf.Builder.NewCondBr(cond, thenBlock, elseBlock)
    cf.Builder = thenBlock
}

func (cf *ControlFlow) FinalizeIf(elseBlock, mergeBlock *ir.Block) {
    cf.Builder.NewBr(mergeBlock)
    cf.Builder = elseBlock
    cf.Builder.NewBr(mergeBlock)
    cf.Builder = mergeBlock
    if len(mergeBlock.Insts) == 0 {
        mergeBlock.NewRet(constant.NewInt(types.I32, 0))
    }
}

func (cf *ControlFlow) StartWhile(condBlock *ir.Block) *LoopContext {
    loop := &LoopContext{
        CondBlock: condBlock,
        BodyBlock: cf.Builder.Parent.NewBlock("while.body"),
        ExitBlock: cf.Builder.Parent.NewBlock("while.exit"),
    }
    cf.LoopStack = append(cf.LoopStack, loop)
    cf.Builder.NewBr(condBlock)
    cf.Builder = condBlock
    return loop
}

func (cf *ControlFlow) EndWhile(cond value.Value, loop *LoopContext) {
    cf.Builder.NewCondBr(cond, loop.BodyBlock, loop.ExitBlock)
    cf.Builder = loop.BodyBlock
    cf.Builder.NewBr(loop.CondBlock)
    cf.Builder = loop.ExitBlock
    cf.LoopStack = cf.LoopStack[:len(cf.LoopStack)-1]
}

func (cf *ControlFlow) CreateFor(init value.Value, condFn func() value.Value, 
    stepFn func(), bodyBlock *ir.Block) *LoopContext {
    
    loop := &LoopContext{
        CondBlock: cf.Builder.Parent.NewBlock("for.cond"),
        BodyBlock: bodyBlock,
        ExitBlock: cf.Builder.Parent.NewBlock("for.exit"),
    }
    if init != nil {
        cf.Builder.NewStore(init, init.(*ir.InstAlloca))
    }
    cf.Builder.NewBr(loop.CondBlock)
    cf.Builder = loop.CondBlock
    cond := condFn()
    loop.CondBlock.NewCondBr(cond, loop.BodyBlock, loop.ExitBlock)
    cf.Builder = loop.BodyBlock
    cf.LoopStack = append(cf.LoopStack, loop)
    return loop
}

func (cf *ControlFlow) EndFor(stepFn func()) {
    if len(cf.LoopStack) == 0 {
        utils.LogError("EndFor called outside loop context")
        return
    }
    
    loop := cf.LoopStack[len(cf.LoopStack)-1]
    stepFn()
    cf.Builder.NewBr(loop.CondBlock)
    cf.Builder = loop.ExitBlock
    cf.LoopStack = cf.LoopStack[:len(cf.LoopStack)-1]
}

func (cf *ControlFlow) HandleBreak() {
    if len(cf.LoopStack) == 0 {
        utils.LogError("break statement outside loop")
        return
    }
    
    currentLoop := cf.LoopStack[len(cf.LoopStack)-1]
    cf.Builder.NewBr(currentLoop.ExitBlock)
    cf.Builder = cf.Builder.Parent.NewBlock("post.break")
}

func handleIf(cf *ControlFlow, comp CompilerInterface, ctx antlr.ParseTree) error {
    ifCtx, ok := ctx.(*parser.IfStatementContext)
    if !ok {
        return fmt.Errorf("expected IfStatementContext, got %T", ctx)
    }
    cond := comp.GetExprGen().GenerateExpression(ifCtx.Expression(0))
    
    thenBlock := comp.GetMainFn().NewBlock("if.then")
    elseBlock := comp.GetMainFn().NewBlock("if.else")
    mergeBlock := comp.GetMainFn().NewBlock("if.merge")
    
    origBuilder := cf.Builder
    cf.Builder = thenBlock
    comp.ProcessBlock(ifCtx.Block(0))
    thenTerminated := thenBlock.Term != nil
    
    var elseTerminated bool
    if ifCtx.Block(1) != nil {
        cf.Builder = elseBlock
        comp.ProcessBlock(ifCtx.Block(1))
        elseTerminated = elseBlock.Term != nil
    } else {
        elseTerminated = false
    }

    cf.Builder = origBuilder
    origBuilder.NewCondBr(cond, thenBlock, elseBlock)
    if !thenTerminated {
        thenBlock.NewBr(mergeBlock)
    }
    if !elseTerminated {
        elseBlock.NewBr(mergeBlock)
    }
    if thenTerminated && elseTerminated {
        filtered := make([]*ir.Block, 0, len(comp.GetMainFn().Blocks)-1)
        for _, blk := range comp.GetMainFn().Blocks {
            if blk != mergeBlock {
                filtered = append(filtered, blk)
            }
        }
        comp.GetMainFn().Blocks = filtered
    } else {
        cf.Builder = mergeBlock
        if mergeBlock.Term == nil {
            mergeBlock.NewUnreachable()
        }
    }
    
    return nil
}

func handleFor(cf *ControlFlow, c CompilerInterface, ctx antlr.ParseTree) error {
    forCtx, ok := ctx.(*parser.ForStatementContext)
    if !ok {
        return fmt.Errorf("expected ForStatementContext, got %T", ctx)
    }
    _ = forCtx
    return fmt.Errorf("for loops not implemented yet")
}

func handleWhile(cf *ControlFlow, c CompilerInterface, ctx antlr.ParseTree) error {
    utils.LogError("while loops not implemented yet")
    return nil
} 