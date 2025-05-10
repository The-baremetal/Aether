
package compiler

/*

please dont delete these comments, the import statements will be used later and the rest is only to be deleted as you implement them. note to luohoa97
import "FLUX/src/aether/ast"
import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/constant"
)

*/


// any code outside a function will be automatically added into main. 
/*
funcMain := m.NewFunc("main", types.I32)
	mainBlock := funcMain.NewBlock("entry") // Added a name to the block
	mainBlock.NewRet(mainBlock.NewCall(funcAdd, constant.NewInt(types.I32, 1), mainBlock.NewLoad(types.I32, globalG)))
*/

/*// Global variable
globalG := m.NewGlobalDef("g", constant.NewInt(types.I32, 2))

// Function add
funcAdd := m.NewFunc("add", types.I32,
	ir.NewParam("x", types.I32),
	ir.NewParam("y", types.I32),
)


*/

// more help:

/*
Control Flow

Before we start, we need to prepare compile function for something like expression and statement that not our target.

type Expr interface{ isExpr() Expr }
type EConstant interface {
	Expr
	isEConstant() EConstant
}
type EVoid struct{ EConstant }
type EBool struct {
	EConstant
	V bool
}
type EI32 struct {
	EConstant
	V int64
}
type EVariable struct {
	Expr
	Name string
}
type EAdd struct {
	Expr
	Lhs, Rhs Expr
}
type ELessThan struct {
	Expr
	Lhs, Rhs Expr
}

And compile functions:

func compileConstant(e EConstant) constant.Constant {
	switch e := e.(type) {
	case *EI32:
		return constant.NewInt(types.I32, e.V)
	case *EBool:
		// we have no boolean in LLVM IR
		if e.V {
			return constant.NewInt(types.I1, 1)
		} else {
			return constant.NewInt(types.I1, 0)
		}
	case *EVoid:
		return nil
	}
	panic("unknown expression")
}

func (ctx *Context) compileExpr(e Expr) value.Value {
	switch e := e.(type) {
	case *EVariable:
		return ctx.lookupVariable(e.Name)
	case *EAdd:
		l, r := ctx.compileExpr(e.Lhs), ctx.compileExpr(e.Rhs)
		return ctx.NewAdd(l, r)
	case *ELessThan:
		l, r := ctx.compileExpr(e.Lhs), ctx.compileExpr(e.Rhs)
		return ctx.NewICmp(enum.IPredSLT, l, r)
	case EConstant:
		return compileConstant(e)
	}
	panic("unimplemented expression")
}

EVariable would need context to remember variable's value. Here is the related definition of Context:

type Context struct {
	*ir.Block
	parent *Context
	vars   map[string]value.Value
}

func NewContext(b *ir.Block) *Context {
	return &Context{
		Block: b,
		parent:   nil,
		vars:     make(map[string]value.Value),
	}
}

func (c *Context) NewContext(b *ir.Block) *Context {
	ctx := NewContext(b)
	ctx.parent = c
	return ctx
}

func (c Context) lookupVariable(name string) value.Value {
	if v, ok := c.vars[name]; ok {
		return v
	} else if c.parent != nil {
		return c.parent.lookupVariable(name)
	} else {
		fmt.Printf("variable: `%s`\n", name)
		panic("no such variable")
	}
}

Finally, we would have some simple statement as placeholder:

type Stmt interface{ isStmt() Stmt }
type SDefine struct {
	Stmt
	Name string
	Typ  types.Type
	Expr Expr
}
type SRet struct {
	Stmt
	Val Expr
}

Then compile:

func (ctx *Context) compileStmt(stmt Stmt) {
	if ctx.Parent != nil {
		return
	}
	f := ctx.Parent
	switch s := stmt.(type) {
	case *SDefine:
		v := ctx.NewAlloca(s.Typ)
		ctx.NewStore(ctx.compileExpr(s.Expr), v)
		ctx.vars[s.Name] = v
	case *SRet:
		ctx.NewRet(ctx.compileExpr(s.Val))
	}
}

If

Since we can let:

if condition {
    // A
} else if condition {
    // B
} else {
    // C
}

became:

if condition {
    // A
} else {
    if condition {
        // B
    } else {
        // C
    }
}

We don't have to convert any else-if pattern. Therefore, our If looks like this:

type SIf struct {
	Stmt
	Cond Expr
	Then Stmt
	Else Stmt
}

Then we can get transformers to generate control flow if. Using conditional jump to generate if statement:

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SIf:
		thenCtx := ctx.NewContext(f.NewBlock("if.then"))
		thenCtx.compileStmt(s.Then)
		elseB := f.NewBlock("if.else")
		ctx.NewContext(elseB).compileStmt(s.Else)
		ctx.NewCondBr(ctx.compileExpr(s.Cond), thenCtx.Block, elseB)
		if !thenCtx.HasTerminator() {
			leaveB := f.NewBlock("leave.if")
			thenCtx.NewBr(leaveB)
		}
	}
}

When generating if, the most important thing is leave block, when if-then block complete, a jump to skip else block required since there has no block in high-level language liked concept in LLVM IR. At the end of a basic-block can be a return and since return would terminate a function, jump after return is a dead code, so we have to check we have to generate leave block or not. Here is a small example as usage:

f := ir.NewFunc("foo", types.Void)
bb := f.NewBlock("")

ctx.compileStmt(&SIf{
    Cond: &EBool{V: true},
    Then: &SRet{Val: &EVoid{}},
    Else: &SRet{Val: &EVoid{}},
})

Finally, we get:

define void @foo() {
0:
	br i1 true, label %if.then, label %if.else

if.then:
	ret void

if.else:
	ret void
}

We didn't support else-if directly at here, then we need to know how to handle this via parsing. First, we handle a sequence of if ( <expr> ) <block>. Ok, we can fill AST with Cond and Then, now we should get a token else, then we expect a <block> or if. When we get a <block> this is a obviously can be use as Else, else a if we keep parsing and use it as Else statement since if for sure is a statement. Of course, with this method, generated IR would have some useless label and jump, but flow analyzing should optimize them later, so it's fine.
Switch

LLVM has switch instruction, hence, we can use it directly.

type SSwitch struct {
	Stmt
	Target   Expr
	CaseList []struct {
		EConstant // LLVM IR only takes constant, if you want advanced switch semantic, then you can't directly use this approach
		Stmt
	}
	DefaultCase Stmt
}

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SSwitch:
		cases := []*ir.Case{}
		for _, ca := range s.CaseList {
			caseB := f.NewBlock("switch.case")
			ctx.NewContext(caseB).compileStmt(ca.Stmt)
			cases = append(cases, ir.NewCase(compileConstant(ca.EConstant), caseB))
		}
		defaultB := f.NewBlock("switch.default")
		ctx.NewContext(defaultB).compileStmt(s.DefaultCase)
		ctx.NewSwitch(ctx.compileExpr(s.Target), defaultB, cases...)
	}
}

For every case, we generate a block, then we can jump to target. Then we put statements into case blocks. Finally, we generate switch for the input block. Notice that, switch instruction of LLVM won't generate break automatically, you can use the same trick in the previous section If to generate auto leave block for each case(Go semantic), or record leave block and introduces break statement(C semantic). Now let's test it:

f := ir.NewFunc("foo", types.Void)
ctx := NewContext(f.NewBlock(""))

ctx.compileStmt(&SSwitch{
	Target: &EBool{V: true},
	CaseList: []struct {
		EConstant
		Stmt
	}{
		{EConstant: &EBool{V: true}, Stmt: &SRet{Val: &EVoid{}}},
	},
	DefaultCase: &SRet{Val: &EVoid{}},
})

And output:

define void @foo() {
0:
	switch i1 true, label %switch.default [
		i1 true, label %switch.case
	]

switch.case:
	ret void

switch.default:
	ret void
}

The switch statement in this section is quite naive, for advanced semantic like pattern matching with extraction or where clause, you would need to do more.
Loop
Break

Break statement needs to extend Context, with a new field called leaveBlock:

type Context struct {
	// ...
	leaveBlock *ir.Block
}

func NewContext(b *ir.Block) *Context {
	return &Context{
		// ...
		leaveBlock: nil,
	}
}

Then it's just a jump:

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SBreak:
		ctx.NewBr(ctx.leaveBlock)
	}
}

Remember to update leave block information(and remove it when needed), and continue can be done in the same way.
Do While

Do while is the simplest loop structure since it's code structure almost same to the IR structure. Here we go:

type SDoWhile struct {
	Stmt
	Cond  Expr
	Block Stmt
}

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SDoWhile:
		doCtx := ctx.NewContext(f.NewBlock("do.while.body"))
		ctx.NewBr(doCtx.Block)
		leaveB := f.NewBlock("leave.do.while")
		doCtx.leaveBlock = leaveB
		doCtx.compileStmt(s.Block)
		doCtx.NewCondBr(doCtx.compileExpr(s.Cond), doCtx.Block, leaveB)
	}
}

Can see that, we jump to do-while body directly. Then we have a leave block, in the end of the do-while body we jump out to leave block or body again depends on condition. Let's test it:

f := ir.NewFunc("foo", types.Void)
ctx := NewContext(f.NewBlock(""))

ctx.compileStmt(&SDoWhile{
	Cond: &EBool{V: true},
	Block: &SDefine{
		Stmt: nil,
		Name: "foo",
		Typ:  types.I32,
		Expr: &EI32{V: 1},
	},
})

f.Blocks[len(f.Blocks)-1].NewRet(nil)

And output:

define void @foo() {
0:
	br label %do.while.body

do.while.body:
	%1 = alloca i32
	store i32 1, i32* %1
	br i1 true, label %do.while.body, label %leave.do.while

leave.do.while:
	ret void
}

For Loop

For-loop would be an interesting case, at here, I only present a for-loop that can only have one initialize variable to reduce complexity, therefore, we have a AST like this:

type SForLoop struct {
	Stmt
	InitName string
	InitExpr Expr
	Step     Expr
	Cond     Expr
	Block    Stmt
}

For example, for (x=0; x=x+1; x<10) {} break down to:

SForLoop {
	InitName: `x`
	InitExpr: `0`
	Step: `x + 1`
	Cond: `x < 10`
	Block: `{}`
}

At first view, people might think for-loop is as easy as do-while, but in SSA form, reuse variable in a loop need a new instruction: phi.

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SForLoop:
		loopCtx := ctx.NewContext(f.NewBlock("for.loop.body"))
		ctx.NewBr(loopCtx.Block)
		firstAppear := loopCtx.NewPhi(ir.NewIncoming(loopCtx.compileExpr(s.InitExpr), ctx.Block))
		loopCtx.vars[s.InitName] = firstAppear
		step := loopCtx.compileExpr(s.Step)
		firstAppear.Incs = append(firstAppear.Incs, ir.NewIncoming(step, loopCtx.Block))
		loopCtx.vars[s.InitName] = step
		leaveB := f.NewBlock("leave.for.loop")
		loopCtx.leaveBlock = leaveB
		loopCtx.compileStmt(s.Block)
		loopCtx.NewCondBr(loopCtx.compileExpr(s.Cond), loopCtx.Block, leaveB)
	}
}

    Create a loop body context
    jump from the previous block
    Put phi into loop body
    Phi would have two incoming, first is InitExpr, the second one is Step result.
    compile step
    compile the conditional branch, jump to loop body or leave block

Testing:

f := ir.NewFunc("foo", types.Void)
ctx := NewContext(f.NewBlock(""))

ctx.compileStmt(&SForLoop{
	InitName: "x",
	InitExpr: &EI32{V: 0},
	Step:     &EAdd{Lhs: &EVariable{Name: "x"}, Rhs: &EI32{V: 1}},
	Cond:     &ELessThan{Lhs: &EVariable{Name: "x"}, Rhs: &EI32{V: 10}},
	Block:    &SDefine{Name: "foo", Typ: types.I32, Expr: &EI32{V: 2}},
})

f.Blocks[len(f.Blocks)-1].NewRet(nil)

The test generates:

define void @foo() {
0:
	br label %for.loop.body

for.loop.body:
	%1 = phi i32 [ 0, %0 ], [ %2, %for.loop.body ]
	%2 = add i32 %1, 1
	%3 = alloca i32
	store i32 2, i32* %3
	%4 = icmp slt i32 %2, 10
	br i1 %4, label %for.loop.body, label %leave.for.loop

leave.for.loop:
	ret void
}

In fact, you can also avoid phi, you can make a try as practice.
While

The last kind of loop we want to present is while loop.

type SWhile struct {
	Stmt
	Cond  Expr
	Block Stmt
}

It looks just like do while, but have different semantic, it might not execute it's body. Here is the implementation.

func (ctx *Context) compileStmt(stmt Stmt) {
	switch s := stmt.(type) {
	case *SWhile:
		condCtx := ctx.NewContext(f.NewBlock("while.loop.cond"))
		ctx.NewBr(condCtx.Block)
		loopCtx := ctx.NewContext(f.NewBlock("while.loop.body"))
		leaveB := f.NewBlock("leave.do.while")
		condCtx.NewCondBr(condCtx.compileExpr(s.Cond), loopCtx.Block, leaveB)
		condCtx.leaveBlock = leaveB
		loopCtx.leaveBlock = leaveB
		loopCtx.compileStmt(s.Block)
		loopCtx.NewBr(condCtx.Block)
	}
}

We would need two blocks since br is a terminator, then the logic is simple:

    while.loop.cond would jump to while.loop.body or leave.do.while by condition
    while.loop.body always back to while.loop.cond.

Finally, test:

f := ir.NewFunc("foo", types.Void)
ctx := NewContext(f.NewBlock(""))

ctx.compileStmt(&SWhile{
	Cond: &EBool{V: true},
	Block: &SDefine{
		Name: "x",
		Typ:  types.I32,
		Expr: &EI32{V: 0},
	},
})

f.Blocks[len(f.Blocks)-1].NewRet(nil)

and output:

define void @foo() {
0:
	br label %while.loop.cond

while.loop.cond:
	br i1 true, label %while.loop.body, label %leave.do.while

while.loop.body:
	%1 = alloca i32
	store i32 0, i32* %1
	br label %while.loop.cond

leave.do.while:
	ret void
}
*/

