@fmt.str = global [4 x i8] c"%d\0A\00"
@fmt.str.1 = global [4 x i8] c"%d\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%a = alloca i32
	store i32 5, i32* %a
	%b = alloca i32
	store i32 10, i32* %b
	%c = alloca i32
	%0 = load i32, i32* %a
	store i32 %0, i32* %c
	%1 = load i32, i32* %c
	%2 = bitcast [4 x i8]* @fmt.str to i8*
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 %1)
	%4 = load i32, i32* %c
	%5 = bitcast [4 x i8]* @fmt.str.1 to i8*
	%6 = call i32 (i8*, ...) @printf(i8* %5, i32 %4)
	ret i32 0
}
