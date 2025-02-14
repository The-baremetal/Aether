@fmt.str = global [4 x i8] c"%d\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%a = alloca i32
	store i32 5, i32* %a
	%b = alloca i32
	store i32 100000, i32* %b
	%c = alloca i32
	%0 = load i32, i32* %a
	%1 = load i32, i32* %b
	%2 = add i32 %0, %1
	store i32 %2, i32* %c
	%3 = load i32, i32* %c
	%4 = bitcast [4 x i8]* @fmt.str to i8*
	%5 = call i32 (i8*, ...) @printf(i8* %4, i32 %3)
	ret i32 0
}
