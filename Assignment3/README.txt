NAME:
=====
	Audrey Yang

Programs Files:
===============
    Part 1:
        - RecursiveHW.kt
        - TailRecursiveHW.kt
    Part 2:
        - NestedFunctions.kt
    Part 3:
        - NthPower.kt
    Part 4: 
        - Filtering.kt
        - GenericFiltering.kt


How to Compile and Run:
=======================
    Part 1:
       - && kotlinc RecursiveHW.kt -include-runtime -d RecursiveHW.jar && java -jar RecursiveHW.jar
       - && kotlinc RecursiveHW.kt -include-runtime -d TailRecursiveHW.jar && java -jar TailRecursiveHW.jar
    Part 2:
       - && kotlinc NestedFunctions.kt -include-runtime -d NestedFunctions.jar && java -jar NestedFunctions.jar
    Part 3:
        - && kotlinc NthPower.kt -include-runtime -d NthPower.jar && java -jar NthPower.jar
    Part 4:
        - && kotlinc Filtering.kt -include-runtime -d Filtering.jar && java -jar Filtering.jar
        - && kotlinc GenericFiltering.kt -include-runtime -d GenericFiltering.jar && java -jar GenericFiltering.jar
       
Questions:
==========
Part 1:
    What is the largest number of time you can print "hello world" without getting a "Stack Overflow" exception?
        - 6388, 6520, 6568, 6373, 6428
        You can print "hello world" roughly 6450 times before you get a "Stack Overflow exception.
        - Additional observations:
            - When I had entered "10000000000" to test the limit of RecursiveHW, the program had a "NumberFormatException" error. 
                I assume it's because I had exceeded the maximum integer value in Kotlin with that.
            - There were many "at RecurseHWKt.recurseHW(RecursiveHW.kt:43) messages after the Stackoverflow error. 
    Does adding "tailrec" change the "stack overflow" limit? What is the limit after adding "tailrec"? 
        I stopped TailRecursiveHW after roughly 20 minutes, possibly more, because the compiler started to struggle, and also because I was getting impatient 
        and felt like the point was being made already. It reached 344152234 before I stopped it.
        This seems to suggest that adding "tailrec" removes the limitation on how much stack the recursive function can use, and thus, virtually no "stack overflow" limit.
        - Fun outputs! "llo, World! 247262639", "Hello, World! 247262649 [\n] Hello, World! 24726265038"
        - Part of the reason why the compiler started struggling is because my recursive hello world function may not actually be tail recursive. Oops.
Part 2:
    Which version of printLst do you consider to be "better"? Why?
        I would consider the single recursive function version of printLst to be better because since it does not have multiple "loops" (recursive functions) in it, the time complexity
        is lower. Both versions have a bunch of if statements, so I don't think the amount of if statements in my single recursive function printLst impacts it much.
        My personal preference, in terms of writing is the nested recursive function because I think it lines up better with how I would have approached the problem than the single recursive funciton. 
Part 3: 
    What is in the closure of the returned function?
        The closure of the returned function is an anonymous function where "powerOf()" is it's "operation." powerBy is the variable the closure captures.

Reflection:
===========
    My greatest enemy is sometimes myself. Which sounds melodramatic, is melodramatic, but is also true whether it be by overcomplicating the question or getting overwhelmed from the information when it's actually
    pretty simple. I think I need to ultiltize the rubber duck debugging more, because someone explaining what I needed to do helped a lot and hopefully that will translate when I'm trying to explain my problems.
    Surprisingly, the recursion was pretty easy and the programs themselves were easy to write once I figured out the desgin for them. I hope that's a sign that my CS skills are improving.

I Worked With:
==============
    N/A

Approximate Hours worked:
=========================
    7 hours 

Special Instructions to the grader:
===================================
    Both versions of the 2D array function are in here and demonstrated in the main of NestedFunctions.kt.

Known Bugs or Limitations:
==========================
    N/A

Other comments:
===============
    I do not know the program naming conventions for Kotlin, and if they're the same naming conventions as in Java I do not think I followed them well.