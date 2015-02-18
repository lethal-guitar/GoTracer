
// func Sqrtf(x float32) float32
TEXT ·Sqrtf(SB),4,$0
	SQRTSS x+0(FP), X0
	MOVSS X0, ret+8(FP)
	RET

// func Minf(a, b float32) float32
TEXT ·Minf(SB),4,$0
	MOVSS a+0(FP), X0
	MINSS b+4(FP), X0
	MOVSS X0, ret+8(FP)
	RET

// func Maxf(a, b float32) float32
TEXT ·Maxf(SB),4,$0
	MOVSS a+0(FP), X0
	MAXSS b+4(FP), X0
	MOVSS X0, ret+8(FP)
	RET
