NAME:
=====
	Audrey Yang

Programs Files:
===============
    - BabyNames.go
    - Essay.pdf (written response)
	
How to Compile & Run:
====================
    - go run BabyNames.go
	
Reflection:
===========
    Converting to Golang was suprisingly simple, but very wordy given my reliance on for-loops in place of many recursive functions. I still feel like there must be
    a better implementation than what I did, but it works and I'm not asked for space efficiency, yet. I did get frustrated since I had to review quite a bit of Go
    to be able to write the program. I also ran into frustration by having to grapple with Golang being pass by reference with regards to modifying the value of slices 
    since that is opposite of how modifying slices intuitively works.
    I had a considerably rougher time debugging compared to my Kotlin implementation. Whether that was because of the imperative programming style, or because I did not
    break things up in chunks to test (which, is probably related to the imperative programming style but I'm fairly sure I did this with my Kotlin program as well), it
    is unclear. I was told "used named variables for constants so that doesn't happen," but given that I had almost left the wrong variable when copying and modifying lines,
    it remains to be seen whether or not that would help preventing bugs from my stupid typos.

I Worked With:
==============
    Neha: for reminding me how sort works in Golang, and for the idea that I should use sort and not try and do whatever madness I did in my Kotlin implementation 
        and emotional support as I tried finding a bug that turned out to be because of one (1) typo. Fun fact: this was exactly the same typo (nameLst[1] instead of 
        nameLst[3] in nameStorer) that appeared in a previous bug!

Approximate Hours worked:
=========================
    6 hours (+ 1 hr for essay writing)

Special Instructions to the grader:
===================================
    N/A

Known Bugs or Limitations:
==========================
    N/A

Other comments:
===============
    - Sacrified comment aesthetic-ness for the legibility of the hover-over-function information :(
    - Used "list" pretty interchangably with "slice." Apologies for the confusion in advance.