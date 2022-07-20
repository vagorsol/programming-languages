/**
* A program which flattens lists of unlimited depth
* @author ayang
* Created: Nov 1, 2021
 */

fun main(args:Array<String>){
    val lst: List<Any?> = listOf(12, 6, listOf(1, listOf(2), listOf(3,4)))
    println(flattenList(lst))
    println(flattenListToo(lst))
}

fun <T> flattenList(lst: List<T>): List<T> {
    // if the given list is empty, just return an empty list
    if(lst.isEmpty()){
        return emptyList()
    }
    
    // create a mutable list, but one that is of type generic
    val ret = mutableListOf<T>()

    fun siftList(lstToo: List<T>, pos: Int = 0){
        // once you reach the end of the list, return
        if(pos > lstToo.lastIndex){
            return
        }

        val testIdx = lstToo.get(pos)

        // if the item at pos is a list, run siftList on that list there
        @Suppress("TYPE_MISMATCH")
        if(testIdx is List<Any?>){
            siftList(testIdx)
        } else {
            // add the item to the return list
            // println(lstToo.get(pos))
            ret.add(lstToo.get(pos))
        }
        
        // go to the next item
        return siftList(lstToo, pos + 1)
    }

    //run the sifter, then return with the flattened list
    siftList(lst)
    return ret
}

// version of flatten that is recursive itself and has an optional parameter 
// i don't know how to fix it
fun <T> flattenListToo(lst: List<T>, pos: Int = 0): List<T> {
    // if the given list is empty, just return an empty list
    if(lst.isEmpty()){
        return emptyList()
    }
   
    // create a mutable list, but one that is of type generic
    val ret = mutableListOf<T>()
    
    // to check internal list
    fun babyList(lstToo: List<T>, index: Int = 0){
        val testIdxToo = lstToo.get(index)
        // if you reach the end of the list, return
        if(index > lstToo.lastIndex){
            return
        } else if(testIdxToo is List<Any?>){
            @Suppress("TYPE_MISMATCH")
            babyList(testIdxToo)
        } else{
            ret.add(lstToo.get(index))
        }
    }

    val testIdx = lst.get(pos)

    @Suppress("TYPE_MISMATCH")
    if(testIdx is List<Any?>){
        babyList(testIdx)
    } else {
        ret.add(lst.get(pos))
    }
    return flattenListToo(lst, pos + 1)
}