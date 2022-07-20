/**
* A demonstration of NestedFunctions in Kotlin. Given a List of Lists, print out the contents in lines 
* based on the order the Inner Lists were given. 
* @author ayang
* Created: Oct 26, 2021
 */

fun main(args:Array<String>){
    val lst = listOf(listOf(1, 2, 3), listOf(4, 5, 6), listOf(7, 8, 9))
    val lstC = listOf(listOf('c', 'd', 'e'), listOf('g','h','i'), listOf('j','l','m','n'))
    val lstS = listOf(listOf("c", "d", "e", "A"), listOf("g","h","i"), listOf("j","l","m","n", "1", "2", "3"))

    // demonstration of printLst (nested recursion) 
    println("printLst\n")  
    printLst(lst)
    println()
    printLst(lstC)
    println()
    printLst(lstS)
    println()

    // demonstration of printLst1R (no nested recursion version)
    println("printLst1R\n")
    printLst1R(lst)
    println()
    printLst1R(lstC)
    println()
    printLst1R(lstS)
    println()
}

/**
* A function that takes a list, and outputs it in lines in the order it is given in. Demonstrates nested recursive 
* functions in Kotlin.
* param: input (list of items to output in an "array"), pos (where you are in the list, default 0/"the beginning")
 */
fun printLst(input: List<List<Any>>, pos: Int = 0) {
    // if given a list  with no elements
    if(input.isEmpty()){
        println("Please give me a list!") // "but what if I want nothing printed out!" well too bad
        return
    }
    // if the user tries to start at a location not in the list
    if(pos < 0 || pos > input.size){
        println("Please start at a point in the list!")
        return
    }
    // exit if finished last line
    if(pos > input.lastIndex){
        return
    }

    // recursively print a "line" of the list
    fun printLstLine(innerList: List<Any>, index: Int = 0){
        // checks if it's past the end of the inner list
        if(index > innerList.lastIndex){
            println()
            return
        }
        print(innerList.get(index))
        print(" ") // I am refusing to figure out how to concatenate at this time
        return printLstLine(innerList, index + 1)
    }

    // invoke printLstLine
    printLstLine(input.get(pos))
    // go to the next line of the List
    return printLst(input, pos + 1)   
}

/**
* A varation of printLst that only uses one recursive function (no nested function)
* the "1R" stands for "one recusion" because it has only one recursive function in it.
 */
fun printLst1R(input: List<List<Any>>, pos: Int = 0, index: Int = 0){
    // if given a list  with no elements
    if(input.isEmpty()){
        println("Please give me a list!") // "but what if I want nothing printed out!" well too bad
        return
    }
    // if the user tries to start at a location not in the list
    if(pos < 0 || pos > input.size){
        println("Please start at a point in the list!")
        return
    }
    
    // val innerList = input.get(pos)
    print(input.get(pos).get(index))
    print(" ")

    // check for if reached the end of the inner list (if the next step is out of bounds)
    if(index + 1 > input.get(pos).lastIndex){
        // if you reached the end of the list, exit
        if(pos  + 1 > input.lastIndex){
            println()
            return
        }
        // otherwise, new line and go to the next element
        println()
        return printLst1R(input, pos + 1)
    }
    return printLst1R(input, pos, index + 1)
}