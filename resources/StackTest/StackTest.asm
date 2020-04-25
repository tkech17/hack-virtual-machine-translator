@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
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
@FOR_JUMP_StackTest.vm0
D;JEQ
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm0)

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@16
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
@FOR_JUMP_StackTest.vm1
D;JEQ
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm1)

@16
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
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
@FOR_JUMP_StackTest.vm2
D;JEQ
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm2)

@892
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
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
@FOR_JUMP_StackTest.vm3
D;JLT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm3)

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@892
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
@FOR_JUMP_StackTest.vm4
D;JLT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm4)

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
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
@FOR_JUMP_StackTest.vm5
D;JLT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm5)

@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
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
@FOR_JUMP_StackTest.vm6
D;JGT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm6)

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32767
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
@FOR_JUMP_StackTest.vm7
D;JGT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm7)

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
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
@FOR_JUMP_StackTest.vm8
D;JGT
@SP
A=M-1
M=0
(FOR_JUMP_StackTest.vm8)

@57
D=A
@SP
A=M
M=D
@SP
M=M+1

@31
D=A
@SP
A=M
M=D
@SP
M=M+1

@53
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

@112
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

@SP
A=M-1
M=-M

@SP
M=M-1
A=M
D=M
A=A-1
M=D&M

@82
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
M=M|D

@SP
A=M-1
M=!M

