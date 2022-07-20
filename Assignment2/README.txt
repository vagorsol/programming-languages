NAME:
=====
	Audrey Yang

Programs Files:
===============
    Part 1:
        - goScope.go
    Part 2:
        - stackSpacing.go
    Part 3: 
        - incrementing.go
        - incrScript.txt
	
How to Compile/Run:
===============
    Part 1:
        - go run goScope.go
    Part 2:
        - go run stackSpacing.go
    Part 3:
        - go run incrementing.go

Part 1 - Scope Description:
==========================
    In static scoping, the returned value of a variable is determined by the value it is assigned in the scope. In dynamic scope,
    it is determined by the most recent assignment of the variable in the stack. The program "goScope" demonstrates how, the two instances of 
    variables named "b" do not affect each other because those values are limited to their respective scopes: because "b" is 7 so long as it is in
    the inner scope, and 5 when "b" is assigned that in the outer scope, it is statically scoped.
    Were it dynamically scoped, the program would take the most recent assignment of the variable named "b" and print out "7" both times. Since it didn't,
    it is not dynamically scoped. 

Part 2 - Questions:
====================
    1. There is roughly 0.5 GB of stack space avaliable by default on my machine. Go allocates 1 GB stack space, and whenever it reaches its stack limit, it tries to double the 
        the space so more is avaliable. Because doubling 0.5 GB will be 1 GB, when it tries to double it hits the limit and throws the runtime error. 
        This value is consistent throughout, as each recursive, when the size taken is calculated (found by multiplying how much space the call takes by how many itterations the function got throughout
        before the error) they are all roughly 0.5 GB.
    2. Because arrays have fixed size in Java, the space the array takes up will be based on how large the arrays are defined to be. This is similar to how arrays are in Go (assuming 
        the question is refering to arrays and not the slice array) so, most likely, the trick of "take up a lot of space creating arrays in a function" will work in Java too.

Reflection:
===========
    When writing the code for Parts 1 and 2, I felt like I was just copying the code shown in class, since my functions were based on those ones, and that I wasn't doing enough to demonstrate
    what was asked for. Writing the code itself went smoothly enough. The instructions at times were unclear to me - particularly, calculating stack space was frustrating and stressful because 
    I did not feel like I knew where to start or what I was doing, as well as what show the unboundedness of the incremeter meant. The real non-"programming languages lesson" takeaway of this assignment
    is, once again, that I need to work on my project managment skills so I don't feel stressed over time, which doesn't help the stress from trying to tackle problems that are challenging for me. 

I Worked With:
==============
    N/A

Approximate Hours worked:
=========================
    ~7-8 hours

Special Instructions to the grader:
===================================
    N/A

Known Bugs or Limitations:
==========================
    - The second time I ran int16 it didn't end in a runtime error.

Other comments:
===============
    - I think my program names are funny.