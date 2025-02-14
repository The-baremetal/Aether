	.text
	.file	"output.ll"
	.globl	main                            # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movl	$5, 20(%rsp)
	movl	$10, 16(%rsp)
	movl	$15, 12(%rsp)
	movq	fmt.str@GOTPCREL(%rip), %rdi
	movl	$15, %esi
	xorl	%eax, %eax
	callq	printf@PLT
	movl	12(%rsp), %esi
	movq	fmt.str.1@GOTPCREL(%rip), %rdi
	xorl	%eax, %eax
	callq	printf@PLT
	xorl	%eax, %eax
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc
                                        # -- End function
	.type	fmt.str,@object                 # @fmt.str
	.data
	.globl	fmt.str
fmt.str:
	.asciz	"%d\n"
	.size	fmt.str, 4

	.type	fmt.str.1,@object               # @fmt.str.1
	.globl	fmt.str.1
fmt.str.1:
	.asciz	"%d\n"
	.size	fmt.str.1, 4

	.section	".note.GNU-stack","",@progbits
