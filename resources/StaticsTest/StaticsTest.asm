@256
D=A
@SP
M=D
@Class1.vm_-_return0
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
(Class1.vm_-_return0)

(Class1.set)

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

@Class1.vm.0
D=A
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
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@Class1.vm.1
D=A
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

(Class1.get)

@Class1.vm.0
D=M
@SP
A=M
M=D
@SP
M=M+1

@Class1.vm.1
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
M=M-D

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

(Class2.set)

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

@Class2.vm.0
D=A
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
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1

@Class2.vm.1
D=A
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

(Class2.get)

@Class2.vm.0
D=M
@SP
A=M
M=D
@SP
M=M+1

@Class2.vm.1
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
M=M-D

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

@6
D=A
@SP
A=M
M=D
@SP
M=M+1

@8
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
@7
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Class1.set
0;JMP
(Sys.init_-_return0)

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

@23
D=A
@SP
A=M
M=D
@SP
M=M+1

@15
D=A
@SP
A=M
M=D
@SP
M=M+1

@Sys.init_-_return1
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
@7
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Class2.set
0;JMP
(Sys.init_-_return1)

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

@Sys.init_-_return2
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
@Class1.get
0;JMP
(Sys.init_-_return2)

@Sys.init_-_return3
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
@Class2.get
0;JMP
(Sys.init_-_return3)

(Sys.init-_-WHILE)

@Sys.init-_-WHILE
0;JMP

