declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
        %0 = call i32 @require()
        %1 = alloca i32
        %2 = load i32, i32* %1
        %3 = load i32, i32* %1
        %4 = load i32, i32* %1
        %5 = alloca i32
        %6 = load i32, i32* %1
        %7 = load i32, i32* %1
        %8 = load i32, i32* %5
        %9 = load i32, i32* %5
        %10 = load i32, i32* %1
        %11 = load i32, i32* %1
        %12 = load i32, i32* %1
        %13 = load i32, i32* %5
        %14 = load i32, i32* %5
        %15 = load i32, i32* %5
        %16 = load i32, i32* %1
        %17 = load i32, i32* %1
        %18 = load i32, i32* %1
        %19 = load i32, i32* %5
        %20 = load i32, i32* %5
        %21 = load i32, i32* %5
        %22 = call i32 @cool()
        %23 = load i32, i32* %1
        %24 = load i32, i32* %1
        %25 = load i32, i32* %1
        %26 = load i32, i32* %5
        %27 = load i32, i32* %5
        %28 = load i32, i32* %5
        %29 = call i32 @gaae()
        ret i32 0
}

declare i32 @require()

declare i32 @cool()

declare i32 @gaae()