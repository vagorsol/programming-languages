/**
* K to M conversion in Kotlin
* @author ayang
* Created: October 18, 2021
 */
fun main(args: Array<String>){
    val dist:Double = args[0]
    val unit:String = args[1]
    // val dist:Double = 10.0
    // val unit:String = "k"

    if(unit.equals("m") || unit.equals("M")){
        // converting from miles to kilometers
        val convDist:Double = dist * 1.609
        println("$convDist k")
    } else {
        // converting from kilometers to miles
        val convDist:Double = dist / 1.609
        println("$convDist m")
    }
}