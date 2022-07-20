/**
* A version of Hello World that attempts to parse an integer, 
* then prints out "hello world" that many times, recursively
* @author ayang
* Created: Oct 22, 2021
 */

fun main(args:Array<String>){
    val enteredValue = args[0]
    // val enteredValue = "1000000000" // this is just so i dont have to write the whole "cd to my external hard drive" and etc. because AHH!!!
    
    // check if it's actually an interger and if so, execute the function. if not, error message.
    try{
        val ittAmount = enteredValue.toInt()
        recurseHW(ittAmount)
    } catch(e: NumberFormatException) {
        println("Please enter an integer!")
    }
}

/**
// pseudo-code for the recursive part of the recusive hello world program
func name (int [default 0] / count, itteration int)
if int == itteration int
return
println hello world
return (count + 1, itteration int)
 */

/**
* The function that prints out "Hello, world!" the given amount of times
* using recursion to keep track of how many times it's printed it out already.
* Once it hits the given limit, it stops
* Outputs an integer with each "Hello, world!" to track of how many itterations it goes for.
 */
fun recurseHW(ittAmount: Int, count: Int = 0) {
    // once it hits the given itteration amount, it stops
    if (count == ittAmount) {
        return
    }
    // base case for in case the user gives a negative number
    if (ittAmount < 0){
        println("Please give me a positive itteration amount!")
    }
    println("Hello, World! $count")
    return recurseHW(ittAmount, count + 1)
}