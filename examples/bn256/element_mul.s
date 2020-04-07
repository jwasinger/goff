// Code generated by goff (v0.1.0) DO NOT EDIT



#include "textflag.h"

// func mulAsmElement(res,y *Element)
// montgomery multiplication of res by y 
// stores the result in res
TEXT ·mulAsmElement(SB), NOSPLIT, $0-16

	// dereference our parameters
	MOVQ res+0(FP), DI
	MOVQ y+8(FP), R8

	// check if we support adx and mulx
	CMPB ·support_adx_Element(SB), $1
	JNE no_adx
	 
	// the algorithm is described here
	// https://hackmd.io/@zkteam/modular_multiplication
	// however, to benefit from the ADCX and ADOX carry chains
	// we split the inner loops in 2:
	// for i=0 to N-1
    // 		for j=0 to N-1
    // 		    (A,t[j])  := t[j] + a[j]*b[i] + A
    // 		m := t[0]*q'[0] mod W
    // 		C,_ := t[0] + m*q[0]
    // 		for j=1 to N-1
    // 		    (C,t[j-1]) := t[j] + m*q[j] + C
    // 		t[N-1] = C + A
	// clear up the carry flags
	XORQ R9 , R9

	// y[0] in R12
	MOVQ 0(R8), R12

	// for j=0 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
		// res[0] in DX
		MOVQ 0(DI), DX
				MULXQ R12, CX ,  R9
		// res[1] in DX
		MOVQ 8(DI), DX
				MOVQ R9, BX
			MULXQ R12, AX,  R9
			ADOXQ AX, BX
		// res[2] in DX
		MOVQ 16(DI), DX
				MOVQ R9, BP
			MULXQ R12, AX,  R9
			ADOXQ AX, BP
		// res[3] in DX
		MOVQ 24(DI), DX
				MOVQ R9, SI
			MULXQ R12, AX,  R9
			ADOXQ AX, SI
	// add the last carries to R9 
	MOVQ $0, DX
	ADCXQ DX, R9 
	ADOXQ DX, R9 
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, DX
	MULXQ CX,R11, DX

	// clear the carry flags
	XORQ DX, DX 

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, DX
	MULXQ R11, AX, R10
	ADCXQ CX ,AX

	// for j=1 to N-1
    //    (C,t[j-1]) := t[j] + m*q[j] + C
		
		MOVQ $0x97816a916871ca8d, DX
		MULXQ R11, AX, DX
		ADCXQ  BX, R10 
		ADOXQ AX, R10
		MOVQ R10, CX
			MOVQ DX, R10
		
		MOVQ $0xb85045b68181585d, DX
		MULXQ R11, AX, DX
		ADCXQ  BP, R10 
		ADOXQ AX, R10
		MOVQ R10, BX
			MOVQ DX, R10
		
		MOVQ $0x30644e72e131a029, DX
		MULXQ R11, AX, DX
		ADCXQ  SI, R10 
		ADOXQ AX, R10
		MOVQ R10, BP
			MOVQ $0, AX
			ADCXQ AX, DX
			ADOXQ DX, R9
			MOVQ R9, SI
	// clear up the carry flags
	XORQ R9 , R9

	// y[1] in R12
	MOVQ 8(R8), R12

	// for j=0 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
		// res[0] in DX
		MOVQ 0(DI), DX
				MULXQ R12, AX,  R9
				ADOXQ AX, CX
		// res[1] in DX
		MOVQ 8(DI), DX
				ADCXQ R9, BX
			MULXQ R12, AX,  R9
			ADOXQ AX, BX
		// res[2] in DX
		MOVQ 16(DI), DX
				ADCXQ R9, BP
			MULXQ R12, AX,  R9
			ADOXQ AX, BP
		// res[3] in DX
		MOVQ 24(DI), DX
				ADCXQ R9, SI
			MULXQ R12, AX,  R9
			ADOXQ AX, SI
	// add the last carries to R9 
	MOVQ $0, DX
	ADCXQ DX, R9 
	ADOXQ DX, R9 
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, DX
	MULXQ CX,R11, DX

	// clear the carry flags
	XORQ DX, DX 

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, DX
	MULXQ R11, AX, R10
	ADCXQ CX ,AX

	// for j=1 to N-1
    //    (C,t[j-1]) := t[j] + m*q[j] + C
		
		MOVQ $0x97816a916871ca8d, DX
		MULXQ R11, AX, DX
		ADCXQ  BX, R10 
		ADOXQ AX, R10
		MOVQ R10, CX
			MOVQ DX, R10
		
		MOVQ $0xb85045b68181585d, DX
		MULXQ R11, AX, DX
		ADCXQ  BP, R10 
		ADOXQ AX, R10
		MOVQ R10, BX
			MOVQ DX, R10
		
		MOVQ $0x30644e72e131a029, DX
		MULXQ R11, AX, DX
		ADCXQ  SI, R10 
		ADOXQ AX, R10
		MOVQ R10, BP
			MOVQ $0, AX
			ADCXQ AX, DX
			ADOXQ DX, R9
			MOVQ R9, SI
	// clear up the carry flags
	XORQ R9 , R9

	// y[2] in R12
	MOVQ 16(R8), R12

	// for j=0 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
		// res[0] in DX
		MOVQ 0(DI), DX
				MULXQ R12, AX,  R9
				ADOXQ AX, CX
		// res[1] in DX
		MOVQ 8(DI), DX
				ADCXQ R9, BX
			MULXQ R12, AX,  R9
			ADOXQ AX, BX
		// res[2] in DX
		MOVQ 16(DI), DX
				ADCXQ R9, BP
			MULXQ R12, AX,  R9
			ADOXQ AX, BP
		// res[3] in DX
		MOVQ 24(DI), DX
				ADCXQ R9, SI
			MULXQ R12, AX,  R9
			ADOXQ AX, SI
	// add the last carries to R9 
	MOVQ $0, DX
	ADCXQ DX, R9 
	ADOXQ DX, R9 
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, DX
	MULXQ CX,R11, DX

	// clear the carry flags
	XORQ DX, DX 

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, DX
	MULXQ R11, AX, R10
	ADCXQ CX ,AX

	// for j=1 to N-1
    //    (C,t[j-1]) := t[j] + m*q[j] + C
		
		MOVQ $0x97816a916871ca8d, DX
		MULXQ R11, AX, DX
		ADCXQ  BX, R10 
		ADOXQ AX, R10
		MOVQ R10, CX
			MOVQ DX, R10
		
		MOVQ $0xb85045b68181585d, DX
		MULXQ R11, AX, DX
		ADCXQ  BP, R10 
		ADOXQ AX, R10
		MOVQ R10, BX
			MOVQ DX, R10
		
		MOVQ $0x30644e72e131a029, DX
		MULXQ R11, AX, DX
		ADCXQ  SI, R10 
		ADOXQ AX, R10
		MOVQ R10, BP
			MOVQ $0, AX
			ADCXQ AX, DX
			ADOXQ DX, R9
			MOVQ R9, SI
	// clear up the carry flags
	XORQ R9 , R9

	// y[3] in R12
	MOVQ 24(R8), R12

	// for j=0 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
		// res[0] in DX
		MOVQ 0(DI), DX
				MULXQ R12, AX,  R9
				ADOXQ AX, CX
		// res[1] in DX
		MOVQ 8(DI), DX
				ADCXQ R9, BX
			MULXQ R12, AX,  R9
			ADOXQ AX, BX
		// res[2] in DX
		MOVQ 16(DI), DX
				ADCXQ R9, BP
			MULXQ R12, AX,  R9
			ADOXQ AX, BP
		// res[3] in DX
		MOVQ 24(DI), DX
				ADCXQ R9, SI
			MULXQ R12, AX,  R9
			ADOXQ AX, SI
	// add the last carries to R9 
	MOVQ $0, DX
	ADCXQ DX, R9 
	ADOXQ DX, R9 
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, DX
	MULXQ CX,R11, DX

	// clear the carry flags
	XORQ DX, DX 

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, DX
	MULXQ R11, AX, R10
	ADCXQ CX ,AX

	// for j=1 to N-1
    //    (C,t[j-1]) := t[j] + m*q[j] + C
		
		MOVQ $0x97816a916871ca8d, DX
		MULXQ R11, AX, DX
		ADCXQ  BX, R10 
		ADOXQ AX, R10
		MOVQ R10, CX
			MOVQ DX, R10
		
		MOVQ $0xb85045b68181585d, DX
		MULXQ R11, AX, DX
		ADCXQ  BP, R10 
		ADOXQ AX, R10
		MOVQ R10, BX
			MOVQ DX, R10
		
		MOVQ $0x30644e72e131a029, DX
		MULXQ R11, AX, DX
		ADCXQ  SI, R10 
		ADOXQ AX, R10
		MOVQ R10, BP
			MOVQ $0, AX
			ADCXQ AX, DX
			ADOXQ DX, R9
			MOVQ R9, SI

