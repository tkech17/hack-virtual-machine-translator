@256
D=A
@SP
M=D
@Main.vm_-_return0
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
@5
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Sys.init
0;JMP
(Main.vm_-_return0)

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
@FOR_JUMP_Main.vm0
D;JLT
@SP
A=M-1
M=0
(FOR_JUMP_Main.vm0)

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
@1
A=D-A
D=M
@THAT
M=D

@R14
D=M
@2
A=D-A
D=M
@THIS
M=D

@R14
D=M
@3
A=D-A
D=M
@ARG
M=D

@R14
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
@LCL
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci_-_return1)

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

@Main.fibonacci_-_return2
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
@LCL
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci_-_return2)

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
@1
A=D-A
D=M
@THAT
M=D

@R14
D=M
@2
A=D-A
D=M
@THIS
M=D

@R14
D=M
@3
A=D-A
D=M
@ARG
M=D

@R14
D=M
@4
A=D-A
D=M
@LCL
M=D

@R15
A=M
0;JMP

(Sys.init)

@4
D=A
@SP
A=M
M=D
@SP
M=M+1

@Sys.init_-_return0
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
@LCL
M=D
@Main.fibonacci
0;JMP
(Sys.init_-_return0)

(Sys.init-_-WHILE)

@Sys.init-_-WHILE
0;JMP