/*
Linkage

The following code shows some linkage can use in IR.

m := ir.NewModule()

add := m.NewFunc("add", types.I64, ir.NewParam("", types.I64))
add.Linkage = enum.LinkageInternal
add1 := m.NewFunc("add1", types.I64, ir.NewParam("", types.I64))
add1.Linkage = enum.LinkageLinkOnce
add2 := m.NewFunc("add2", types.I64, ir.NewParam("", types.I64))
add2.Linkage = enum.LinkagePrivate
add3 := m.NewFunc("add3", types.I64, ir.NewParam("", types.I64))
add3.Linkage = enum.LinkageWeak
add4 := m.NewFunc("add4", types.I64, ir.NewParam("", types.I64))
add4.Linkage = enum.LinkageExternal

The code would produce:

declare internal i64 @add(i64)

declare linkonce i64 @add1(i64)

declare private i64 @add2(i64)

declare weak i64 @add3(i64)

declare external i64 @add4(i64)

For further information about linkage, refer to LLVM doc and pkg.go.dev.
*/

/*
Variant Argument (a.k.a. VAArg)

One example of a variadic function is printf. This is how to create a function prototype for printf:

m := ir.NewModule()

printf := m.NewFunc(
	"printf",
	types.I32,
	ir.NewParam("", types.NewPointer(types.I8)),
)
printf.Sig.Variadic = true

The above code would produce the following IR:

declare i32 @printf(i8*, ...)
*/

