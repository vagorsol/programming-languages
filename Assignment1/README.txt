NAME:
=====
	Audrey Yang

Programs Files:
===============
    - flightStructs.go
	
How to Compile/Run:
===============
    - go run flightStructs.go
     
Questions:
===========
    1. In Go, the copy and original are equal. In Java, they are not. In both, the modified copy and the original are not equal. In Java, when using the "==" operand, it checks what
        address each object is being used, and when a copy is made, the copy has it's own distinct address. In Go, until modifications are made to the copy (or original), they
        use the same address, so the equals operand returns true. 
    2. The item at position 1 of the original slice and the slice with the first 5 items are equal. When I change the value of a field
        in the original slice, that value is also changed in the subslice, so the two remained equal. This seems to fit with how, in the previous question,
        I came to the conclusion that the copy just referred to the same address as the original until a change to the copy was made. Considering how, the first time
        I wrote code to help me figure out this question, I changed the copy instead of the original and got the same result, it may be that the original slice and the subslice
        both point to the same address, so changing the value in one changes the value in the other.  
	
Reflection:
===========
    When writing, I felt rather frustrated because of my unfamiliarity with the language, feeling as though I was doing things incorrectly/inefficiently, and doubting my ability
    the entire time (imposter syndrome, yay). Aside from that, the assignment felt relatively straightforward - although with moments of frustration with structures and how to use them.
    Answering the questions was especially difficult for me, as knowing why programs work like they work is a weak point for me. Trying to find answers, or ones that would point me in the right
    direction was virtually impossible, so I did my best. 
    Overall, technically not very difficult, but personally very frustrating. 

I Worked With:
==============
    N/A

Approximate Hours worked:
=========================
    ~6 hours

Special Instructions to the grader:
===================================
    N/A

Known Bugs or Limitations:
==========================
    N/A

Other comments:
===============
    After I first finished Part 1, I wrote a note to myself that said "modified D and C are still equal. they would not be in java". When
    I tested this later, it was not the case. I don't know what happened before that made me write that conclusion, which made answering the questions
    difficult, and in general I am confused.