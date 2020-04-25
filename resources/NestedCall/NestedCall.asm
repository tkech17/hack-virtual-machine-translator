(Sys.init)

@4000
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
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

@5000
D=A
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
@5
D=D-A
@ARG
M=D
@SP
D=M
@0
D=D-A
@LCL
M=D
@Sys.main
0;JMP
(Sys.init_-_return0)

@1
D=A
@R5
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

(Sys.init-_-LOOP)

@Sys.init-_-LOOP
0;JMP

(Sys.main)
@0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
D=A
@SP
A=M
M=D
@SP
M=M+1


@4001
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
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

@5001
D=A
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

@200
D=A
@SP
A=M
M=D
@SP
M=M+1

@1
D=A
@LCL
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

@40
D=A
@SP
A=M
M=D
@SP
M=M+1

@2
D=A
@LCL
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

@6
D=A
@SP
A=M
M=D
@SP
M=M+1

@3
D=A
@LCL
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

@123
D=A
@SP
A=M
M=D
@SP
M=M+1

@Sys.main_-_return1
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
@Sys.add12
0;JMP
(Sys.main_-_return1)

@0
D=A
@R5
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
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@1
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@2
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@3
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@4
D=A
@LCL
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

@SP
M=M-1
A=M
D=M
A=A-1
M=D+M

@SP
M=M-1
A=M
D=M
A=A-1
M=D+M

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

(Sys.add12)

@4002
D=A
@SP
A=M
M=D
@SP
M=M+1

@0
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

@5002
D=A
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
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@12
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