/*
Function Overloading

There is no overloading in LLVM IR. One solution is to create one function per function signature, where each LLVM IR function would have a unique name (this is why C++ compilers do name mangling).
*/

/*

First-class Function(Closure)
Naive Implementation

Create a closure(a.k.a. first-class function) requires a place to store captured variables. In LLVM, the best way is create a structure for such case:

m := ir.NewModule()

zero := constant.NewInt(types.I32, 0)
one := constant.NewInt(types.I32, 1)

captureStruct := m.NewTypeDef("id_capture", types.NewStruct(
	types.I32,
))
captureTyp := types.NewPointer(captureStruct)
idFn := m.NewFunc("id", types.I32, ir.NewParam("capture", captureTyp))
idB := idFn.NewBlock("")
v := idB.NewGetElementPtr(captureStruct, idFn.Params[0], zero, zero)
idB.NewRet(idB.NewLoad(types.I32, v))
idClosureTyp := m.NewTypeDef("id_closure", types.NewStruct(
	captureTyp,
	idFn.Type(),
))

mainFn := m.NewFunc("main", types.I32)
b := mainFn.NewBlock("")
// define a local variable `i`
i := b.NewAlloca(types.I32)
b.NewStore(constant.NewInt(types.I32, 10), i)
// use alloca at here to simplify code, in real case should be `malloc` or `gc_malloc`
captureInstance := b.NewAlloca(captureStruct)
ptrToCapture := b.NewGetElementPtr(captureStruct, captureInstance, zero, zero)
// capture variable
b.NewStore(b.NewLoad(types.I32, i), ptrToCapture)
// prepare closure
idClosure := b.NewAlloca(idClosureTyp)
ptrToCapturePtr := b.NewGetElementPtr(idClosureTyp, idClosure, zero, zero)
b.NewStore(captureInstance, ptrToCapturePtr)
ptrToFuncPtr := b.NewGetElementPtr(idClosureTyp, idClosure, zero, one)
b.NewStore(idFn, ptrToFuncPtr)
// assuming we transfer closure into another context
accessCapture := b.NewGetElementPtr(idClosureTyp, idClosure, zero, zero)
accessFunc := b.NewGetElementPtr(idClosureTyp, idClosure, zero, one)
result := b.NewCall(b.NewLoad(idFn.Type(), accessFunc), b.NewLoad(captureTyp, accessCapture))

printIntegerFormat := m.NewGlobalDef("tmp", irutil.NewCString("%d\n"))
pointerToString := b.NewGetElementPtr(types.NewArray(3, types.I8), printIntegerFormat, zero, zero)
// ignore printf
b.NewCall(printf, pointerToString, result)

b.NewRet(constant.NewInt(types.I32, 0))

This is a huge example, I understand it's hard to read, but concept is clean. It would generate below LLVM IR:

%id_capture = type { i32 }
%id_closure = type { %id_capture*, i32 (%id_capture*)* }

@tmp = global [3 x i8] c"%d\0A"

declare i32 @printf(i8* %format, ...)

define i32 @id(%id_capture* %capture) {
; <label>:0
	%1 = getelementptr %id_capture, %id_capture* %capture, i32 0, i32 0
	%2 = load i32, i32* %1
	ret i32 %2
}

define i32 @main() {
; <label>:0
	%1 = alloca i32
	store i32 10, i32* %1
	%2 = alloca %id_capture
	%3 = getelementptr %id_capture, %id_capture* %2, i32 0, i32 0
	%4 = load i32, i32* %1
	store i32 %4, i32* %3
	%5 = alloca %id_closure
	%6 = getelementptr %id_closure, %id_closure* %5, i32 0, i32 0
	store %id_capture* %2, %id_capture** %6
	%7 = getelementptr %id_closure, %id_closure* %5, i32 0, i32 1
	store i32 (%id_capture*)* @id, i32 (%id_capture*)** %7
	%8 = getelementptr %id_closure, %id_closure* %5, i32 0, i32 0
	%9 = getelementptr %id_closure, %id_closure* %5, i32 0, i32 1
	%10 = load i32 (%id_capture*)*, i32 (%id_capture*)** %9
	%11 = load %id_capture*, %id_capture** %8
	%12 = call i32 %10(%id_capture* %11)
	%13 = getelementptr [3 x i8], [3 x i8]* @tmp, i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i32 %12)
	ret i32 0
}

Our id function captures an Integer and return it. To reach that id_capture was introduced for storing captured value. For passing whole closure in convenience, id_closure was introduced and stored capture structure and function pointer. When invoke a closure, get captured structure and function pointer from id_closure structure, then apply function with captured structure and additional arguments(if there's any). In this example omit the part about memory management, all structures allocated in the stack, this won't work in most real world case. Must notice this problem.
Improvements

The naive implementation is not good enough, we have several ways can improve it, but instead of implementing them I'm going to list what can we do:

    Laziness function: Arity would be a thing in case
    Access cross asynchronous model
    If language has copy capture and reference capture, e.g. C++?
    What if working with a GC?

Return Structure

When meet program that return structure by value, compiler has chance to remove such cloning. That's storing return structure into a reference passed by the caller. Which means, if we get:

struct Foo {
    // ...
};

Foo foo() {
    Foo f;
    // ...
    return f;
}

should compile to:

define void @foo(%Foo* noalias sret f) {
    // ...
}

    sret hints this is a return value.
    noalias hints other arguments won't point to the same place, LLVM optimizer might rely on such fact, so don't add it everywhere.

Add parameter attributes

Here is example shows how to add parameter attributes:

m := ir.NewModule()

fooTyp := m.NewTypeDef("Foo", types.NewStruct(
	types.I32,
))
retS := ir.NewParam("result", fooTyp)
retS.Attrs = append(retS.Attrs, enum.ParamAttrNoAlias)
retS.Attrs = append(retS.Attrs, enum.ParamAttrSRet)
m.NewFunc("foo", types.Void, retS)

*/

