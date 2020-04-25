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
@0
D=D-A
@LCL
M=D
@Main.fibonacci
0;JMP
(Sys.init_-_return0)

(Sys.init-_-WHILE)

@Sys.init-_-WHILE
0;JMP

