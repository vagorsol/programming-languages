/**
* Fizz Buzz in Kotlin
* @author ayang
* Created: October 18, 2021
 */

fun main(args: Array<String>) {
    p(100)
} 

tailrec fun p(i:Int) {
    if(i == 0){
        // base case
        return
    }
    
    if(i%3 == 0 && i%5 ==0){
        // if it is divisible by 3 AND 5
        println("FizzBuzz")
    } else if(i%3 == 0 && i%5 != 0) {
        // if it is divisible by 3 but NOT 5
        println("Fizz")
    } else if(i%3 != 0 && i%5 == 0){
        // if it is divisible by 5 but NOT 5
        println("Buzz")
    } else {
        println(i)
    }
    return p(i - 1)
}