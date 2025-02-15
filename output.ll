@str.0 = global [9 x i8] c"if works\00"
@fmt.str.1 = global [4 x i8] c"%s\0A\00"
@str.2 = global [83 x i8] c"if this doesn't work, this confirms if is being called but its just a string issue\00"
@fmt.str.3 = global [4 x i8] c"%s\0A\00"
@str.4 = global [49 x i8] c"Welp if that didn't work atleast you have this!!\00"
@fmt.str.5 = global [4 x i8] c"%s\0A\00"
@fmt.str.6 = global [4 x i8] c"%d\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = add i32 3, 2
	%1 = mul i32 5, %0
	%x = alloca i32
	store i32 %1, i32* %x
	%2 = getelementptr [9 x i8], [9 x i8]* @str.0, i32 0, i32 0
	%3 = bitcast i8* %2 to i8*
	%y = alloca i8*
	store i8* %3, i8** %y
	%z = alloca i32
	store i32 3, i32* %z
	%4 = load i32, i32* %x
	%5 = load i32, i32* %z
	%6 = icmp sgt i32 %4, %5
	%7 = load i8*, i8** %y
	%8 = bitcast [4 x i8]* @fmt.str.1 to i8*
	%9 = call i32 (i8*, ...) @printf(i8* %8, i8* %7)
	%10 = getelementptr [83 x i8], [83 x i8]* @str.2, i32 0, i32 0
	%11 = bitcast i8* %10 to i8*
	%12 = bitcast [4 x i8]* @fmt.str.3 to i8*
	%13 = call i32 (i8*, ...) @printf(i8* %12, i8* %11)
	%14 = load i32, i32* %x
	%15 = load i32, i32* %x
	%16 = icmp eq i32 %14, %15
	%17 = getelementptr [49 x i8], [49 x i8]* @str.4, i32 0, i32 0
	%18 = bitcast i8* %17 to i8*
	%19 = bitcast [4 x i8]* @fmt.str.5 to i8*
	%20 = call i32 (i8*, ...) @printf(i8* %19, i8* %18)
	%21 = load i32, i32* %x
	%22 = bitcast [4 x i8]* @fmt.str.6 to i8*
	%23 = call i32 (i8*, ...) @printf(i8* %22, i32 %21)
	ret i32 0

if.then:
	br label %if.merge

if.else:
	br label %if.merge

if.merge:
	br i1 %16, label %if.then, label %if.else

if.then:
	br label %if.merge

if.else:
	br label %if.merge

if.merge:
	unreachable

unreachable_after_return:
	ret i32 0
}
