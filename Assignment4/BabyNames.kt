/**
    A program that, when the user gives a name or set of names, checks a "database" 
    (a set of files with information on names for many years) and returns information about
    that name if it finds it. In addition, returns results for names that contain that given name
    as a prefix.
    
    @author ayang
    Created: Nov 15, 2021
 */

import java.io.*
import java.net.*

/**
    Defines a class NameHolder that contains all the information for a given name
    @param: name (the name), uses (how many babies have that name)
 */
class NameHolder (val name: String, val amount: String) {
    var years = 1 
    var totalAmount: Int = amount.toInt() // making it easier to just parse now bc everything in the excel is string. fun!

    // add a year to the list of years
    fun addYear() {
        years = years + 1
    }

    // adds to the total
    fun addAmount(numToAdd: String) {
        totalAmount = totalAmount + numToAdd.toInt()
    }
}

/**
    Stores the names given in the list for each year into two lists: "GirlNames" and "BoyNames"
    @param: name (row of names to store info about), boyLst and girlLst (lists to store the name information about)
*/
fun nameStorer(name: String, boyLst: MutableList<NameHolder>, girlLst: MutableList<NameHolder>) {
    // turn the name into like. something that I can call indexes on.
    val nameLst: List<String> = name.split(",").map { it.trim() }

    // if the list is empty, just add it. checking for both because. security. could honestly probably just "&&" but. "security"
    if(!boyLst.isNotEmpty()){
        boyLst.add(NameHolder(nameLst[1], nameLst[2]))
    }
    if(!girlLst.isNotEmpty()){
        girlLst.add(NameHolder(nameLst[3], nameLst[4]))
    }

    // do for the boy list
    if(containsName(nameLst[1], boyLst) > -1){
        // if the name is in the list, just update the metrics
        val idx = containsName(nameLst[1], boyLst)
        boyLst[idx].addYear()
        boyLst[idx].addAmount(nameLst[2])
    } 
    else {
        // otherwise, create a new item and add it to the list
        addAlphabetically(NameHolder(nameLst[1], nameLst[2]), boyLst)
    }
    // do for the girl list
    if(containsName(nameLst[3], girlLst) > -1){
        // if the name is in the list, just update the metrics
        val idx = containsName(nameLst[3], girlLst)
        girlLst[idx].addYear()
        girlLst[idx].addAmount(nameLst[4])
    } 
    else {
        // otherwise, create a new item and add it to the list
        addAlphabetically(NameHolder(nameLst[3], nameLst[4]), girlLst)
    }
}

/**
    A function that searches for a given name, and if it or ones that contains it as a prefix
*/
fun nameSearch(name: String, lst: MutableList<NameHolder>, gender: String) {
    // search for the name & names that contain it as a prefix
    val prefixLst = containsPrefix(name, lst)

    // if the name is not in the given list and has no prefixes, do not say anything. simply leave. do not say anything. pass go, do not collect $200.
    if(!prefixLst.isNotEmpty()){
        return
    }

    // print out what the user searched for
    println(name)
    // if there are names that have the prefix, then print out the information for those names
    prefixLst.forEach(){
        returnNameInfo(lst, gender, it)
    }
}

/**
    checks to see if the name is in the list. 
    @param: name (the name you're checking for), lst (list your checking in), count (to itterate)
    @return: count iff true (i.e., the name occurs in the list, so return the index for ease)
             else, -1 (an arbitray number that count probably can't be)
*/ 
fun containsName(name: String, lst: MutableList<NameHolder>, count: Int = 0): Int{
    // if you reach the end, it's not in the list and return false (-1)
    if(count > lst.lastIndex){
        return -1
    }

    // if the name is equal to the name at the index, return the index
    if(name.equals(lst[count].name, ignoreCase = true)){
        return count
    }

    return containsName(name, lst, count + 1)
}

/**
    Returns a list of indexes of where there are names with the prefix
    @param: name (name/prefix searching for), lst (list searching from)
    @return: the list of indexes that contain the name/prefix. else, return an empty list
 */ 
fun containsPrefix(name: String, lst: MutableList<NameHolder>): MutableList<Int> {
    // the list of indexes that you will return
    val retLst: MutableList<Int> = mutableListOf()

    // does the get list of indexes 
    fun checkPrefixGetIndx(name: String, lst: MutableList<NameHolder>, count: Int = 0){
        if(count > lst.lastIndex){
            return
        }
        if(lst[count].name.length >= name.length){
            val prefixName = (lst[count].name).substring(0, name.length)
            // println(prefixName)
            if(prefixName.contains(name, ignoreCase = true)){
                // println("check")
                retLst.add(count)
            }
        }
        return checkPrefixGetIndx(name, lst, count + 1)
    }
    checkPrefixGetIndx(name, lst)
    // println(retLst)
    return retLst
}

