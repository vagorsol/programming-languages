/**
* A program demonstrating closures in Kotlin. 
* Demonstrates a function that returns a function that computes the [given integer]
* power of something
* @author ayang
* Created: Oct, 27, 2021
 */

// typealias Power = (Int) -> Int;

fun main(args:Array<String>){
    val pow4 = nthpower(4)
    println(pow4(2))
    println(pow4(3))
    println(pow4(4))
    println(pow4(5))
    
    println(nthpower(5) (2))
    println(nthpower(8) (2))
    // println(powerOf(2,5))
}

/**
* A function that, given a exponent to do a power by, returns a function that
* takes a base, and takes the power of the base using the previously given exponent
* @param: powerBy (the power)
* @return: the exponent of base to power
 */
fun nthpower(powerBy: Int): (Int) -> Int {
    return{ base: Int -> powerOf(base, powerBy)}
}

/**
* I counldn't figure out Kotlin's native exponent function so I made my own.
* Given a base and it's exponent, find the power.
* 10/29 Note: I figured out what I was doing wrong with my notation but I already wrote this program
* and I don't feel like doing probably very simple changes so it is staying.
 */
fun powerOf(base: Int, power: Int, count: Int = 1): Int {
    // base cases: power of 0, power of 1
    if(power == 0){
        return 1
    }
    if(power == 1){
        return base
    }
    // when it reaches the exponent, stop
    if(count == power){
        return base
    }
    return base * powerOf(base, power, count + 1)
}