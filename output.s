	.text
	.file	"output.ll"
	.globl	main                            # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	movb	$1, %al
	.p2align	4, 0x90
.LBB0_1:                                # %while.cond_4
                                        # =>This Inner Loop Header: Depth=1
	testb	%al, %al
	jne	.LBB0_1
# %bb.2:                                # %while.end_6
	xorl	%eax, %eax
	.p2align	4, 0x90
.LBB0_3:                                # %repeat.loop_7
                                        # =>This Inner Loop Header: Depth=1
	testb	%al, %al
	jne	.LBB0_3
# %bb.4:                                # %repeat.end_9
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc
                                        # -- End function
	.type	str_0,@object                   # @str_0
	.section	.rodata,"a",@progbits
	.globl	str_0
str_0:
	.ascii	"hello"
	.size	str_0, 5

	.type	str_1,@object                   # @str_1
	.globl	str_1
str_1:
	.ascii	"world"
	.size	str_1, 5

	.type	str_2,@object                   # @str_2
	.globl	str_2
str_2:
	.ascii	"Overflow"
	.size	str_2, 8

	.type	str_3,@object                   # @str_3
	.globl	str_3
str_3:
	.ascii	"Arthur"
	.size	str_3, 6

	.type	str_4,@object                   # @str_4
	.globl	str_4
str_4:
	.ascii	"sword"
	.size	str_4, 5

	.type	str_5,@object                   # @str_5
	.globl	str_5
str_5:
	.ascii	"shield"
	.size	str_5, 6

	.type	str_6,@object                   # @str_6
	.globl	str_6
str_6:
	.ascii	"N/A"
	.size	str_6, 3

	.type	str_7,@object                   # @str_7
	.globl	str_7
str_7:
	.ascii	"Error occurred:"
	.size	str_7, 15

	.section	".note.GNU-stack","",@progbits
