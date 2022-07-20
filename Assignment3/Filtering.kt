/**
* A program that demonstrates filters in Kotlin. 
* Given a list, filter and return only the numbers that
* fufill the specified quality.
* Limitations: only works with positive numbers
* @author ayang
* Created: Oct 28, 2021
 */

fun main(args:Array<String>){
    val ll = listOf(1,2,3,4,5,6,7,8,9)
    
    println(ll.filter(isOdd))
    println(myFilter(isOdd, ll))
    println(myFilter(isEven, ll))
    println(myFilter(isPrime, ll))

    // checking that my "catch empty list" case works or not
    // val testEmpty: List<Int> = emptyList()
    // println(myFilter(isOdd, testEmpty))
    // println(primeChecker(12)) // testing if my prime tester Actually Works or not
}

// is the number odd or not
val isOdd = {n: Int -> (n%2 != 0)}
// is the number even or not
val isEven = {n: Int -> (n%2 == 0)}
// is the number prime or not
val isPrime = {n: Int -> primeChecker(n)}

/**
* A function that, given a function by which to filter a list with, runs the given list through
* and returns a list with only the elements of the given list that pass the filter
* @param: funcc (the function by which to filter a list with), lst (the list to filter with)
* @ret: ret (a list containing the elements of lst that passed funcc)
 */
fun myFilter(funcc: (Int) -> Boolean, lst:List<Int>) : List<Int> {
    // if given a list with no elements, give an "error message"
    if(lst.isEmpty()){
        println("Please give me a list with items in it!")
        // Kotlin makes me return something here >:(
        return emptyList() 
    }

    // a list for the items that fit the given filter factor
    val ret = mutableListOf<Int>()

    /**
    * A function to do the filtering. Goes through each element of the given list, checks it with the 
    * given function, and if it is true adds it to the filtered items list (to be returned later)
     */
    fun recFilter(listToFilter:List<Int>, pos: Int = 0) {
        // once you reach the end, return the list with the filtered items
        if(pos > listToFilter.lastIndex){
            return
        }
        if(funcc(listToFilter.get(pos))){
            ret.add(listToFilter.get(pos))
        }
        return recFilter(listToFilter, pos + 1)
    }

    // run the filter
    recFilter(lst)
    return ret
} 

/**
* A function that tests to see if a (positive) number is a prime number or not.
* @param: numCheck (the number you want to know is prime or not), 
*         divisor (number used to test if numCheck is prime or not)
* @return: true iff numCheck is prime
 */
fun primeChecker(numCheck: Int, divisor: Int = 5): Boolean {
    // if the number is one or negative, return false
    if(numCheck <= 1){
        return false
    }
    // if the number is 2 or 3, just return true
    if(numCheck == 2 || numCheck == 3){
        return true
    }
    // check two of the most common divisors (4 is 2 * 2 so we can just start at 5)
    if(numCheck % 2 == 0 || numCheck % 3 == 0) {
        return false
    }
    // if it has reached this far without being false, it is a prime number
    if (divisor * divisor > numCheck) {
        return true
    }
    // run the primality test
    if (numCheck % divisor == 0 || numCheck % (divisor + 2) == 0){
        if(numCheck != divisor){
            return false
        }
    }
    return primeChecker(numCheck, divisor + 1)
}