/*
Exception

In this section, you will see how to reuse c++ exception in LLVM.

First, we need to setup a set of function from c++ ABI:

type ModuleWithException struct {
    *ir.Module
    _ZTIi                    *ir.Global
    __cxa_allocate_exception *ir.Func
    __cxa_throw              *ir.Func
    __cxa_begin_catch        *ir.Func
    __cxa_end_catch          *ir.Func
    __cxa_call_unexpected    *ir.Func
    llvm_eh_typeid_for       *ir.Func
}

func NewModuleWithException() *ModuleWithException {
    m := ir.NewModule()
    mWithE := &ModuleWithException{
        Module: m,
        _ZTIi:  m.NewGlobal("_ZTIi", TPtr(TI8)),
        __cxa_allocate_exception: m.NewFunc("__cxa_allocate_exception", TPtr(TI8),
            ir.NewParam("", TI64),
        ),
        __cxa_throw: m.NewFunc("__cxa_throw", TVoid,
            ir.NewParam("exception_header", TPtr(TI8)),
            ir.NewParam("", TPtr(TI8)),
            ir.NewParam("", TPtr(TI8)),
        ),
        __cxa_begin_catch:     m.NewFunc("__cxa_begin_catch", TPtr(TI8), ir.NewParam("", TPtr(TI8))),
        __cxa_end_catch:       m.NewFunc("__cxa_end_catch", TVoid),
        __cxa_call_unexpected: m.NewFunc("__cxa_call_unexpected", TVoid, ir.NewParam("", TPtr(TI8))),
        llvm_eh_typeid_for:    m.NewFunc("llvm.eh.typeid.for", TI32, ir.NewParam("", TPtr(TI8))),
    }
    mWithE._ZTIi.Linkage = enum.LinkageExternal
    return mWithE
}

And a helper for throw exception from a block:

func throwException(m *ModuleWithException, bb *ir.Block) {
    // C++ requires one allocate an exception first
    payload := bb.NewCall(m.__cxa_allocate_exception, CI64(4))
    // now we stores I32 `1` into payload
    bb.NewStore(CI32(1), bb.NewBitCast(payload, TPtr(TI32)))
    // finally, we call `__cxa_throw` to throw exception
    bb.NewCall(m.__cxa_throw,
        payload,
        constant.NewBitCast(m._ZTIi, TPtr(TI8)),
        constant.NewNull(TPtr(TI8)),
    )
}

Finally, is our full example

m := NewModuleWithException()

exceptionThrower := m.NewFunc("I throw exception!", TI32)
bb := exceptionThrower.NewBlock("")
throwException(m, bb)
bb.NewRet(CI32(1))

main := m.NewFunc("main", TI32)
main.Personality = constant.True
mainB := main.NewBlock("")
normalRetB := main.NewBlock("normalRet")
exceptionRetB := main.NewBlock("exceptionRet")
// we must use invoke when a function might throw exception, we need to give
// 1. normal return block for function returns normally
// 2. exception return block for function throws an exception
mainB.NewInvoke(exceptionThrower, []value.Value{}, normalRetB, exceptionRetB)
normalRetB.NewRet(CI32(0))
// landingpad stands for catch and cleanup
exc := exceptionRetB.NewLandingPad(types.NewStruct(TPtr(TI8), TI32),
    ir.NewClause(enum.ClauseTypeCatch, constant.NewBitCast(m._ZTIi, TPtr(TI8))),
    ir.NewClause(enum.ClauseTypeFilter,
        constant.NewArray(types.NewArray(1, TPtr(TI8)), constant.NewBitCast(m._ZTIi, TPtr(TI8))),
    ),
)
exc.Cleanup = true
exc_ptr := exceptionRetB.NewExtractValue(exc, 0)
exc_sel := exceptionRetB.NewExtractValue(exc, 1)
tid_int := exceptionRetB.NewCall(m.llvm_eh_typeid_for, constant.NewBitCast(m._ZTIi, TPtr(TI8)))
tid_int.Tail = enum.TailTail
catchintB := main.NewBlock("catchint")
cleanupB := main.NewBlock("cleanup")
cleanupB.NewResume(exc)
// we check typeinfo is as expected
// 1. if type info is same as our expection, we going to catchint block to handle threw I32
// 2. if not, we cleanup exception
tst_int := exceptionRetB.NewICmp(enum.IPredEQ, exc_sel, tid_int)
exceptionRetB.NewCondBr(tst_int, catchintB, cleanupB)
// in catchint block, we
// 1. call `__cxa_begin_catch` to begin catching
// 2. load payload to get threw I32
// 3. call `__cxa_end_catch` to end catching
payload := catchintB.NewCall(m.__cxa_begin_catch, exc_ptr)
payload.Tail = enum.TailTail
payload_int := catchintB.NewBitCast(payload, TPtr(TI32))
retval := catchintB.NewLoad(TI32, payload_int)
end_catch := catchintB.NewCall(m.__cxa_end_catch)
end_catch.Tail = enum.TailTail

// Finally, we return threw value from main
returnB := main.NewBlock("return")
catchintB.NewBr(returnB)
returnB.NewRet(retval)
*/

/*
High level types
Structure

Structure is quite common and basic type in programming language. Here focus on how to create an equal LLVM mapping for structure.

LLVM has the
