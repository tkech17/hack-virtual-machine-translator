(Main.fibonacci)

@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@2
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
M=-1
@FOR_JUMP_Main.fibonacci0
D;JLT
@SP
A=M-1
M=0
(FOR_JUMP_Main.fibonacci0)

@SP
M=M-1
A=M
D=M
@Main.fibonacci-_-IF_TRUE
D;JNE

@Main.fibonacci-_-IF_FALSE
0;JMP

(Main.fibonacci-_-IF_TRUE)

@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@LCL
D=M
@R14
M=D
@R14
D=M
D=M
@5
A=D-A
D=M
@R15
M=D

@0
D=A
@ARG
D=D+M
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D

@ARG
D=M
@1
D=D+A
@SP
M=D
@R14
D=M
D=M
@1
A=D-A
D=M
@THAT
M=D

@R14
D=M
D=M
@2
A=D-A
D=M
@THIS
M=D

@R14
D=M
D=M
@3
A=D-A
D=M
@ARG
M=D

@R14
D=M
D=M
@4
A=D-A
D=M
@LCL
M=D

@R15
A=M
0;JMP

(Main.fibonacci-_-IF_FALSE)

@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@2
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
A=A-1
M=M-D

@Main.fibonacci_-_return0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R2
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R3
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R4
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@SP
D=M
@6
D=D-A
@ARG
M=D
@SP
D=M
@0
D=D-A
@LCL
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci_-_return0)

@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@1
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
M=M-1
A=M
D=M
A=A-1
M=M-D

@Main.fibonacci_-_return1
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R2
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R3
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@R4
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

@SP
D=M
@6
D=D-A
@ARG
M=D
@SP
D=M
@0
D=D-A
@LCL
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci_-_return1)

@SP
M=M-1
A=M
D=M
A=A-1
M=D+M

@LCL
D=M
@R14
M=D
@R14
D=M
D=M
@5
A=D-A
D=M
@R15
M=D

@0
D=A
@ARG
D=D+M
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D

@ARG
D=M
@1
D=D+A
@SP
M=D
@R14
D=M
D=M
@1
A=D-A
D=M
@THAT
M=D

@R14
D=M
D=M
@2
A=D-A
D=M
@THIS
M=D

@R14
D=M
D=M
@3
A=D-A
D=M
@ARG
M=D

@R14
D=M
D=M
@4
A=D-A
D=M
@LCL
M=D

@R15
A=M
0;JMP

