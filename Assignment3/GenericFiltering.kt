/**
* A program that demonstrates filters in Kotlin, but with Generics. 
* @author ayang
* Created: Oct 28, 2021
 */

fun main(args:Array<String>){
    // i did not know how to make a list/variable of type generic so i just made them any
    val ll: List<Any> = listOf(2,2.2,5,9,10,"hi","how do i supress val unused warnings","69", "nvm figured it out", "whoops made this harder than i needed to")
    println(myFilterGeneric(isNumeric, ll))
    println(myFilterGeneric(isInteger, ll))
    
    println()
    println("The ACTUAL demonstrations of the program working the way it was asked for:")

    // lists that actually fit the prompt that i misread.
    val listThatFitsThePrompt = listOf(1, 2, 3, 4 ,5 ,6 ,7, 8 ,9 ,10)
    val listThatFitsThePrompt2 = listOf(0.1, 2.2, 3.3, 4.4, 5.5, 0.9)
    println(myFilterGeneric(isOdd, listThatFitsThePrompt))
    println(myFilterGeneric(isDoubleBig, listThatFitsThePrompt2))

}

// checking if the element is a number or not
val isNumeric = {n: Any -> isNums(n)}
// checking if the element is an *integer* or not
val isInteger = {n: Any -> n is Int}
// is the number odd or not
val isOdd = {n: Int -> (n%2 != 0)}
// is the double greater than one
val isDoubleBig= {n: Double -> (n > 1.0)}

/**
* A function that, given a function by which to filter a list with, runs the given list through
* and returns a list with only the elements of the given list that pass the filter. Basically, myFilter
* in Filtering.kt, but is generic!
* @param: funcc (the function by which to filter a list with), lst (the list to filter with)
* @ret: ret (a list containing the elements of lst that passed funcc)
 */
fun <Q> myFilterGeneric(funcc: (Q) -> Boolean, lst:List<Q>) : List<Q> {
    // if the given list is empty, just return an empty list
    if(lst.isEmpty()){
        return emptyList()
    }
    
    // create a mutable list, but one that is of type generic
    val ret = mutableListOf<Q>()

    fun recFilter(listToFilter: List<Q>, pos: Int = 0)  {
        // once you reach the end, return the list with the filtered items
        if(pos > listToFilter.lastIndex){
            return
        }
        if(funcc(listToFilter.get(pos))){
            ret.add(listToFilter.get(pos))
        }
        return recFilter(listToFilter, pos + 1)
    }

    recFilter(lst)
    return ret
}

/**
* Assuming it's just primary types, checks to see whether the value is a number/decimal or not
* @param: questionableNums (the value to check)
* @return: true iff the value holds a number of some kind
 */
fun <Q> isNums(questionableNums: Q): Boolean {
    // check if it is a "number" and if it is just return true
    if(questionableNums is Float || questionableNums is Double || questionableNums is Int){
        return true
    }
    // if it is a boolean, return flase
    if(questionableNums is Boolean){
        return false
    }

    // if you reached this point without exiting, assume you're a string
    val questionableString = questionableNums as String

    // now, check the strings
    // and also ignore the compiler telling me i have variables i don't use. 
    // i know, but i don't know how to check without creating unused variables
    @Suppress("UNUSED_VARIABLE") 
    try{
        // check if it holds a double or not
        val doubleNum = questionableString.toDouble()
    } catch(e: NumberFormatException){
        // if it doesn't have a double, then check float
        try{
            val floatNum = questionableString.toFloat()
        } catch(e: NumberFormatException) {
            // if it doesn't have a double or float, then check int
            try{
                val intNum = questionableString.toInt()
            } catch(e: NumberFormatException){
                // if it's not an int, double, or float, return false
                return false
            }
        }
    }

    // if the string passed the above convoluted test, return true 
    return true
}