reduce:
	// reduce, constant time version
	// first we copy registers storing t in a separate set of registers
	// as SUBQ modifies the 2nd operand
	MOVQ CX, DX
	MOVQ BX, R8
	MOVQ BP, R9
	MOVQ SI, R10
	MOVQ $0x3c208c16d87cfd47, R11
	SUBQ  R11, DX
	MOVQ $0x97816a916871ca8d, R11
	SBBQ  R11, R8
	MOVQ $0xb85045b68181585d, R11
	SBBQ  R11, R9
	MOVQ $0x30644e72e131a029, R11
	SBBQ  R11, R10
	JCS t_is_smaller // no borrow, we return t
	MOVQ DX, (DI)
	MOVQ R8, 8(DI)
	MOVQ R9, 16(DI)
	MOVQ R10, 24(DI)
	RET
t_is_smaller:
	MOVQ CX, 0(DI)
	MOVQ BX, 8(DI)
	MOVQ BP, 16(DI)
	MOVQ SI, 24(DI)
	RET

no_adx:
	// (A,t[0]) := t[0] + x[0]*y[0]
	MOVQ (DI), AX // x[0]
	MOVQ 0(R8), R12
	MULQ R12 // x[0] * y[0]	
	MOVQ DX, R9
	MOVQ AX, CX
	
	
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, R11
	IMULQ CX , R11

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, AX
	MULQ R11
	ADDQ CX ,AX
	ADCQ $0, DX
	MOVQ  DX, R10

	// for j=1 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
    //    (C,t[j-1]) := t[j] + m*q[j] + C
	MOVQ 8(DI), AX
	MULQ R12 // x[1] * y[0]
	MOVQ R9, BX
	ADDQ AX, BX
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x97816a916871ca8d, AX
	MULQ R11
	ADDQ  BX, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, CX
	MOVQ DX, R10
	MOVQ 16(DI), AX
	MULQ R12 // x[2] * y[0]
	MOVQ R9, BP
	ADDQ AX, BP
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0xb85045b68181585d, AX
	MULQ R11
	ADDQ  BP, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BX
	MOVQ DX, R10
	MOVQ 24(DI), AX
	MULQ R12 // x[3] * y[0]
	MOVQ R9, SI
	ADDQ AX, SI
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x30644e72e131a029, AX
	MULQ R11
	ADDQ  SI, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BP
	MOVQ DX, R10

	ADDQ R10, R9
	MOVQ R9, SI
	// (A,t[0]) := t[0] + x[0]*y[1]
	MOVQ (DI), AX // x[0]
	MOVQ 8(R8), R12
	MULQ R12 // x[0] * y[1]
	ADDQ AX, CX 
	ADCQ $0, DX	
	MOVQ DX, R9
	
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, R11
	IMULQ CX , R11

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, AX
	MULQ R11
	ADDQ CX ,AX
	ADCQ $0, DX
	MOVQ  DX, R10

	// for j=1 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
    //    (C,t[j-1]) := t[j] + m*q[j] + C
	MOVQ 8(DI), AX
	MULQ R12 // x[1] * y[1]
	ADDQ R9, BX
	ADCQ $0, DX
	ADDQ AX, BX
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x97816a916871ca8d, AX
	MULQ R11
	ADDQ  BX, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, CX
	MOVQ DX, R10
	MOVQ 16(DI), AX
	MULQ R12 // x[2] * y[1]
	ADDQ R9, BP
	ADCQ $0, DX
	ADDQ AX, BP
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0xb85045b68181585d, AX
	MULQ R11
	ADDQ  BP, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BX
	MOVQ DX, R10
	MOVQ 24(DI), AX
	MULQ R12 // x[3] * y[1]
	ADDQ R9, SI
	ADCQ $0, DX
	ADDQ AX, SI
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x30644e72e131a029, AX
	MULQ R11
	ADDQ  SI, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BP
	MOVQ DX, R10

	ADDQ R10, R9
	MOVQ R9, SI
	// (A,t[0]) := t[0] + x[0]*y[2]
	MOVQ (DI), AX // x[0]
	MOVQ 16(R8), R12
	MULQ R12 // x[0] * y[2]
	ADDQ AX, CX 
	ADCQ $0, DX	
	MOVQ DX, R9
	
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, R11
	IMULQ CX , R11

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, AX
	MULQ R11
	ADDQ CX ,AX
	ADCQ $0, DX
	MOVQ  DX, R10

	// for j=1 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
    //    (C,t[j-1]) := t[j] + m*q[j] + C
	MOVQ 8(DI), AX
	MULQ R12 // x[1] * y[2]
	ADDQ R9, BX
	ADCQ $0, DX
	ADDQ AX, BX
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x97816a916871ca8d, AX
	MULQ R11
	ADDQ  BX, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, CX
	MOVQ DX, R10
	MOVQ 16(DI), AX
	MULQ R12 // x[2] * y[2]
	ADDQ R9, BP
	ADCQ $0, DX
	ADDQ AX, BP
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0xb85045b68181585d, AX
	MULQ R11
	ADDQ  BP, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BX
	MOVQ DX, R10
	MOVQ 24(DI), AX
	MULQ R12 // x[3] * y[2]
	ADDQ R9, SI
	ADCQ $0, DX
	ADDQ AX, SI
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x30644e72e131a029, AX
	MULQ R11
	ADDQ  SI, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BP
	MOVQ DX, R10

	ADDQ R10, R9
	MOVQ R9, SI
	// (A,t[0]) := t[0] + x[0]*y[3]
	MOVQ (DI), AX // x[0]
	MOVQ 24(R8), R12
	MULQ R12 // x[0] * y[3]
	ADDQ AX, CX 
	ADCQ $0, DX	
	MOVQ DX, R9
	
	
	// m := t[0]*q'[0] mod W
	MOVQ $0x87d20782e4866389, R11
	IMULQ CX , R11

	// C,_ := t[0] + m*q[0]
	MOVQ $0x3c208c16d87cfd47, AX
	MULQ R11
	ADDQ CX ,AX
	ADCQ $0, DX
	MOVQ  DX, R10

	// for j=1 to N-1
	//    (A,t[j])  := t[j] + x[j]*y[i] + A
    //    (C,t[j-1]) := t[j] + m*q[j] + C
	MOVQ 8(DI), AX
	MULQ R12 // x[1] * y[3]
	ADDQ R9, BX
	ADCQ $0, DX
	ADDQ AX, BX
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x97816a916871ca8d, AX
	MULQ R11
	ADDQ  BX, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, CX
	MOVQ DX, R10
	MOVQ 16(DI), AX
	MULQ R12 // x[2] * y[3]
	ADDQ R9, BP
	ADCQ $0, DX
	ADDQ AX, BP
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0xb85045b68181585d, AX
	MULQ R11
	ADDQ  BP, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BX
	MOVQ DX, R10
	MOVQ 24(DI), AX
	MULQ R12 // x[3] * y[3]
	ADDQ R9, SI
	ADCQ $0, DX
	ADDQ AX, SI
	ADCQ $0, DX
	MOVQ DX, R9

	MOVQ $0x30644e72e131a029, AX
	MULQ R11
	ADDQ  SI, R10
	ADCQ $0, DX
	ADDQ AX, R10
	ADCQ $0, DX
	
	MOVQ R10, BP
	MOVQ DX, R10

	ADDQ R10, R9
	MOVQ R9, SI

	JMP reduce
