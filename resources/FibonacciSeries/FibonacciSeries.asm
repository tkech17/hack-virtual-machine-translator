@1
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
@R3
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D

@0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@THAT
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

@1
D=A
@SP
A=M
M=D
@SP
M=M+1

@1
D=A
@THAT
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

(FibonacciSeries.vm-_-MAIN_LOOP_START)

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

@SP
M=M-1
A=M
D=M
@FibonacciSeries.vm-_-COMPUTE_ELEMENT
D;JNE

@FibonacciSeries.vm-_-END_PROGRAM
0;JMP

(FibonacciSeries.vm-_-COMPUTE_ELEMENT)

@0
D=A
@THAT
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@1
D=A
@THAT
A=D+M
D=M
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
M=D+M

@2
D=A
@THAT
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

@1
D=A
@R3
A=D+A
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
M=D+M

@1
D=A
@R3
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D

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

@FibonacciSeries.vm-_-MAIN_LOOP_START
0;JMP

(FibonacciSeries.vm-_-END_PROGRAM)