/**
    Adds an item in a list so that it preserves an alphabetical order
    @param nameChunk (the name that wants to be checked, and the item to be added), lst (the list to be added to)
*/
fun addAlphabetically(nameChunk: NameHolder, lst: MutableList<NameHolder>, count: Int = 0){
    // if you reach the end, add the name at the end
    if(count > lst.lastIndex){
        lst.add(nameChunk)
        return
    }
    // println(name.compareTo(lst[count]))
    // println(count)

    // if the name is alphabetically before a name at a given index, put it before (i.e. at, insert) that index
    if(nameChunk.name.compareTo(lst[count].name) < 0){
        lst.add(count, nameChunk)
        return
    }
    // otherwise, KEEP GOING
    return addAlphabetically(nameChunk, lst, count + 1)
}

/**
    A function that returns information about a given name (aka a more fancy toString but function)
    @param: lst (that you're getting information about the name from), gender (gender of the name), indx (the index of the name)
 */
fun returnNameInfo(lst: MutableList<NameHolder>, gender: String, indx: Int){
    // get the denomenator for the calc of how many kids are in the lists
    val totalKids = lst.map{it.totalAmount}.fold(0) {total, next -> total + next }

    // calculate the percentage
    val percent = ((lst[indx].totalAmount.toDouble()) / (totalKids.toDouble())) * 100.0

    val rank = indx + 1
    
    // print out the findings
    // println(lst[indx].name) // to match the formating on the site
    println(String.format("%s: %s \t %s: %s \t %s: %d", "Name", lst[indx].name, "Gender",gender, "# of Years",lst[indx].years))
    println(String.format("%s: %d", "Alphabetical Position", rank))
    println(String.format("%s: %d", "Occurences", lst[indx].totalAmount)) 
    println(String.format("%s %s %s: %f","Percent of All", gender, "Names", percent))
    println() // because FORMATING
}

fun main(args:Array<String>){
    // "do it so it's mutable only w/in a function" geoff i'm not that clever idk how to do that
    val BoyNames: MutableList<NameHolder> = mutableListOf()
    val GirlNames: MutableList<NameHolder> = mutableListOf()
    val NameSearchFor: MutableList<String> = mutableListOf()

    // get the user inputs
    args.forEach(){
        NameSearchFor.add(it)
    }
    
    // function to read the url
    val readURL = { url:String -> 
        val r = mutableListOf<String>();
        try {
            val oracle = URL(url);
            val br = BufferedReader(InputStreamReader(oracle.openStream()));
            loop@ while (true) {
                val line = br.readLine();
                if (line==null) {
                    break@loop
                }
                r.add(line);
            }
        } catch (ee: Exception) {
            println("Problem ${ee}");
        }
        r
    } 
    
    // would it be easier to just manually write it out? probably. am I doing that? no. please help.
    val filesLst = listOf("https://cs.brynmawr.edu/cs245/Data/names2000.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2001.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2002.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2003.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2004.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2005.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2006.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2007.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2008.csv",
        "https://cs.brynmawr.edu/cs245/Data/names2009.csv"
    )   

    // list to hold the read files
    val readFiles: MutableList<MutableList<String>> = mutableListOf()
    
    // read all the files
    filesLst.forEach(){
        readFiles.add(readURL(it))
    }

    // put all the name info into the two lists
    readFiles.forEach(){
        it.map{nameStorer(it, BoyNames, GirlNames)}
    }

    // run the search for all the names
    NameSearchFor.forEach(){
        nameSearch(it, BoyNames, "Boy")
        nameSearch(it, GirlNames, "Girl")
    }

    /*
        // test to see how it works
        val lst2000Names = readURL("https://cs.brynmawr.edu/cs245/Data/names2000.csv")
        val lst2001Names = readURL("https://cs.brynmawr.edu/cs245/Data/names2001.csv")

        // println(lst2000Names[0]) // is a string   
        lst2000Names.map{nameStorer(it, BoyNames, GirlNames)} // -> foreach or itt over for them all and then at the end: metrics
        lst2001Names.map{nameStorer(it, BoyNames, GirlNames)}

        // testiiiing 
        nameSearch("aaron", BoyNames, "Boy")
        nameSearch("dAN", BoyNames, "Boy")
        nameSearch("Marlen", GirlNames, "Girl")
        nameSearch("syd", GirlNames, "Girl")
    */
}